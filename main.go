
package main

import (
	_ "net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"lab.castawaylabs.com/orderchef/routes"
	"lab.castawaylabs.com/orderchef/database"
)

func main() {
	db := database.Mysql()
	defer db.Close()

	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Extensions: []string{".html"},
		Delims: render.Delims{"{[{", "}]}"},
		IndentJSON: true,
	}))

	m.Group("/api", routes.Route)

	m.Use(martini.Static("templates", martini.StaticOptions{}))
	m.Use(martini.Static("public", martini.StaticOptions{}))

	m.Run()
}
