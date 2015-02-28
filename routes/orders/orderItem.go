
package orders

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

func addOrderItem(c *gin.Context) {
	order, err := getOrder(c)
	if err != nil {
		return
	}

	orderItem := models.OrderItem{}
	c.Bind(&orderItem)

	orderItem.Quantity = 1
	orderItem.OrderId = order.Id

	found, err := models.FindExistingOrder(orderItem)

	if err != nil {
		utils.ServeError(c, err)
		return
	}

	if orderItem != found {
		found.Quantity++
	}

	if err := found.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(201, gin.H{})
}

func getOrderItem(c *gin.Context) {
	item_id, err := utils.GetIntParam("item_id", c)
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	item := models.OrderItem{Id: item_id}
	if err := item.Get(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Set("orderItem", item)
	c.Next()
}

func saveOrderItem(c *gin.Context) {
	orderItem_, err := c.Get("orderItem")
	if err != nil {
		return
	}

	orderItem := orderItem_.(models.OrderItem)

	newOrderItem := models.OrderItem{}
	c.Bind(&newOrderItem)

	orderItem.Quantity = newOrderItem.Quantity
	orderItem.Notes = newOrderItem.Notes

	if err := orderItem.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}

func deleteOrderItem(c *gin.Context) {
	orderItem_, err := c.Get("orderItem")
	if err != nil {
		return
	}

	orderItem := orderItem_.(models.OrderItem)
	if err := orderItem.Remove(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}
