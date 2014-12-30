package routes

import (
	"github.com/go-martini/martini"
)

func Route(r martini.Router) {
	categoryRouter(r)
}

func categoryRouter(r martini.Router) {
	r.Get("/tables", func() {
	})
}