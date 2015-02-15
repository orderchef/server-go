
package models

import (
	"database/sql"
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

func (orderItem *OrderItem) Save() error {
	db := database.Mysql()

	var err error
	if orderItem.Id <= 0 {
		err = db.Insert(orderItem)
	} else {
		_, err = db.Update(orderItem)
	}

	return err
}

func (orderItem *OrderItem) Remove() error {
	db := database.Mysql()

	if _, err := db.Delete(orderItem); err != nil {
		return err
	}

	return nil
}

func FindExistingOrder(existing OrderItem) (OrderItem, error) {
	db := database.Mysql()

	var found OrderItem
	err := db.SelectOne(&found, "select * from order__item where order_id=? and item_id=?", existing.OrderId, existing.ItemId)
	if err == sql.ErrNoRows {
		return existing, nil
	} else if err != nil {
		return OrderItem{}, err
	}

	return found, nil
}
