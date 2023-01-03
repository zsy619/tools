package maths

import (
	"math"
	"math/bits"

	"golang.org/x/exp/constraints"
	"haedu.gov.cn/tools/xgeneric"
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
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 最小值
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// 绝对值
func Abs[T constraints.Float | constraints.Signed](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

// Log2
// 修改更加高效算法
func Log2[T constraints.Integer | constraints.Float](a T) T {
	return T(math.Log2(float64(a)))
}

// IsPowOf2 check if a value is a power of two
// Determine whether some value is a power of two, where zero is
// not considered a power of two.
func IsPowOf2[T constraints.Unsigned](n T) bool {
	return n != 0 && ((n & (n - 1)) == 0)
}

// RoundUpPowOf2 round up to nearest power of two
func RoundUpPowOf2[T constraints.Unsigned](n T) T {
	return 1 << FindLastBitSet(n-1)
}

// RoundDownPowOf2 round down to nearest power of two
func RoundDownPowOf2[T constraints.Unsigned](n T) T {
	return 1 << (FindLastBitSet(n) - 1)
}

// FindLastBitSet find last (most-significant) bit set
// This is defined the same way as ffs.
// Note FindLastBitSet(0) = 0, FindLastBitSet(1) = 1, FindLastBitSet(0x80000000) = 32.
func FindLastBitSet[T constraints.Unsigned](x T) int {
	return bits.Len64(uint64(x))
}

// Range creates an array of numbers (positive and/or negative) with given length.
// Play: https://go.dev/play/p/0r6VimXAi9H
func Range(elementNum int) []int {
	length := xgeneric.If(elementNum < 0, -elementNum).Else(elementNum)
	result := make([]int, length)
	step := xgeneric.If(elementNum < 0, -1).Else(1)
	for i, j := 0, 0; i < length; i, j = i+1, j+step {
		result[i] = j
	}
	return result
}

// RangeFrom creates an array of numbers from start with specified length.
// Play: https://go.dev/play/p/0r6VimXAi9H
func RangeFrom[T constraints.Integer | constraints.Float](start T, elementNum int) []T {
	length := xgeneric.If(elementNum < 0, -elementNum).Else(elementNum)
	result := make([]T, length)
	step := xgeneric.If(elementNum < 0, -1).Else(1)
	for i, j := 0, start; i < length; i, j = i+1, j+T(step) {
		result[i] = j
	}
	return result
}

// RangeWithSteps creates an array of numbers (positive and/or negative) progressing from start up to, but not including end.
// step set to zero will return empty array.
// Play: https://go.dev/play/p/0r6VimXAi9H
func RangeWithSteps[T constraints.Integer | constraints.Float](start, end, step T) []T {
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

// Clamp clamps number within the inclusive lower and upper bounds.
// Play: https://go.dev/play/p/RU4lJNC2hlI
func Clamp[T constraints.Ordered](value T, min T, max T) T {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

// Sum sums the values in a collection. If collection is empty 0 is returned.
// Play: https://go.dev/play/p/upfeJVqs4Bt
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	var sum T = 0
	for _, val := range collection {
		sum += val
	}
	return sum
}

// SumBy summarizes the values in a collection using the given return value from the iteration function. If collection is empty 0 is returned.
// Play: https://go.dev/play/p/Dz_a_7jN_ca
func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](collection []T, iteratee func(item T) R) R {
	var sum R = 0
	for _, item := range collection {
		sum = sum + iteratee(item)
	}
	return sum
}
