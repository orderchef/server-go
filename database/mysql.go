
package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var mysql_db *sql.DB

func Mysql() *sql.DB {
	if mysql_db == nil {
		db, err := sql.Open("mysql", "root:@/orderchef?parseTime=true")

		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		mysql_db = db
	}

	return mysql_db
}
