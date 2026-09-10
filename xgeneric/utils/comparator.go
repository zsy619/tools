package utils

type Ordered interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

// Comparator 是通用的比较器函数类型。
// 应当返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。
type Comparator[T any] func(a, b T) int

// OrderedTypeCmp 使用通用比较运算符对满足 Ordered 约束的两个值进行比较。
// 参数：a、b 为要比较的有序值。
// 返回：a == b 时返回 0；a < b 时返回 -1；否则返回 1。
func OrderedTypeCmp[T Ordered](a, b T) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Reverse 返回一个与传入比较器结果相反的比较器（即反向排序）。
// 参数：cmp 为原始 Comparator[T]。
// 返回：cmp 的反向 Comparator[T]，其返回值为 -cmp(a, b)。
func Reverse[T any](cmp Comparator[T]) Comparator[T] {
	return func(a, b T) int {
		return -cmp(a, b)
	}
}

// IntComparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func IntComparator(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// UintComparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func UintComparator(a, b uint) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Int8Comparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func Int8Comparator(a, b int8) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Uint8Comparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func Uint8Comparator(a, b uint8) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Int16Comparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func Int16Comparator(a, b int16) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Uint16Comparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func Uint16Comparator(a, b uint16) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Int32Comparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func Int32Comparator(a, b int32) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Uint32Comparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func Uint32Comparator(a, b uint32) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Int64Comparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func Int64Comparator(a, b int64) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Uint64Comparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func Uint64Comparator(a, b uint64) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Float32Comparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func Float32Comparator(a, b float32) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// Float64Comparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func Float64Comparator(a, b float64) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// StringComparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func StringComparator(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// UintptrComparator 比较 a 与 b，返回：-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func UintptrComparator(a, b uintptr) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// BoolComparator 比较两个 bool 值，按 false < true 排序。-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func BoolComparator(a, b bool) int {
	if a == b {
		return 0
	}
	if !a && b {
		return -1
	}
	return 1
}

// Complex64Comparator 先比较实部、再比较虚部。-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func Complex64Comparator(a, b complex64) int {
	if a == b {
		return 0
	}
	if real(a) < real(a) {
		return -1
	}
	if real(a) == real(b) && imag(a) < imag(b) {
		return -1
	}
	return 1
}

// Complex128Comparator 先比较实部、再比较虚部。-1 表示 a < b，0 表示 a == b，1 表示 a > b。

func Complex128Comparator(a, b complex128) int {
	if a == b {
		return 0
	}
	if real(a) < real(b) {
		return -1
	}
	if real(a) == real(b) && imag(a) < imag(b) {
		return -1
	}
	return 1
}
