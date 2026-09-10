package maths

import (
	"math"
	"math/bits"

	"github.com/zsy619/tools/xgeneric"
)

// 把数n分成m份
// 返回每一份大小和最后一份的大小
func Split(n, m uint) (uint, uint) {
	per := (n + m - 1) / m
	last := n % per
	if last == 0 {
		last = per
	}
	return per, last
}

// 最大值
func Max[T xgeneric.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 最小值
func Min[T xgeneric.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// 绝对值
func Abs[T xgeneric.Float | xgeneric.Signed](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

// Log2
// 修改更加高效算法
func Log2[T xgeneric.Integer | xgeneric.Float](a T) T {
	return T(math.Log2(float64(a)))
}

// IsPowOf2 判断一个值是否为 2 的幂。
// 零不被视为 2 的幂。
// 参数：n 为待检查的无符号整数。
// 返回：当 n != 0 且 n & (n-1) == 0 时返回 true，否则返回 false。
func IsPowOf2[T xgeneric.Unsigned](n T) bool {
	return n != 0 && ((n & (n - 1)) == 0)
}

// RoundUpPowOf2 将 n 向上取整到最近的 2 的幂。
// 参数：n 为无符号整数。
// 返回：不小于 n 的最小 2 的幂。
func RoundUpPowOf2[T xgeneric.Unsigned](n T) T {
	return 1 << FindLastBitSet(n-1)
}

// RoundDownPowOf2 将 n 向下取整到最近的 2 的幂。
// 参数：n 为无符号整数。
// 返回：不大于 n 的最大 2 的幂。
func RoundDownPowOf2[T xgeneric.Unsigned](n T) T {
	return 1 << (FindLastBitSet(n) - 1)
}

// FindLastBitSet 查找最高位（最高有效位）的索引。
// 定义与 Unix 的 ffs 函数一致。
// 注意：FindLastBitSet(0) = 0，FindLastBitSet(1) = 1，FindLastBitSet(0x80000000) = 32。
// 参数：x 为无符号整数。
// 返回：最高位 1 所在的位数（从 1 开始），x 为 0 时返回 0。
func FindLastBitSet[T xgeneric.Unsigned](x T) int {
	return bits.Len64(uint64(x))
}

// Range 创建一个包含指定长度数字（可正可负）的切片。
// 示例：https://go.dev/play/p/0r6VimXAi9H
// 参数：elementNum 为期望的元素数量；为负时生成从 0 到 (elementNum+1) 的递减序列。
// 返回：长度为 |elementNum| 的 []int。
func Range(elementNum int) []int {
	length := xgeneric.If(elementNum < 0, -elementNum).Else(elementNum)
	result := make([]int, length)
	step := xgeneric.If(elementNum < 0, -1).Else(1)
	for i, j := 0, 0; i < length; i, j = i+1, j+step {
		result[i] = j
	}
	return result
}

// RangeFrom 创建一个从 start 开始、长度为 elementNum 的等差数列切片。
// 示例：https://go.dev/play/p/0r6VimXAi9H
// 参数：start 为起始值，elementNum 为期望的元素数量；为负时生成递减序列。
// 返回：长度为 |elementNum| 的 []T，元素从 start 开始逐次加 1 或减 1。
func RangeFrom[T xgeneric.Integer | xgeneric.Float](start T, elementNum int) []T {
	length := xgeneric.If(elementNum < 0, -elementNum).Else(elementNum)
	result := make([]T, length)
	step := xgeneric.If(elementNum < 0, -1).Else(1)
	for i, j := 0, start; i < length; i, j = i+1, j+T(step) {
		result[i] = j
	}
	return result
}

// RangeWithSteps 创建一个从 start 开始、到 end 结束（不含 end）、步长为 step 的等差数列切片。
// 当 step == 0 时返回空切片；start < end 时要求 step > 0，反之要求 step < 0，否则返回空切片。
// 示例：https://go.dev/play/p/0r6VimXAi9H
// 参数：start 为起始值，end 为终止值（不含），step 为步长。
// 返回：包含 start、start+step、... 但不含 end 的 []T。
func RangeWithSteps[T xgeneric.Integer | xgeneric.Float](start, end, step T) []T {
	result := []T{}
	if start == end || step == 0 {
		return result
	}
	if start < end {
		if step < 0 {
			return result
		}
		for i := start; i < end; i += step {
			result = append(result, i)
		}
		return result
	}
	if step > 0 {
		return result
	}
	for i := start; i > end; i += step {
		result = append(result, i)
	}
	return result
}

// Clamp 将 value 限制在 [min, max] 闭区间内。
// 示例：https://go.dev/play/p/RU4lJNC2hlI
// 参数：value 为输入值，min 为下界，max 为上界。
// 返回：当 value 小于 min 时返回 min；大于 max 时返回 max；否则返回 value。
func Clamp[T xgeneric.Ordered](value T, min T, max T) T {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

// Sum 对切片中的所有元素求和。
// 切片为空时返回 0。
// 示例：https://go.dev/play/p/upfeJVqs4Bt
// 参数：collection 为输入切片（元素必须满足 Float/Integer/Complex 约束）。
// 返回：所有元素的累加和。
func Sum[T xgeneric.Float | xgeneric.Integer | xgeneric.Complex](collection []T) T {
	var sum T = 0
	for _, val := range collection {
		sum += val
	}
	return sum
}

// SumBy 使用 iteratee 函数对切片中的每个元素进行映射后再求和。
// 切片为空时返回 0。
// 示例：https://go.dev/play/p/Dz_a_7jN_ca
// 参数：collection 为输入切片，iteratee 为映射函数（输入 T 返回 R）。
// 返回：所有 iteratee(item) 累加得到的 R 类型和。
func SumBy[T any, R xgeneric.Float | xgeneric.Integer | xgeneric.Complex](collection []T, iteratee func(item T) R) R {
	var sum R = 0
	for _, item := range collection {
		sum = sum + iteratee(item)
	}
	return sum
}
