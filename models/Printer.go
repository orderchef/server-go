package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type Printer struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Location string `db:"location"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(Printer{}, "printer").SetKeys(true, "id")
}
