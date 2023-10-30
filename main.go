package main

import (
	"log"
	"middle-developer-test/config"
	"middle-developer-test/controller"
)

func main() {
	log.Println("Starting mitrais middle-developer-test")
	cfg := config.InitConfig()

	controller.HandleAPI(cfg)
}
