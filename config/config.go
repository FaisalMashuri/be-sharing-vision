package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBName         string `envconfig:"DBNAME"`
	HOSTDB         string `envconfig:"HOSTDB" required:"true"`
	DBPORT         string `envconfig:"DBPORT" required:"true"`
	DBUSER         string `envconfig:"DBUSER" required:"true"`
	DBPASSWORD     string `envconfig:"DBPASSWORD" required:"true"`
	APPPORT        string `envconfig:"APPPORT"`
	APPHTTPDEBUG   bool   `envconfig:"APPHTTPDEBUG" required:"true"`
	SERVER_TIMEOUT int    `envconfig:"SERVER_TIMEOUT"`
}

func Load() (Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	return c, err
}
