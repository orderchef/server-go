package orders

import (
	"time"
	"bytes"
	"text/template"
	// "encoding/json"
	"database/sql"
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/util"
	"lab.castawaylabs.com/orderchef/database"
	"github.com/garyburd/redigo/redis"
	"github.com/matejkramny/gopos"
)

var kitchenReceipt = template.Must(template.New("kitchenReceipt").Parse(`[[lf]][[justify 0]]
{{.time}}
Table: {{.table_name}}. Order #{{.order.Id}}
{{range .items}}---------------
{{.itemObject.Name}}
{{range .modifiers}} - {{.group.Name}} ({{.modifier.Name}})
{{end}}---------------
{{end}}
[[lf]]
[[lf]]
[[lf]]
[[cut]]`))

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
				"group": modifierGroup,
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

		for printerID, _ := range final_printers {
			if printers[printerID] == nil {
				printers[printerID] = map[string]interface{}{
					"items": []map[string]interface{}{},
					"modifiers": []map[string]interface{}{},
					"time": time.Now().Format("15:04:05"),
					"order": order,
					"table_name": table_name,
				}
			}

			printers[printerID]["items"] = append(printers[printerID]["items"].([]map[string]interface{}), map[string]interface{}{
				"item": item,
				"modifiers": modifiers,
				"itemObject": itemObject,
			})
		}
	}

	for printer, o := range printers {
		buf := new(bytes.Buffer)
		kitchenReceipt.Execute(buf, o)
		buffer := gopos.RenderTemplate(buf.String())

		// data := map[string]interface{}{
		// 	"print": buf.String(),
		// }

		// jsonBlob, err := json.Marshal(data)
		// if err != nil {
		// 	panic(err)
		// }

		// print it!
		num_clients, err := redis.Int(redis_c.Do("PUBLISH", "oc_print." + printer, buffer.String()))
		if err != nil {
			panic(err)
		}

		if num_clients == 0 {
			// NOT Printed
			not_printed = append(not_printed, printer)
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