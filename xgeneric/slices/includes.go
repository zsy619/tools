package slices

// Includes 判断切片中是否包含指定元素。
// 参数：slice 为输入切片，item 为要查找的元素。
// 返回：找到则返回 true，否则返回 false。
func Includes[T comparable](slice []T, item T) bool {
	for _, _item := range slice {
		if _item == item {
			return true
		}
	}
	return false
}
