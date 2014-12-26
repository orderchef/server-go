
package config

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"lab.castawaylabs.com/orderchef/models"
)

func Router(r martini.Router) {
	// config
	r.Group("/config", func (r martini.Router) {
		r.Get("/table-types", getAllTableTypes)
		r.Post("/table-types", binding.Bind(models.ConfigTableType{}), createTableType)
		r.Get("/table-type/:table_type_id", getTableTypeHandler, getTableType)
		r.Delete("/table-type/:table_type_id", getTableTypeHandler, deleteTableType)
	})
}
