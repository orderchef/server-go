
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type ItemModifier struct {
	Item int `db:"item_id" json:"item_id"`
	ModifierGroup int `db:"modifier_group_id" json:"modifier_group_id"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(Item{}, "item__modifier")
}
