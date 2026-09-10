package xgeneric

// FromPtr 解引用指针，若指针为 nil 则返回 T 的零值。
// 参数：x 为待解引用的指针。
// 返回：当 x 非 nil 时返回 *x，否则返回 Empty[T]() 即 T 的零值。
func FromPtr[T any](x *T) T {
	if x == nil {
		return Empty[T]()
	}

	return *x
}

// FromPtrOr 解引用指针，若指针为 nil 则返回给定的回退值。
// 参数：x 为待解引用的指针，fallback 为 x 为 nil 时返回的值。
// 返回：当 x 非 nil 时返回 *x，否则返回 fallback。
func FromPtrOr[T any](x *T, fallback T) T {
	if x == nil {
		return fallback
	}

	return *x
}

// ToAnySlice 将切片中所有元素转换为 any 类型，生成一个新的 []any。
// 参数：collection 为输入切片。
// 返回：长度与 collection 相同、元素类型为 any 的新切片。
func ToAnySlice[T any](collection []T) []any {
	result := make([]any, len(collection))
	for i, item := range collection {
		result[i] = item
	}
	return result
}

// FromAnySlice 将 []any 切片中的元素全部转换为类型 T。
// 参数：in 为输入的 []any 切片。
// 返回：转换后的 []T 切片与 true；若任一元素的实际类型与 T 不匹配而触发 panic 会被 recover 捕获，返回空切片与 false。
func FromAnySlice[T any](in []any) (out []T, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			out = []T{}
			ok = false
		}
	}()

	result := make([]T, len(in))
	for i, item := range in {
		result[i] = item.(T)
	}
	return result, true
}

// IsEmpty 判断值是否为该类型的零值。
// 参数：v 为待检查的值（必须可比较）。
// 返回：当 v 等于其类型零值时返回 true，否则返回 false。
func IsEmpty[T comparable](v T) bool {
	var zero T
	return zero == v
}

// IsNotEmpty 判断值是否不是该类型的零值。
// 参数：v 为待检查的值（必须可比较）。
// 返回：当 v 不等于其类型零值时返回 true，否则返回 false。
func IsNotEmpty[T comparable](v T) bool {
	var zero T
	return zero != v
}
