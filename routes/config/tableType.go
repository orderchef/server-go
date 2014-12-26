
package config

import (
	"log"
	"strconv"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"lab.castawaylabs.com/orderchef/models"
)

func getAllTableTypes(res http.ResponseWriter, r render.Render) {
	tableTypes, err := models.GetAllTableTypes()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.JSON(200, tableTypes)
}

func getTableTypeHandler(res http.ResponseWriter, c martini.Context, params martini.Params) {
	id, err := strconv.Atoi(params["table_type_id"])
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	obj := models.ConfigTableType{Id: id}
	if err := obj.Get(); err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	c.Map(obj)
	c.Next()
}

func getTableType(r render.Render, tableType models.ConfigTableType) {
	r.JSON(200, tableType)
}

func createTableType(res http.ResponseWriter, r render.Render, tableType models.ConfigTableType) {
	log.Println(tableType)

	err := tableType.Save()
	if err != nil {
		log.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	r.JSON(200, tableType)
}

func deleteTableType(res http.ResponseWriter, r render.Render, tableType models.ConfigTableType) {
	if err := tableType.Remove(); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(200)
}
