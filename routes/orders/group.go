package orders

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/util"
)

type order struct {
	models.Order
	Items []orderItem `json:"items"`
}
type orderItem struct {
	models.OrderItem
	Modifiers []models.OrderItemModifier `json:"modifiers"`
}

func getGroupById(c *gin.Context) (models.OrderGroup, error) {
	group_id, err := util.GetIntParam("order_group_id", c)
	if err != nil {
		return models.OrderGroup{}, err
	}

	group := models.OrderGroup{Id: group_id}
	if err := group.Get(); err != nil {
		util.ServeError(c, err)
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
		util.ServeError(c, err)
		return
	}

	ordersWithItems := make([]order, len(orders))
	for i, orderObject := range orders {
		items, _ := orderObject.GetItems()
		ordersWithItems[i] = order{orderObject, make([]orderItem, len(items))}

		for itemIndex, orderItemObject := range items {
			modifiers, _ := orderItemObject.GetModifiers()
			ordersWithItems[i].Items[itemIndex] = orderItem{
				OrderItem: orderItemObject,
				Modifiers: modifiers,
			}
		}
	}

	c.JSON(200, ordersWithItems)
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
		util.ServeError(c, err)
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
	group.Covers = temp.Covers

	if err := group.Save(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}
