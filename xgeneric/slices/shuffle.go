package slices

import "math/rand"

// Shuffle 原地打乱切片中元素的顺序。
// 参数：slice 为要打乱的切片（会被就地修改）。
// 副作用：使用 math/rand 中的 Shuffle 实现，每次运行结果随机。
func Shuffle[T any](slice []T) {
	ll := len(slice)
	for i := 0; i < ll; i++ {
		rand.Shuffle(ll, func(i, j int) {
			slice[i], slice[j] = slice[j], slice[i]
		})
	}
}
