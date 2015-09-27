package orders

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"text/template"
	"time"
	// "encoding/json"
	"database/sql"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/matejkramny/gopos"
	"lab.castawaylabs.com/orderchef/database"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/util"
)

var billReceipt *template.Template

func init() {
	db := database.Mysql()

	billTemplate, err := db.SelectStr("select value from config where name='customer_bill'")
	if err != nil {
		fmt.Println("Cannot find customer_bill in config table")
		return
	}

	tmpl, err := template.New("customer_bill").Parse(billTemplate)
	if err != nil {
		fmt.Println("Cannot compile customer_bill template")
		return
	}

	billReceipt = tmpl
}

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

	methods, _ := db.Select(models.OrderBill{}, "select sum(total) as paid_amount, payment_method_id from order__bill_payment join order__bill as bill on bill.id = order__bill_payment.bill_id where bill.group_id=? group by payment_method_id", group.Id)

	c.JSON(200, map[string]interface{}{
		"paid":  methods,
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

	// if bill.Paid = false; roundPrice(float64(bill.PaidAmount)) == roundPrice(float64(bill.Total)) && bill.Total > 0 {
	// 	bill.Paid = true
	// }

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
	if billReceipt == nil {
		c.AbortWithStatus(500)
		return
	}

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

	var postData struct {
		Covers   int     `json:"covers"`
		PerCover float32 `json:"perCover"`
	}

	if err := c.Bind(&postData); err != nil {
		c.JSON(400, err)
		return
	}

	for i := 0; i < postData.Covers; i++ {
		bill := models.OrderBill{GroupID: group.Id, CreatedAt: time.Now(), Total: postData.PerCover}
		if err := db.Insert(&bill); err != nil {
			panic(err)
		}

		item := models.OrderBillItem{
			BillID:    bill.ID,
			ItemName:  "-",
			ItemPrice: postData.PerCover,
			Discount:  0,
		}
		if err := db.Insert(&item); err != nil {
			panic(err)
		}
	}

	c.AbortWithStatus(204)
}

func setBillPayment(c *gin.Context) {
	var payment models.OrderBillPayment
	if err := c.Bind(&payment); err != nil {
		c.JSON(400, payment)
		return
	}

	if _, err := database.Mysql().Exec("replace into order__bill_payment set bill_id=?, payment_method_id=?, amount=?", payment.BillID, payment.PaymentMethodID, payment.Amount); err != nil {
		panic(err)
	}

	c.AbortWithStatus(204)
}

func getBillPayments(c *gin.Context) {
	bill := c.MustGet("bill").(models.OrderBill)
	var payments []models.OrderBillPayment

	if _, err := database.Mysql().Select(&payments, "select * from order__bill_payment where bill_id=?", bill.ID); err != nil {
		panic(err)
	}

	c.JSON(200, payments)
}
