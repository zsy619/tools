package xstring

import "hash/crc32"

// ToHashcode 函数将一个字符串s转换为一个哈希码（整数）
// 参数s是需要转换的字符串
// 返回值是一个整数，表示字符串s的哈希码
func ToHashcode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
