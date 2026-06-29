// Package utils provides utility functions for the x-HanJin framework.
package utils

import (
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"time"

	"x-HanJin/pkg/log"

	"github.com/tjfoc/gmsm/sm3"

	"go.uber.org/zap"
)

// CalcHash 计算消息的SM3哈希值，支持指定编码类型
func CalcHash(message string, encoding string) string {
	hashBytes := sm3.Sm3Sum([]byte(message))

	switch encoding {
	case "hex":
		return hex.EncodeToString(hashBytes)
	case "base64":
		return base64.StdEncoding.EncodeToString(hashBytes)
	default:
		log.Error("unsupported encoding type, using hex as default", zap.String("encoding", encoding))
		return hex.EncodeToString(hashBytes)
	}
}

// CalcHashWithSalt 计算带盐值的SM3哈希，支持指定编码类型
func CalcHashWithSalt(message, salt string, encoding string) string {
	data := []byte(message + salt)
	hashBytes := sm3.Sm3Sum(data)

	switch encoding {
	case "hex":
		return hex.EncodeToString(hashBytes)
	case "base64":
		return base64.StdEncoding.EncodeToString(hashBytes)
	default:
		log.Error("unsupported encoding type, using hex as default", zap.String("encoding", encoding))
		return hex.EncodeToString(hashBytes)
	}
}

// SM3GenerateSalt 生成指定长度的随机盐值
func SM3GenerateSalt(length int) string {
	if length <= 0 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"

	salt := make([]byte, length)
	for i := range salt {
		salt[i] = charset[rand.Intn(len(charset))]
	}

	return string(salt)
}
