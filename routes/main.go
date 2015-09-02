package routes

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/routes/categories"
	"lab.castawaylabs.com/orderchef/routes/config"
	"lab.castawaylabs.com/orderchef/routes/items"
	"lab.castawaylabs.com/orderchef/routes/orders"
	"lab.castawaylabs.com/orderchef/routes/tables"
	"lab.castawaylabs.com/orderchef/routes/reports"
)

func Route(r *gin.RouterGroup) {
	r.GET("/ping", pong)

	r.GET("/datapack", getDatapack)

	tables.Router(r)
	config.Router(r.Group("/config"))
	orders.Router(r)
	categories.Router(r)
	items.Router(r)
	reports.Router(r.Group("/reports"))
}

func pong(c *gin.Context) {
	c.String(200, "pong")
}