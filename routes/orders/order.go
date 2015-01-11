
package orders

import (
	"log"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

func getOrderById(res http.ResponseWriter, params martini.Params) (models.Order, error) {
	order_id, err := utils.GetIntParam("order_id", params, res)
	if err != nil {
		return models.Order{}, err
	}

	order := models.Order{Id: order_id}
	if err := order.Get(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return order, err
	}

	return order, nil
}

func GetOrder(res http.ResponseWriter, params martini.Params, r render.Render) {
	order, err := getOrderById(res, params)
	if err != nil {
		return
	}

	r.JSON(200, order)
}

func GetOrderItems(res http.ResponseWriter, params martini.Params, r render.Render) {
	order, err := getOrderById(res, params)
	if err != nil {
		return
	}

	items, err := order.GetItems()
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	r.JSON(200, items)
}