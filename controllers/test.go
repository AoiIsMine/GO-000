package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "pong")
}
func TestName(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("nameIs %s", c.Param("name")))
}
