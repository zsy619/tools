package xcrypto

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5Hash 函数计算给定字符串的MD5哈希值，并返回其十六进制表示形式
//
// 参数：
// text string - 需要计算哈希值的字符串
//
// 返回值：
// string - 字符串的MD5哈希值的十六进制表示形式
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
