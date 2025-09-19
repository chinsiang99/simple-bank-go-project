package routers

import (
	"github.com/chinsiang99/simple-bank-go-project/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(router *gin.Engine, handler handlers.IAccountHandler) {
	users := router.Group("/api/v1/accounts")
	{
		users.POST("", handler.CreateAccount)
		users.GET("/:id", handler.GetAccount)
		// users.GET("", handler.GetAllUsers)
		// users.GET("/:id", handler.GetUserByID)
		// users.PUT("/:id", handler.UpdateUser)
		// users.DELETE("/:id", handler.DeleteUser)
	}
}
