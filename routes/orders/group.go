
package orders

import (
	"log"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

func getGroupById(res http.ResponseWriter, params martini.Params) (models.OrderGroup, error) {
	group_id, err := utils.GetIntParam("order_group_id", params, res)
	if err != nil {
		return models.OrderGroup{}, err
	}

	group := models.OrderGroup{Id: group_id}
	if err := group.Get(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return group, err
	}

	return group, nil
}

func GetGroup(res http.ResponseWriter, r render.Render, params martini.Params) {
	group, err := getGroupById(res, params)
	if err != nil {
		return
	}

	r.JSON(200, group)
}

func GetGroupOrders(res http.ResponseWriter, r render.Render, params martini.Params) {
	group, err := getGroupById(res, params)
	if err != nil {
		return
	}

	orders, err := group.GetOrders()
	if err != nil {
		res.WriteHeader(500)
		return
	}

	r.JSON(200, orders)
}

func AddOrderToGroup(res http.ResponseWriter, order models.Order, params martini.Params, r render.Render) {
	group, err := getGroupById(res, params)
	if err != nil {
		return
	}

	order.GroupId = group.Id
	if err := order.Save(); err != nil {
		res.WriteHeader(500)
		return
	}

	r.JSON(201, order)
}
