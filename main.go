package main

import (
	"log"

	"gocashier.db/server"
)


// @title GoCashier API
// @version 1.0
// @description API for GoCashier
// @host localhost:8080
// @BasePath /api

func main() {
	if err := server.RunServer(); err != nil {
		log.Fatalf("Error running the server : %s", err)
	}
}
