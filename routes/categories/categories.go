
package categories

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/utils"
)

// Get all categories
func GetAll(c *gin.Context) {
	categories, err := models.GetAllCategories()
	// if there's an error, its a server error..
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, categories)
}

// Fetch a single category
func GetSingle(c *gin.Context) {
	category_id, err := utils.GetIntParam("category_id", c)
	if err != nil {
		return
	}

	// Create category with Id property.
	category := models.Category{Id: category_id}
	// Fetch the other properties from the database (uses ID)
	if err := category.Get(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(200, category)
}

// Create new category
// the 'category' property is injected using the Bind service (from routes/main.go)
func Add(c *gin.Context) {
	category := models.Category{}
	c.Bind(&category)

	if err := category.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.JSON(201, gin.H{})
}

// Update category
func Save(c *gin.Context) {
	category_id, err := utils.GetIntParam("category_id", c)
	if err != nil {
		utils.ServeError(c, err)
		return
	}

	category := models.Category{Id: category_id}
	c.Bind(&category)

	if err := category.Save(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Abort(204)
}

// Delete category
func Delete(c *gin.Context) {
	category_id, err := utils.GetIntParam("category_id", c)
	if err != nil {
		return
	}

	// Gorp looks at the Id (primary key) to delete the document
	category := models.Category{Id: category_id}
	if err := category.Remove(); err != nil {
		utils.ServeError(c, err)
		return
	}

	c.Abort(204)
}
