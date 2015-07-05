package util

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type mySQLConfig struct {
	Hostname string `json:"host"`
	Username string `json:"user"`
	Password string `json:"pass"`
	DbName   string `json:"name"`
}

type mandrillConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type configType struct {
	MySQL     mySQLConfig    `json:"mysql"`
	SessionDb string         `json:"session_db"`
	Mandrill  mandrillConfig `json:"mandrill"`
	Port      string         `json:"port"`
}

var Config configType

func init() {
	configFile, err := Asset("config.json")
	if err != nil {
		fmt.Println("Cannot Find configuration.")
	}

	if err := json.Unmarshal(configFile, &Config); err != nil {
		fmt.Println("Could not decode configuration! (you're a moron)")
		panic(err)
	}

	Config.SessionDb = strings.Replace(Config.SessionDb, "tcp://", "", 1)

	if len(os.Getenv("PORT")) > 0 {
		Config.Port = ":" + os.Getenv("PORT")
	}

	if len(Config.Port) == 0 {
		Config.Port = ":3000"
	}
}
