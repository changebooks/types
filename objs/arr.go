package objs

// 找到？下标 : -1
func InArr(value interface{}, list []interface{}) int {
	if len(list) == 0 {
		return -1
	}

	for k, v := range list {
		if value == v {
			return k
		}
	}

	return -1
}

// 合并
func Merge(list ...[]interface{}) []interface{} {
	if list == nil {
		return nil
	}

	size := 0
	for _, v := range list {
		if v != nil {
			size += len(v)
		}
	}

	if size == 0 {
		return nil
	}

	r := make([]interface{}, 0, size)
	for _, v := range list {
		if v != nil {
			r = append(r, v...)
		}
	}

	return r
}

// 反转
func Reverse(list []interface{}) []interface{} {
	if len(list) == 0 {
		return nil
	}

	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

	return list
}
