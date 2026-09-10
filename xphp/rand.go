package xphp

import "crypto/rand"

// RandomBytes 生成加密安全的伪随机字节切片。
func RandomBytes(length int) []byte {
	b := make([]byte, length)
	rand.Read(b)
	return b
}
