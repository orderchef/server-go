package reports

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/database"
	"lab.castawaylabs.com/orderchef/models"
)

func getBillsReport(c *gin.Context) {
	db := database.Mysql()

	start, end := getDates(c)
	if start == nil || end == nil {
		return
	}

	var bills []struct {
		models.OrderBill
		Covers int `db:"covers" json:"covers"`
	}
	if _, err := db.Select(&bills, "select order__bill.*, order__group.covers from order__bill join order__group on order__group.id=order__bill.group_id where order__group.cleared=1 and (order__group.cleared_when between ? and ?)", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05")); err != nil {
		panic(err)
	}

	var total float32
	for _, bill := range bills {
		total += bill.Total
	}

	unclearedTables, err := db.SelectInt("select count(1) as ex from order__group where order__group.cleared=0")
	if err != nil {
		panic(err)
	}

	c.JSON(200, map[string]interface{}{
		"bills":           bills,
		"total":           total,
		"unclearedTables": unclearedTables,
	})
}

func getPopularItems(c *gin.Context) {
	start, end := getDates(c)
	if start == nil || end == nil {
		return
	}

	var popularItems []struct {
		Quantity int    `db:"ex"`
		Name     string `db:"name"`
		Category string `db:"category"`
	}
	if _, err := database.Mysql().Select(&popularItems, "select count(1) as ex, item.name, category.name as category from order__group join order__group_member on order__group_member.group_id=order__group.id join order__item on order__item.order_id=order__group_member.id join item on item.id=order__item.item_id join category on category.id=item.category_id where order__group.cleared=1 and (order__group.cleared_when between ? and ?) group by item.id order by ex DESC", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05")); err != nil {
		panic(err)
	}

	c.JSON(200, popularItems)
}
