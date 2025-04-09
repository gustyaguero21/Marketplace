package main

import (
	"log"
	"marketplace/config"
	"marketplace/internal/infrastructure/http/server"
)

func main() {

	router := server.StartServer()

	if err := router.Run(config.Port); err != nil {
		log.Fatal("error starting server")
	}

}
