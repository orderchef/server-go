
package items

import (
	"fmt"
	"database/sql"
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
	"lab.castawaylabs.com/orderchef/database"
)

func Router(r *gin.RouterGroup) {
	all := r.Group("/items")
	{
		all.GET("", GetAll)
		all.POST("", Add)
	}

	single := r.Group("/item/:item_id")
	{
		single.Use(getItem)

		single.GET("", GetSingle)
		single.PUT("", Save)
		single.DELETE("", Delete)

		// Modifiers
		single.GET("/modifiers", getItemModifiers)
		single.POST("/modifiers", addItemModifier)
		single.DELETE("/modifiers", removeItemModifiers)
		single.DELETE("/modifier/:modifier_group_id", removeItemModifier)
	}
}

func GetAll(c *gin.Context) {
	items, err := models.GetAllItems()
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, items)
}

func getItem(c *gin.Context) {
	item_id, err := utils.GetIntParam("item_id", c)
	if err != nil {
		return
	}

	item := models.Item{Id: item_id}
	if err := item.Get(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Set("item", item)
	c.Set("itemId", item_id)

	c.Next()
}

func GetSingle(c *gin.Context) {
	c.JSON(200, c.MustGet("item"))
}

func Add(c *gin.Context) {
	item := models.Item{}
	c.Bind(&item)

	if err := item.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(201, item)
}

func Save(c *gin.Context) {
	item := c.MustGet("item").(models.Item)
	c.Bind(&item)

	if err := item.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(201, item)
}

func Delete(c *gin.Context) {
	item := c.MustGet("item").(models.Item)

	if err := item.Remove(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}

func getItemModifiers(c *gin.Context) {
	item := c.MustGet("item").(models.Item)
	modifiers, err := item.GetModifiers()

	if err != nil {
		utils.ServeError(c, err)
		return
	}
	fmt.Println(modifiers)

	c.JSON(200, modifiers)
}

func addItemModifier(c *gin.Context) {
	item := c.MustGet("item").(models.Item)
	modifiers, err := item.GetModifiers()

	if err != nil {
		utils.ServeError(c, err)
		return
	}

	new_modifier := models.ItemModifier{}
	c.Bind(&new_modifier)
	new_modifier.Item = item.Id

	found := false
	for _, modifier := range modifiers {
		if modifier == new_modifier.ModifierGroup {
			found = true
			break
		}
	}

	if found == true {
		c.AbortWithStatus(204)
		return
	}

	err = new_modifier.Save()
	if err != nil {
		utils.ServeError(c, err)
		return
	}
	fmt.Println(new_modifier)
	c.AbortWithStatus(201)
}

func removeItemModifier(c *gin.Context) {
	item := c.MustGet("item").(models.Item)
	group_id, err := utils.GetIntParam("modifier_group_id", c)
	if err != nil {
		return
	}

	modifier := models.ItemModifier{Item: item.Id, ModifierGroup: group_id}
	err = modifier.Remove()
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.AbortWithStatus(204)
}

func removeItemModifiers(c *gin.Context) {
	item := c.MustGet("item").(models.Item)

	db := database.Mysql()
	if _, err := db.Exec("delete from item__modifier where item_id=?", item.Id); err != nil && err != sql.ErrNoRows {
		utils.ServeError(c, err)
	} else {
		c.AbortWithStatus(204)
	}
}
