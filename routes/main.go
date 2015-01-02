package routes

import (
	"log"
	"strconv"
	"net/http"
	"github.com/go-martini/martini"
)

func Route(r martini.Router) {
	tableRouter(r)
}

func getIntParam(name string, params martini.Params, res http.ResponseWriter) (int, error) {
	intParam, err := strconv.Atoi(params[name])
	if err != nil {
		log.Println(err)
		res.WriteHeader(400)
		return 0, err
	}

	return intParam, nil
}
