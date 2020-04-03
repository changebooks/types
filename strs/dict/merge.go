package dict

// 合并
func Merge(list ...map[string]string) map[string]string {
	if list == nil {
		return nil
	}

	r := make(map[string]string)
	for _, m := range list {
		if m != nil {
			for k, v := range m {
				r[k] = v
			}
		}
	}

	return r
}
