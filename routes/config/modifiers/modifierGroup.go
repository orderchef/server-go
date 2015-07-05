package modifiers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

// Get modifier group
// `modifierGroup` (id :modifier_id)
func getModifierGroupMiddleware(c *gin.Context) {
	modifier_id, err := utils.GetIntParam("modifier_id", c)
	if err != nil {
		return
	}

	modifier := models.ConfigModifierGroup{Id: modifier_id}
	err = modifier.Get()
	if err == sql.ErrNoRows {
		utils.ServeNotFound(c)
		return
	} else if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Set("modifierGroup", modifier)
	c.Set("modifierGroupId", modifier_id)

	c.Next()
}

// GET /modifiers
func getModifierGroups(c *gin.Context) {
	modifierGroups, err := models.GetAllModifierGroups()
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, modifierGroups)
}

// POST /modifiers
func addModifierGroup(c *gin.Context) {
	modifier := models.ConfigModifierGroup{}
	c.Bind(&modifier)

	if err := modifier.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(201, modifier)
}

// GET /modifier/:modifier_id
func getModifierGroup(c *gin.Context) {
	modifier := c.MustGet("modifierGroup").(models.ConfigModifierGroup)
	c.JSON(200, modifier)
}

// PUT /modifier/:modifier_id
func saveModifierGroup(c *gin.Context) {
	modifier := c.MustGet("modifierGroup").(models.ConfigModifierGroup)

	c.Bind(&modifier)
	modifier.Id = c.MustGet("modifierGroupId").(int)

	if err := modifier.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(201, modifier)
}

// DELETE /modifer/:modifer_id
func removeModifierGroup(c *gin.Context) {
	modifier := c.MustGet("modifierGroup").(models.ConfigModifierGroup)

	if err := modifier.Remove(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.AbortWithStatus(204)
}

// GET /modifier/:modifier_id/items
func getGroupModifiers(c *gin.Context) {
	modifier := c.MustGet("modifierGroup").(models.ConfigModifierGroup)

	modifiers, err := modifier.GetModifiers()
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, modifiers)
}

// POST /modifier/:modifier_id/items
func addGroupModifier(c *gin.Context) {
	modifier := c.MustGet("modifierGroup").(models.ConfigModifierGroup)

	item := models.ConfigModifier{}
	c.Bind(&item)
	item.GroupId = modifier.Id

	if err := item.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(201, item)
}
