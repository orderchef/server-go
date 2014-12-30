
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type OrderGroup struct {
	Id int `db:"id"`

	Table Table `db:"-"`
	TableId int `db:"table_id"`

	Cleared bool `db:"cleared"`
	ClearedWhen int `db:"cleared_when"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(OrderGroup{}, "order__group").SetKeys(true, "id")
}
