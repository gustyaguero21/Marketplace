package http

import "github.com/gin-gonic/gin"

func SetupServer() *gin.Engine {
	router := gin.Default()

	UrlMapping(router)

	return router
}
