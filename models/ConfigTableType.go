package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type ConfigTableType struct {
	Id   int    `db:"id" json:"id" form:"id"`
	Name string `db:"name" json:"name" form:"name" binding:"required"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(ConfigTableType{}, "config__table_type").SetKeys(true, "id")
}

func GetAllTableTypes() ([]ConfigTableType, error) {
	db := database.Mysql()

	var objs []ConfigTableType
	if _, err := db.Select(&objs, "select * from config__table_type order by name"); err != nil {
		return nil, err
	}

	return objs, nil
}

func (t *ConfigTableType) Get() error {
	db := database.Mysql()

	if err := db.SelectOne(&t, "select * from config__table_type where id=?", t.Id); err != nil {
		return err
	}

	return nil
}

func (t *ConfigTableType) Save() error {
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

func (t *ConfigTableType) Remove() error {
	db := database.Mysql()

	if _, err := db.Delete(t); err != nil {
		return err
	}

	return nil
}
