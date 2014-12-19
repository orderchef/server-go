
package models

var ConfigTable = "config"

type Config struct {
	Version string
	IsSetup bool
	ClientId string
	ApiKey string
}
