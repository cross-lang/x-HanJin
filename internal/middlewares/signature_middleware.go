package middlewares

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"x-HanJin/internal/config"
)

const (
	Ver                = "TestGin-1"              // 签名算法版本
	HeaderAuthorization = "TestGin-Authorization"  // 签名头名称
	HeaderDate         = "TestGin-Date"            // 时间戳头名称
)

// sha256Hash 计算数据的 SHA256 哈希值，返回十六进制编码字符串
func sha256Hash(body []byte) string {
	hasher := sha256.New()
	hasher.Write(body)
	return hex.EncodeToString(hasher.Sum(nil))
}

// generateSignature 生成 HMAC-SHA256 签名。
// 将版本号、请求方法、URI、Content-Type、时间戳和请求体哈希拼接后，
// 使用 AppKey 进行 HMAC-SHA256 加密。
func generateSignature(method, uri, contentType, timestamp, bodyHash string) string {
	signStr := Ver + method + uri + contentType + timestamp + bodyHash
	mac := hmac.New(sha256.New, []byte(config.Cfg.AppKey))
	mac.Write([]byte(signStr))
	return hex.EncodeToString(mac.Sum(nil))
}

// SignatureMiddleware 签名验证中间件。
// 校验请求头中的 Authorization 字段，确保请求来自合法客户端。
// Swagger 文档路径跳过签名验证。
func SignatureMiddleware(c *gin.Context) {
	// Swagger 路径跳过签名验证
	path := c.Request.URL.Path
	if len(path) >= 9 && path[:9] == "/swagger/" {
		c.Next()
		return
	}

	method := c.Request.Method
	uri := c.Request.URL.Path
	contentType := c.GetHeader("Content-Type")
	date := c.GetHeader(HeaderDate)
	authorization := c.GetHeader(HeaderAuthorization)

	// 验证必要的 header 字段
	if contentType == "" || date == "" || authorization == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required headers"})
		c.Abort()
		return
	}

	// 计算请求体的 SHA256 哈希值
	var bodyHash string
	if method == "POST" || method == "PUT" {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read request body"})
			c.Abort()
			return
		}
		bodyHash = sha256Hash(bodyBytes)

		// 将读取的 Body 重新写回，供后续 handler 使用
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// 生成服务端签名并验证
	serverSignature := generateSignature(method, uri, contentType, date, bodyHash)
	expectedAuth := fmt.Sprintf("%s %s:%s", Ver, config.Cfg.AppId, serverSignature)

	// 使用 hmac.Equal 进行时间安全的比较，防止时序攻击
	if !hmac.Equal([]byte(authorization), []byte(expectedAuth)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid signature"})
		c.Abort()
		return
	}

	c.Next()
}
