package nums

import "math"

// 取最大值
func Max(nums ...float64) float64 {
	size := len(nums)
	if size == 0 {
		panic("invalid argument")
	}

	max := nums[0]
	if size == 1 {
		return max
	}

	for i := 1; i < size; i++ {
		max = math.Max(max, nums[i])
	}

	return max
}

// 取最小值
func Min(nums ...float64) float64 {
	size := len(nums)
	if size == 0 {
		panic("invalid argument")
	}

	min := nums[0]
	if size == 1 {
		return min
	}

	for i := 1; i < size; i++ {
		min = math.Min(min, nums[i])
	}

	return min
}
