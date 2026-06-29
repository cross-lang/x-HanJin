// Package utils provides utility functions for the x-HanJin framework.
package utils

import (
	"fmt"

	"x-HanJin/pkg/log"

	"go.uber.org/zap"
)

// GetUserInput 提示用户输入一个值并返回
func GetUserInput(prompt string) string {
	var input string

	log.Info("❓", zap.String("prompt", FGreen(prompt)))
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Info("读取用户输入时出错", zap.Error(err))
		return ""
	}
	return input
}

// ConfirmUserChoice 提示用户进行是/否选择并返回结果
func ConfirmUserChoice(prompt string) bool {
	input := GetUserInput(prompt + " [y/n]:")
	return input == "" || input == "y"
}

// SelectFromOptions 提示用户从一组选项中选择一个并返回所选选项
func SelectFromOptions(prompt string, options []string) string {
	for i, option := range options {
		log.Info("option", zap.Int("index", i+1), zap.String("value", option))
	}
	input := GetUserInput(prompt)
	var index int
	_, err := fmt.Sscanf(input, "%d", &index)
	if err != nil || index > len(options) || index < 1 {
		log.Info("无效的选择", zap.String("input", input))
		return ""
	}
	return options[index-1]
}
