package models

import (
	"fmt"
	"lab.castawaylabs.com/orderchef/database"
	"time"
)

type OrderBill struct {
	ID int `db:"id" json:"id"`

	GroupID int     `db:"group_id" json:"group_id"`
	Paid    bool    `db:"paid" json:"paid"`
	Total   float32 `db:"total" json:"total"`

	PrintedAt *time.Time `db:"printed_at" json:"printed_at"`
	CreatedAt time.Time  `db:"created" json:"created"`

	Items      []OrderBillItem  `db:"-" json:"bill_items" form:"items"`
	BillExtras []OrderBillExtra `db:"-" json:"bill_extras"`
}

type OrderBillItem struct {
	ID                 int     `db:"id" json:"-" form:"-"`
	BillID             int     `db:"bill_id" json:"bill_id"`
	OrderItemID        *int    `db:"order_item_id" json:"order_item_id"`
	ItemName           string  `db:"item_name" json:"item_name"`
	ItemPrice          float32 `db:"item_price" json:"item_price"`
	ItemPriceFormatted string  `db:"-" json:"-" form:"-"`
	Deleted            bool    `db:"deleted" json:"deleted"`
	Discount           float32 `db:"discount" json:"discount"`
}

type OrderBillExtra struct {
	ID         int `db:"id" json:"-"`
	BillID     int `db:"bill_id" json:"bill_id"`
	BillItemID int `db:"bill_item_id" json:"bill_item_id" binding:"required"`
	Quantity   int `db:"quantity" json:"quantity"`

	ItemName           *string  `db:"item_name" json:"item_name"`
	ItemPrice          *float32 `db:"item_price" json:"item_price"`
	ItemPriceFormatted string   `db:"-" json:"-" form:"-"`
}

type OrderBillPayment struct {
	BillID          int     `db:"bill_id" json:"bill_id" binding:"required"`
	PaymentMethodID int     `db:"payment_method_id" json:"payment_method_id" binding:"required"`
	Amount          float32 `db:"amount" json:"amount" binding:"required"`
}

func (bill *OrderBill) GetItems() error {
	db := database.Mysql()

	bill.Items = []OrderBillItem{}
	_, err := db.Select(&bill.Items, "select * from order__bill_item where bill_id=?", bill.ID)

	for i, item := range bill.Items {
		bill.Items[i].ItemPriceFormatted = fmt.Sprintf("%.2f", item.ItemPrice)
	}

	bill.BillExtras = []OrderBillExtra{}
	_, err = db.Select(&bill.BillExtras, "select * from order__bill_extra where bill_id=?", bill.ID)
	for i, extra := range bill.BillExtras {
		bill.BillExtras[i].ItemPriceFormatted = fmt.Sprintf("%.2f", extra.ItemPrice)
	}

	return err
}

func init() {
	db := database.Mysql()

	db.AddTableWithName(OrderBill{}, "order__bill").SetKeys(true, "id")
	db.AddTableWithName(OrderBillItem{}, "order__bill_item").SetKeys(true, "id")
	db.AddTableWithName(OrderBillExtra{}, "order__bill_extra").SetKeys(true, "id")
}
