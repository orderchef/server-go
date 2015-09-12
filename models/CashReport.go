package models

import (
	"time"
	"lab.castawaylabs.com/orderchef/database"
)

type CashReport struct {
	ID int `db:"id" json:"id"`

	Category string `db:"category" json:"category"`
	Amount float32 `db:"amount" json:"amount"`
	Date time.Time `db:"date" json:"date"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(CashReport{}, "report__cash").SetKeys(true, "id")
}

func GetCashReportCategories() []string {
	db := database.Mysql()

	var cats []string
	if _, err := db.Select(&cats, "select distinct(category) from report__cash order by category asc"); err != nil {
		panic(err)
	}

	return cats
}