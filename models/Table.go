
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type Table struct {
	Id int `db:"id" form:"id" json:"id"`

	TypeId int `db:"type_id" form:"type_id" json:"type_id" binding:"required"`

	Name *string `db:"name" form:"name" json:"name" binding:"required"`
	TableNumber *string `db:"table_number" form:"table_number" json:"table_number"`
	Location *string `db:"location" form:"location" json:"location"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(Table{}, "table__items").SetKeys(true, "id")
}
