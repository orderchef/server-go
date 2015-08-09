package models

import (
	"time"
	"lab.castawaylabs.com/orderchef/database"
)

type OrderBill struct {
	ID int `db:"id"`
	Paid bool `db:"paid"`
	PaidAmount float32 `db:"paid_amount"`
	Total float32 `db:"total"`

	PrintedAt *time.Time `db:"printed_at"`

	Created time.Time `db:"created"`
}

type OrderBillItem struct {
	ID int `db:"id"`
	BillID int `db:"bill_id"`
	OrderItemID *int `db:"order_item_id"`
	ItemName string `db:"item_name"`
	ItemPrice float32 `db:"item_price"`
	Discount float32 `db:"discount"`
}

func init() {
	db := database.Mysql()

	db.AddTableWithName(OrderBill{}, "order__bill").SetKeys(true, "id")
	db.AddTableWithName(OrderBillItem{}, "order__bill_item").SetKeys(true, "id")
}
