package orders

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	all := r.Group("/order-group/:order_group_id")
	{
		all.GET("", GetGroup)
		all.PUT("", updateOrderGroup)
		all.GET("/orders", GetGroupOrders)
		all.POST("/orders", AddOrderToGroup)
	}

	order := r.Group("/order/:order_id")
	{
		order.Use(getOrderById)

		order.GET("", GetOrder)
		order.GET("/items", GetOrderItems)
		order.POST("/items", addOrderItem)

		item := order.Group("/item/:item_id")
		{
			item.Use(getOrderItem)

			item.PUT("", saveOrderItem)
			item.DELETE("", deleteOrderItem)

			item.GET("/modifiers", getOrderItemModifiers)
			item.POST("/modifiers", addOrderItemModifier)
			item.DELETE("/modifier/:order_modifier_id", removeOrderItemModifier)
		}
	}
}
