package slices

// Shift 移除切片中的第一个元素，返回剩余切片与被移除的元素。
// 参数：slice 为输入切片。
// 返回：去除首元素后的切片，以及被移除的首元素。
// 注意：当 slice 为空时访问 slice[0] 会引发运行时 panic。
func Shift[T any](slice []T) ([]T, T) {
	t := slice[0]
	return slice[1:], t
}
