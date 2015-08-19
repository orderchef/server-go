package orders

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	func (api *gin.RouterGroup) {
		// /order-group/:order_group_id
		api.GET("", GetGroup)
		api.PUT("", updateOrderGroup)
		api.GET("/orders", GetGroupOrders)
		api.POST("/orders", AddOrderToGroup)
		api.POST("/clear", clearGroup)

		func (api *gin.RouterGroup) {
			// /order-group/:order_group_id/bills
			api.GET("", getAllBills)
			api.POST("", makeBill)
			api.GET("/totals", getBillTotals)
		}(api.Group("/bills"))

		func (api *gin.RouterGroup) {
			api.Use(getBill)

			api.GET("", serveBill)
			api.PUT("", updateBill)
			api.DELETE("", deleteBill)
			api.POST("/print", printBill)
		}(api.Group("/bill/:bill_id"))
	}(r.Group("/order-group/:order_group_id"))

	func (api *gin.RouterGroup) {
		// /order/:order_id
		api.Use(getOrderById)

		api.GET("", GetOrder)
		api.GET("/items", GetOrderItems)
		api.POST("/items", addOrderItem)
		api.DELETE("", DeleteOrder)
		api.POST("/print", PrintOrder)

		func (api *gin.RouterGroup) {
			// /order/:order_id/item/:item_id
			api.Use(getOrderItem)

			api.PUT("", saveOrderItem)
			api.DELETE("", deleteOrderItem)

			api.GET("/modifiers", getOrderItemModifiers)
			api.POST("/modifiers", addOrderItemModifier)
			api.DELETE("/modifier/:order_modifier_id", removeOrderItemModifier)
		}(api.Group("/item/:item_id"))
	}(r.Group("/order/:order_id"))
}
