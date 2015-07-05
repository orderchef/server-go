package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type OrderReceipt struct {
	Id int `db:"id"`

	OrderGroup   OrderGroup `db:"-"`
	OrderGroupId int        `db:"order_group_id"`

	Printer   Printer `db:"-"`
	PrinterId int     `db:"printer_id"`

	Employee   Employee `db:"-"`
	EmployeeId int      `db:"employee_id"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(OrderReceipt{}, "order__receipt").SetKeys(true, "id")
}
