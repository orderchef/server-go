
package table

import (
	"log"
	"strconv"
	"net/http"
	"lab.castawaylabs.com/orderchef/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"
)

func Router(r martini.Router) {
	r.Group("/tables", tablesRouter)
	r.Group("/table/:table_id", tableRouter)
}

func tablesRouter(r martini.Router) {
	r.Get("", getAllTables)
	r.Post("", binding.Bind(models.Table{}), createTable)
}

func tableRouter(r martini.Router) {
	r.Get("", getTableHandler, getTable)
}

func getTableHandler(res http.ResponseWriter, c martini.Context, params martini.Params) {
	id, err := strconv.Atoi(params["table_id"])
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	table := models.Table{Id: id}
	if err := table.Get(); err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	c.Map(table)
	c.Next()
}

func getAllTables(res http.ResponseWriter, r render.Render) {
	tables, err := models.GetAllTables()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(tables)
	r.JSON(200, tables)
}

func createTable(res http.ResponseWriter, r render.Render, table models.Table) {
	log.Println(table)

	err := table.Save()
	if err != nil {
		log.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	r.JSON(200, table)
}

func getTable(r render.Render, table models.Table) {
	r.JSON(200, table)
}
