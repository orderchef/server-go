
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type Category struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(Category{}, "category").SetKeys(true, "id")
}

func GetAllCategories() ([]Category, error) {
	db := database.Mysql()

	var types []Category
	if _, err := db.Select(&types, "select * from category"); err != nil {
		return nil, err
	}

	return types, nil
}

func (category *Category) Get() error {
	db := database.Mysql()

	if err := db.SelectOne(&category, "select * from category where id=?", category.Id); err != nil {
		return err
	}

	return nil
}

func (category *Category) Save() error {
	db := database.Mysql()

	var err error
	if category.Id <= 0 {
		err = db.Insert(category)
	} else {
		_, err = db.Update(category)
	}

	if err != nil {
		return err
	}

	return nil
}

func (category *Category) Remove() error {
	db := database.Mysql()

	if _, err := db.Delete(category); err != nil {
		return err
	}

	return nil
}
