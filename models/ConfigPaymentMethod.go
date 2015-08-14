package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type ConfigPaymentMethod struct {
	ID int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(ConfigPaymentMethod{}, "config__payment_method").SetKeys(true, "id")
}
