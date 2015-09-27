package config

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"lab.castawaylabs.com/orderchef/database"
	"lab.castawaylabs.com/orderchef/models"
	"lab.castawaylabs.com/orderchef/routes/config/modifiers"
	"lab.castawaylabs.com/orderchef/routes/config/orderType"
	"lab.castawaylabs.com/orderchef/routes/config/tableType"
	"lab.castawaylabs.com/orderchef/util"
	"log"
)

func Router(r *gin.RouterGroup) {
	r.GET("/settings", GetConfig)
	r.POST("/settings", UpdateConfig)
	r.GET("/settings/:settings_key", getConfigByKey)
	r.PUT("/settings/:settings_key", saveConfigByKey)
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

func getConfigByKey(c *gin.Context) {
	db := database.Mysql()

	var key models.DBConfig
	if err := db.SelectOne(&key, "select * from config where name=?", c.Params.ByName("settings_key")); err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, key)
}

func saveConfigByKey(c *gin.Context) {
	db := database.Mysql()

	var key models.DBConfig
	if err := c.Bind(&key); err != nil {
		c.AbortWithStatus(400)
		return
	}

	if _, err := db.Exec("insert into config (name, value) values (?, ?) on duplicate key update value=?", key.Name, key.Value, key.Value); err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.AbortWithStatus(204)
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
