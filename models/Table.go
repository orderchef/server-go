
package models

import (
	"lab.castawaylabs.com/orderchef/database"
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

	TypeId int `form:"type_id" json:"type_id" binding:"required"`

	Name *string `form:"name" json:"name" binding:"required"`
	TableNumber *string `form:"table_number" json:"table_number"`
	Location *string `form:"location" json:"location"`
}

type TableExport struct {
	Table
	Type ConfigTableType `json:"type"`
}

func GetAllTables() ([]TableExport, error) {
	db := database.Mysql()

	rows, err := db.Query(`
		 select t.id, t.name, t.type_id, t.table_number, t.location, t_type.id, t_type.name
		 from ` + TableTable + ` as t
		 join ` + ConfigTableTypeTable + ` as t_type on t_type.id=t.type_id
	`)

	if err != nil {
		return []TableExport{}, err
	}

	defer rows.Close()

	tables := []TableExport{}
	for rows.Next() {
		table := TableExport{}

		err := rows.Scan(&table.Id,
			&table.Name,
			&table.TypeId,
			&table.TableNumber,
			&table.Location,
			&table.Type.Id,
			&table.Type.Name,
		)

		if err != nil {
			return []TableExport{}, err
		}

		tables = append(tables, table)
	}

	return tables, rows.Err()
}

func (t *TableExport) Get() error {
	db := database.Mysql()

	query := `
	 select t.id, t.name, t.type_id, t.table_number, t.location, t_type.id, t_type.name
	 from ` + TableTable + ` as t
	 join ` + ConfigTableTypeTable + ` as t_type on t_type.id=t.type_id
	 where t.id = ?`

	err := db.QueryRow(query, t.Id).Scan(
		&t.Id,
		&t.Name,
		&t.TypeId,
		&t.TableNumber,
		&t.Location,
		&t.Type.Id,
		&t.Type.Name,
	)

	if err != nil {
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

func (t *Table) Remove() error {
	db := database.Mysql()

	if _, err := db.Exec("delete from " + TableTable + " where id = ?", t.Id); err != nil {
		return err
	}

	return nil
}

func (t *TableExport) GetTableType() error {
	t.Type.Id = t.TypeId
	if err := t.Type.Get(); err != nil {
		return err
	}

	return nil
}
