// Package utils provides utility functions for the x-HanJin framework.
package utils

// IfInt 根据给定的布尔条件返回两个整数中的一个。
// condition: 用于判断的布尔条件。
// trueVal: 当条件为 true 时返回的整数值。
// falseVal: 当条件为 false 时返回的整数值。
// 返回值: 根据条件返回 trueVal 或 falseVal。
func IfInt(condition bool, trueVal, falseVal int) int {
	if condition {
		return trueVal
	}
	return falseVal
}

// IfInt64 根据给定的布尔条件返回两个 64 位整数中的一个。
// condition: 用于判断的布尔条件。
// trueVal: 当条件为 true 时返回的 64 位整数值。
// falseVal: 当条件为 false 时返回的 64 位整数值。
// 返回值: 根据条件返回 trueVal 或 falseVal。
func IfInt64(condition bool, trueVal, falseVal int64) int64 {
	if condition {
		return trueVal
	}
	return falseVal
}
