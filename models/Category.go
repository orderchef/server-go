
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

var CategoryTable = "category"

type Category struct {
	Id int
	Name string
	Description string
}

func GetAllCategories() ([]Category, error) {
	db := database.Mysql()
	objs := []Category{}

	rows, err := db.Query("select id, name, description from " + CategoryTable)
	if err != nil {
		return objs, err
	}

	defer rows.Close()

	for rows.Next() {
		obj := Category{}

		if err := rows.Scan(&obj.Id, &obj.Name, &obj.Description); err != nil {
			return objs, err
		}

		objs = append(objs, obj)
	}

	return objs, rows.Err()
}

func (category *Category) Get() error {
	db := database.Mysql()

	if err := db.QueryRow("select id, name, description from " + CategoryTable + " where id = ?", category.Id).Scan(&category.Id, &category.Name, &category.Description); err != nil {
		return err
	}

	return nil
}

func (category *Category) GetPrinters() ([]Printer, error) {
	db := database.Mysql()
	objs := []Printer{}

	rows, err := db.Query("select printer_id from " + CategoryPrinterTable + " where category_id=?", category.Id)
	if err != nil {
		return objs, err
	}

	defer rows.Close()

	for rows.Next() {
		obj := Printer{}

		if err := rows.Scan(&obj.Id); err != nil {
			return objs, err
		}

		objs = append(objs, obj)
	}

	for i := 0; i < len(objs); i++ {
		if err := objs[i].Get(); err != nil {
			return []Printer{}, err
		}
	}

	return objs, rows.Err()
}

func (category *Category) Save() error {
	db := database.Mysql()

	query := "update " + CategoryTable + " set name = ?, description = ?"
	if category.Id == 0 {
		query = "insert into " + CategoryTable + " (name, description) values (?, ?)"
	}

	result, err := db.Exec(query, category.Name, category.Description)
	if err != nil {
		return err
	}

	if category.Id == 0 {
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		category.Id = int(id)
	}

	return nil
}
