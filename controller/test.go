package controller

import (
	"go-battle/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	data := service.GetTestService().Ping()
	c.IndentedJSON(http.StatusOK, data)
}
func TestName(c *gin.Context) {
	data := service.GetTestService().TestName(c.Param("name"))
	c.String(http.StatusOK, data)
}
