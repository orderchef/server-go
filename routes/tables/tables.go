package tables

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/util"
)

func GetAll(c *gin.Context) {
	tables, err := models.GetAllTables()
	if err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(200, tables)
}

func GetAllSorted(c *gin.Context) {
	types, err := models.GetAllTablesSorted()
	if err != nil {
		util.ServeError(c, err)
		return
	}

	// for _, ttype := range types {
	// 	log.Println(len(ttype.Tables))
	// }

	c.JSON(200, types)
}

func GetOpenTables(c *gin.Context) {
	tables, err := models.GetOpenTables()
	if err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(200, tables)
}

func GetSingle(c *gin.Context) {
	table_id, err := util.GetIntParam("table_id", c)
	if err != nil {
		return
	}

	table := models.Table{Id: table_id}
	if err := table.Get(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(200, table)
}

func Add(c *gin.Context) {
	table := models.Table{}

	c.Bind(&table)

	if err := table.Save(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(201, table)
}

func Save(c *gin.Context) {
	table_id, err := util.GetIntParam("table_id", c)
	if err != nil {
		return
	}

	table := models.Table{}

	c.Bind(&table)
	table.Id = table_id

	if err := table.Save(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}

func Delete(c *gin.Context) {
	table_id, err := util.GetIntParam("table_id", c)
	if err != nil {
		return
	}

	table := models.Table{Id: table_id}

	if err := table.Remove(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}

func GetOrderGroup(c *gin.Context) {
	table_id, err := util.GetIntParam("table_id", c)
	if err != nil {
		return
	}

	table := models.Table{Id: table_id}
	orderGroup := models.OrderGroup{TableId: table.Id}
	err = orderGroup.GetByTableId()

	statusCode := 200

	if err == sql.ErrNoRows {
		// no results, create new group
		if err := orderGroup.Save(); err != nil {
			util.ServeError(c, err)
			return
		}

		statusCode = 201
	} else if err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(statusCode, orderGroup)
}
