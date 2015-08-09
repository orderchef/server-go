package items

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/util"
	"lab.castawaylabs.com/orderchef/database"
)

// Get all printers associated to an item
func getPrinters(c *gin.Context) {
	db := database.Mysql()
	item_id, err := util.GetIntParam("item_id", c)
	if err != nil {
		return
	}

	var printers []models.CategoryPrinter
	if _, err := db.Select(&printers, "select printer_id from category_printer where item_id=?", item_id); err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(200, printers)
}

// Link printer to category
func addPrinter(c *gin.Context) {
	db := database.Mysql()

	item_id, _ := util.GetIntParam("item_id", c)
	printer_id := c.Params.ByName("printer_id")

	count, _ := db.SelectInt("select count(1) from category_printer where item_id=? and printer_id=?", item_id, printer_id)

	if count == 0 {
		// create it
		if _, err := db.Exec("insert into category_printer (printer_id, item_id) values (?, ?)", printer_id, item_id); err != nil {
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

	item_id, err := util.GetIntParam("item_id", c)
	printer_id := c.Params.ByName("printer_id")
	if err != nil {
		return
	}

	if _, err := db.Exec("delete from category_printer where item_id=? and printer_id=?", item_id, printer_id); err != nil {
		panic(err)
	}

	c.Writer.WriteHeader(204)
}
