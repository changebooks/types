package nums

// 格式：数字、浮点、负数、16进制、科学计算法
func Format(s string) (isNumeric bool, isFloat bool, isNegative bool, isBase16 bool, isScientific bool) {
	// Last index
	last := len(s) - 1
	if last < 0 {
		// Empty
		return
	}

	//	Current index, Used for comparison and autoincrement
	i := 0

	// Skip [-|+] signed character
	if s[i] == '-' || s[i] == '+' {
		if last == 0 {
			// Just "+"、"-"
			return
		}

		if s[i] == '-' {
			isNegative = true
		}

		i++
	}

	// First char is 0, Maybe base16
	if s[i] == '0' {
		if last == i {
			// last == 0 || last == 1
			// i == 0 || i == 1
			// Maybe "0"、"-0"、"+0"
			isNumeric = true
			return
		}

		i++
		// base 16
		if s[i] == 'x' || s[i] == 'X' {
			isBase16 = true

			if last == i {
				// last == 1 || last == 2
				// i == 1 || i == 2
				// Just "0x"、"0X"、"-0x"、"-0X"、"+0x"、"+0X"，End of 'x'\'X'
				isNumeric = false
				return
			}

			i++
			for _, c := range s[i:] {
				if IsBase10(c) {
					// '0' <= c <= '9'
					continue
				}

				if IsBase16(c) {
					// ('a' <= c <= 'f') || ('A' <= c <= 'F')
					continue
				}

				// Not just ('0' <= c <= '9') || ('a' <= c <= 'f') || ('A' <= c <= 'F')
				isNumeric = false
				return
			}

			isNumeric = true
			return
		}
	} else if s[i] == 'e' || s[i] == 'E' {
		// 'e'\'E' is first character
		// Maybe "e***"、"E***"、"-e***"、"-E***"、"+e***"、"+E***"
		return
	}

	// base 10

	// '.' index
	p := -1
	// 'e'\'E' index
	e := -1
	for i, c := range s[i:] {
		if IsBase10(c) {
			// '0' <= c <= '9'
			continue
		}

		if c == '.' {
			if isFloat {
				// Too many '.'
				isNumeric = false
				return
			}

			p = i
			isFloat = true
			continue
		}

		if c == 'e' || c == 'E' {
			if isScientific {
				// Too many 'e'\'E'
				isNumeric = false
				return
			}

			e = i
			isScientific = true
			continue
		}

		// Not just ('0' <= c <= '9') || '.'
		isNumeric = false
		return
	}

	if isScientific {
		if s[last] == 'e' || s[last] == 'E' {
			// 'e'\'E' is last character
			isNumeric = false
			return
		}
	}

	if isFloat {
		if s[last] == '.' {
			// '.' is last character
			isNumeric = false
			return
		}
	}

	if isScientific && isFloat {
		if p >= e {
			// '.' is behind 'e'\'E'
			isNumeric = false
			return
		}
	}

	isNumeric = true
	return
}
