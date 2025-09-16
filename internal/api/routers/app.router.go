package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAppRoutes sets up application-level routes
func RegisterAppRoutes(router *gin.Engine) {
	app := router.Group("/app")
	{
		// Health check endpoint
		app.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "service is healthy",
			})
		})
	}
}
