package controller

import (
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRouter(router *gin.Engine) {
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		service.Test()
		c.String(http.StatusOK, "Hello %s", name)
	})
}
