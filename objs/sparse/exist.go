package sparse

// 键存在？
func KeyExist(key int, m map[int]interface{}) bool {
	if len(m) == 0 {
		return false
	}

	_, ok := m[key]
	return ok
}

// 值存在？键 : -1
func ValueExist(value interface{}, m map[int]interface{}) (bool, int) {
	if len(m) == 0 {
		return false, -1
	}

	for k, v := range m {
		if value == v {
			return true, k
		}
	}

	return false, -1
}
