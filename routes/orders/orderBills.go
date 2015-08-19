package orders

import (
	"math"
	"log"
	"time"
	"text/template"
	"bytes"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"database/sql"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/util"
	"lab.castawaylabs.com/orderchef/database"
	"github.com/gin-gonic/gin"
)

var billReceipt = template.Must(template.New("billReceipt").Parse(`[[justify 1]][[font 1]]Bill receipt[[lf]][[justify 2]]
justified just right

[[justify 0]]{{range .items}} {{.ItemName}}[[justify 2]]£{{.ItemPrice}}
[[justify 0]]{{end}}

[[emphesize 1]]Total:[[justify 2]]£{{.total}}

[[cut]]`))

// get totals - items that are paid, amounts
func getBillTotals(c *gin.Context) {
	db := database.Mysql()
	group, err := getGroupById(c)
	if err != nil {
		return
	}

	total, err := db.SelectFloat("select sum(item.price) from order__group_member join order__item as oi on oi.order_id=order__group_member.id join item on item.id=oi.item_id where group_id=?", group.Id)
	if err != nil {
		total = 0
	}

	totalModifiers, err := db.SelectFloat("select sum(cm.price) from order__group_member join order__item as oi on oi.order_id=order__group_member.id join order__item_modifier as oim on oim.order_item_id=oi.id join config__modifier as cm on cm.id=oim.modifier_id where order__group_member.group_id=?", group.Id)
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

	if err := bill.GetItems(); err != nil {
		panic(err)
	}

	printData := map[string]interface{}{}
	printData["total"] = bill.Total
	printData["items"] = bill.Items

	buf := new(bytes.Buffer)
	billReceipt.Execute(buf, printData)

	data := map[string]interface{}{
		"print": buf.String(),
	}

	jsonBlob, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	num_clients, err := redis.Int(redis_c.Do("PUBLISH", "oc_print.receipt", string(jsonBlob)))
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