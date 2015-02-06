
package utils

import (
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetIntParam(name string, c *gin.Context) (int, error) {
	intParam, err := strconv.Atoi(c.Params.ByName(name))
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{})
		return 0, err
	}

	return intParam, nil
}

func ServeError(c *gin.Context, err error) {
	log.Println(err)
	c.JSON(500, gin.H{})
	return
}
