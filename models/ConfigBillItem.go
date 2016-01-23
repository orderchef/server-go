package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type ConfigBillItem struct {
	ID        int     `db:"id" json:"id" form:"-"`
	Name      string  `db:"name" json:"name"`
	Price     float64 `db:"price" json:"price"`
	IsPercent bool    `db:"is_percent" json:"is_percent"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(ConfigBillItem{}, "config__bill_item").SetKeys(true, "id")
}
