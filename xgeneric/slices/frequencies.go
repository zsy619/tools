package slices

// Frequencies 统计切片中每个不同元素出现的次数。
// 参数：slice 为输入切片（元素必须可比较）。
// 返回：键为元素、值为出现次数的 map[T]int。
func Frequencies[T comparable](slice []T) map[T]int {
	frequencies := make(map[T]int)
	for _, item := range slice {
		frequencies[item]++
	}
	return frequencies
}
