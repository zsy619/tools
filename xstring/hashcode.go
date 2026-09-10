package xstring

import "hash/crc32"

// ToHashcode 将字符串转换为稳定的 32 位哈希码。
//
// 内部使用 CRC-32 IEEE 算法，相同输入总会得到相同的非负整数。
//
// 返回值保证 >= 0，避免 crc32 转换成 int 时出现负数（含 MinInt 的
// 特殊情况已特殊处理）。
//
// 典型用途：
//   - 简易的字符串到桶的映射，例如对 URL 做 sharding；
//   - 在内存中作为 map 的 key（注意哈希冲突）。
//
// 注意：本函数不保证抗碰撞，如需密码学场景请使用 crypto/sha256 等
// 安全哈希。
func ToHashcode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == math.MinInt 时取反仍为 MinInt，需要兜底返回 0。
	return 0
}
