
package configTableType

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
	tableTypes, err := models.GetAllTableTypes()
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	r.JSON(200, tableTypes)
}

func GetSingle(res http.ResponseWriter, r render.Render, params martini.Params) {
	type_id, err := utils.GetIntParam("table_type_id", params, res)
	if err != nil {
		return
	}

	tableType := models.ConfigTableType{Id: type_id}
	if err := tableType.Get(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	r.JSON(200, tableType)
}

func Add(res http.ResponseWriter, tableType models.ConfigTableType) {
	if err := tableType.Save(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(201)
	}
}

func Save(res http.ResponseWriter, tableType models.ConfigTableType, params martini.Params) {
	type_id, err := strconv.Atoi(params["table_type_id"])
	if err != nil {
		log.Println(err)
		res.WriteHeader(400)
		return
	}

	tableType.Id = type_id

	if err := tableType.Save(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(204)
	}
}

func Delete(res http.ResponseWriter, params martini.Params) {
	type_id, err := utils.GetIntParam("table_type_id", params, res)
	if err != nil {
		return
	}

	tableType := models.ConfigTableType{Id: type_id}
	if err := tableType.Remove(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(204)
	}
}
