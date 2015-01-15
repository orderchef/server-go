
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

// Get all categories
func GetAll(res http.ResponseWriter, r render.Render) {
	categories, err := models.GetAllCategories()
	// if there's an error, its a server error..
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	r.JSON(200, categories)
}

// Fetch a single category
func GetSingle(res http.ResponseWriter, r render.Render, params martini.Params) {
	category_id, err := utils.GetIntParam("category_id", params, res)
	if err != nil {
		return
	}

	// Create category with Id property.
	category := models.Category{Id: category_id}
	// Fetch the other properties from the database (uses ID)
	if err := category.Get(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	r.JSON(200, category)
}

// Create new category
// the 'category' property is injected using the Bind service (from routes/main.go)
func Add(res http.ResponseWriter, category models.Category) {
	if err := category.Save(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(201)
	}
}

// Update category
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

// Delete category
func Delete(res http.ResponseWriter, params martini.Params) {
	category_id, err := utils.GetIntParam("category_id", params, res)
	if err != nil {
		return
	}

	// Gorp looks at the Id (primary key) to delete the document
	category := models.Category{Id: category_id}
	if err := category.Remove(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(204)
	}
}
