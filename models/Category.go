
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type Category struct {
	Id int `db:"id"`
	Name string `db:"name"`
	Description string `db:"description"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(Category{}, "category").SetKeys(true, "id")
}
