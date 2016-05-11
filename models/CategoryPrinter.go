package models

type CategoryPrinter struct {
	PrinterID  string `db:"printer_id" json:"printer_id"`
	CategoryID *int   `db:"category_id" json:"category_id"`
	ItemID     *int   `db:"item_id" json:"item_id"`
	Order      int    `db:"order" json:"order"`
}
