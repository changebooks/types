package dict

// 合并
func Merge(list ...map[string]interface{}) map[string]interface{} {
	if list == nil {
		return nil
	}

	r := make(map[string]interface{})
	for _, m := range list {
		if m != nil {
			for k, v := range m {
				r[k] = v
			}
		}
	}

	return r
}
