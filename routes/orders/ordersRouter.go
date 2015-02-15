
package orders

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	all := r.Group("/order-group/:order_group_id")
	{
		all.GET("", GetGroup)
		all.GET("/orders", GetGroupOrders)
		all.POST("/orders", AddOrderToGroup)
	}

	single := r.Group("/order/:order_id")
	{
		single.GET("", GetOrder)
		single.GET("/items", GetOrderItems)
	}
}
