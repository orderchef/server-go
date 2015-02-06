
package items

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

func GetAll(c *gin.Context) {
	items, err := models.GetAllItems()
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, items)
}

func GetSingle(c *gin.Context) {
	item_id, err := utils.GetIntParam("item_id", c)
	if err != nil {
		return
	}

	item := models.Item{Id: item_id}
	if err := item.Get(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, item)
}

func Add(c *gin.Context) {
	item := models.Item{}
	c.Bind(&item)

	if err := item.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(201, gin.H{})
}

func Save(c *gin.Context) {
	item_id, err := utils.GetIntParam("item_id", c)
	if err != nil {
		return
	}

	item := models.Item{Id: item_id}
	c.Bind(&item)

	if err := item.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Abort(204)
}

func Delete(c *gin.Context) {
	item_id, err := utils.GetIntParam("item_id", c)
	if err != nil {
		return
	}

	item := models.Item{Id: item_id}

	if err := item.Remove(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Abort(204)
}
