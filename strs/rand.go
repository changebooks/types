package strs

import (
	"bytes"
	"math/rand"
	"time"
)

// 随机size长度的子串
func Rand(s string, size int) string {
	if s == "" || size <= 0 {
		return ""
	}

	var bucket bytes.Buffer
	num := len(s)
	for i := 0; i < size; i++ {
		bucket.WriteByte(s[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(num)])
	}

	return bucket.String()
}
