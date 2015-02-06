
package tables

import (
	"log"
	"database/sql"
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

func GetAll(c *gin.Context) {
	tables, err := models.GetAllTables()
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, tables)
}

func GetAllSorted(c *gin.Context) {
	types, err := models.GetAllTablesSorted()
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	for _, ttype := range types {
		log.Println(len(ttype.Tables))
	}

	c.JSON(200, types)
}

func GetSingle(c *gin.Context) {
	table_id, err := utils.GetIntParam("table_id", c)
	if err != nil {
		return
	}

	table := models.Table{Id: table_id}
	if err := table.Get(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, table)
}

func Add(c *gin.Context) {
	table := models.Table{}

	c.Bind(&table)

	if err := table.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(201, gin.H{})
}

func Save(c *gin.Context) {
	table_id, err := utils.GetIntParam("table_id", c)
	if err != nil {
		return
	}

	table := models.Table{}

	c.Bind(&table)
	table.Id = table_id

	if err := table.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Abort(204)
}

func Delete(c *gin.Context) {
	table_id, err := utils.GetIntParam("table_id", c)
	if err != nil {
		return
	}

	table := models.Table{Id: table_id}

	if err := table.Remove(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Abort(204)
}

func GetOrderGroup(c *gin.Context) {
	table_id, err := utils.GetIntParam("table_id", c)
	if err != nil {
		return
	}

	table := models.Table{Id: table_id}
	orderGroup := models.OrderGroup{TableId: table.Id}
	err = orderGroup.GetByTableId()

	statusCode := 200

	if err == sql.ErrNoRows {
		if err := orderGroup.Save(); err != nil {
			utils.ServeError(c, err)
			return
		}

		statusCode = 201
	} else if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(statusCode, orderGroup)
}
