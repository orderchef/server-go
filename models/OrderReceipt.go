
package models

var OrderReceiptTable = "order__receipt"

type OrderReceipt struct {
	Id int

	OrderGroup OrderGroup
	OrderGroupId int

	Printer Printer
	PrinterId int

	Employee Employee
	EmployeeId int
}
