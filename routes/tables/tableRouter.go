package tables

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("/tables", GetAll)
	r.POST("/tables", Add)
	r.GET("/tables/sorted", GetAllSorted)
	r.GET("/tables/open", GetOpenTables)

	table := r.Group("/table/:table_id")
	{
		table.GET("", GetSingle)
		table.PUT("", Save)
		table.DELETE("", Delete)

		// Order Group
		table.GET("/group", GetOrderGroup)
	}
}
