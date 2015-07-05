package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type Order struct {
	Id      int `db:"id" json:"id"`
	TypeId  int `db:"type_id" json:"type_id"`
	GroupId int `db:"group_id" json:"group_id"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(Order{}, "order__group_member").SetKeys(true, "id")
}

func (order *Order) Get() error {
	db := database.Mysql()

	if err := db.SelectOne(&order, "select * from order__group_member where id=?", order.Id); err != nil {
		return err
	}

	return nil
}

func (order *Order) GetItems() ([]OrderItem, error) {
	db := database.Mysql()

	var items []OrderItem
	if _, err := db.Select(&items, "select * from order__item where order_id=?", order.Id); err != nil {
		return nil, err
	}

	return items, nil
}

func (order *Order) Save() error {
	db := database.Mysql()

	var err error
	if order.Id <= 0 {
		err = db.Insert(order)
	} else {
		_, err = db.Update(order)
	}

	if err != nil {
		return err
	}

	return nil
}
