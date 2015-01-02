
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

func tableRouter(r martini.Router) {
	r.Group("/tables", func(tableRouter martini.Router) {
		tableRouter.Get("", getAllTables)
		tableRouter.Post("", binding.Bind(models.Table{}), addTable)
		tableRouter.Get("/:table_id", getTable)
		tableRouter.Put("/:table_id", binding.Bind(models.Table{}), saveTable)
		tableRouter.Delete("/:table_id", deleteTable)
	})
}

func getAllTables(res http.ResponseWriter, r render.Render) {
	tables, err := models.GetAllTables()
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	r.JSON(200, tables)
}

func getTable(res http.ResponseWriter, r render.Render, params martini.Params) {
	table_id, err := getIntParam("table_id", params, res)
	if err != nil {
		return
	}

	table := models.Table{Id: table_id}
	if err := table.Get(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	r.JSON(200, table)
}

func addTable(res http.ResponseWriter, table models.Table) {
	if err := table.Save(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(201)
	}
}

func saveTable(res http.ResponseWriter, table models.Table, params martini.Params) {
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
}

func deleteTable(res http.ResponseWriter, params martini.Params) {
	table_id, err := getIntParam("table_id", params, res)
	if err != nil {
		return
	}

	table := models.Table{Id: table_id}
	if err := table.Remove(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(204)
	}
}
