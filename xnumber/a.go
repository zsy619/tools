package xnumber

//go:generate goption -p . -c Worker -w
//go:generate gofmt -w .

type (
	IntAll interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}
	UintAll interface {
		~uint | ~uint8 | ~uint16 | ~uint32
	}
	FloatAll  interface{ ~float32 | ~float64 }
	NumberAll interface{ IntAll | UintAll | FloatAll }
)
