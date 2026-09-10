package xgeneric

// ToPtr 返回指向值 x 副本的指针。
// 参数：x 为任意值。
// 返回：指向 x 副本（位于新分配的地址）的 *T。调用方可以安全地持有该指针而不会影响原值。
func ToPtr[T any](x T) *T {
	return &x
}

// Empty 返回类型 T 的零值。
// 参数：无（仅使用泛型参数 T）。
// 返回：类型 T 的零值（与 var t T 等价）。
func Empty[T any]() T {
	var t T
	return t
}

// Coalesce 返回第一个非零值参数。
// 参数：v 为可变数量的可比较值。
// 返回：第一个与其类型零值不等的元素，以及 true；若所有参数均为零值则返回零值与 false。
func Coalesce[T comparable](v ...T) (result T, ok bool) {
	for _, e := range v {
		if e != result {
			result = e
			ok = true
			return
		}
	}

	return
}
