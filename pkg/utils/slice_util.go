// Package utils provides utility functions for the x-HanJin framework.
package utils

import (
	"strconv"
	"strings"
)

// IsSlicesEqual 判断两个由字符串组成的切片是否相等。
// 不考虑元素顺序，只要元素及其数量相同则认为相等。
// a: 第一个字符串切片
// b: 第二个字符串切片
// 返回值: 如果两个切片相等则返回 true，否则返回 false
func IsSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	countA := make(map[string]int)
	countB := make(map[string]int)

	for _, v := range a {
		countA[v]++
	}

	for _, v := range b {
		countB[v]++
	}

	for k, v := range countA {
		if countB[k] != v {
			return false
		}
	}

	return true
}

// Int64SliceToString 将 int64 切片转换为字符串，元素之间用指定分隔符分隔。
// 示例: [1,2,3] => "1,2,3"
// ids: 要转换的 int64 切片
// sep: 分隔符
// 返回值: 转换后的字符串
func Int64SliceToString(ids []int64, sep string) string {
	if len(ids) == 0 {
		return ""
	}
	str := ""
	for _, id := range ids {
		str = str + strconv.FormatInt(id, 10) + sep
	}
	return str[0 : len(str)-len(sep)]
}

// IdsStringToInt64Slice 将由分隔符分隔的 id 字符串转换为 int64 切片。
// 示例："1,2,3" => [1,2,3]
// ids: 要转换的字符串
// sep: 分隔符
// 返回值: 转换后的 int64 切片和可能出现的错误
func IdsStringToInt64Slice(ids string, sep string) ([]int64, error) {
	iIds := make([]int64, 0)
	split := strings.Split(ids, sep)
	for _, item := range split {
		iItem := strings.TrimSpace(item)
		id, err := strconv.ParseInt(iItem, 10, 64)
		if err != nil {
			return nil, err
		}
		iIds = append(iIds, id)
	}
	return iIds, nil
}

// StrSliceToString 将字符串切片转换为字符串，元素之间用指定分隔符分隔。
// 示例:["1","2","3"] => "1,2,3"
// ids: 要转换的字符串切片
// sep: 分隔符
// 返回值: 转换后的字符串
func StrSliceToString(ids []string, sep string) string {
	if len(ids) == 0 {
		return ""
	}
	str := ""
	for _, id := range ids {
		str = str + id + sep
	}
	return str[0 : len(str)-len(sep)]
}

// StringSliceToString 将字符串切片转换为字符串，每个元素用单引号包裹，元素之间用指定分隔符分隔。
// 示例:["1","2","3"] => "'1','2','3'"
// ids: 要转换的字符串切片
// sep: 分隔符
// 返回值: 转换后的字符串
func StringSliceToString(ids []string, sep string) string {
	if len(ids) == 0 {
		return ""
	}
	str := ""
	for _, id := range ids {
		str = str + "'" + id + "'" + sep
	}
	return str[0 : len(str)-len(sep)]
}

// UniqueSliceInt64 对字符串切片进行去重操作。
// ss: 要去重的字符串切片
// 返回值: 去重后的字符串切片
func UniqueSliceInt64(ss []string) []string {
	newSS := make([]string, 0)  // 返回的新切片
	m1 := make(map[string]byte) // 用来去重的临时 map
	for _, v := range ss {
		if _, ok := m1[v]; !ok {
			m1[v] = 1
			newSS = append(newSS, v)
		}
	}
	return newSS
}

// Contains 检查切片中是否包含指定元素。
// slice: 要检查的字符串切片
// elem: 要查找的元素
// 返回值: 如果切片中包含该元素则返回 true，否则返回 false
func Contains(slice []string, elem string) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

// IfSlice 根据给定条件返回两个字节切片中的一个。
// condition: 用于判断的布尔条件
// trueVal: 当条件为 true 时返回的字节切片
// falseVal: 当条件为 false 时返回的字节切片
// 返回值: 根据条件返回 trueVal 或 falseVal
func IfSlice(condition bool, trueVal, falseVal []byte) []byte {
	if condition {
		return trueVal
	}
	return falseVal
}
