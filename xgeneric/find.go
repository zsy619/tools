package xgeneric

import (
	"fmt"
	"math"
	"math/rand"
)

// IndexOf 返回指定元素在切片中首次出现的下标；若不存在则返回 -1。
// 参数：collection 为待搜索的切片，element 为要查找的元素。
// 返回：首个匹配元素的下标（从 0 开始），未找到时返回 -1。
func IndexOf[T comparable](collection []T, element T) int {
	for i, item := range collection {
		if item == element {
			return i
		}
	}

	return -1
}

// LastIndexOf 返回指定元素在切片中最后一次出现的下标；若不存在则返回 -1。
// 参数：collection 为待搜索的切片，element 为要查找的元素。
// 返回：最后一个匹配元素的下标，未找到时返回 -1。
func LastIndexOf[T comparable](collection []T, element T) int {
	length := len(collection)

	for i := length - 1; i >= 0; i-- {
		if collection[i] == element {
			return i
		}
	}

	return -1
}

// Find 根据谓词函数在切片中查找第一个满足条件的元素。
// 参数：collection 为待搜索的切片，predicate 为判断元素是否满足条件的函数。
// 返回：找到的元素及 true；若未找到则返回 T 的零值与 false。
func Find[T any](collection []T, predicate func(T) bool) (T, bool) {
	for _, item := range collection {
		if predicate(item) {
			return item, true
		}
	}

	var result T
	return result, false
}

// FindIndexOf 根据谓词函数在切片中查找第一个满足条件的元素，并返回其下标。
// 参数：collection 为待搜索的切片，predicate 为判断元素是否满足条件的函数。
// 返回：匹配的元素、其下标以及 true；若未找到则返回 T 的零值、-1 与 false。
func FindIndexOf[T any](collection []T, predicate func(T) bool) (T, int, bool) {
	for i, item := range collection {
		if predicate(item) {
			return item, i, true
		}
	}

	var result T
	return result, -1, false
}

// FindLastIndexOf 根据谓词函数在切片中从尾部向前查找最后一个满足条件的元素，并返回其下标。
// 参数：collection 为待搜索的切片，predicate 为判断元素是否满足条件的函数。
// 返回：匹配的元素、其下标以及 true；若未找到则返回 T 的零值、-1 与 false。
func FindLastIndexOf[T any](collection []T, predicate func(T) bool) (T, int, bool) {
	length := len(collection)

	for i := length - 1; i >= 0; i-- {
		if predicate(collection[i]) {
			return collection[i], i, true
		}
	}

	var result T
	return result, -1, false
}

// FindOrElse 根据谓词函数在切片中查找第一个满足条件的元素；若未找到则返回给定的回退值。
// 参数：collection 为待搜索的切片，fallback 为未匹配时返回的值，predicate 为判断函数。
// 返回：首个匹配的元素；若没有任何匹配项则返回 fallback。
func FindOrElse[T any](collection []T, fallback T, predicate func(T) bool) T {
	for _, item := range collection {
		if predicate(item) {
			return item
		}
	}

	return fallback
}

// Min 返回切片中的最小值。
// 参数：collection 为输入切片（元素必须满足 Ordered 约束）。
// 返回：切片中的最小值；当切片为空时返回 T 的零值。
func Min[T Ordered](collection []T) T {
	var min T

	if len(collection) == 0 {
		return min
	}

	min = collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if item < min {
			min = item
		}
	}

	return min
}

// MinBy 使用给定的比较函数在切片中查找最小值。
// 参数：collection 为输入切片，comparison 为比较函数，当第一个参数小于第二个参数时返回 true。
// 返回：切片中的最小元素；当存在多个相同最小值时返回最早出现的那个；切片为空时返回 T 的零值。
func MinBy[T any](collection []T, comparison func(T, T) bool) T {
	var min T

	if len(collection) == 0 {
		return min
	}

	min = collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if comparison(item, min) {
			min = item
		}
	}

	return min
}

// Max 返回切片中的最大值。
// 参数：collection 为输入切片（元素必须满足 Ordered 约束）。
// 返回：切片中的最大值；当切片为空时返回 T 的零值。
func Max[T Ordered](collection []T) T {
	var max T

	if len(collection) == 0 {
		return max
	}

	max = collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if item > max {
			max = item
		}
	}

	return max
}

// MaxBy 使用给定的比较函数在切片中查找最大值。
// 参数：collection 为输入切片，comparison 为比较函数，当第一个参数大于第二个参数时返回 true。
// 返回：切片中的最大元素；当存在多个相同最大值时返回最早出现的那个；切片为空时返回 T 的零值。
func MaxBy[T any](collection []T, comparison func(T, T) bool) T {
	var max T

	if len(collection) == 0 {
		return max
	}

	max = collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if comparison(item, max) {
			max = item
		}
	}

	return max
}

// Last 返回切片的最后一个元素。
// 参数：collection 为输入切片。
// 返回：最后一个元素及 nil 错误；当切片为空时返回 T 的零值与 error（信息为 "last: cannot extract the last element of an empty slice"）。
func Last[T any](collection []T) (T, error) {
	length := len(collection)

	if length == 0 {
		var t T
		return t, fmt.Errorf("last: cannot extract the last element of an empty slice")
	}

	return collection[length-1], nil
}

// Nth 返回切片中下标为 nth 的元素。
// 参数：collection 为输入切片，nth 为目标下标；当 nth 为负数时表示从尾部倒数（-1 表示最后一个元素）。
// 返回：该下标处的元素及 nil 错误；当 |nth| 超出切片范围时返回 T 的零值与 error。
func Nth[T any](collection []T, nth int) (T, error) {
	if int(math.Abs(float64(nth))) >= len(collection) {
		var t T
		return t, fmt.Errorf("nth: %d out of slice bounds", nth)
	}

	length := len(collection)

	if nth >= 0 {
		return collection[nth], nil
	}

	return collection[length+nth], nil
}

// Sample 从切片中随机返回一个元素。
// 参数：collection 为输入切片。
// 返回：随机选中的元素；当切片为空时返回 Empty[T]()（即 T 的零值）。使用 math/rand 进行随机抽样。
func Sample[T any](collection []T) T {
	size := len(collection)
	if size == 0 {
		return Empty[T]()
	}

	return collection[rand.Intn(size)]
}

// Samples 从切片中随机返回 count 个不重复的元素。
// 参数：collection 为输入切片，count 为期望返回的元素个数。
// 返回：长度最多为 count 的新切片，每个元素互不相同；当 count 大于切片长度时仅返回切片全部元素。使用 math/rand 进行抽样，会修改内部副本。
func Samples[T any](collection []T, count int) []T {
	size := len(collection)

	// put values into a map, for faster deletion
	cOpy := make([]T, 0, size)
	cOpy = append(cOpy, collection...)

	results := []T{}

	for i := 0; i < size && i < count; i++ {
		copyLength := size - i

		index := rand.Intn(size - i)
		results = append(results, cOpy[index])

		// Removes element.
		// It is faster to swap with last element and remove it.
		cOpy[index] = cOpy[copyLength-1]
		cOpy = cOpy[:copyLength-1]
	}

	return results
}
