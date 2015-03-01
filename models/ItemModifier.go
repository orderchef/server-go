
package models

import (
	"database/sql"
	"lab.castawaylabs.com/orderchef/database"
)

type ItemModifier struct {
	Item int `db:"item_id" json:"item_id"`
	ModifierGroup int `db:"modifier_group_id" json:"modifier_group_id" binding:"required"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(ItemModifier{}, "item__modifier")
}

func (modifier *ItemModifier) Save() error {
	db := database.Mysql()

	count, err := db.SelectInt("select count(*) from item__modifier where item_id=? and modifier_group_id=?", modifier.Item, modifier.ModifierGroup)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if count > 0 {
		return nil
	}

	err = db.Insert(modifier)

	return err
}

func (modifier *ItemModifier) Remove() error {
	db := database.Mysql()

	if _, err := db.Delete(modifier); err != nil {
		return err
	}

	return nil
}
