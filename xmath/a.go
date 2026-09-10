package xmath

//go:generate goption -p . -c Worker -w
//go:generate gofmt -w .

type (
	// IntAll 所有有符号整型（包括类型定义的别名，如 type MyInt int）。
	IntAll interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}
	// UintAll 所有无符号整型（包括类型定义的别名）。
	UintAll interface {
		~uint | ~uint8 | ~uint16 | ~uint32
	}
	// FloatAll 所有浮点类型（float32 / float64）。
	FloatAll interface{ ~float32 | ~float64 }
	// NumberAll 通用数值类型，聚合了 IntAll、UintAll、FloatAll。
	NumberAll interface{ IntAll | UintAll | FloatAll }
)
