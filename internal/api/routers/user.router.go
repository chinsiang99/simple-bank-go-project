package routers

import (
	"github.com/chinsiang99/simple-bank-go-project/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, handler handlers.IUserHandler) {
	users := router.Group("/api/v1/users")
	{
		users.POST("", handler.CreateUser)
		users.GET("", handler.GetAllUsers)
		users.GET("/:id", handler.GetUserByID)
		users.PUT("/:id", handler.UpdateUser)
		users.DELETE("/:id", handler.DeleteUser)
	}
}
