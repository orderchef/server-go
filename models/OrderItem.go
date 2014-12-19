
package models

var OrderItemTable = "order__item"

type OrderItem struct {
	Id uint

	Item Item
	ItemId uint

	Quantity uint
	Notes string
}