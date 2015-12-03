package orders

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/database"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/util"
)

func addOrderItem(c *gin.Context) {
	order := c.MustGet("order").(models.Order)

	orderItem := models.OrderItem{}
	c.Bind(&orderItem)

	orderItem.Quantity = 1
	orderItem.OrderId = order.Id

	if orderItem.ItemId <= 0 {
		util.ServeError(c, errors.New("Invalid Item ID"))
		return
	}

	if len(orderItem.Notes) == 0 {
		exists, err := database.Mysql().SelectInt("select order__item.id from order__item left join order__item_modifier on order__item_modifier.order_item_id=order__item.id where order__item_modifier.order_item_id is null and item_id=? and order_id=? and char_length(order__item.notes) = 0 limit 1", orderItem.ItemId, order.Id)
		if err != nil {
			panic(err)
		}
		if exists > 0 {
			if _, err := database.Mysql().Exec("update order__item set quantity = quantity + ? where id=?", orderItem.Quantity, exists); err != nil {
				panic(err)
			}

			c.AbortWithStatus(204)
			return
		}
	}

	if err := orderItem.Save(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(201, orderItem)
}

func getOrderItem(c *gin.Context) {
	item_id, err := util.GetIntParam("item_id", c)
	if err != nil {
		util.ServeError(c, err)
		return
	}

	item := models.OrderItem{Id: item_id}
	err = item.Get()
	if err == sql.ErrNoRows {
		util.ServeNotFound(c)
		return
	} else if err != nil {
		util.ServeError(c, err)
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
	orderItem.Quantity = newOrderItem.Quantity
	if orderItem.Quantity <= 0 {
		orderItem.Quantity = 1
	}

	if err := orderItem.Save(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}

func deleteOrderItem(c *gin.Context) {
	orderItem := c.MustGet("orderItem").(models.OrderItem)
	if err := orderItem.Remove(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}
