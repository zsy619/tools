package slices

// FindLast 从尾部向前查找切片中最后一个满足 condition 的元素。
// 参数：slice 为输入切片，condition 为判断函数。
// 返回：最后一个匹配的元素与 true；若未找到则返回 T 的零值与 false。
func FindLast[T any](slice []T, condition func(item T, index int, slice []T) bool) (T, bool) {
	ll := len(slice)
	for i := ll - 1; i >= 0; i-- {
		if condition(slice[i], i, slice) {
			return slice[i], true
		}
	}
	var t T
	return t, false
}
