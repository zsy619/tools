package xgeneric

// Filter 根据谓词函数过滤切片，保留所有满足条件的元素，保持原有顺序。
// 参数：f 为判断函数，src 为输入切片。
// 返回：仅包含满足 f 为 true 的元素的新切片；当没有元素匹配时返回 nil。
func Filter[T any](f func(T) bool, src []T) []T {
	var dst []T
	for _, v := range src {
		if f(v) {
			dst = append(dst, v)
		}
	}
	return dst
}

// IFF 根据 yes 条件选择返回 a 或 b。
// 参数：yes 为判断条件，a 为条件为真时返回的值，b 为条件为假时返回的值。
// 返回：yes 为 true 时返回 a，否则返回 b。
func IFF[T any](yes bool, a, b T) T {
	if yes {
		return a
	}
	return b
}

// IFN 根据 yes 条件选择调用 a() 或 b()。
// 参数：yes 为判断条件，a、b 为返回 T 的函数。
// 返回：yes 为 true 时返回 a() 的结果，否则返回 b() 的结果。仅调用被选中的函数。
func IFN[T any](yes bool, a, b func() T) T {
	if yes {
		return a()
	}
	return b()
}

// IsSame 判断两个值是否相等（基于 == 操作符）。
// 参数：a、b 为要比较的两个值（必须可比较）。
// 返回：相等时返回 true，否则返回 false。
func IsSame[T comparable](a T, b T) bool {
	return a == b
}
