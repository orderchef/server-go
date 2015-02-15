
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type ConfigModifierGroup struct {
	Id int `db:"id"`

	Name string `db:"name"`
	Required bool `db:"choice_required"`
}

func init() {
	database.Mysql().AddTableWithName(ConfigModifierGroup{}, "config__modifier_group").SetKeys(true, "id")
}
