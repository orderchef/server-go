
package tableType

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

func GetAll(c *gin.Context) {
	tableTypes, err := models.GetAllTableTypes()
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, tableTypes)
}

func GetSingle(c *gin.Context) {
	type_id, err := utils.GetIntParam("table_type_id", c)
	if err != nil {
		return
	}

	tableType := models.ConfigTableType{Id: type_id}
	if err := tableType.Get(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, tableType)
}

func Add(c *gin.Context) {
	tableType := models.ConfigTableType{}

	c.Bind(&tableType)

	if err := tableType.Save(); err != nil {
		utils.ServeError(c, err)
	}

	c.JSON(201, tableType)
}

func Save(c *gin.Context) {
	type_id, err := utils.GetIntParam("table_type_id", c)
	if err != nil {
		return
	}

	tableType := models.ConfigTableType{}
	c.Bind(&tableType)

	tableType.Id = type_id

	if err := tableType.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Abort(204)
}

func Delete(c *gin.Context) {
	type_id, err := utils.GetIntParam("table_type_id", c)
	if err != nil {
		return
	}

	tableType := models.ConfigTableType{Id: type_id}

	if err := tableType.Remove(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Abort(204)
}
