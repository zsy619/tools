package xgeneric

func Filter[T any](f func(T) bool, src []T) []T {
	var dst []T
	for _, v := range src {
		if f(v) {
			dst = append(dst, v)
		}
	}
	return dst
}

// IFF if yes return a else b
func IFF[T any](yes bool, a, b T) T {
	if yes {
		return a
	}
	return b
}

// IFN if yes return func, a() else b().
func IFN[T any](yes bool, a, b func() T) T {
	if yes {
		return a()
	}
	return b()
}

func IsSame[T comparable](a T, b T) bool {
	return a == b
}
