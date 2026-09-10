package slices

// FindLastIndex 从尾部向前查找切片中最后一个满足 condition 的元素下标。
// 参数：slice 为输入切片，condition 为判断函数。
// 返回：最后一个匹配元素的下标；未找到时返回 -1。
func FindLastIndex[T any](slice []T, condition func(item T, index int, slice []T) bool) int {
	ll := len(slice)
	for i := ll - 1; i >= 0; i-- {
		if condition(slice[i], i, slice) {
			return i
		}
	}
	return -1
}
