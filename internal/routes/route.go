// Package routes 负责注册所有 HTTP 路由，
// 按业务模块组织路由分组，并配置 Swagger 文档入口。
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"x-HanJin/internal/controllers"
	_ "x-HanJin/docs"
	"x-HanJin/internal/services"
)

// InitRoutes 注册所有路由到 Gin 引擎
func InitRoutes(r *gin.Engine) {
	// Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 路由组
	v1 := r.Group("/api/v1")

	// 用户模块路由
	userSvr := services.NewUserService()
	usrCtrl := controllers.NewUserController(userSvr)
	userRoutes := v1.Group("/users")
	{
		userRoutes.POST("", usrCtrl.CreateUser)
		userRoutes.GET("", usrCtrl.GetAllUsers)
		userRoutes.GET("/:id", usrCtrl.GetUserById)
		userRoutes.PUT("/:id", usrCtrl.UpdateUser)
		userRoutes.DELETE("/:id", usrCtrl.DeleteUser)
	}

	// 404 兜底路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "resource not found",
		})
	})
}
