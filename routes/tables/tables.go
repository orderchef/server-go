
package tables

import (
	"log"
	"strconv"
	"net/http"
	"database/sql"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

func GetAll(res http.ResponseWriter, r render.Render) {
	tables, err := models.GetAllTables()
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	r.JSON(200, tables)
}

func GetAllSorted(res http.ResponseWriter, r render.Render) {
	types, err := models.GetAllTablesSorted()
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	for _, ttype := range types {
		log.Println(len(ttype.Tables))
	}

	r.JSON(200, types)
}

func GetSingle(res http.ResponseWriter, r render.Render, params martini.Params) {
	table_id, err := utils.GetIntParam("table_id", params, res)
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

func Add(res http.ResponseWriter, table models.Table) {
	if err := table.Save(); err != nil {
		log.Println(err)
		res.WriteHeader(500)
	} else {
		res.WriteHeader(201)
	}
}

func Save(res http.ResponseWriter, table models.Table, params martini.Params) {
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

func Delete(res http.ResponseWriter, params martini.Params) {
	table_id, err := utils.GetIntParam("table_id", params, res)
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

func GetOrderGroup(res http.ResponseWriter, r render.Render, params martini.Params) {
	table_id, err := utils.GetIntParam("table_id", params, res)
	if err != nil {
		return
	}

	table := models.Table{Id: table_id}
	orderGroup := models.OrderGroup{TableId: table.Id}
	err = orderGroup.GetByTableId()

	statusCode := 200

	if err == sql.ErrNoRows {
		if err := orderGroup.Save(); err != nil {
			log.Println(err)
			res.WriteHeader(500)
			return
		}

		statusCode = 201
	} else if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	r.JSON(statusCode, orderGroup)
}
