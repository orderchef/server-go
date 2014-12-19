
package models

var OrderTable = "order__group_member"

type Order struct {
	Id uint
	Type ConfigOrderType
	TypeId uint
	Group OrderGroup
	GroupId uint
}
