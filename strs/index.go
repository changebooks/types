package strs

import "strings"

// 查找子串，忽略大小写
func CaseIndex(s string, substr string) int {
	return strings.Index(strings.ToLower(s), strings.ToLower(substr))
}

// 查找子串，忽略大小写
func LastCaseIndex(s string, substr string) int {
	return strings.LastIndex(strings.ToLower(s), strings.ToLower(substr))
}
