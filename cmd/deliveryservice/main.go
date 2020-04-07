package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/akraskovski/go-delivery-service/internal/app/apiserver"
	"github.com/akraskovski/go-delivery-service/internal/app/apiserver/config"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/deliveryservice.toml", "path to config file")
}

func main() {
	flag.Parse()

	appConfig := config.New()
	_, err := toml.DecodeFile(configPath, appConfig)
	if err != nil {
		log.Fatal("Cannot read config path from configuration\n", err)
	}

	log.Fatal("Cannot start API Server\n", apiserver.Start(appConfig))
}
