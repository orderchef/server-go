
package orders

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

func getOrderById(c *gin.Context) (models.Order, error) {
	order_id, err := utils.GetIntParam("order_id", c)
	if err != nil {
		return models.Order{}, err
	}

	order := models.Order{Id: order_id}
	if err := order.Get(); err != nil {
		utils.ServeError(c, err)
		return order, err
	}

	return order, nil
}

func GetOrder(c *gin.Context) {
	order, err := getOrderById(c)
	if err != nil {
		return
	}

	c.JSON(200, order)
}

func GetOrderItems(c *gin.Context) {
	order, err := getOrderById(c)
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