package main

import (
	"log"

	"github.com/alekseiapa/mini-go-projects/go-domain-checker/server"
)

const serverAddress = "0.0.0.0:8081"

func main() {
	server := server.NewServer()

	err := server.Start(serverAddress)
	if err != nil {
		log.Fatal(err)
	}
}
