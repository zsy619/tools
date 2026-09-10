package slices

// ToSlice 遍历 map 的键值对并使用 transform 函数转换为新切片。
// 参数：dict 为输入 map，transform 为转换函数（接收键与值返回 T）。
// 返回：按 map 迭代顺序（随机）生成的 []T。
func ToSlice[K comparable, V any, T any](dict map[K]V, transform func(key K, value V) T) []T {
	slice := make([]T, 0, len(dict))
	for key, value := range dict {
		slice = append(slice, transform(key, value))
	}
	return slice
}
