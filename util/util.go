package util

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
)

func AuthFailed(c *gin.Context) {
	accept := c.Request.Header.Get("Accept")
	if strings.Contains(accept, "text/html") {
		// c.Redirect(302, "/account/home")
	}

	c.AbortWithStatus(401)

	return
}

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
