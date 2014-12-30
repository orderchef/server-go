
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type ConfigTableType struct {
	Id int `db:"id" json:"id" form:"id"`
	Name string `db:"name" json:"name" form:"name" binding:"required"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(ConfigTableType{}, "config__table_type").SetKeys(true, "id")
}
