
package modifiers

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	all := r.Group("/modifiers")
	{
		all.GET("", getAllModifiers)
	}
}
