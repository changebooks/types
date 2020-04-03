package dict

// 键存在？
func KeyExist(key string, m map[string]int) bool {
	if len(m) == 0 {
		return false
	}

	_, ok := m[key]
	return ok
}

// 值存在？键 ：""
func ValueExist(value int, m map[string]int) (bool, string) {
	if len(m) == 0 {
		return false, ""
	}

	for k, v := range m {
		if value == v {
			return true, k
		}
	}

	return false, ""
}
