package main

import (
	"log"
	"marketplace/config"
	"marketplace/internal/infrastructure/http/server"
)

func main() {

	if loadErr := config.LoadEnv(); loadErr != nil {
		log.Fatal(loadErr)
	}

	router := server.StartServer()

	if err := router.Run(config.Port); err != nil {
		log.Fatal("error starting server")
	}

}
