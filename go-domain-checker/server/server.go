package server

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

func (s *Server) index(ctx *gin.Context) {

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func NewServer() *Server {
	server := &Server{}
	router := gin.Default()

	router.GET("/", server.index)
	server.router = router
	return server
}
