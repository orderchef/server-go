package models

import (
	"time"
	"lab.castawaylabs.com/orderchef/database"
)

type OrderGroup struct {
	Id int `db:"id" json:"id"`

	TableId int `db:"table_id" json:"table_id"`

	Cleared     bool `db:"cleared" json:"cleared"`
	ClearedWhen *time.Time  `db:"cleared_when" json:"cleared_when"`

	Covers int `db:"covers" json:"covers"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(OrderGroup{}, "order__group").SetKeys(true, "id")
}

func (group *OrderGroup) Get() error {
	db := database.Mysql()

	if err := db.SelectOne(&group, "select * from order__group where id=?", group.Id); err != nil {
		return err
	}

	return nil
}

func (group *OrderGroup) GetOrders() ([]Order, error) {
	db := database.Mysql()

	var orders []Order
	if _, err := db.Select(&orders, "select * from order__group_member where group_id=?", group.Id); err != nil {
		return nil, err
	}

	return orders, nil
}

func (group *OrderGroup) GetByTableId() error {
	db := database.Mysql()

	if err := db.SelectOne(&group, "select * from order__group where cleared=? and table_id=? limit 1", group.Cleared, group.TableId); err != nil {
		return err
	}

	return nil
}

func (group *OrderGroup) Save() error {
	db := database.Mysql()

	var err error
	if group.Id <= 0 {
		err = db.Insert(group)
	} else {
		_, err = db.Update(group)
	}

	if err != nil {
		return err
	}

	return nil
}

func (group *OrderGroup) Remove() error {
	db := database.Mysql()

	if _, err := db.Delete(group); err != nil {
		return err
	}

	return nil
}
