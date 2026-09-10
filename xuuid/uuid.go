package xuuid

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

// NewUUID 使用密码学安全随机数生成一个新的 UUID（形如 xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx）字符串。
func NewUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

// GetAuth 返回去掉连字符后的 UUID 的前 16 位，用于生成认证凭据。
func GetAuth() string {
	auth := NewUUID()
	auth = strings.Replace(auth, "-", "", -1)
	auth = auth[:16]
	return auth
}

// GetAes 返回去掉连字符后的完整 UUID（32 位十六进制字符），常用于 AES 等密钥场景。
func GetAes() string {
	aes := NewUUID()
	aes = strings.Replace(aes, "-", "", -1)
	return aes
}

// GenerateRandomBytes 返回安全生成的随机字节。
// 若系统的安全随机数生成器未能正常工作，则会返回错误，
// 此时调用方不应继续执行后续操作。
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// 注意：只有当读满 len(b) 字节时 err 才为 nil。
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomKey 生成指定长度的随机字节并对结果进行 base64 编码，返回字符串形式的密钥。
func GenerateRandomKey(keyLen int) (string, error) {
	key, err := GenerateRandomBytes(keyLen)
	if err != nil {
		return "", err
	}
	fmt.Println("GenerateRandomKey：", fmt.Sprintf("%x", key))
	return base64.StdEncoding.EncodeToString(key), nil
}

// GenerateAESSecret16Key 生成 16 字节（128 位）的 AES 密钥字符串。
func GenerateAESSecret16Key() (string, error) {
	return GenerateRandomKey(16)
}
