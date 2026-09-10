package slices

// RemoveDuplicate 去除切片中的重复元素（基于 == 比较，保留首次出现的元素）。
// 参数：slice 为输入切片。
// 返回：去重后的新切片；空切片时返回 nil。
func RemoveDuplicate[T comparable](slice []T) []T {
	var filtered []T
	contains := make(map[T]bool, len(slice))
	for _, item := range slice {
		if !contains[item] {
			contains[item] = true
			filtered = append(filtered, item)
		}
	}
	return filtered
}
