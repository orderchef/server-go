
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type OrderItem struct {
	Id int `db:"id" json:"id"`

	ItemId int `db:"item_id" json:"item_id" binding:"required"`
	OrderId int `db:"order_id" json:"order_id" binding:"required"`

	Quantity int `db:"quantity" json:"quantity"`
	Notes string `db:"notes" json:"notes"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(OrderItem{}, "order__item").SetKeys(true, "id")
}

func (orderItem *OrderItem) Get() error {
	db := database.Mysql()

	if err := db.SelectOne(&orderItem, "select * from order__item where id=?", orderItem.Id); err != nil {
		return err
	}

	return nil
}
