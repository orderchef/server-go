
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type CategoryPrinter struct {
	PrinterId int `db:"printer_id" json:"printer_id"`
	CategoryId int `db:"category_id" json:"category_id"`

	Printer Printer `db:"-"`
	Category Category `db:"-"`
}

func init() {
	database.Mysql().AddTableWithName(CategoryPrinter{}, "category_printer")
}
