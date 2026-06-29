// Package main 是 x-HanJin 应用的启动入口
package main

import (
	"fmt"

	"x-HanJin/internal/config"
	"x-HanJin/internal/databases"
	"x-HanJin/internal/message_queues"
	"x-HanJin/internal/routes"
	"x-HanJin/pkg/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


// @title           x-HanJin API
// @description     汉津应用 API 接口文档
// @termsOfService  https://gitee.com/cross-lang/x-HanJin

// @contact.name   yeyushilai-team
// @contact.url    https://gitee.com/cross-lang/x-HanJin

// @license.name  MIT
// @license.url   https://gitee.com/cross-lang/x-HanJin/blob/main/LICENSE

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		panic(fmt.Sprintf("配置初始化失败: %v", err))
	}

	// 初始化日志
	if err := log.InitLogger(config.Cfg.Logger); err != nil {
		panic(fmt.Sprintf("日志初始化失败: %v", err))
	}

	// 初始化数据库
	databases.Init()

	// 初始化消息队列
	message_queues.Init()

	// 初始化 Gin 路由
	r := gin.Default()
	
	routes.InitRoutes(r)

	// 启动服务
	addr := fmt.Sprintf("%s:%d", config.Cfg.WebHost, config.Cfg.WebPort)
	log.Info("服务启动中", zap.String("address", addr))
	if err := r.Run(addr); err != nil {
		panic(fmt.Sprintf("服务启动失败: %v", err))
	}
}
