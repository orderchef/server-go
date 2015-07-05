package modifiers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/util"
)

// Get modifier item
// `modifier` (id :modifier_item_id)
func getModifierMiddleware(c *gin.Context) {
	item_id, err := util.GetIntParam("modifier_item_id", c)
	if err != nil {
		return
	}

	modifierGroup := c.MustGet("modifierGroup").(models.ConfigModifierGroup)
	modifier := models.ConfigModifier{Id: item_id, GroupId: modifierGroup.Id}
	err = modifier.Get()
	if err == sql.ErrNoRows {
		util.ServeNotFound(c)
		return
	} else if err != nil {
		util.ServeError(c, err)
		return
	}

	c.Set("modifier", modifier)
	c.Set("modifierId", item_id)

	c.Next()
}

// GET /modifier/:modifier_id/item/:modifier_item_id
func getGroupModifier(c *gin.Context) {
	c.JSON(200, c.MustGet("modifier"))
}

// PUT /modifier/:modifier_id/item/:modifier_item_id
func saveGroupModifier(c *gin.Context) {
	modifier := c.MustGet("modifier").(models.ConfigModifier)

	c.Bind(&modifier)
	modifier.Id = c.MustGet("modifierId").(int)

	if err := modifier.Save(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(201, modifier)
}

// DELETE /modifier/:modifier_id/item/:modifier_item_id
func removeGroupModifier(c *gin.Context) {
	modifier := c.MustGet("modifier").(models.ConfigModifier)

	if err := modifier.Remove(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.AbortWithStatus(204)
}
