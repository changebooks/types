package sparse

// 合并
func Merge(list ...map[int]int) map[int]int {
	if list == nil {
		return nil
	}

	r := make(map[int]int)
	for _, m := range list {
		if m != nil {
			for k, v := range m {
				r[k] = v
			}
		}
	}

	return r
}
