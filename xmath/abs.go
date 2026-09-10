package xmath

// AbsInt8 获取 int8 的绝对值。
//
// 算法说明：通过算术右移取得符号位，再利用 XOR 与减法组合规避分支判断。
//
// 示例 1：AbsInt8(-5)
// -5 在内存中的表示如下：
// 原码		1000,0101
// 反码		1111,1010
// 补码		1111,1011
// 负数在内存中以补码表示。
// shifted = n >> 7 = (1111,1011) >> 7 = 1111,1111 = -1（十进制）（负数右移，左补 1）
//
//	1111,1011
//
// n xor shifted =  ----------- = 0000,0100 = 4（十进制）
//
//	1111,1111
//
// (n ^ shifted) - shifted = 4 - (-1) = 5
//
// 示例 2：AbsInt8(5)
// 5 在内存中的表示如下：
// 原码 0000,0101
// 正数在内存中以原码表示，正数与 0 做 XOR 等于其自身。
// shifted = n >> 7 = 0
//
//	0000,0101
//
// n xor shifted =  ----------- = 0000,0101 = 5（十进制）
//
//	0000,0000
//
// (n ^ shifted) - shifted = 5 - 0 = 5
//
// 参数：
//   - n: 待求绝对值的 int8 数值。
//
// 返回值：n 的绝对值。注意：math.MinInt8 没有正数对应值，结果仍为 math.MinInt8。
func AbsInt8(n int8) int8 {
	shifted := n >> 7
	return (n ^ shifted) - shifted
}

// AbsInt16 获取 int16 的绝对值（参考 AbsInt8 的位运算思路）。
//
// 参数：
//   - n: 待求绝对值的 int16 数值。
//
// 返回值：n 的绝对值。
func AbsInt16(n int16) int16 {
	shifted := n >> 15
	return (n ^ shifted) - shifted
}

// AbsInt32 获取 int32 的绝对值（参考 AbsInt8 的位运算思路）。
//
// 参数：
//   - n: 待求绝对值的 int32 数值。
//
// 返回值：n 的绝对值。
func AbsInt32(n int32) int32 {
	shifted := n >> 31
	return (n ^ shifted) - shifted
}

// AbsInt64 获取 int64 的绝对值（参考 AbsInt8 的位运算思路）。
//
// 参数：
//   - n: 待求绝对值的 int64 数值。
//
// 返回值：n 的绝对值。
func AbsInt64(n int64) int64 {
	shifted := n >> 63
	return (n ^ shifted) - shifted
}
