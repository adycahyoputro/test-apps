package main

import (
	"backend/server"

	"log"
)

func main() {
	// run the server
	if err := server.Run(); err != nil {
		log.Fatalf("error running the server: %s", err)
	}
}
