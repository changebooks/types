package alphabet

const Letter = Upper + Lower

// 英文串序号总和
func Sum(s string) uint {
	var r uint = 0
	for _, c := range s {
		r += Num(c)
	}

	return r
}

// 字母表中序号
func Num(r rune) uint {
	if i, ok := LowerNum[r]; ok {
		return i
	}

	if i, ok := UpperNum[r]; ok {
		return i
	}

	return 0
}
