
package orders

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

func getOrderById(c *gin.Context) {
	order_id, err := utils.GetIntParam("order_id", c)
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	order := models.Order{Id: order_id}
	if err := order.Get(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Set("order", order)
	c.Next()
}

func getOrder(c *gin.Context) (models.Order, error) {
	order, err := c.Get("order")
	if err != nil {
		return models.Order{}, nil
	}

	return order.(models.Order), nil
}

func GetOrder(c *gin.Context) {
	order, err := getOrder(c)
	if err != nil {
		return
	}

	c.JSON(200, order)
}

func GetOrderItems(c *gin.Context) {
	order, err := getOrder(c)
	if err != nil {
		return
	}

	items, err := order.GetItems()
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, items)
}
