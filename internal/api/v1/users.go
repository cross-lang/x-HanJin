// Package v1 注册 API v1 版本下的业务路由。
package v1

import (
	"x-HanJin/internal/controllers"
	"x-HanJin/internal/services"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes 注册用户模块 API。
func RegisterUserRoutes(v1Router *gin.RouterGroup) {
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)
	userRoutes := v1Router.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.GET("", userController.GetAllUsers)
		userRoutes.GET("/:id", userController.GetUserById)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}
}
