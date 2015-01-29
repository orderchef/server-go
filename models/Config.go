
package models

import (
	"lab.castawaylabs.com/orderchef/database"
)

type DBConfig struct {
	Name string `db:"name"`
	Value string `db:"value"`
}

type Config struct {
	IsSetup bool `json:"is_setup"`
	VenueName string `json:"venue_name"`
	ClientId string `json:"client_id"`
	ApiKey string `json:"-"`
}

func init() {
	database.Mysql().AddTableWithName(DBConfig{}, "config")
}

func GetConfig() (Config, error) {
	db := database.Mysql()

	var rows []DBConfig
	config := Config{}

	if _, err := db.Select(&rows, "select * from config"); err != nil {
		return config, err
	}

	for _, row := range rows {
		switch row.Name {
		case "is_setup":
			config.IsSetup = false
			if row.Value == "1" {
				config.IsSetup = true
			}
		case "venue_name":
			config.VenueName = row.Value
		case "client_id":
			config.ClientId = row.Value
		case "api_key":
			config.ApiKey = row.Value
		}
	}

	return config, nil
}

func (config *Config) Save() error {
	// db := database.Mysql()

	return nil
}

func (config *Config) SaveKey() error {
	return nil
}
