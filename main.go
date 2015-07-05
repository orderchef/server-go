package main

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/database"
	"lab.castawaylabs.com/orderchef/routes"
)

func main() {
	db := database.Mysql()
	if err := db.CreateTablesIfNotExists(); err != nil {
		panic(err)
	}
	// defer db.Close()

	r := gin.Default()

	api := r.Group("/api")
	routes.Route(api)

	r.Run(":3001")
}
