
package models

import (
	_ "lab.castawaylabs.com/orderchef/database"
)

var CategoryPrinterTable = "category_printer"

type CategoryPrinter struct {
	PrinterId int
	CategoryId int

	Printer Printer
	Category Category
}
