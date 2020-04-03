package sparse

// 合并
func Merge(list ...map[int]float64) map[int]float64 {
	if list == nil {
		return nil
	}

	r := make(map[int]float64)
	for _, m := range list {
		if m != nil {
			for k, v := range m {
				r[k] = v
			}
		}
	}

	return r
}
