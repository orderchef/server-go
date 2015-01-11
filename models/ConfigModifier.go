
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type ConfigModifier struct {
	Id int `db:"id"`

	Group ConfigModifierGroup `db:"-"`
	GroupId int `db:"group_id"`

	Name string `db:"name"`
	Price float32 `db:"price"`
}

func init() {
	database.Mysql().AddTableWithName(ConfigModifier{}, "config__modifier").SetKeys(true, "id")
}
