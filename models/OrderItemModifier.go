
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type OrderItemModifier struct {
	Id int `db:"id" json:"id"`

	OrderItem OrderItem `db:"-"`
	OrderItemId int `db:"order_item_id" json:"order_item_id"`

	ModifierGroup ConfigModifierGroup `db:"-" json:"-"`
	ModifierGroupId int `db:"modifier_group_id" json:"modifier_group_id"`
	ModifierId int `db:"modifier_id" json:"modifier_id"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(OrderItemModifier{}, "order__item_modifier").SetKeys(true, "id")
}

func (modifier *OrderItemModifier) Save() error {
	db := database.Mysql()

	var err error
	if modifier.Id <= 0 {
		err = db.Insert(modifier)
	} else {
		_, err = db.Update(modifier)
	}

	return err
}

func (modifier *OrderItemModifier) Remove() error {
	db := database.Mysql()

	if _, err := db.Delete(modifier); err != nil {
		return err
	}

	return nil
}
