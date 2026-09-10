package slices

// ForEach 对切片中的每个元素按顺序执行 action。
// 参数：slice 为输入切片，action 为要执行的回调函数。
// 副作用：依赖 action 的具体行为；不会中断或修改切片遍历。
func ForEach[T any](slice []T, action func(item T, index int, slice []T)) {
	for index, item := range slice {
		action(item, index, slice)
	}
}
