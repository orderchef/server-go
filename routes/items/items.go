
package items

import (
	"log"
	"strconv"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

func GetAll(res http.ResponseWriter, r render.Render) {
	items, err := models.GetAllItems()
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	r.JSON(200, items)
}

func GetSingle(res http.ResponseWriter, r render.Render, params martini.Params) {
	item_id, err := utils.GetIntParam("item_id", params, res)
	if err != nil {
		return
	}

	item := models.Item{Id: item_id}
	if err := item.Get(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	r.JSON(200, item)
}

func Add(res http.ResponseWriter, item models.Item) {
	if err := item.Save(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(201)
	}
}

func Save(res http.ResponseWriter, item models.Item, params martini.Params) {
	item_id, err := strconv.Atoi(params["item_id"])
	if err != nil {
		log.Println(err)
		res.WriteHeader(400)
		return
	}

	item.Id = item_id

	if err := item.Save(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(204)
	}
}

func Delete(res http.ResponseWriter, params martini.Params) {
	item_id, err := utils.GetIntParam("item_id", params, res)
	if err != nil {
		return
	}

	item := models.Item{Id: item_id}
	if err := item.Remove(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(204)
	}
}
