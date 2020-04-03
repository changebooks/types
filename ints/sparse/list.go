package sparse

// 键列表，无序
func Keys(m map[int]int) []int {
	if len(m) == 0 {
		return nil
	}

	r := make([]int, len(m))
	i := 0
	for k := range m {
		r[i] = k
		i++
	}

	return r
}

// 值列表，无序
func Values(m map[int]int) []int {
	if len(m) == 0 {
		return nil
	}

	r := make([]int, len(m))
	i := 0
	for _, v := range m {
		r[i] = v
		i++
	}

	return r
}
