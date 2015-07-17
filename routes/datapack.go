package routes

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
)

type modifierGroup struct {
	models.ConfigModifierGroup
	Modifiers []models.ConfigModifier `json:"modifiers"`
}

type table struct {
	models.Table
	TableType models.ConfigTableType `json:"table_type"`
}

type item struct {
	models.Item
	Category models.Category `json:"category"`
}

func getDatapack(c *gin.Context) {
	config, _ := models.GetConfig()
	modifier_groups := getModifierGroups()
	order_types, _ := models.GetAllOrderTypes()
	table_types, _ := models.GetAllTableTypes()
	categories, _ := models.GetAllCategories()
	tables := getTables()
	items := getItems()

	c.JSON(200, &gin.H{
		"config": config,
		"modifiers": modifier_groups,
		"order_types": order_types,
		"table_types": table_types,
		"tables": tables,
		"items": items,
		"categories": categories,
	})
}

func getModifierGroups() []*modifierGroup {
	modifier_groups, _ := models.GetAllModifierGroups()
	groups := make([]*modifierGroup, len(modifier_groups))

	for i, group := range modifier_groups {
		modifiers, _ := group.GetModifiers()
		groups[i] = &modifierGroup{
			group,
			modifiers,
		}
	}

	return groups
}

func getTables() []*table {
	tablesObjects, _ := models.GetAllTables()
	tables := make([]*table, len(tablesObjects))

	for i, tableObject := range tablesObjects {
		tables[i] = &table{
			Table: tableObject,
			TableType: models.ConfigTableType{Id: tableObject.TypeId},
		}

		tables[i].TableType.Get()
	}

	return tables
}

func getItems() []*item {
	itemObjects, _ := models.GetAllItems()
	items := make([]*item, len(itemObjects))

	for i, itemObject := range itemObjects {
		items[i] = &item{
			Item: itemObject,
			Category: models.Category{Id: itemObject.CategoryId},
		}

		items[i].Category.Get()
	}

	return items
}