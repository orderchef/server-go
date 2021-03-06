package orderType

import (
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/util"
)

func Router(r *gin.RouterGroup) {
	all := r.Group("/order-types")
	{
		all.GET("", GetAll)
		all.POST("", Add)
	}

	single := r.Group("/order-type/:order_type_id")
	{
		single.GET("", GetSingle)
		single.PUT("", Save)
		single.DELETE("", Delete)
	}
}

func GetAll(c *gin.Context) {
	orderTypes, err := models.GetAllOrderTypes()
	if err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(200, orderTypes)
}

func GetSingle(c *gin.Context) {
	type_id, err := util.GetIntParam("order_type_id", c)
	if err != nil {
		return
	}

	orderType := models.ConfigOrderType{Id: type_id}
	if err := orderType.Get(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(200, orderType)
}

func Add(c *gin.Context) {
	orderType := models.ConfigOrderType{}

	c.Bind(&orderType)

	if err := orderType.Save(); err != nil {
		util.ServeError(c, err)
	}

	c.JSON(201, orderType)
}

func Save(c *gin.Context) {
	type_id, err := util.GetIntParam("order_type_id", c)
	if err != nil {
		return
	}

	orderType := models.ConfigOrderType{}
	c.Bind(&orderType)

	orderType.Id = type_id

	if err := orderType.Save(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}

func Delete(c *gin.Context) {
	type_id, err := util.GetIntParam("order_type_id", c)
	if err != nil {
		return
	}

	orderType := models.ConfigOrderType{Id: type_id}

	if err := orderType.Remove(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}
