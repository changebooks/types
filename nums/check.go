package nums

// 负数？
func IsNegative(f float64) bool {
	return f < 0
}

// 正数？
func IsPositive(f float64) bool {
	return f > 0
}

// 非负数？
func IsNonNegative(f float64) bool {
	return f >= 0
}

// 非正数？
func IsNonPositive(f float64) bool {
	return f <= 0
}

// 16进制的字母？
func IsBase16(r rune) bool {
	return ('a' <= r && r <= 'f') || ('A' <= r && r <= 'F')
}

// 10进制？
func IsBase10(r rune) bool {
	return '0' <= r && r <= '9'
}
