// Package api 负责聚合并注册所有 HTTP API 路由。
package api

import (
	"net/http"

	_ "x-HanJin/docs"
	v1 "x-HanJin/internal/api/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRoutes 注册所有路由到 Gin 引擎。
func InitRoutes(r *gin.Engine) {
	// Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 路由组，由各业务模块分别注册具体接口。
	v1Router := r.Group("/api/v1")
	v1.RegisterUserRoutes(v1Router)

	// 404 兜底路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "resource not found",
		})
	})
}
