package database

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql" // imports mysql driver
	"lab.castawaylabs.com/orderchef/util"
)

var mysqlDb *gorp.DbMap

// Mysql database
func Mysql() *gorp.DbMap {
	if mysqlDb == nil {
		url := util.Config.MySQL.Username + ":" + util.Config.MySQL.Password + "@tcp(" + util.Config.MySQL.Hostname + ":3306)/" + util.Config.MySQL.DbName + "?parseTime=true"
		if url == ":@tcp(:3306)/?parseTime=true" {
			url = "root@/orderchef?parseTime=true"
		}

		db, err := sql.Open("mysql", url)

		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		mysqlDb = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	}

	return mysqlDb
}
