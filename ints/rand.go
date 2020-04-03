package ints

import (
	"math/rand"
	"time"
)

// [min, max)
func Rand(min int, max int) int {
	if min < 0 {
		min = 0
	}

	if max <= 0 {
		panic("invalid argument")
	}

	if min >= max {
		return max
	}

	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max-min) + min
}
