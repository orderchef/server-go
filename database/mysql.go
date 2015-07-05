package database

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var mysql_db *gorp.DbMap

func Mysql() *gorp.DbMap {
	if mysql_db == nil {
		db, err := sql.Open("mysql", "root:@/orderchef?parseTime=true")

		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		mysql_db = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	}

	return mysql_db
}
