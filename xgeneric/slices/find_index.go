package slices

// FindIndex 返回切片中第一个满足 condition 的元素的下标。
// 参数：slice 为输入切片，condition 为判断函数。
// 返回：首个匹配元素的下标；未找到时返回 -1。
func FindIndex[T any](slice []T, condition func(item T, index int, slice []T) bool) int {
	for index, item := range slice {
		if condition(item, index, slice) {
			return index
		}
	}
	return -1
}
