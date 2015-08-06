package models

import (
	"database/sql"
	"lab.castawaylabs.com/orderchef/database"
)

type Item struct {
	Id int `db:"id" json:"id"`

	Name        string  `db:"name" json:"name"`
	Description string  `db:"description" json:"description"`
	Price       float32 `db:"price" json:"price"`
	CategoryId  int     `db:"category_id" json:"category_id"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(Item{}, "item").SetKeys(true, "id")
}

func GetAllItems() ([]Item, error) {
	db := database.Mysql()

	var types []Item
	if _, err := db.Select(&types, "select * from item"); err != nil {
		return nil, err
	}

	return types, nil
}

func (item *Item) Get() error {
	db := database.Mysql()

	if err := db.SelectOne(&item, "select * from item where id=?", item.Id); err != nil {
		return err
	}

	return nil
}

func (item *Item) Save() error {
	db := database.Mysql()

	var err error
	if item.Id <= 0 {
		err = db.Insert(item)
	} else {
		_, err = db.Update(item)
	}

	return err
}

func (item *Item) Remove() error {
	db := database.Mysql()

	if _, err := db.Delete(item); err != nil {
		return err
	}

	return nil
}

func (item *Item) GetModifiers() ([]int, error) {
	db := database.Mysql()

	var modifiers_items []ItemModifier
	_, err := db.Select(&modifiers_items, "select * from item__modifier where item_id=?", item.Id)

	modifiers := make([]int, len(modifiers_items))
	if err != nil && err != sql.ErrNoRows {
		return modifiers, err
	}

	for i, m := range modifiers_items {
		modifiers[i] = m.ModifierGroup
	}

	return modifiers, nil
}
