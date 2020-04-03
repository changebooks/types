package strs

import "unicode/utf8"

// 汉字 == 1
func MultiByteLen(s string) int {
	return utf8.RuneCountInString(s)
}

// 汉字 == 1
func MultiByteCut(s string, size int) string {
	if s == "" || size <= 0 {
		return s
	}

	rs := []rune(s)
	if size >= len(rs) {
		return ""
	}

	return string(rs[size:])
}
