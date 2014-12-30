
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type Config struct {
	Version string `db:"version"`
	IsSetup bool `db:"isSetup"`
	ClientId string `db:"client_id"`
	ApiKey string `db:"api_key"`
}

func init() {
	database.Mysql().AddTableWithName(Config{}, "config")
}
