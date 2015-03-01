
package orders

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

// GET /modifiers
func getOrderItemModifiers(c *gin.Context) {
	orderItem := c.MustGet("orderItem").(models.OrderItem)

	modifiers, err := orderItem.GetModifiers()
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, modifiers)
}

// POST /modifiers
func addOrderItemModifier(c *gin.Context) {
	orderItem := c.MustGet("orderItem").(models.OrderItem)

	modifier := models.OrderItemModifier{}
	c.Bind(&modifier)

	modifier.OrderItemId = orderItem.Id

	if err := modifier.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(201, modifier)
}

// DELETE /modifier/:order_modifier_id
func removeOrderItemModifier(c *gin.Context) {
	orderItem := c.MustGet("orderItem").(models.OrderItem)

	modifier_id, err := utils.GetIntParam("order_modifier_id", c)
	if err != nil {
		return
	}

	modifier := models.OrderItemModifier{Id: modifier_id, OrderItemId: orderItem.Id}
	if err := modifier.Remove(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.AbortWithStatus(204)
}
