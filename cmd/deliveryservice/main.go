package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/akraskovski/go-delivery-service/internal/app/apiserver"
	"github.com/akraskovski/go-delivery-service/internal/app/apiserver/config"
	"github.com/kelseyhightower/envconfig"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/deliveryservice.toml", "path to config file")
}

func main() {
	flag.Parse()

	appConfig := config.New()
	readConfigFromFile(appConfig)
	readConfigFromEnv(appConfig)

	log.Fatal("Cannot start API Server\n", apiserver.Start(appConfig))
}

func readConfigFromFile(conf *config.Config) {
	_, err := toml.DecodeFile(configPath, conf)
	if err != nil {
		log.Fatal("Cannot read config path from configuration\n", err)
	}
}

func readConfigFromEnv(conf *config.Config) {
	err := envconfig.Process("delivery-service", conf)
	if err != nil {
		log.Fatal("Cannot read config path from environment\n", err)
	}
}
