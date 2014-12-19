
package models

var CustomerTable = "customer"

type Customer struct {
	Id int
	Name string
	Email string
	Telephone string
	Postcode string
}
