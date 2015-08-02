package models

import (
	"database/sql"
	"lab.castawaylabs.com/orderchef/database"
)

type OrderItem struct {
	Id int `db:"id" json:"id"`

	ItemId  int `db:"item_id" json:"item_id" binding:"required"`
	OrderId int `db:"order_id" json:"order_id" binding:"required"`

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

func (orderItem *OrderItem) GetModifiers() ([]OrderItemModifier, error) {
	db := database.Mysql()

	var objs []OrderItemModifier
	_, err := db.Select(&objs, "select * from order__item_modifier where order_item_id=?", orderItem.Id)
	if err == sql.ErrNoRows {
		return objs, nil
	} else if err != nil {
		return nil, err
	}

	return objs, nil
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

	_, err := db.Exec("delete from order__item_modifier where order_item_id=?", orderItem.Id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if _, err := db.Delete(orderItem); err != nil {
		return err
	}

	return nil
}
