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

	var bills []models.OrderBill
	if _, err := db.Select(&bills, "select order__bill.* from order__bill join order__group on order__group.id=order__bill.group_id where order__group.cleared=1 and (order__group.cleared_when between ? and ?)", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05")); err != nil {
		panic(err)
	}

	var total float32
	for _, bill := range bills {
		total += bill.Total
	}

	c.JSON(200, map[string]interface{}{
		"bills": bills,
		"total": total,
	})
}