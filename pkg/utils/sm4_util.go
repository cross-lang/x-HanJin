// Package utils provides utility functions for the x-HanJin framework.
package utils

//import (
//	"bytes"
//	"crypto/cipher"
//	"encoding/base64"
//	"encoding/hex"
//	"fmt"
//	"math/rand"
//	"time"
//	"x-HanJin/constants"
//
//	"github.com/tjfoc/gmsm/sm4"
//
//)
//
//type SM4Crypto struct {
//	blockSize int
//}
//
//func NewSM4Crypto() *SM4Crypto {
//	return &SM4Crypto{
//		blockSize: sm4.BlockSize, // SM4块大小为16字节
//	}
//}
//
//// Encrypt 加密函数
//func (s *SM4Crypto) Encrypt(plaintext, key, iv string, mode, padding, encoding string) (string, error) {
//	// 设置默认参数
//	if mode == "" {
//		mode = constants.ModeECB
//	}
//	if padding == "" {
//		padding = constants.PKCS7Padding
//	}
//	if encoding == "" {
//		encoding = constants.EncodingBase64
//	}
//
//	// 验证参数
//	if mode != constants.ModeECB && mode != constants.ModeCBC {
//		return "", fmt.Errorf("<<<<<<<< unsupported mode:%s, only ECB and CBC are supported", mode)
//	}
//
//	if padding != constants.PKCS7Padding && padding != constants.ISO10126Padding && padding != constants.NoPadding {
//		return "", fmt.Errorf("<<<<<<<< unsupported padding:%s, supported paddings are PKCS7, ISO10126, constants.NoPadding", padding)
//	}
//
//	if encoding != constants.EncodingBase64 && encoding != constants.EncodingHex {
//		return "", fmt.Errorf("<<<<<<<< unsupported encoding:%s, only base64 and hex are supported", encoding)
//	}
//
//	// 验证密钥长度
//	if len(key) != 16 {
//		return "", fmt.Errorf("<<<<<<<< key length must be 16 bytes, current length:%d", len(key))
//	}
//
//	// 验证IV长度（仅CBC模式需要）
//	if mode == constants.ModeCBC && len(iv) != s.blockSize {
//		return "", fmt.Errorf("<<<<<<<< IV length must be %d bytes for CBC mode, current length:%d", s.blockSize, len(iv))
//	}
//
//	// 创建加密块
//	block, err := sm4.NewCipher([]byte(key))
//	if err != nil {
//		return "", err
//	}
//
//	// 处理明文填充
//	plaintextBytes := []byte(plaintext)
//	if padding != constants.NoPadding {
//		plaintextBytes, err = s.applyPadding(plaintextBytes, padding)
//		if err != nil {
//			return "", err
//		}
//	}
//
//	// 验证填充后的长度
//	if len(plaintextBytes)%s.blockSize != 0 {
//		return "", fmt.Errorf("<<<<<<<< plaintext length must be a multiple of block size (%d bytes) after padding", s.blockSize)
//	}
//
//	// 加密
//	ciphertext := make([]byte, len(plaintextBytes))
//
//	switch mode {
//	case constants.ModeECB:
//		// 修复：将 s.blockSize() 改为 s.blockSize
//		for bs, be := 0, s.blockSize; bs < len(plaintextBytes); bs, be = bs+s.blockSize, be+s.blockSize {
//			block.Encrypt(ciphertext[bs:be], plaintextBytes[bs:be])
//		}
//	case constants.ModeCBC:
//		modeCBC := cipher.NewCBCEncrypter(block, []byte(iv))
//		modeCBC.CryptBlocks(ciphertext, plaintextBytes)
//	}
//
//	// 编码
//	var encoded string
//	switch encoding {
//	case constants.EncodingBase64:
//		encoded = base64.StdEncoding.EncodeToString(ciphertext)
//	case constants.EncodingHex:
//		encoded = hex.EncodeToString(ciphertext)
//	}
//
//	log.Infof(">>>>>>>> plaintext:%s", plaintext)
//	log.Infof(">>>>>>>> encrypted:%s", encoded)
//	return encoded, nil
//}
//
//// Decrypt 解密函数
//func (s *SM4Crypto) Decrypt(ciphertext, key, iv string, mode, padding string) (string, error) {
//	// 设置默认参数
//	if mode == "" {
//		mode = constants.ModeECB
//	}
//	if padding == "" {
//		padding = constants.PKCS7Padding
//	}
//
//	// 验证参数
//	if mode != constants.ModeECB && mode != constants.ModeCBC {
//		return "", fmt.Errorf("<<<<<<<< unsupported mode:%s, only ECB and CBC are supported", mode)
//	}
//
//	if padding != constants.PKCS7Padding && padding != constants.ISO10126Padding && padding != constants.NoPadding {
//		return "", fmt.Errorf("<<<<<<<< unsupported padding:%s, supported paddings are PKCS7, ISO10126, constants.NoPadding", padding)
//	}
//
//	// 验证密钥长度
//	if len(key) != 16 {
//		return "", fmt.Errorf("<<<<<<<< key length must be 16 bytes, current length:%d", len(key))
//	}
//
//	// 验证IV长度（仅CBC模式需要）
//	if mode == constants.ModeCBC && len(iv) != s.blockSize {
//		return "", fmt.Errorf("<<<<<<<< IV length must be %d bytes for CBC mode, current length:%d", s.blockSize, len(iv))
//	}
//
//	// 解码
//	var cipherBytes []byte
//	var err error
//
//	// 尝试Base64解码
//	cipherBytes, err = base64.StdEncoding.DecodeString(ciphertext)
//	if err != nil {
//		// 尝试Hex解码
//		cipherBytes, err = hex.DecodeString(ciphertext)
//		if err != nil {
//			return "", fmt.Errorf("<<<<<<<< failed to decode ciphertext:%v", err)
//		}
//	}
//
//	// 验证密文长度
//	if len(cipherBytes)%s.blockSize != 0 {
//		return "", fmt.Errorf("<<<<<<<< ciphertext length must be a multiple of block size (%d bytes), current length:%d", s.blockSize, len(cipherBytes))
//	}
//
//	// 创建解密块
//	block, err := sm4.NewCipher([]byte(key))
//	if err != nil {
//		return "", err
//	}
//
//	// 解密
//	plaintext := make([]byte, len(cipherBytes))
//
//	switch mode {
//	case constants.ModeECB:
//		// 修复：将 s.blockSize() 改为 s.blockSize
//		for bs, be := 0, s.blockSize; bs < len(cipherBytes); bs, be = bs+s.blockSize, be+s.blockSize {
//			block.Decrypt(plaintext[bs:be], cipherBytes[bs:be])
//		}
//	case constants.ModeCBC:
//		modeCBC := cipher.NewCBCDecrypter(block, []byte(iv))
//		modeCBC.CryptBlocks(plaintext, cipherBytes)
//	}
//
//	// 去除填充
//	if padding != constants.NoPadding {
//		plaintext, err = s.removePadding(plaintext, padding)
//		if err != nil {
//			return "", err
//		}
//	}
//
//	log.Infof(">>>>>>>> ciphertext:%s", ciphertext)
//	log.Infof(">>>>>>>> decrypted:%s", string(plaintext))
//	return string(plaintext), nil
//}
//
//// applyPadding 应用填充
//func (s *SM4Crypto) applyPadding(data []byte, padding string) ([]byte, error) {
//	blockSize := s.blockSize
//	paddingNeeded := blockSize - (len(data) % blockSize)
//
//	if paddingNeeded == 0 && padding != constants.NoPadding {
//		paddingNeeded = blockSize // 整个块都是填充
//	}
//
//	switch padding {
//	case constants.PKCS7Padding:
//		paddingBytes := bytes.Repeat([]byte{byte(paddingNeeded)}, paddingNeeded)
//		return append(data, paddingBytes...), nil
//
//	case constants.ISO10126Padding:
//		// 最后一个字节是填充长度，其余用随机字节填充
//		paddingBytes := make([]byte, paddingNeeded)
//		paddingBytes[paddingNeeded-1] = byte(paddingNeeded)
//
//		// 使用当前时间作为随机数种子
//		rand.Seed(time.Now().UnixNano())
//		for i := 0; i < paddingNeeded-1; i++ {
//			paddingBytes[i] = byte(rand.Intn(256))
//		}
//
//		return append(data, paddingBytes...), nil
//
//	case constants.NoPadding:
//		if len(data)%blockSize != 0 {
//			return nil, fmt.Errorf("<<<<<<<< data length must be a multiple of block size when using constants.NoPadding")
//		}
//		return data, nil
//
//	default:
//		return nil, fmt.Errorf("<<<<<<<< unsupported padding type:%s", padding)
//	}
//}
//
//// removePadding 移除填充
//func (s *SM4Crypto) removePadding(data []byte, padding string) ([]byte, error) {
//	if len(data) == 0 {
//		return data, nil
//	}
//
//	switch padding {
//	case constants.PKCS7Padding:
//		paddingLen := int(data[len(data)-1])
//		if paddingLen <= 0 || paddingLen > s.blockSize {
//			return nil, fmt.Errorf("<<<<<<<< invalid PKCS#7 padding length")
//		}
//
//		for i := len(data) - paddingLen; i < len(data); i++ {
//			if data[i] != byte(paddingLen) {
//				return nil, fmt.Errorf("<<<<<<<< invalid PKCS#7 padding bytes")
//			}
//		}
//
//		return data[:len(data)-paddingLen], nil
//
//	case constants.ISO10126Padding:
//		paddingLen := int(data[len(data)-1])
//		if paddingLen <= 0 || paddingLen > s.blockSize {
//			return nil, fmt.Errorf("<<<<<<<< invalid ISO 10126 padding length")
//		}
//
//		return data[:len(data)-paddingLen], nil
//
//	case constants.NoPadding:
//		return data, nil
//
//	default:
//		return nil, fmt.Errorf("<<<<<<<< unsupported padding type:%s", padding)
//	}
//}
