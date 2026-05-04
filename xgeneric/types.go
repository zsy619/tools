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

// Unbox 函数将Tuple2中的两个元素拆分成两个单独的变量返回
// A和B是Tuple2中两个元素的类型
// 返回值是Tuple2中的两个元素
func (t Tuple2[A, B]) Unbox() (A, B) {
	return t.A, t.B
}

// NewTuple2 是一个泛型函数，用于创建一个Tuple2类型的实例
// A和B是泛型类型，表示Tuple2元素可以是任意类型
// a和b是Tuple2的两个元素，分别对应A和B类型
// 返回值是一个Tuple2类型的实例，包含a和b两个元素
func NewTuple2[A any, B any](a A, b B) Tuple2[A, B] {
	return Tuple2[A, B]{A: a, B: b}
}

// Tuple3 is a group of 3 elements.
type Tuple3[A any, B any, C any] struct {
	A A
	B B
	C C
}

// Unbox 方法用于将 Tuple3[A, B, C] 类型的元组解包为三个独立的值，
// 并返回这三个值。其中，A、B、C 分别代表元组中的三个类型。
func (t Tuple3[A, B, C]) Unbox() (A, B, C) {
	return t.A, t.B, t.C
}

func NewTuple3[A any, B any, C any](a A, b B, c C) Tuple3[A, B, C] {
	return Tuple3[A, B, C]{A: a, B: b, C: c}
}

// Tuple4 is a group of 4 elements.
type Tuple4[A any, B any, C any, D any] struct {
	A A
	B B
	C C
	D D
}

// Unbox 返回 Tuple4[A, B, C, D] 类型的值中的四个元素 A, B, C, D
// Unbox 返回 Tuple4[A, B, C, D] 类型的值中的四个元素 A, B, C, D
func (t Tuple4[A, B, C, D]) Unbox() (A, B, C, D) {
	return t.A, t.B, t.C, t.D
}

func NewTuple4[A any, B any, C any, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	return Tuple4[A, B, C, D]{A: a, B: b, C: c, D: d}
}

// Tuple5 is a group of 5 elements.
type Tuple5[A any, B any, C any, D any, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

// Unbox 方法用于将Tuple5[A, B, C, D, E]类型的值解包成五个独立的值A、B、C、D、E
// 并返回这五个值
func (t Tuple5[A, B, C, D, E]) Unbox() (A, B, C, D, E) {
	return t.A, t.B, t.C, t.D, t.E
}

func NewTuple5[A any, B any, C any, D any, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	return Tuple5[A, B, C, D, E]{A: a, B: b, C: c, D: d, E: e}
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

// Unbox 返回 Tuple6 中的所有元素
// 返回值为 Tuple6 中包含的 A、B、C、D、E、F 类型的元素
func (t Tuple6[A, B, C, D, E, F]) Unbox() (A, B, C, D, E, F) {
	return t.A, t.B, t.C, t.D, t.E, t.F
}

// NewTuple6 是一个用于创建Tuple6结构体的函数
// A, B, C, D, E, F 是Tuple6结构体中元素的类型
// a, b, c, d, e, f 是用于初始化Tuple6结构体中元素的参数
// 返回值是一个Tuple6类型的结构体
func NewTuple6[A any, B any, C any, D any, E any, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F] {
	return Tuple6[A, B, C, D, E, F]{A: a, B: b, C: c, D: d, E: e, F: f}
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

// Unbox 返回 Tuple7 实例中的 A, B, C, D, E, F, G 类型的值
func (t Tuple7[A, B, C, D, E, F, G]) Unbox() (A, B, C, D, E, F, G) {
	return t.A, t.B, t.C, t.D, t.E, t.F, t.G
}

// NewTuple7 是一个用于创建Tuple7类型实例的函数
// 它接受七个参数，分别为A、B、C、D、E、F、G类型的值
// 并返回一个Tuple7类型的实例，其中包含了传入的七个参数的值
func NewTuple7[A any, B any, C any, D any, E any, F any, G any](a A, b B, c C, d D, e E, f F, g G) Tuple7[A, B, C, D, E, F, G] {
	return Tuple7[A, B, C, D, E, F, G]{A: a, B: b, C: c, D: d, E: e, F: f, G: g}
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

// Unbox 方法将Tuple8类型中的元素拆分成8个独立的变量返回
// 返回值类型分别为A, B, C, D, E, F, G, H
func (t Tuple8[A, B, C, D, E, F, G, H]) Unbox() (A, B, C, D, E, F, G, H) {
	return t.A, t.B, t.C, t.D, t.E, t.F, t.G, t.H
}

// NewTuple8 函数返回一个Tuple8类型，它包含了八个类型分别为A, B, C, D, E, F, G, H的泛型元素。
// 参数a, b, c, d, e, f, g, h分别对应Tuple8类型中的A, B, C, D, E, F, G, H元素。
// 函数返回值为Tuple8[A, B, C, D, E, F, G, H]类型。
func NewTuple8[A any, B any, C any, D any, E any, F any, G any, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tuple8[A, B, C, D, E, F, G, H] {
	return Tuple8[A, B, C, D, E, F, G, H]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h}
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

// Unbox 将Tuple9对象中的元素解包成单个的值并返回
// 返回值类型为Tuple9中定义的泛型类型A, B, C, D, E, F, G, H, I
func (t Tuple9[A, B, C, D, E, F, G, H, I]) Unbox() (A, B, C, D, E, F, G, H, I) {
	return t.A, t.B, t.C, t.D, t.E, t.F, t.G, t.H, t.I
}

// NewTuple9 是一个泛型函数，用于创建一个包含9个元素的元组Tuple9
// A, B, C, D, E, F, G, H, I 是元组中元素的类型
// a, b, c, d, e, f, g, h, i 是元组中对应的元素值
// 函数返回一个Tuple9类型的元组，其中包含了传入的9个元素
func NewTuple9[A any, B any, C any, D any, E any, F any, G any, H any, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Tuple9[A, B, C, D, E, F, G, H, I] {
	return Tuple9[A, B, C, D, E, F, G, H, I]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h, I: i}
}

// Tuple10 is a group of 9 elements.
type Tuple10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
	J J
}

// Unbox 方法用于将Tuple10类型对象中的元素分别返回。
// 返回值类型为Tuple10中定义的类型参数A, B, C, D, E, F, G, H, I, J。
func (t Tuple10[A, B, C, D, E, F, G, H, I, J]) Unbox() (A, B, C, D, E, F, G, H, I, J) {
	return t.A, t.B, t.C, t.D, t.E, t.F, t.G, t.H, t.I, t.J
}

// NewTuple10 是一个泛型函数，用于创建一个包含10个元素的元组
// 参数a、b、c、d、e、f、g、h、i、j分别为元组的10个元素，类型分别为A、B、C、D、E、F、G、H、I、J
// 返回值为一个Tuple10类型的元组，包含10个元素，元素类型与传入的参数类型一致
func NewTuple10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) Tuple10[A, B, C, D, E, F, G, H, I, J] {
	return Tuple10[A, B, C, D, E, F, G, H, I, J]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h, I: i, J: j}
}
