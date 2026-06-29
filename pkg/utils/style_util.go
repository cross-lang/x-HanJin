// Package utils provides utility functions for the x-HanJin framework.
package utils

import (
	"fmt"

	"x-HanJin/pkg/log"

	"go.uber.org/zap"
)

// FBlack 黑色
func FBlack(str string) string {
	return "\033[30m" + str + "\033[0m"
}

// FRed 红色
func FRed(str string) string {
	return "\033[31m" + str + "\033[0m"
}

// FGreen 绿色
func FGreen(str string) string {
	return "\033[32m" + str + "\033[0m"
}

// FYellow 黄色
func FYellow(str string) string {
	return "\033[33m" + str + "\033[0m"
}

// FBlue 蓝色
func FBlue(str string) string {
	return "\033[34m" + str + "\033[0m"
}

// FPurple 紫色
func FPurple(str string) string {
	return "\033[35m" + str + "\033[0m"
}

// FCyan 青色
func FCyan(str string) string {
	return "\033[36m" + str + "\033[0m"
}

// FWhite 白色
func FWhite(str string) string {
	return "\033[37m" + str + "\033[0m"
}

func ErrorLog(format string, v ...any) {
	log.Error("❌ "+format, zap.Any("values", v))
}

func AwaitLog(format string, v ...any) {
	log.Info("⏳ "+format, zap.Any("values", v))
}

func SuccessLog(format string, v ...any) {
	log.Info("✔️ "+format, zap.Any("values", v))
}

func GenErrorf(format string, v ...any) error {
	return fmt.Errorf("❌ "+format, v...)
}
