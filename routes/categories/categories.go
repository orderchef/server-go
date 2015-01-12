
package categories

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
	categories, err := models.GetAllCategories()
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	r.JSON(200, categories)
}

func GetSingle(res http.ResponseWriter, r render.Render, params martini.Params) {
	category_id, err := utils.GetIntParam("category_id", params, res)
	if err != nil {
		return
	}

	category := models.Category{Id: category_id}
	if err := category.Get(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	r.JSON(200, category)
}

func Add(res http.ResponseWriter, category models.Category) {
	if err := category.Save(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(201)
	}
}

func Save(res http.ResponseWriter, category models.Category, params martini.Params) {
	category_id, err := strconv.Atoi(params["category_id"])
	if err != nil {
		log.Println(err)
		res.WriteHeader(400)
		return
	}

	category.Id = category_id

	if err := category.Save(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(204)
	}
}

func Delete(res http.ResponseWriter, params martini.Params) {
	category_id, err := utils.GetIntParam("category_id", params, res)
	if err != nil {
		return
	}

	category := models.Category{Id: category_id}
	if err := category.Remove(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(204)
	}
}
