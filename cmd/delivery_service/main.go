package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/akraskovski/go-delivery-service/internal/app/delivery_service"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/delivery_service.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := delivery_service.NewDeliveryServiceConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal("Cannot read config path from configuration\n", err)
	}

	apiServer := delivery_service.NewAPIServer(config)
	if err := apiServer.Start(); err != nil {
		log.Fatal("Cannot start API Server\n", err)
	}

}
