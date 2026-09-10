package slices

// IndexOfAll 返回切片中所有等于 item 的元素的下标列表。
// 参数：slice 为输入切片，item 为目标元素。
// 返回：包含所有匹配下标的 []int（按升序）；没有匹配时返回 nil。
func IndexOfAll[T comparable](slice []T, item T) []int {
	var indexes []int
	ll := len(slice)
	for i := 0; i < ll; i++ {
		if slice[i] == item {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
