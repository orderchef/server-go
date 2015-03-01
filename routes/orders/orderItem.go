
package orders

import (
	"database/sql"
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

	orderItem.OrderId = order.Id
	if err := orderItem.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(201, orderItem)
}

func getOrderItem(c *gin.Context) {
	item_id, err := utils.GetIntParam("item_id", c)
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	item := models.OrderItem{Id: item_id}
	err = item.Get()
	if err == sql.ErrNoRows {
		utils.ServeNotFound(c)
		return
	} else if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Set("orderItem", item)
	c.Set("orderItemId", item_id)
	c.Next()
}

func saveOrderItem(c *gin.Context) {
	orderItem := c.MustGet("orderItem").(models.OrderItem)

	newOrderItem := models.OrderItem{}
	c.Bind(&newOrderItem)

	orderItem.Notes = newOrderItem.Notes

	if err := orderItem.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}

func deleteOrderItem(c *gin.Context) {
	orderItem := c.MustGet("orderItem").(models.OrderItem)
	if err := orderItem.Remove(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}
