package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

import (
	"github.com/Uttkarsh-Raj/Proxie/routes"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger()) // Helps you use the inbuilt logger from the gin framework
	routes.IncomingRoutes(router)
	log.Fatal(router.Run(":7000"))
}
