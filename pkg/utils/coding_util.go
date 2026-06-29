// Package utils provides utility functions for the x-HanJin framework.
package utils

import (
	"encoding/base64"
	"encoding/hex"

	"x-HanJin/pkg/log"

	"go.uber.org/zap"
)

// Base64Encode 将字节切片编码为Base64字符串
// data: 要进行编码的字节切片
// 返回值: 编码后的Base64字符串和可能出现的错误
// 若编码失败，会使用 log 记录错误信息
func Base64Encode(data []byte) (string, error) {
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded, nil
}

// Base64Decode 将Base64字符串解码为字节切片
// str: 要进行解码的Base64字符串
// 返回值: 解码后的字节切片和可能出现的错误
// 若解码失败，会使用 log 记录错误信息，包括错误详情和传入的字符串
func Base64Decode(str string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Error("<<<<<<<< Failed to decode Base64 str", zap.Error(err), zap.String("str", str))
		return []byte(nil), err
	}
	return decoded, nil
}

// HexEncode 将字节切片编码为十六进制字符串
// data: 要进行编码的字节切片
// 返回值: 编码后的十六进制字符串和可能出现的错误
// 若编码失败，会使用 log 记录错误信息
func HexEncode(data []byte) (string, error) {
	encoded := hex.EncodeToString(data)
	return encoded, nil
}

// HexDecode 将十六进制字符串解码为字节切片
// str: 要进行解码的十六进制字符串
// 返回值: 解码后的字节切片和可能出现的错误
// 若解码失败，会使用 log 记录错误信息，包括错误详情和传入的字符串
func HexDecode(str string) ([]byte, error) {
	decoded, err := hex.DecodeString(str)
	if err != nil {
		log.Error("<<<<<<<< Failed to decode Hex str", zap.Error(err), zap.String("str", str))
		return []byte(nil), err
	}
	return decoded, nil
}
