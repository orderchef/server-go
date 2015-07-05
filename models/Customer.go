package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type Customer struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Telephone string `db:"telephone"`
	Postcode  string `db:"postcode"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(Customer{}, "customer").SetKeys(true, "id")
}
