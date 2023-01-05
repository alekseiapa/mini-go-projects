package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func NewServer() *Server {
	server := &Server{}
	router := gin.Default()
	router.Static("/public", "server/public")
	router.GET("/", func(c *gin.Context) {
		server.HomePage(c, "index.gohtml")
	})
	router.POST("/", func(c *gin.Context) {
		server.HomePage(c, "index.gohtml")
	})
	server.router = router
	return server
}
