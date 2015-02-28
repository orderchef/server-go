
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
	isSetup := "0"
	if config.IsSetup == true {
		isSetup = "1"
	}

	keys := [4]DBConfig{
		DBConfig{Name: "is_setup", Value: isSetup},
		DBConfig{Name: "venue_name", Value: config.VenueName},
		DBConfig{Name: "client_id", Value: config.ClientId},
		DBConfig{Name: "api_key", Value: config.ApiKey},
	}

	for _, key := range keys {
		if len(key.Value) == 0 {
			continue
		}

		if err := key.Save(); err != nil {
			return err
		}
	}

	return nil
}

func (config *DBConfig) Save() error {
	db := database.Mysql()

	_, err := db.Exec("insert into config (name, value) values (?, ?) on duplicate key update value=?", config.Name, config.Value, config.Value)
	if err != nil {
		return err
	}

	return nil
}
