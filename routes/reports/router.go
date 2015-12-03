package reports

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Router(r *gin.RouterGroup) {
	// bills report
	r.GET("/bills", getBillsReport)

	// cash report
	r.GET("/cash", getCashReport)
	r.POST("/cash", createCashReport)
	r.GET("/cash/categories", getCashReportCategories)

	// Popular items report
	r.GET("/popularItems", getPopularItems)

	r.GET("/sales", getSalesReport)
	r.POST("/sales", createSalesReport)
}

func getDate(dateString string) (*time.Time, error) {
	date, err := strconv.ParseInt(dateString, 10, 64)
	if err != nil {
		return nil, err
	}

	unixTime := time.Unix(date, 0)
	return &unixTime, nil
}

func getDates(c *gin.Context) (*time.Time, *time.Time) {
	start, errStart := getDate(c.Query("start"))
	end, errEnd := getDate(c.Query("end"))

	if errStart != nil || errEnd != nil {
		fmt.Println(errStart, errEnd)
		c.AbortWithStatus(400)

		return nil, nil
	}

	return start, end
}
