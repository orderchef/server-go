
package models

import (
	"log"
	"lab.castawaylabs.com/orderchef/database"
)

var ConfigTableTypeTable = "config__table_type"

type ConfigTableType struct {
	Id int `json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required"`
}

func GetAllTableTypes() ([]ConfigTableType, error) {
	db := database.Mysql()
	objs := []ConfigTableType{}

	rows, err := db.Query("select id, name from " + ConfigTableTypeTable)
	if err != nil {
		return objs, err
	}

	defer rows.Close()

	for rows.Next() {
		obj := ConfigTableType{}

		if err := rows.Scan(&obj.Id, &obj.Name); err != nil {
			return objs, err
		}

		objs = append(objs, obj)
	}

	log.Println("Table Types:", objs)

	return objs, rows.Err()
}

func (tableType *ConfigTableType) Get() error {
	db := database.Mysql()

	if err := db.QueryRow("select id, name from " + ConfigTableTypeTable + " where id = ?", tableType.Id).Scan(&tableType.Id, &tableType.Name); err != nil {
		return err
	}

	return nil
}

func (tableType *ConfigTableType) Save() error {
	db := database.Mysql()

	query := "update " + ConfigTableTypeTable + " set name = ?"
	if tableType.Id == 0 {
		query = "insert into " + ConfigTableTypeTable + " (name) values (?)"
	}

	result, err := db.Exec(query, tableType.Name)
	if err != nil {
		return err
	}

	if tableType.Id == 0 {
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		tableType.Id = int(id)
	}

	return nil
}
