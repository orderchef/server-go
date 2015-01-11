package routes

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/routes/tables"
	"lab.castawaylabs.com/orderchef/routes/configTableType"
)

func Route(r martini.Router) {
	r.Group("/tables", tableRouter)
	r.Group("/config", configRouter)
}

func tableRouter(r martini.Router) {
	r.Get("", tables.GetAll)
	r.Get("/sorted", tables.GetAllSorted)
	r.Post("", binding.Bind(models.Table{}), tables.Add)

	r.Get("/:table_id", tables.GetSingle)
	r.Put("/:table_id", binding.Bind(models.Table{}), tables.Save)
	r.Delete("/:table_id", tables.Delete)
}

func configRouter(r martini.Router) {
	r.Group("/table-types", func (table_types martini.Router) {
		table_types.Get("", configTableType.GetAll)
		table_types.Post("", binding.Bind(models.ConfigTableType{}), configTableType.Add)

		table_types.Get("/:table_type_id", configTableType.GetSingle)
		table_types.Put("/:table_type_id", binding.Bind(models.ConfigTableType{}), configTableType.Save)
		table_types.Delete("/:table_type_id", configTableType.Delete)
	})
}
