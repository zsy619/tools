package slices

// Fill 使用 itemFunc 函数对切片中的每个元素进行重新赋值（就地修改）。
// 参数：slice 为输入切片，itemFunc 为生成新值的函数（接收当前值、下标、整个切片）。
// 返回：原切片本身（已就地修改）。
func Fill[T any](slice []T, itemFunc func(item T, index int, slice []T) T) []T {
	ll := len(slice)
	for index := 0; index < ll; index++ {
		slice[index] = itemFunc(slice[index], index, slice)
	}
	return slice
}
