package slices

// ReduceRight 按反向顺序将切片归约为单个值。
// 参数：slice 为输入切片，reduce 为归约函数（接收累计值、当前元素、下标与切片），init 为累计值初值。
// 返回：所有元素从尾部向头部累积后的最终值；空切片时直接返回 init。
func ReduceRight[T any, R any](slice []T, reduce func(total R, item T, index int, slice []T) R, init R) R {
	ll := len(slice)
	for index := ll - 1; index >= 0; index-- {
		init = reduce(init, slice[index], index, slice)
	}
	return init
}
