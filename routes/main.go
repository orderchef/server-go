
package routes

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/routes/tables"
	"lab.castawaylabs.com/orderchef/routes/configTableType"
	"lab.castawaylabs.com/orderchef/routes/orders"
	"lab.castawaylabs.com/orderchef/routes/categories"
	"lab.castawaylabs.com/orderchef/routes/items"
)

func Route(r martini.Router) {
	r.Group("/tables", tableRouter)
	r.Group("/config", configRouter)
	r.Group("/order-groups", orderGroupRouter)
	r.Group("/orders", ordersRouter)
	r.Group("/categories", categoriesRouter)
	r.Group("/items", itemsRouter)
}

func tableRouter(r martini.Router) {
	r.Get("", tables.GetAll)
	r.Get("/sorted", tables.GetAllSorted)
	r.Post("", binding.Bind(models.Table{}), tables.Add)

	r.Get("/:table_id", tables.GetSingle)
	r.Put("/:table_id", binding.Bind(models.Table{}), tables.Save)
	r.Delete("/:table_id", tables.Delete)

	// Order Group
	r.Get("/:table_id/group", tables.GetOrderGroup)
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

func orderGroupRouter(r martini.Router) {
	r.Group("/:order_group_id", func (groupRouter martini.Router) {
		groupRouter.Get("", orders.GetGroup)
		groupRouter.Get("/orders", orders.GetGroupOrders)
		groupRouter.Post("/orders", binding.Bind(models.Order{}), orders.AddOrderToGroup)
	})
}

func ordersRouter(r martini.Router) {
	r.Group("/:order_id", func (orderRouter martini.Router) {
		orderRouter.Get("", orders.GetOrder)
		orderRouter.Get("/items", orders.GetOrderItems)
	})
}

func categoriesRouter(r martini.Router) {
	// GET /categories -> Get all categories
	r.Get("", categories.GetAll)
	// bidning.Bind takes JSON/Argument POST
	r.Post("", binding.Bind(models.Category{}), categories.Add)

	// :category_id is a parameter
	r.Get("/:category_id", categories.GetSingle)
	r.Put("/:category_id", binding.Bind(models.Category{}), categories.Save)
	r.Delete("/:category_id", categories.Delete)
}

func itemsRouter(r martini.Router) {
	r.Get("", items.GetAll)
	r.Post("", binding.Bind(models.Item{}), items.Add)

	r.Get("/:item_id", items.GetSingle)
	r.Put("/:item_id", binding.Bind(models.Item{}), items.Save)
	r.Delete("/:item_id", items.Delete)
}
