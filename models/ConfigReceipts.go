package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type ConfigReceipts struct {
	// Printer   Printer `db:"-"`
	PrinterId int     `db:"printer_id"`

	Receipt string `db:"receipt"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(ConfigReceipts{}, "config__receipt")
}
