package controller

import "github.com/gin-gonic/gin"

func SlbRouter(router *gin.Engine) {
	router.GET("/slb.html", func(c *gin.Context) {
		c.String(200, "SLB check OK!")
	})
}
