
package routes

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/routes/tables"
	"lab.castawaylabs.com/orderchef/routes/config"
	"lab.castawaylabs.com/orderchef/routes/config/orderType"
	"lab.castawaylabs.com/orderchef/routes/config/tableType"
	"lab.castawaylabs.com/orderchef/routes/orders"
	"lab.castawaylabs.com/orderchef/routes/categories"
	"lab.castawaylabs.com/orderchef/routes/items"
)

func Route(r *gin.RouterGroup) {
	r.GET("/ping", pong)

	tableRouter(r)
	configRouter(r.Group("/config"))
	orderGroupRouter(r.Group("/order-groups"))
	ordersRouter(r.Group("/orders"))
	categoriesRouter(r)
	itemsRouter(r)
}

func tableRouter(r *gin.RouterGroup) {
	ts := r.Group("/tables")
	{
		ts.GET("", tables.GetAll)
		ts.GET("/sorted", tables.GetAllSorted)
		ts.POST("", tables.Add)
	}

	t := r.Group("/table/:table_id")
	{
		t.GET("", tables.GetSingle)
		t.PUT("", tables.Save)
		t.DELETE("", tables.Delete)

		// Order Group
		t.GET("/group", tables.GetOrderGroup)
	}
}

func configRouter(r *gin.RouterGroup) {
	r.GET("/settings", config.GetConfig)
	r.POST("/settings", config.UpdateConfig)

	ts := r.Group("/table-types")
	{
		ts.GET("", tableType.GetAll)
		ts.POST("", tableType.Add)
	}

	t := r.Group("/table-type/:table_type_id")
	{
		t.GET("", tableType.GetSingle)
		t.PUT("", tableType.Save)
		t.DELETE("", tableType.Delete)
	}

	os := r.Group("/order-types")
	{
		os.GET("", orderType.GetAll)
		os.POST("", orderType.Add)
	}

	o := r.Group("/order-type/:order_type_id")
	{
		o.GET("", orderType.GetSingle)
		o.PUT("", orderType.Save)
		o.DELETE("", orderType.Delete)
	}
}

func orderGroupRouter(r *gin.RouterGroup) {
	gs := r.Group("/:order_group_id")
	{
		gs.GET("", orders.GetGroup)
		gs.GET("/orders", orders.GetGroupOrders)
		gs.POST("/orders", orders.AddOrderToGroup)
	}
}

func ordersRouter(r *gin.RouterGroup) {
	os := r.Group("/:order_id")
	{
		os.GET("", orders.GetOrder)
		os.GET("/items", orders.GetOrderItems)
	}
}

func categoriesRouter(r *gin.RouterGroup) {
	cs := r.Group("/categories")
	{
		// GET /categories -> Get all categories
		cs.GET("", categories.GetAll)
		cs.POST("", categories.Add)
	}

	c := r.Group("/category/:category_id")
	{
		// :category_id is a parameter
		c.GET("", categories.GetSingle)
		c.PUT("", categories.Save)
		c.DELETE("", categories.Delete)
	}
}

func itemsRouter(r *gin.RouterGroup) {
	ts := r.Group("/items")
	{
		ts.GET("", items.GetAll)
		ts.POST("", items.Add)
	}

	t := r.Group("/item/:item_id")
	{
		t.GET("", items.GetSingle)
		t.PUT("", items.Save)
		t.DELETE("", items.Delete)
	}
}

func pong(c *gin.Context) {
	c.String(200, "pong")
}
