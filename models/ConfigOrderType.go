package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type ConfigOrderType struct {
	Id          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(ConfigOrderType{}, "config__order_type").SetKeys(true, "id")
}

func GetAllOrderTypes() ([]ConfigOrderType, error) {
	db := database.Mysql()

	var objs []ConfigOrderType
	if _, err := db.Select(&objs, "select * from config__order_type order by name"); err != nil {
		return nil, err
	}

	return objs, nil
}

func (t *ConfigOrderType) Get() error {
	db := database.Mysql()

	if err := db.SelectOne(&t, "select * from config__order_type where id=?", t.Id); err != nil {
		return err
	}

	return nil
}

func (t *ConfigOrderType) Save() error {
	db := database.Mysql()

	var err error
	if t.Id <= 0 {
		err = db.Insert(t)
	} else {
		_, err = db.Update(t)
	}

	if err != nil {
		return err
	}

	return nil
}

func (t *ConfigOrderType) Remove() error {
	db := database.Mysql()

	if _, err := db.Delete(t); err != nil {
		return err
	}

	return nil
}
