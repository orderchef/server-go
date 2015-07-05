package routes

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/routes/categories"
	"lab.castawaylabs.com/orderchef/routes/config"
	"lab.castawaylabs.com/orderchef/routes/items"
	"lab.castawaylabs.com/orderchef/routes/orders"
	"lab.castawaylabs.com/orderchef/routes/tables"
)

func Route(r *gin.RouterGroup) {
	r.GET("/ping", pong)

	tables.Router(r)
	config.Router(r.Group("/config"))
	orders.Router(r)
	categories.Router(r)
	items.Router(r)
}

func pong(c *gin.Context) {
	c.String(200, "pong")
}
