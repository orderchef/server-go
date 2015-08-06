package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type Employee struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	Manager   bool   `db:"manager"`
	Passkey   string `db:"passkey"`
	LastLogin int    `db:"last_login"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(Employee{}, "employee").SetKeys(true, "id")
}
