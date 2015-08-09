package categories

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/util"
	"lab.castawaylabs.com/orderchef/database"
)

// Get all printers associated to category
func getPrinters(c *gin.Context) {
	db := database.Mysql()
	category_id, err := util.GetIntParam("category_id", c)
	if err != nil {
		return
	}

	var printers []models.CategoryPrinter
	if _, err := db.Select(&printers, "select printer_id from category_printer where category_id=?", category_id); err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(200, printers)
}

// Link printer to category
func addPrinter(c *gin.Context) {
	db := database.Mysql()

	category_id, _ := util.GetIntParam("category_id", c)
	printer_id := c.Params.ByName("printer_id")

	count, _ := db.SelectInt("select count(1) from category_printer where category_id=? and printer_id=?", category_id, printer_id)

	if count == 0 {
		// create it
		if _, err := db.Exec("insert into category_printer (printer_id, category_id) values (?, ?)", printer_id, category_id); err != nil {
			panic(err)
		}

		c.AbortWithStatus(201)
		return
	}

	c.AbortWithStatus(204)
}

// Remove printer link
func deletePrinter(c *gin.Context) {
	db := database.Mysql()

	category_id, err := util.GetIntParam("category_id", c)
	printer_id := c.Params.ByName("printer_id")
	if err != nil {
		return
	}

	if _, err := db.Exec("delete from category_printer where category_id=? and printer_id=?", category_id, printer_id); err != nil {
		panic(err)
	}

	c.Writer.WriteHeader(204)
}
