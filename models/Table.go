
package models

import (
	"lab.castawaylabs.com/orderchef/database"
	"database/sql"
	"log"
)

var (
	TableTable = "table__items"
	TableSchema = `
		create table table__items (
			id int auto_increment not null,
			name varchar(255) not null,
			type_id int not null,
			table_number varchar(255) null,
			location varchar(255) null,

			primary key (id),
			foreign key (type_id) references config__table_type (id)
		);
		create index table_type_idx on table__items (type_id) using hash;
	`
)

type Table struct {
	Id int `form:"id" json:"id"`

	Type ConfigTableType
	TypeId int `form:"type_id" json:"type_id" binding:"required"`

	Name string `form:"name" json:"name" binding:"required"`
	TableNumber sql.NullString `form:"table_number" json:"table_number"`
	Location sql.NullString `form:"location" json:"location"`
}

func GetAllTables() ([]Table, error) {
	db := database.Mysql()

	rows, err := db.Query("select id, name, type_id, table_number, location from " + TableTable)
	if err != nil {
		return []Table{}, err
	}

	defer rows.Close()

	tables := []Table{}
	for rows.Next() {
		table := Table{}

		err := rows.Scan(&table.Id, &table.Name, &table.TypeId, &table.TableNumber, &table.Location)
		if err != nil {
			return []Table{}, err
		}

		tables = append(tables, table)
		log.Println(table)
	}

	log.Println("Tables:", tables)

	return tables, rows.Err()
}

func (t *Table) Get() error {
	db := database.Mysql()

	if err := db.QueryRow("select id, name, type_id, table_number, location from " + TableTable + " where id = ?", t.Id).Scan(&t.Id, &t.Name, &t.TypeId, &t.TableNumber, &t.Location); err != nil {
		return err
	}

	return nil
}

func (t *Table) Save() error {
	db := database.Mysql()

	query := "update " + TableTable + " set name = ?, type_id = ?, table_number = ?, location = ?"
	if t.Id == 0 {
		query = "insert into " + TableTable + " (name, type_id, table_number, location) values (?, ?, ?, ?)"
	}

	result, err := db.Exec(query, t.Name, t.TypeId, t.TableNumber, t.Location)
	if err != nil {
		return err
	}

	if t.Id == 0 {
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		t.Id = int(id)
	}

	return nil
}
