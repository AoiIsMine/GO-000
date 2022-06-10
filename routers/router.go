package routers

import (
	"go-battle/controllers"

	"github.com/gin-gonic/gin"
)

func RoutersInit() *gin.Engine {
	router := gin.Default()
	//route group
	v1 := router.Group("/v1")
	v1.GET("/ping", controllers.Ping)
	//route param
	v1.GET("/test/:name", controllers.TestName)
	return router
}
