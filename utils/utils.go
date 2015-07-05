package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func GetIntParam(name string, c *gin.Context) (int, error) {
	intParam, err := strconv.Atoi(c.Params.ByName(name))
	if err != nil {
		c.AbortWithStatus(400)
		return 0, err
	}

	return intParam, nil
}

func ServeError(c *gin.Context, err error) {
	log.Println(err)
	c.AbortWithStatus(500)

	return
}

func ServeNotFound(c *gin.Context) {
	c.AbortWithStatus(404)
	return
}
