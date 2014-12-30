
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type ConfigOrderType struct {
	Id int `db:"id"`
	Name string `db:"name"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(ConfigOrderType{}, "config__order_type").SetKeys(true, "id")
}
