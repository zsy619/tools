package slices

// Any 判断切片中是否存在满足 condition 的元素。
// 参数：slice 为输入切片，condition 为判断函数（接收元素、其下标与整个切片）。
// 返回：存在则返回 true，否则返回 false。空切片始终返回 false。
func Any[T any](slice []T, condition func(item T, index int, slice []T) bool) bool {
	for index, item := range slice {
		if condition(item, index, slice) {
			return true
		}
	}
	return false
}
