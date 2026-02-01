package main

import (
	"log"

	"gocashier.db/server"
)

func main() {
if err := server.RunServer(); err != nil {
	log.Fatalf("Error running the server : %s", err)
}
}
