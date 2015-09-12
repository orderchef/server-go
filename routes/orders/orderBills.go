package orders

import (
	"math"
	"log"
	"fmt"
	"time"
	"text/template"
	"bytes"
	// "encoding/json"
	"github.com/garyburd/redigo/redis"
	"database/sql"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/util"
	"lab.castawaylabs.com/orderchef/database"
	"github.com/gin-gonic/gin"
	"github.com/matejkramny/gopos"
)

var billReceipt = template.Must(template.New("billReceipt").Parse(`[[justify 1]]    [[emphesize true]]Address:[[emphesize false]] 100 Cowley Road
             OX4 1JE Oxford
[[emphesize true]]Telephone:[[emphesize false]] 01865 434100
 [[emphesize true]]Facebook:[[emphesize false]] TaberuOxford

Printed [[emphesize true]]{{.time}}[[emphesize false]]
Bill [[emphesize true]]#{{.billID}}[[emphesize false]]
Table [[emphesize true]]{{.table_name}}[[emphesize false]]
[[lf]]
[[justify 0]]
{{range .items}}{{.ItemName}}[[spaces "{{.ItemName}}" "{{.ItemPriceFormatted}}"]][[at]]{{.ItemPriceFormatted}}
[[justify 0]]{{end}}
[[justify 2]][[emphesize true]]Total:[[emphesize false]] [[at]]{{.totalFormatted}}
[[lf]]
[[lf]]
[[justify 1]]Service charge not included
[[justify 0]][[lf]]
[[lf]]
[[lf]]
[[cut]]`))

// get totals - items that are paid, amounts
func getBillTotals(c *gin.Context) {
	db := database.Mysql()
	group, err := getGroupById(c)
	if err != nil {
		return
	}

	total, err := db.SelectFloat("select sum(item.price * oi.quantity) from order__group_member join order__item as oi on oi.order_id=order__group_member.id join item on item.id=oi.item_id where group_id=?", group.Id)
	if err != nil {
		total = 0
	}

	totalModifiers, err := db.SelectFloat("select sum(cm.price * oi.quantity) from order__group_member join order__item as oi on oi.order_id=order__group_member.id join order__item_modifier as oim on oim.order_item_id=oi.id join config__modifier as cm on cm.id=oim.modifier_id where order__group_member.group_id=?", group.Id)
	if err != nil {
		totalModifiers = 0
	}

	total += totalModifiers

	methods, _ := db.Select(models.OrderBill{}, "select sum(total) as paid_amount, payment_method_id from order__bill where paid=? and group_id=? group by payment_method_id", true, group.Id)

	c.JSON(200, map[string]interface{}{
		"paid": methods,
		"total": total,
	})
}

 // get all bills
func getAllBills(c *gin.Context) {
	db := database.Mysql()
	group, err := getGroupById(c)
	if err != nil {
		return
	}

	var bills []*models.OrderBill
	if _, err := db.Select(&bills, "select * from order__bill where group_id=?", group.Id); err != nil {
		log.Println(err)
		util.ServeError(c, err)

		return
	}

	for _, bill := range bills {
		if err := bill.GetItems(); err != nil {
			util.ServeError(c, err)
			return
		}
	}

	c.JSON(200, bills)
}

func getBill(c *gin.Context) {
	db := database.Mysql()
	bill := models.OrderBill{}

	err := db.SelectOne(&bill, "select * from order__bill where id=?", c.Params.ByName("bill_id"))

	if err == sql.ErrNoRows {
		c.AbortWithStatus(404)
		return
	} else if err != nil {
		util.ServeError(c, err)
		return
	}

	if err := bill.GetItems(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.Set("bill", bill)
	c.Next()
}

func serveBill(c *gin.Context) {
	c.JSON(200, c.MustGet("bill"))
}

// create new bill
func makeBill(c *gin.Context) {
	db := database.Mysql()
	group, err := getGroupById(c)
	if err != nil {
		return
	}

	bill := models.OrderBill{GroupID: group.Id, CreatedAt: time.Now()}
	if err := db.Insert(&bill); err != nil {
		panic(err)
	}

	c.JSON(200, bill)
}

func roundPrice(price float64) float64 {
	price *= 100

	decimals, _ := math.Modf(price)
	if decimals > 0.5 {
		price = math.Ceil(price)
	} else {
		price = math.Floor(price)
	}

	return price / 100
}

// update bill
func updateBill(c *gin.Context) {
	db := database.Mysql()
	bill := c.MustGet("bill").(models.OrderBill)

	if err := c.Bind(&bill); err != nil {
		c.JSON(400, err)
		return
	}

	if _, err := db.Exec("delete from order__bill_item where bill_id=?", c.Params.ByName("bill_id")); err != nil {
		panic(err)
	}

	for _, item := range bill.Items {
		item.BillID = bill.ID
		if err := db.Insert(&item); err != nil {
			panic(err)
		}
	}

	if bill.Paid = false; roundPrice(float64(bill.PaidAmount)) == roundPrice(float64(bill.Total)) && bill.Total > 0 {
		bill.Paid = true
	}

	if _, err := db.Update(&bill); err != nil {
		panic(err)
	}

	c.AbortWithStatus(204)
}

func deleteBill(c *gin.Context) {
	db := database.Mysql()

	bill := c.MustGet("bill").(models.OrderBill)

	if _, err := db.Exec("delete from order__bill_item where bill_id=?", bill.ID); err != nil {
		panic(err)
	}

	if _, err := db.Delete(&bill); err != nil {
		panic(err)
	}

	c.AbortWithStatus(204)
}

func printBill(c *gin.Context) {
	db := database.Mysql()
	redis_c := database.Redis.Get()
	defer redis_c.Close()

	bill := c.MustGet("bill").(models.OrderBill)

	var table models.Table
	if err := db.SelectOne(&table, "select name from table__items join order__group on order__group.table_id=table__items.id where order__group.id=?", bill.GroupID); err != nil {
		table = models.Table{}
	}

	if err := bill.GetItems(); err != nil {
		panic(err)
	}

	printData := map[string]interface{}{}
	printData["time"] = time.Now().Format("02-01-2006 15:04")
	printData["billID"] = bill.ID
	printData["total"] = bill.Total
	printData["totalFormatted"] = fmt.Sprintf("%.2f", bill.Total)
	printData["items"] = bill.Items
	printData["table_name"] = table.Name

	buf := new(bytes.Buffer)
	billReceipt.Execute(buf, printData)

	buffer := gopos.RenderTemplate(buf.String())

	// data := map[string]interface{}{
	// 	"print": buffer.String(),
	// }

	// jsonBlob, err := json.Marshal(data)
	// if err != nil {
	// 	panic(err)
	// }

	num_clients, err := redis.Int(redis_c.Do("PUBLISH", "oc_print.receipt", buffer.String()))
	if err != nil {
		panic(err)
	}

	now := time.Now()
	bill.PrintedAt = &now

	if _, err := db.Update(&bill); err != nil {
		panic(err)
	}

	if num_clients == 0 {
		// NOT Printed
		// send err response.
		// 503 = service unavailable
		c.AbortWithStatus(503)
		return
	}

	c.Writer.WriteHeader(204)
}

func splitBills(c *gin.Context) {
	db := database.Mysql()
	group, _ := getGroupById(c)

	var postData struct{
		Covers int `json:"covers"`
		PerCover float32 `json:"perCover"`
	}

	if err := c.Bind(&postData); err != nil {
		c.JSON(400, err)
		return
	}

	for i := 0; i < postData.Covers; i++ {
		bill := models.OrderBill{GroupID: group.Id, CreatedAt: time.Now(), BillType: "amount", Total: postData.PerCover}
		if err := db.Insert(&bill); err != nil {
			panic(err)
		}

		item := models.OrderBillItem{
			BillID: bill.ID,
			ItemName: "-",
			ItemPrice: postData.PerCover,
			Discount: 0,
		}
		if err := db.Insert(&item); err != nil {
			panic(err)
		}
	}

	c.AbortWithStatus(204)
}