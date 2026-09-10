package slices

// Filter 保留切片中满足 condition 的元素，按原顺序组成新切片。
// 参数：slice 为输入切片，condition 为判断函数。
// 返回：仅包含满足条件的元素的新切片；空切片时返回 nil。
func Filter[T any](slice []T, condition func(item T, index int, slice []T) bool) []T {
	var filtered []T
	for index, item := range slice {
		if condition(item, index, slice) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
