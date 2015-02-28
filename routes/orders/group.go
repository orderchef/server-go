
package orders

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

func getGroupById(c *gin.Context) (models.OrderGroup, error) {
	group_id, err := utils.GetIntParam("order_group_id", c)
	if err != nil {
		return models.OrderGroup{}, err
	}

	group := models.OrderGroup{Id: group_id}
	if err := group.Get(); err != nil {
		utils.ServeError(c, err)
		return group, err
	}

	return group, nil
}

func GetGroup(c *gin.Context) {
	group, err := getGroupById(c)
	if err != nil {
		return
	}

	c.JSON(200, group)
}

func GetGroupOrders(c *gin.Context) {
	group, err := getGroupById(c)
	if err != nil {
		return
	}

	orders, err := group.GetOrders()
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, orders)
}

func AddOrderToGroup(c *gin.Context) {
	group, err := getGroupById(c)
	if err != nil {
		return
	}

	order := models.Order{}
	c.Bind(&order)

	order.GroupId = group.Id

	if err := order.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(201, order)
}

func updateOrderGroup(c *gin.Context) {
	group, err := getGroupById(c)
	if err != nil {
		return
	}

	temp := models.OrderGroup{}
	c.Bind(&temp)

	group.TableId = temp.TableId

	if err := group.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}
