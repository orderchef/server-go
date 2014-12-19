
package models

var OrderItemModifierTable = "order__item_modifier"

type OrderItemModifier struct {
	Id uint

	OrderItem OrderItem
	OrderItemId uint

	ModifierGroup ConfigModifierGroup
	ModifierGroupId uint
}
