package controller

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	router := gin.Default()
	SlbRouter(router)
	UserRouter(router)
	return router
}

