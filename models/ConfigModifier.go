
package models

var ConfigModifierTable = "config__modifier"

type ConfigModifier struct {
	Id int

	Group ConfigModifierGroup
	GroupId int

	Name string
	Price float32
}
