package sparse

// 合并
func Merge(list ...map[int]string) map[int]string {
	if list == nil {
		return nil
	}

	r := make(map[int]string)
	for _, m := range list {
		if m != nil {
			for k, v := range m {
				r[k] = v
			}
		}
	}

	return r
}
