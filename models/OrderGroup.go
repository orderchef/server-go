
package models

var OrderGroupTable = "order__group"

type OrderGroup struct {
	Id int

	Table Table
	TableId int

	Cleared bool
	ClearedWhen int
}
