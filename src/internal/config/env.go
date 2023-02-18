package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Host         string `default:"0.0.0.0"`
	Port         int16  `default:"8000"`
	MongoDBUri   string `required:"true"`
	MongoDBName  string `required:"true"`
	AuthUsername string `required:"false"`
	AuthPassword string `required:"false"`
}

func initConfig() Config {
	var c Config

	err := envconfig.Process("app", &c)

	if err != nil {
		log.Fatal(err.Error())
	}

	return c
}

var Env Config = initConfig()
