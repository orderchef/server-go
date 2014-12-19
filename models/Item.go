
package models

var ItemTable = "item"

type Item struct {
	Id int
	Name string
	Description string
	Price float32

	Category Category
	CategoryId int
}
