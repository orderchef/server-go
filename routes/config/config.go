
package config

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

func UpdateConfig(c *gin.Context) {

}

func GetConfig(c *gin.Context) {
	config, err := models.GetConfig()
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, config)
}
