package xgeneric

// 数字类型
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Signed is a constraint that permits any signed integer type.
// If future releases of Go add new predeclared signed integer types,
// this constraint will be modified to include them.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
// If future releases of Go add new predeclared unsigned integer types,
// this constraint will be modified to include them.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint that permits any integer type.
// If future releases of Go add new predeclared integer types,
// this constraint will be modified to include them.
type Integer interface {
	Signed | Unsigned
}

// Float is a constraint that permits any floating-point type.
// If future releases of Go add new predeclared floating-point types,
// this constraint will be modified to include them.
type Float interface {
	~float32 | ~float64
}

// Complex is a constraint that permits any complex numeric type.
// If future releases of Go add new predeclared complex numeric types,
// this constraint will be modified to include them.
type Complex interface {
	~complex64 | ~complex128
}

// Entry defines a key/value pairs.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// Tuple2 is a group of 2 elements (pair).
type Tuple2[A any, B any] struct {
	A A
	B B
}

// Tuple3 is a group of 3 elements.
type Tuple3[A any, B any, C any] struct {
	A A
	B B
	C C
}

// Tuple4 is a group of 4 elements.
type Tuple4[A any, B any, C any, D any] struct {
	A A
	B B
	C C
	D D
}

// Tuple5 is a group of 5 elements.
type Tuple5[A any, B any, C any, D any, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

// Tuple6 is a group of 6 elements.
type Tuple6[A any, B any, C any, D any, E any, F any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
}

// Tuple7 is a group of 7 elements.
type Tuple7[A any, B any, C any, D any, E any, F any, G any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
}

// Tuple8 is a group of 8 elements.
type Tuple8[A any, B any, C any, D any, E any, F any, G any, H any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
}

// Tuple9 is a group of 9 elements.
type Tuple9[A any, B any, C any, D any, E any, F any, G any, H any, I any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
}
