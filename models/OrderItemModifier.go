
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type OrderItemModifier struct {
	Id int `db:"id"`

	OrderItem OrderItem `db:"-"`
	OrderItemId int `db:"order_item_id"`

	ModifierGroup ConfigModifierGroup `db:"-"`
	ModifierGroupId int `db:"modifier_group_id"`
	ModifierId int `db:"modifier_id"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(OrderItemModifier{}, "order__item_modifier").SetKeys(true, "id")
}
