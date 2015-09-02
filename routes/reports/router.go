package reports

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"strconv"
)

func Router(r *gin.RouterGroup) {
	r.GET("/bills", getBillsReport)
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