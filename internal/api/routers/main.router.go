package routers

import (
	"github.com/chinsiang99/simple-bank-go-project/internal/api/handlers"
	"github.com/chinsiang99/simple-bank-go-project/internal/api/middlewares"
	"github.com/chinsiang99/simple-bank-go-project/internal/config"
	_ "github.com/chinsiang99/simple-bank-go-project/internal/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouterManager(handlers *handlers.HandlerManager, appConfig *config.AppConfig, securityConfig *config.SecurityConfig) *gin.Engine {
	// Set Gin mode if it is production
	if appConfig.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Swagger endpoint (disable in production)
	// if config.Environment != "production" {
	// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// }

	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middlewares.Log())
	router.Use(middlewares.SecurityMiddleware(appConfig, securityConfig))

	// --------------------------
	// Swagger endpoint
	// --------------------------
	// Only expose Swagger in non-production, or protect in production
	if appConfig.Environment != "hehe" {
		url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	RegisterAuthRoutes(router, handlers.AuthHandler)
	RegisterAppRoutes(router)
	RegisterUserRoutes(router, handlers.UserHandler)
	// RegisterTicketRoutes(router, h.TicketHandler)

	return router
}
