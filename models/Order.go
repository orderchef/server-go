
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type Order struct {
	Id int `db:"id"`

	Type ConfigOrderType `db:"-"`
	TypeId int `db:"type_id"`

	Group OrderGroup `db:"-"`
	GroupId int `db:"group_id"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(Order{}, "order__group_member").SetKeys(true, "id")
}
