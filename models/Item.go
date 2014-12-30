
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type Item struct {
	Id int `db:"id"`
	Name string `db:"name"`
	Description string `db:"description"`
	Price float32 `db:"price"`

	Category Category `db:"-"`
	CategoryId int `db:"category_id"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(Item{}, "item").SetKeys(true, "id")
}
