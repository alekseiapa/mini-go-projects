package api

import (
	db "github.com/alekseiapa/mini-go-projects/book-store/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/api/books", server.createBook)
	router.GET("/api/books/:id", server.getBook)
	router.GET("'/api/books", server.listBook)
	router.PUT("/api/books/:id", server.updateBookName)
	router.DELETE("/api/books/:id", server.deleteBook)

	server.router = router
	return server
}
