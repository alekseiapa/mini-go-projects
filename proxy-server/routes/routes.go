package routes

import (
	"github.com/alekseiapa/mini-go-projects/proxy-server/controller"
	"github.com/gin-gonic/gin"
)

func IncomingRoutes(router *gin.Engine) {
	router.GET("/", controller.ProxyServer())
}
