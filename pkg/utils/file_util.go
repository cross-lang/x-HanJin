// Package utils provides utility functions for the x-HanJin framework.
package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"x-HanJin/pkg/log"

	"go.uber.org/zap"
)

// ReadFileByLine 逐行读取指定文件的内容，并将每行内容存储在一个字符串切片中返回。
// filePath: 要读取的文件的路径。
// 返回值: 包含文件每行内容的字符串切片和可能出现的错误。
func ReadFileByLine(filePath string) ([]string, error) {
	// 以只读模式打开指定路径的文件
	file, err := os.Open(filePath)
	if err != nil {
		// 若打开文件失败，直接返回错误
		return nil, fmt.Errorf("<<<<<<<< Failed to open file %s:%w", filePath, err)
	}
	// 确保在函数结束时关闭文件，避免资源泄漏
	defer func() {
		err := file.Close()
		if err != nil {
			log.Error("<<<<<<<< Error closing file", zap.String("filePath", filePath), zap.Error(err))
		}
	}()

	// 创建一个字符串切片，用于存储文件的每一行内容
	var lines []string
	// 创建一个扫描器，用于逐行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 将扫描到的当前行内容添加到切片中
		lines = append(lines, scanner.Text())
	}

	// 检查扫描过程中是否出现错误
	if err := scanner.Err(); err != nil {
		log.Error("<<<<<<<< Error scanning file", zap.String("filePath", filePath), zap.Error(err))
		return nil, err
	}

	return lines, nil
}

// CreateDir 创建指定路径的目录。如果目录已经存在，则不进行任何操作；
// 如果目录不存在，则递归创建该目录。
// dir: 要创建的目录的路径。
// 返回值: 可能出现的错误。
func CreateDir(dir string) error {
	// 检查指定目录是否存在
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// 若目录不存在，则使用 MkdirAll 递归创建该目录
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			// 若创建目录失败，返回错误信息
			return fmt.Errorf("<<<<<<<< Failed to create directory %s:%w", dir, err)
		}
		log.Info(">>>>>>>> Directory has been successfully created", zap.String("dir", dir))
	} else if err != nil {
		// 若检查目录时出现除目录不存在之外的其他错误，返回错误信息
		return fmt.Errorf("<<<<<<<< Failed to check directory %s:%w", dir, err)
	}

	return nil
}

// IsOfficeFile 检查文件是否为办公文件类型
// fileName: 文件名
// 返回值: 如果是办公文件类型返回true，否则返回false
func IsOfficeFile(fileName string) bool {
	// 定义需要检查的文件后缀
	extensions := []string{".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx"}

	// 将文件名转为小写，实现大小写不敏感的判断
	lowerFileName := strings.ToLower(fileName)

	// 检查是否以任何一个指定后缀结尾
	for _, ext := range extensions {
		if strings.HasSuffix(lowerFileName, ext) {
			return true
		}
	}
	return false
}
