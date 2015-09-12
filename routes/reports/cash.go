package reports

import (
	"time"
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/database"
	"lab.castawaylabs.com/orderchef/models"
)

func getCashReport(c *gin.Context) {
	db := database.Mysql()

	start, end := getDates(c)
	if start == nil || end == nil {
		return
	}

	var reports []struct{
		Category string `db:"category"`
		Amount float32 `db:"amount"`
	}
	if _, err := db.Select(&reports, "select SUM(amount) as amount, category from report__cash where (date between ? and ?) group by category", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05")); err != nil {
		panic(err)
	}

	report := map[string]float32{}

	for _, r := range reports {
		report[r.Category] += r.Amount
	}

	c.JSON(200, report)
}

func getCashReportCategories(c *gin.Context) {
	c.JSON(200, models.GetCashReportCategories())
}

func createCashReport(c *gin.Context) {
	db := database.Mysql()

	report := models.CashReport{}
	if err := c.Bind(&report); err != nil {
		c.JSON(400, err)
		return
	}

	report.Date = time.Now()

	if err := db.Insert(&report); err != nil {
		panic(err)
	}

	c.AbortWithStatus(204)
}