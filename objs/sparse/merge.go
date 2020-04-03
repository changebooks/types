package sparse

// 合并
func Merge(list ...map[int]interface{}) map[int]interface{} {
	if list == nil {
		return nil
	}

	r := make(map[int]interface{})
	for _, m := range list {
		if m != nil {
			for k, v := range m {
				r[k] = v
			}
		}
	}

	return r
}
