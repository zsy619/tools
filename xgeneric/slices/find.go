package slices

// Find 返回切片中第一个满足 condition 的元素。
// 参数：slice 为输入切片，condition 为判断函数。
// 返回：首个匹配的元素与 true；若未找到则返回 T 的零值与 false。
func Find[T any](slice []T, condition func(item T, index int, slice []T) bool) (T, bool) {
	for index, item := range slice {
		if condition(item, index, slice) {
			return item, true
		}
	}
	var t T
	return t, false
}
