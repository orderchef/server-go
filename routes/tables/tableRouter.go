
package tables

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	ts := r.Group("/tables")
	{
		ts.GET("", GetAll)
		ts.GET("/sorted", GetAllSorted)
		ts.POST("", Add)
	}

	t := r.Group("/table/:table_id")
	{
		t.GET("", GetSingle)
		t.PUT("", Save)
		t.DELETE("", Delete)

		// Order Group
		t.GET("/group", GetOrderGroup)
	}
}
