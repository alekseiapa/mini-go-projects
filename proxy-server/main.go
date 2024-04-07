package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

import (
	"github.com/alekseiapa/mini-go-projects/proxy-server/routes"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger()) // Helps you use the inbuilt logger from the gin framework
	routes.IncomingRoutes(router)
	log.Fatal(router.Run(":7000"))
}
