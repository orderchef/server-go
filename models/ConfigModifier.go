package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type ConfigModifier struct {
	Id int `db:"id" json:"id"`

	GroupId int `db:"group_id" json:"group_id"`

	Name  string  `db:"name" json:"name"`
	Price float32 `db:"price" json:"price"`

	Deleted bool `db:"deleted" json:"-"`
}

func init() {
	database.Mysql().AddTableWithName(ConfigModifier{}, "config__modifier").SetKeys(true, "id")
}

func (modifier *ConfigModifier) Get() error {
	db := database.Mysql()

	if err := db.SelectOne(&modifier, "select * from config__modifier where id=?", modifier.Id); err != nil {
		return err
	}

	return nil
}

func (modifier *ConfigModifier) Save() error {
	db := database.Mysql()

	var err error
	if modifier.Id <= 0 {
		err = db.Insert(modifier)
	} else {
		_, err = db.Update(modifier)
	}

	if err != nil {
		return err
	}

	return nil
}

func (modifier *ConfigModifier) Remove() error {
	db := database.Mysql()

	modifier.Deleted = true
	if _, err := db.Update(modifier); err != nil {
		return err
	}

	return nil
}
