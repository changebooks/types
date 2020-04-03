package dict

// 合并
func Merge(list ...map[string]uint64) map[string]uint64 {
	if list == nil {
		return nil
	}

	r := make(map[string]uint64)
	for _, m := range list {
		if m != nil {
			for k, v := range m {
				r[k] = v
			}
		}
	}

	return r
}
