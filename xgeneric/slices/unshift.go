package slices

// Unshift 将 items 中的元素添加到 slice 头部并返回结果切片。
// 参数：slice 为原切片，items 为要前置的元素（可变数量）。
// 返回：包含 items 全部元素后再追加 slice 的新切片。
func Unshift[T any](slice []T, items ...T) []T {
	return append(items, slice...)
}
