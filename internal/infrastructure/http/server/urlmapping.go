package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UrlMapping(r *gin.Engine) {
	api := r.Group("/api/marketplace/v1")

	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
