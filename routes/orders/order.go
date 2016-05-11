package orders

import (
	"bytes"
	"fmt"
	"strconv"
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

var kitchenReceipt *template.Template
var kitchenReceiptQuantity int

func init() {
	CompileKitchenReceipt()
}

func CompileKitchenReceipt() {
	db := database.Mysql()
	kitchenReceiptQuantity = 1

	templateQuantity, _ := db.SelectStr("select value from config where name='kitchen_receipt_quantity'")
	q, err := strconv.Atoi(templateQuantity)
	if err == nil && q > 0 {
		kitchenReceiptQuantity = q
	}

	templateString, err := db.SelectStr("select value from config where name='kitchen_receipt'")
	if err != nil {
		fmt.Println("Cannot find kitchen_receipt in config table")
		return
	}

	tmpl, err := template.New("kitchen_receipt").Parse(templateString)
	if err != nil {
		fmt.Println("Cannot compile kitchen_receipt template")
		return
	}

	kitchenReceipt = tmpl
}

func getOrderById(c *gin.Context) {
	order_id, err := util.GetIntParam("order_id", c)
	if err != nil {
		util.ServeError(c, err)
		return
	}

	order := models.Order{Id: order_id}
	err = order.Get()
	if err == sql.ErrNoRows {
		util.ServeNotFound(c)
		return
	} else if err != nil {
		util.ServeError(c, err)
		return
	}

	c.Set("order", order)
	c.Set("orderId", order_id)
	c.Next()
}

func GetOrder(c *gin.Context) {
	order := c.MustGet("order")

	c.JSON(200, order)
}

func GetOrderItems(c *gin.Context) {
	order := c.MustGet("order").(models.Order)

	items, err := order.GetItems()
	if err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(200, items)
}

func DeleteOrder(c *gin.Context) {
	order := c.MustGet("order").(models.Order)

	if err := order.Remove(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}

func PrintOrder(c *gin.Context) {
	if kitchenReceipt == nil {
		c.AbortWithStatus(500)
		return
	}

	db := database.Mysql()
	redis_c := database.Redis.Get()
	defer redis_c.Close()

	// get Order
	order := c.MustGet("order").(models.Order)

	// get table name
	table_name := order.GetTableName()

	// get order items
	orderItems, err := order.GetItems()
	if err != nil {
		util.ServeError(c, err)
		return
	}

	// printers are mapped by [id]{items, modifiers....}
	printers := map[string]map[string]interface{}{}
	not_printed := []string{}

	for _, item := range orderItems {
		// get the actual item
		itemObject := models.Item{Id: item.ItemId}
		_ = itemObject.Get()

		// item modifiers
		modifiers := []map[string]interface{}{}

		orderModifiers, _ := item.GetModifiers()
		for _, orderModifier := range orderModifiers {
			// get actual modifiers
			modifier := models.ConfigModifier{Id: orderModifier.ModifierId}
			_ = modifier.Get()
			modifierGroup := models.ConfigModifierGroup{Id: orderModifier.ModifierGroupId}
			_ = modifierGroup.Get()

			modifiers = append(modifiers, map[string]interface{}{
				"modifier": modifier,
				"group":    modifierGroup,
			})
		}

		// Find printers to print to
		printer_models := []models.CategoryPrinter{}
		if _, err := db.Select(&printer_models, "select * from category_printer where (category_id=? or item_id=?)", itemObject.CategoryId, itemObject.Id); err != nil {
			if err == sql.ErrNoRows {
				continue
			}

			panic(err)
		}

		// find if printers have been overridden by item
		itemOverride, _ := db.SelectInt("select count(1) from category_printer where item_id=?", itemObject.Id)

		// map is used as a hash bucket to prevent duplicates in slices
		final_printers := map[string]int{}
		for _, printerModel := range printer_models {
			if (itemOverride > 0 && printerModel.ItemID != nil) || (itemOverride == 0 && printerModel.CategoryID != nil) {
				final_printers[printerModel.PrinterID] = 0
			}
		}

		for printerID := range final_printers {
			if printers[printerID] == nil {
				printers[printerID] = map[string]interface{}{
					"items":      []map[string]interface{}{},
					"modifiers":  []map[string]interface{}{},
					"time":       time.Now().Format("15:04:05"),
					"order":      order,
					"table_name": table_name,
				}
			}

			printers[printerID]["items"] = append(printers[printerID]["items"].([]map[string]interface{}), map[string]interface{}{
				"item":       item,
				"modifiers":  modifiers,
				"itemObject": itemObject,
			})
		}
	}

	for printer, o := range printers {
		buf := new(bytes.Buffer)
		kitchenReceipt.Execute(buf, o)
		buffer := gopos.RenderTemplate(buf.String())

		for i := 0; i < kitchenReceiptQuantity; i++ {
			// print it!
			num_clients, err := redis.Int(redis_c.Do("PUBLISH", "oc_print."+printer, buffer.String()))
			if err != nil {
				panic(err)
			}

			if num_clients == 0 {
				// NOT Printed
				not_printed = append(not_printed, printer)
			}
		}
	}

	now := time.Now()
	order.PrintedAt = &now
	order.Save()

	if len(not_printed) > 0 {
		// send err response.
		// 503 = service unavailable
		c.JSON(503, not_printed)
		return
	}

	c.Writer.WriteHeader(204)
}
