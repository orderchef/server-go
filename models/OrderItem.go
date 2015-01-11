
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type OrderItem struct {
	Id int `db:"id"`

	ItemId int `db:"item_id"`
	OrderId int `db:"order_id"`

	Quantity int `db:"quantity"`
	Notes string `db:"notes"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(OrderItem{}, "order__item").SetKeys(true, "id")
}
