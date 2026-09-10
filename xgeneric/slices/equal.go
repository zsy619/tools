package slices

// Equal 判断两个切片是否完全相等（基于 == 操作符）。
// 参数：slice1、slice2 为待比较的两个切片。
// 返回：长度相同且逐元素相等时返回 true，否则返回 false。
func Equal[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for index, item := range slice1 {
		if item != slice2[index] {
			return false
		}
	}
	return true
}

// EqualFunc 使用自定义比较函数判断两个切片是否完全相等。
// 参数：slice1、slice2 为待比较的两个切片，f 为相等性比较函数。
// 返回：长度相同且每对元素经 f 判断相等时返回 true，否则返回 false。
func EqualFunc[T any](slice1, slice2 []T, f func(item1, item2 T) bool) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for index, item := range slice1 {
		if !f(item, slice2[index]) {
			return false
		}
	}
	return true
}
