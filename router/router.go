package router

import (
	"go-battle/controller"

	"github.com/gin-gonic/gin"
)

func RoutersInit() *gin.Engine {
	router := gin.Default()
	//route group
	v1 := router.Group("/v1")
	v1.GET("/ping", controller.Ping)
	//route param
	v1.GET("/test/:name", controller.TestName)
	return router
}
