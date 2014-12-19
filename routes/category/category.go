
package category

import (
	"log"
	"strconv"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"
	"lab.castawaylabs.com/orderchef/models"
)

func Router(r martini.Router) {
	r.Get("/categories", getAll)
	r.Post("/category", binding.Bind(models.category{}), create)
	r.Get("/category/:category_id", getModelHandler, get)

	r.Get("/category/:category_id/printers", getModelHandler, getPrinters)
	r.Put("/category/:category_id/printer/:printer_id", getModelHandler, getPrinterModelHandler, getPrinter)
	r.Delete("/category/:category_id/printer/:printer_id", getModelHandler, getPrinterModelHandler, removePrinter)
}

func getAll(res http.ResponseWriter, r render.Render) {
	objs, err := models.GetAllCategories()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.JSON(200, objs)
}

func getModelHandler(res http.ResponseWriter, c martini.Context, params martini.Params) {
	id, err := strconv.Atoi(params["category_id"])
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	obj := models.Category{Id: id}
	if err := obj.Get(); err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	c.Map(obj)
	c.Next()
}

func get(r render.Render, category models.Category) {
	r.JSON(200, category)
}

func create(res http.ResponseWriter, r render.Render, category models.Category) {
	err := tableType.Save()
	if err != nil {
		log.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	r.JSON(200, tableType)
}

func getPrinters(res http.ResponseWriter, r render.Render, category Category) {
	printers, err := category.getPrinters()
	if err != nil {
		log.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	r.JSON(200, printers)
}

func getPrinter(res http.ResponseWriter, r render.Render, category Category, printer CategoryPrinter) {
	r.JSON(200, printer)
}

func removePrinter(res http.ResponseWriter, printer CategoryPrinter) {
	if err := printer.Remove(); err != nil {
		http.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOk)
}

func getPrinterModelHandler(res http.ResponseWriter, c martini.Context, params martini.Params) {
	id, err := strconv.Atoi(params["printer_id"])
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	obj := models.CategoryPrinter{Id: id}
	if err := obj.Get(); err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	c.Map(obj)
	c.Next()
}
