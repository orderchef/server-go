package orders

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/util"
)

func getOrderById(c *gin.Context) {
	order_id, err := util.GetIntParam("order_id", c)
	if err != nil {
		util.ServeError(c, err)
		return
	}

	order := models.Order{Id: order_id}
	err = order.Get()
	if err == sql.ErrNoRows {
		util.ServeNotFound(c)
		return
	} else if err != nil {
		util.ServeError(c, err)
		return
	}

	c.Set("order", order)
	c.Set("orderId", order_id)
	c.Next()
}

func getOrder(c *gin.Context) (models.Order, error) {
	order, exists := c.Get("order")
	if exists == false {
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
		util.ServeError(c, err)
		return
	}

	c.JSON(200, items)
}
