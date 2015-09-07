package models

import (
	"time"
	"fmt"
	"lab.castawaylabs.com/orderchef/database"
)

type OrderBill struct {
	ID int `db:"id" json:"id"`

	GroupID int `db:"group_id" json:"group_id"`
	Paid bool `db:"paid" json:"paid"`
	PaidAmount float32 `db:"paid_amount" json:"paid_amount"`
	Total float32 `db:"total" json:"total"`
	PaymentMethodID int `db:"payment_method_id" json:"payment_method_id"`
	BillType string `db:"bill_type" json:"bill_type"`

	PrintedAt *time.Time `db:"printed_at" json:"printed_at"`
	CreatedAt time.Time `db:"created" json:"created"`

	Items []OrderBillItem `db:"-" json:"bill_items" form:"items"`
}

type OrderBillItem struct {
	ID int `db:"id" json:"-" form:"-"`
	BillID int `db:"bill_id" json:"bill_id" form:"bill_id"`
	OrderItemID *int `db:"order_item_id" json:"order_item_id" form:"order_item_id"`
	ItemName string `db:"item_name" json:"item_name" form:"item_name"`
	ItemPrice float32 `db:"item_price" json:"item_price" form:"item_price"`
	ItemPriceFormatted string `db:"-" json:"-" form:"-"`
	Discount float32 `db:"discount" json:"discount" form:"discount"`
}

func (bill *OrderBill) GetItems() error {
	db := database.Mysql()

	bill.Items = []OrderBillItem{}
	_, err := db.Select(&bill.Items, "select * from order__bill_item where bill_id=?", bill.ID)

	for i, item := range bill.Items {
		bill.Items[i].ItemPriceFormatted = fmt.Sprintf("%.2f", item.ItemPrice)
	}

	return err
}

func init() {
	db := database.Mysql()

	db.AddTableWithName(OrderBill{}, "order__bill").SetKeys(true, "id")
	db.AddTableWithName(OrderBillItem{}, "order__bill_item").SetKeys(true, "id")
}
