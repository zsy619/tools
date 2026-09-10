package slices

// LastIndexOf 返回切片中最后一个等于 item 的元素的下标。
// 参数：slice 为输入切片，item 为目标元素。
// 返回：最后一个匹配元素的下标；未找到时返回 -1。
func LastIndexOf[T comparable](slice []T, item T) int {
	ll := len(slice)
	for i := ll - 1; i >= 0; i-- {
		if slice[i] == item {
			return i
		}
	}
	return -1
}
