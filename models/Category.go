package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type Category struct {
	Id          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	PrintOrder  int    `db:"print_order" json:"print_order"`
}

func init() {
	db := database.Mysql()
	db.AddTableWithName(Category{}, "category").SetKeys(true, "id")
}

// Fetch all categories from the database
func GetAllCategories() ([]Category, error) {
	db := database.Mysql()

	var types []Category
	if _, err := db.Select(&types, "select * from category"); err != nil {
		return nil, err
	}

	return types, nil
}

// Get single category
func (category *Category) Get() error {
	db := database.Mysql()

	if err := db.SelectOne(&category, "select * from category where id=?", category.Id); err != nil {
		return err
	}

	return nil
}

// Save / update category
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

// Remove category
func (category *Category) Remove() error {
	db := database.Mysql()

	if _, err := db.Delete(category); err != nil {
		return err
	}

	return nil
}
