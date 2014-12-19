
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

var PrinterTable = "printer"

type Printer struct {
	Id string
	Name string
	Location string
}

func (printer *Printer) Get() error {
	db := database.Mysql()

	if err := db.QueryRow("select id, name, location from " + PrinterTable + " where id = ?, name = ?, location = ?", printer.Id).Scan(&printer.Id, &printer.Name, &printer.Location); err != nil {
		return err
	}

	return nil
}
