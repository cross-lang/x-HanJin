// Package middlewares 提供 HTTP 中间件，
// 包括异常恢复、签名验证等全局中间件功能。
package middlewares

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"x-HanJin/pkg/log"
	"go.uber.org/zap"
)

// RecoveryMiddleware 异常恢复中间件。
// 捕获 handler 中的 panic，记录错误日志并返回 500 响应，
// 防止单个请求的异常导致整个服务崩溃。
func RecoveryMiddleware(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 记录 panic 详细信息和堆栈
			log.Error("请求处理发生 panic",
				zap.Any("error", err),
				zap.String("path", c.Request.URL.Path),
				zap.String("method", c.Request.Method),
				zap.String("stack", string(debug.Stack())),
			)
			fmt.Printf("panic recovered: %v\n%s\n", err, debug.Stack())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "内部服务器错误"})
			c.Abort()
		}
	}()

	c.Next()
}
