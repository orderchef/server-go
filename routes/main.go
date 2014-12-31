package routes

import (
	"log"
	"strconv"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"
	"lab.castawaylabs.com/orderchef/models"
)

func Route(r martini.Router) {
	categoryRouter(r)
}

func categoryRouter(r martini.Router) {
	r.Get("/tables", func(res http.ResponseWriter, r render.Render) {
		tables, err := models.GetAllTables()
		if err != nil {
			log.Println(err)
			res.WriteHeader(500)
			return
		}

		r.JSON(200, tables)
	})

	r.Post("/tables", binding.Bind(models.Table{}), func(res http.ResponseWriter, table models.Table) {
		if err := table.Save(); err != nil {
			log.Println(err)
			res.WriteHeader(500)
		} else {
			res.WriteHeader(201)
		}
	})

	r.Put("/tables/:table_id", binding.Bind(models.Table{}), func(res http.ResponseWriter, table models.Table, params martini.Params) {
		table_id, err := strconv.Atoi(params["table_id"])
		if err != nil {
			log.Println(err)
			res.WriteHeader(400)
			return
		}

		table.Id = table_id

		if err := table.Save(); err != nil {
			log.Println(err)
			res.WriteHeader(500)
		} else {
			res.WriteHeader(204)
		}
	})

	r.Delete("/tables/:table_id", func(res http.ResponseWriter, params martini.Params) {
		table_id, err := strconv.Atoi(params["table_id"])
		if err != nil {
			log.Println(err)
			res.WriteHeader(400)
			return
		}

		table := models.Table{Id: table_id}
		if err := table.Remove(); err != nil {
			log.Println(err)
			res.WriteHeader(500)
		} else {
			res.WriteHeader(204)
		}
	})
}
