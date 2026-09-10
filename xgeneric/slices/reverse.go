package slices

// Reverse 返回一个元素顺序反转的新切片。
// 参数：slice 为输入切片（不会修改原切片）。
// 返回：长度与 slice 相同、元素顺序反转的新切片。
func Reverse[T any](slice []T) []T {
	ll := len(slice)
	reversed := make([]T, 0, ll)
	for i := ll - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}
	return reversed
}
