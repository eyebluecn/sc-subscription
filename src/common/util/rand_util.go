package util

import (
	"math/rand"
)

// 随机获取true或者false.
func RandomBoolean() bool {
	return rand.Intn(2) == 1
}

// 随机获取长度为n的字符串
func RandomString(n int) string {
	if n <= 0 {
		return ""
	}
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := ""

	for i := 0; i < n; i++ {
		index := rand.Intn(len(characters))
		result += string(characters[index])
	}

	return result
}

// 随机生成[a,b]数字. 输入有误返回-1
func RandomInt(a int, b int) int {
	if a > b {
		return -1
	}
	return rand.Intn(b-a+1) + a
}
