
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type ConfigModifierGroup struct {
	Id int `db:"id"`

	Name string `db:"name"`
	NumberRequired int `db:"number_required"`
}

func init() {
	database.Mysql().AddTableWithName(ConfigModifierGroup{}, "config__modifier_group")
}
