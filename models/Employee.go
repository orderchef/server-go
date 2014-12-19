
package models

var EmployeeTable = "employee"

type Employee struct {
	Id  int
	Name string
	Manager bool
	Passkey string
	LastLogin int
}
