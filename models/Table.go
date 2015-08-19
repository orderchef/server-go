package models

import (
	"time"
	"lab.castawaylabs.com/orderchef/database"
)

type Table struct {
	Id int `db:"id" form:"id" json:"id"`

	TypeId int `db:"type_id" form:"type_id" json:"type_id" binding:"required"`

	Name        *string `db:"name" form:"name" json:"name" binding:"required"`
	TableNumber *string `db:"table_number" form:"table_number" json:"table_number"`
	Location    *string `db:"location" form:"location" json:"location"`
}

type TableTypeExport struct {
	Type_name string  `json:"type_name" db:"name"`
	Type_id   int     `json:"type_id" db:"id"`
	Tables    []Table `json:"tables" db:"tables"`
}

type OpenTable struct {
	Table

	Covers *int `json:"covers" db:"covers"`
	PrintedOrders int `json:"printedOrders" db:"printedOrders"`
	Orders int `json:"orders" db:"orders"`
	LastPrintedOrder *time.Time `json:"last_printed_order" db:"last_printed_order"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(Table{}, "table__items").SetKeys(true, "id")
}

func GetAllTables() ([]Table, error) {
	db := database.Mysql()

	var tables []Table
	if _, err := db.Select(&tables, "select * from table__items order by name"); err != nil {
		return nil, err
	}

	return tables, nil
}

func GetAllTablesSorted() ([]TableTypeExport, error) {
	db := database.Mysql()

	var types []TableTypeExport
	if _, err := db.Select(&types, "select * from config__table_type"); err != nil {
		return nil, err
	}

	for i, t := range types {
		if _, err := db.Select(&types[i].Tables, "select * from table__items where type_id=?", t.Type_id); err != nil {
			return nil, err
		}
	}

	return types, nil
}

func GetOpenTables() ([]OpenTable, error) {
	var tables []OpenTable
	db := database.Mysql()

	_, err := db.Select(&tables, "select ti.*, og.covers, count(ogm.printed_at) as printedOrders, count(ogm.id) as orders, max(ogm.printed_at) as last_printed_order from table__items as ti join order__group as og on og.table_id = ti.id left join order__group_member as ogm on ogm.group_id=og.id where og.cleared=0 group by ti.id")

	return tables, err
}

func (t *Table) Get() error {
	db := database.Mysql()
	err := db.SelectOne(&t, "select * from table__items where id=?", t.Id)

	return err
}

func (t *Table) Save() error {
	db := database.Mysql()

	var err error
	if t.Id <= 0 {
		err = db.Insert(t)
	} else {
		_, err = db.Update(t)
	}

	return err
}

func (t *Table) Remove() error {
	db := database.Mysql()

	_, err := db.Delete(t)

	return err
}
