package slices

// Distinct 根据 condition 比较函数去除切片中的重复元素（保留首次出现）。
// 参数：slice 为输入切片，condition 为相等性比较函数。
// 返回：仅包含去重后元素的新切片；输入为 nil 或空时返回 nil。
func Distinct[T any](slice []T, condition func(item1, item2 T) bool) []T {
	var newSlice []T
	ll := len(slice)
	for i := 0; i < ll; i++ {
		for j := 0; j <= i; j++ {
			if j == i {
				newSlice = append(newSlice, slice[i])
				break
			}
			if condition(slice[j], slice[i]) {
				break
			}
		}
	}
	return newSlice
}
