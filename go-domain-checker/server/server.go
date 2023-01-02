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
	router.GET("/", server.HomePage)
	server.router = router
	return server
}
