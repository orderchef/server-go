package modifiers

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	all := r.Group("/modifiers")
	{
		all.GET("", getModifierGroups)
		all.POST("", addModifierGroup)
	}

	single := r.Group("/modifier/:modifier_id")
	{
		single.Use(getModifierGroupMiddleware)
		single.GET("", getModifierGroup)
		single.PUT("", saveModifierGroup)
		single.DELETE("", removeModifierGroup)

		single.GET("/items", getGroupModifiers)
		single.POST("/items", addGroupModifier)

		modifier := single.Group("/item/:modifier_item_id")
		{
			modifier.Use(getModifierMiddleware)
			modifier.GET("", getGroupModifier)
			modifier.PUT("", saveGroupModifier)
			modifier.DELETE("", removeGroupModifier)
		}
	}
}
