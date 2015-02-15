
package config

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

func UpdateConfig(c *gin.Context) {
	config := models.Config{}
	c.Bind(&config)

	if err := config.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}

func GetConfig(c *gin.Context) {
	config, err := models.GetConfig()
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, config)
}
