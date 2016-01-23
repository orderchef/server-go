package reports

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/database"
	"lab.castawaylabs.com/orderchef/models"
)

func getExtrasReport(c *gin.Context) {
	db := database.Mysql()

	start, end := getDates(c)
	if start == nil || end == nil {
		return
	}

	var extras []struct {
		ID int `db:"id" json:"id"`
		models.ConfigBillItem
		models.OrderBillExtra
	}
	if _, err := db.Select(&extras, `select
			order__bill_extra.quantity, order__bill_extra.item_price, order__bill_extra.id,
			config__bill_item.name, config__bill_item.is_percent, config__bill_item.price
			from order__bill
			join order__group on order__group.id=order__bill.group_id
			join order__bill_extra on order__bill_extra.bill_id=order__bill.id
			join config__bill_item on config__bill_item.id=order__bill_extra.bill_item_id
			where order__group.cleared=1
			and (order__group.cleared_when between ? and ?)`,
		start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05")); err != nil {
		panic(err)
	}

	unclearedTables, err := db.SelectInt("select count(1) as ex from order__group where order__group.cleared=0")
	if err != nil {
		panic(err)
	}

	c.JSON(200, map[string]interface{}{
		"extras":          extras,
		"unclearedTables": unclearedTables,
	})
}
