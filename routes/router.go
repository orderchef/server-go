
package routes

import (
	"github.com/go-martini/martini"
	"lab.castawaylabs.com/orderchef/routes/table"
	"lab.castawaylabs.com/orderchef/routes/config"
)

func Route (r martini.Router) {
	config.Router(r)
	table.Router(r)
}
