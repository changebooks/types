package strs

import "unicode/utf8"

func Chr(ascii int) string {
	return string(ascii)
}

func Ord(c string) int {
	r, _ := utf8.DecodeRune([]byte(c))
	return int(r)
}
