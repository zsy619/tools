package slices

// Every 判断切片中是否所有元素都满足 condition。
// 参数：slice 为输入切片，condition 为判断函数。
// 返回：所有元素都满足时返回 true，否则返回 false。空切片始终返回 true。
func Every[T any](slice []T, condition func(item T, index int, slice []T) bool) bool {
	for index, item := range slice {
		if !condition(item, index, slice) {
			return false
		}
	}
	return true
}
