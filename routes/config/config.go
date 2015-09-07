package config

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/garyburd/redigo/redis"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/routes/config/modifiers"
	"lab.castawaylabs.com/orderchef/routes/config/orderType"
	"lab.castawaylabs.com/orderchef/routes/config/tableType"
	"lab.castawaylabs.com/orderchef/util"
	"lab.castawaylabs.com/orderchef/database"
)

func Router(r *gin.RouterGroup) {
	r.GET("/settings", GetConfig)
	r.POST("/settings", UpdateConfig)
	r.GET("/printers", getPrinters)
	r.GET("/payment-methods", getPaymentMethods)
	r.GET("/bill-items", getBillItems)

	tableType.Router(r)
	orderType.Router(r)
	modifiers.Router(r)
}

func UpdateConfig(c *gin.Context) {
	config := models.Config{}
	c.Bind(&config)

	if err := config.Save(); err != nil {
		util.ServeError(c, err)
		return
	}

	c.Writer.WriteHeader(204)
}

func GetConfig(c *gin.Context) {
	config, err := models.GetConfig()
	if err != nil {
		util.ServeError(c, err)
		return
	}

	c.JSON(200, config)
}

func getPrinters(c *gin.Context) {
	redis_c := database.Redis.Get()
	defer redis_c.Close()

	members, err := redis.Strings(redis_c.Do("SMEMBERS", "oc_printers"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(500)

		return
	}

	c.JSON(200, members)
}

func getPaymentMethods(c *gin.Context) {
	db := database.Mysql()
	pm := []models.ConfigPaymentMethod{}

	if _, err := db.Select(&pm, "select * from config__payment_method"); err != nil {
		panic(err)
	}

	c.JSON(200, pm)
}

func getBillItems(c *gin.Context) {
	db := database.Mysql()
	items := []models.ConfigBillItem{}

	if _, err := db.Select(&items, "select * from config__bill_item"); err != nil {
		panic(err)
	}

	c.JSON(200, items)
}