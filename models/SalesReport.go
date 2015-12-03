package models

import (
	"lab.castawaylabs.com/orderchef/database"
	"time"
)

type SalesReport struct {
	ID int `db:"id" json:"id"`

	Category string    `db:"category" json:"category" binding:"required"`
	Amount   float32   `db:"amount" json:"amount"`
	Date     time.Time `db:"date" json:"date" binding:"required"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(SalesReport{}, "report__sales").SetKeys(true, "id")
}

func GetSalesReportCategories() []string {
	db := database.Mysql()

	var cats []string
	if _, err := db.Select(&cats, "select distinct(category) from report__sales order by category asc"); err != nil {
		panic(err)
	}

	return cats
}
