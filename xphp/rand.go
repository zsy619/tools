package xphp

import "crypto/rand"

// RandomBytes generates cryptographically secure pseudo-random bytes
func RandomBytes(length int) []byte {
	b := make([]byte, length)
	rand.Read(b)
	return b
}
