package config

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/routes/config/modifiers"
	"lab.castawaylabs.com/orderchef/routes/config/orderType"
	"lab.castawaylabs.com/orderchef/routes/config/tableType"
	"lab.castawaylabs.com/orderchef/util"
)

func Router(r *gin.RouterGroup) {
	r.GET("/settings", GetConfig)
	r.POST("/settings", UpdateConfig)

	tableType.Router(r)
	orderType.Router(r)
	modifiers.Router(r)
}

func UpdateConfig(c *gin.Context) {
	config := models.Config{}
	c.Bind(&config)

	if err := config.Save(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}

func GetConfig(c *gin.Context) {
	config, err := models.GetConfig()
	if err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(200, config)
}
