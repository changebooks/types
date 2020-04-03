package sparse

// 合并
func Merge(list ...map[int]uint64) map[int]uint64 {
	if list == nil {
		return nil
	}

	r := make(map[int]uint64)
	for _, m := range list {
		if m != nil {
			for k, v := range m {
				r[k] = v
			}
		}
	}

	return r
}
