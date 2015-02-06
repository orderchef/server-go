
package main

import (
	_ "net/http"
	_ "lab.castawaylabs.com/orderchef/models"

	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/routes"
	"lab.castawaylabs.com/orderchef/database"
)

func main() {
	db := database.Mysql()
	if err := db.CreateTablesIfNotExists(); err != nil {
		panic(err)
	}
	// defer db.Close()

	r := gin.Default()

	r.Static("/site", "./templates")
	r.Static("/assets", "./public")

	routes.Route(r.Group("/api"))

	r.Run(":3001")
}
