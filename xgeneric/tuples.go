package xgeneric

// T2 由两个值构造一个 Tuple2 元组。
// 参数：a、b 为元组的两个元素。
// 返回：包含 a、b 的 Tuple2[A, B]。
func T2[A any, B any](a A, b B) Tuple2[A, B] {
	return Tuple2[A, B]{A: a, B: b}
}

// T3 由三个值构造一个 Tuple3 元组。
// 参数：a、b、c 为元组的三个元素。
// 返回：包含 a、b、c 的 Tuple3[A, B, C]。
func T3[A any, B any, C any](a A, b B, c C) Tuple3[A, B, C] {
	return Tuple3[A, B, C]{A: a, B: b, C: c}
}

// T4 由四个值构造一个 Tuple4 元组。
// 参数：a、b、c、d 为元组的四个元素。
// 返回：包含 a、b、c、d 的 Tuple4[A, B, C, D]。
func T4[A any, B any, C any, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	return Tuple4[A, B, C, D]{A: a, B: b, C: c, D: d}
}

// T5 由五个值构造一个 Tuple5 元组。
// 参数：a、b、c、d、e 为元组的五个元素。
// 返回：包含 a、b、c、d、e 的 Tuple5[A, B, C, D, E]。
func T5[A any, B any, C any, D any, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	return Tuple5[A, B, C, D, E]{A: a, B: b, C: c, D: d, E: e}
}

// T6 由六个值构造一个 Tuple6 元组。
// 参数：a、b、c、d、e、f 为元组的六个元素。
// 返回：包含 a、b、c、d、e、f 的 Tuple6[A, B, C, D, E, F]。
func T6[A any, B any, C any, D any, E any, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F] {
	return Tuple6[A, B, C, D, E, F]{A: a, B: b, C: c, D: d, E: e, F: f}
}

// T7 由七个值构造一个 Tuple7 元组。
// 参数：a、b、c、d、e、f、g 为元组的七个元素。
// 返回：包含 a、b、c、d、e、f、g 的 Tuple7[A, B, C, D, E, F, G]。
func T7[A any, B any, C any, D any, E any, F any, G any](a A, b B, c C, d D, e E, f F, g G) Tuple7[A, B, C, D, E, F, G] {
	return Tuple7[A, B, C, D, E, F, G]{A: a, B: b, C: c, D: d, E: e, F: f, G: g}
}

// T8 由八个值构造一个 Tuple8 元组。
// 参数：a、b、c、d、e、f、g、h 为元组的八个元素。
// 返回：包含 a、b、c、d、e、f、g、h 的 Tuple8[A, B, C, D, E, F, G, H]。
func T8[A any, B any, C any, D any, E any, F any, G any, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tuple8[A, B, C, D, E, F, G, H] {
	return Tuple8[A, B, C, D, E, F, G, H]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h}
}

// T9 由九个值构造一个 Tuple9 元组。
// 参数：a、b、c、d、e、f、g、h、i 为元组的九个元素。
// 返回：包含 a、b、c、d、e、f、g、h、i 的 Tuple9[A, B, C, D, E, F, G, H, I]。
func T9[A any, B any, C any, D any, E any, F any, G any, H any, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Tuple9[A, B, C, D, E, F, G, H, I] {
	return Tuple9[A, B, C, D, E, F, G, H, I]{A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h, I: i}
}

// Unpack2 解包 Tuple2 元组，返回其中的两个独立值。
// 参数：tuple 为要解包的 Tuple2[A, B]。
// 返回：元组中的 A、B 两个值。
func Unpack2[A any, B any](tuple Tuple2[A, B]) (A, B) {
	return tuple.A, tuple.B
}

// Unpack3 解包 Tuple3 元组，返回其中的三个独立值。
// 参数：tuple 为要解包的 Tuple3[A, B, C]。
// 返回：元组中的 A、B、C 三个值。
func Unpack3[A any, B any, C any](tuple Tuple3[A, B, C]) (A, B, C) {
	return tuple.A, tuple.B, tuple.C
}

// Unpack4 解包 Tuple4 元组，返回其中的四个独立值。
// 参数：tuple 为要解包的 Tuple4[A, B, C, D]。
// 返回：元组中的 A、B、C、D 四个值。
func Unpack4[A any, B any, C any, D any](tuple Tuple4[A, B, C, D]) (A, B, C, D) {
	return tuple.A, tuple.B, tuple.C, tuple.D
}

// Unpack5 解包 Tuple5 元组，返回其中的五个独立值。
// 参数：tuple 为要解包的 Tuple5[A, B, C, D, E]。
// 返回：元组中的 A、B、C、D、E 五个值。
func Unpack5[A any, B any, C any, D any, E any](tuple Tuple5[A, B, C, D, E]) (A, B, C, D, E) {
	return tuple.A, tuple.B, tuple.C, tuple.D, tuple.E
}

// Unpack6 解包 Tuple6 元组，返回其中的六个独立值。
// 参数：tuple 为要解包的 Tuple6[A, B, C, D, E, F]。
// 返回：元组中的 A、B、C、D、E、F 六个值。
func Unpack6[A any, B any, C any, D any, E any, F any](tuple Tuple6[A, B, C, D, E, F]) (A, B, C, D, E, F) {
	return tuple.A, tuple.B, tuple.C, tuple.D, tuple.E, tuple.F
}

// Unpack7 解包 Tuple7 元组，返回其中的七个独立值。
// 参数：tuple 为要解包的 Tuple7[A, B, C, D, E, F, G]。
// 返回：元组中的 A、B、C、D、E、F、G 七个值。
func Unpack7[A any, B any, C any, D any, E any, F any, G any](tuple Tuple7[A, B, C, D, E, F, G]) (A, B, C, D, E, F, G) {
	return tuple.A, tuple.B, tuple.C, tuple.D, tuple.E, tuple.F, tuple.G
}

// Unpack8 解包 Tuple8 元组，返回其中的八个独立值。
// 参数：tuple 为要解包的 Tuple8[A, B, C, D, E, F, G, H]。
// 返回：元组中的 A、B、C、D、E、F、G、H 八个值。
func Unpack8[A any, B any, C any, D any, E any, F any, G any, H any](tuple Tuple8[A, B, C, D, E, F, G, H]) (A, B, C, D, E, F, G, H) {
	return tuple.A, tuple.B, tuple.C, tuple.D, tuple.E, tuple.F, tuple.G, tuple.H
}

// Unpack9 解包 Tuple9 元组，返回其中的九个独立值。
// 参数：tuple 为要解包的 Tuple9[A, B, C, D, E, F, G, H, I]。
// 返回：元组中的 A、B、C、D、E、F、G、H、I 九个值。
func Unpack9[A any, B any, C any, D any, E any, F any, G any, H any, I any](tuple Tuple9[A, B, C, D, E, F, G, H, I]) (A, B, C, D, E, F, G, H, I) {
	return tuple.A, tuple.B, tuple.C, tuple.D, tuple.E, tuple.F, tuple.G, tuple.H, tuple.I
}

// Zip2 将多个切片按下标位置打包为 Tuple2 切片。
// 参数：a、b 为输入切片。
// 返回：长度为 max(len(a), len(b)) 的 Tuple2 切片；当输入切片长度不一致时，缺失位置以类型零值填充。
func Zip2[A any, B any](a []A, b []B) []Tuple2[A, B] {
	size := Max([]int{len(a), len(b)})

	result := make([]Tuple2[A, B], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)

		result = append(result, Tuple2[A, B]{
			A: _a,
			B: _b,
		})
	}

	return result
}

// Zip3 将多个切片按下标位置打包为 Tuple3 切片。
// 参数：a、b、c 为输入切片。
// 返回：长度为 max(len(a), len(b), len(c)) 的 Tuple3 切片；当输入切片长度不一致时，缺失位置以类型零值填充。
func Zip3[A any, B any, C any](a []A, b []B, c []C) []Tuple3[A, B, C] {
	size := Max([]int{len(a), len(b), len(c)})

	result := make([]Tuple3[A, B, C], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)

		result = append(result, Tuple3[A, B, C]{
			A: _a,
			B: _b,
			C: _c,
		})
	}

	return result
}

// Zip4 将多个切片按下标位置打包为 Tuple4 切片。
// 参数：a、b、c、d 为输入切片。
// 返回：长度为 max(len(a), len(b), len(c), len(d)) 的 Tuple4 切片；当输入切片长度不一致时，缺失位置以类型零值填充。
func Zip4[A any, B any, C any, D any](a []A, b []B, c []C, d []D) []Tuple4[A, B, C, D] {
	size := Max([]int{len(a), len(b), len(c), len(d)})

	result := make([]Tuple4[A, B, C, D], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)

		result = append(result, Tuple4[A, B, C, D]{
			A: _a,
			B: _b,
			C: _c,
			D: _d,
		})
	}

	return result
}

// Zip5 将多个切片按下标位置打包为 Tuple5 切片。
// 参数：a、b、c、d、e 为输入切片。
// 返回：长度为 max(len(a), len(b), len(c), len(d), len(e)) 的 Tuple5 切片；当输入切片长度不一致时，缺失位置以类型零值填充。
func Zip5[A any, B any, C any, D any, E any](a []A, b []B, c []C, d []D, e []E) []Tuple5[A, B, C, D, E] {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e)})

	result := make([]Tuple5[A, B, C, D, E], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)

		result = append(result, Tuple5[A, B, C, D, E]{
			A: _a,
			B: _b,
			C: _c,
			D: _d,
			E: _e,
		})
	}

	return result
}

// Zip6 将多个切片按下标位置打包为 Tuple6 切片。
// 参数：a、b、c、d、e、f 为输入切片。
// 返回：长度为 max(len(a)..len(f)) 的 Tuple6 切片；当输入切片长度不一致时，缺失位置以类型零值填充。
func Zip6[A any, B any, C any, D any, E any, F any](a []A, b []B, c []C, d []D, e []E, f []F) []Tuple6[A, B, C, D, E, F] {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e), len(f)})

	result := make([]Tuple6[A, B, C, D, E, F], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)
		_f, _ := Nth(f, index)

		result = append(result, Tuple6[A, B, C, D, E, F]{
			A: _a,
			B: _b,
			C: _c,
			D: _d,
			E: _e,
			F: _f,
		})
	}

	return result
}

// Zip7 将多个切片按下标位置打包为 Tuple7 切片。
// 参数：a、b、c、d、e、f、g 为输入切片。
// 返回：长度为 max(len(a)..len(g)) 的 Tuple7 切片；当输入切片长度不一致时，缺失位置以类型零值填充。
func Zip7[A any, B any, C any, D any, E any, F any, G any](a []A, b []B, c []C, d []D, e []E, f []F, g []G) []Tuple7[A, B, C, D, E, F, G] {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e), len(f), len(g)})

	result := make([]Tuple7[A, B, C, D, E, F, G], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)
		_f, _ := Nth(f, index)
		_g, _ := Nth(g, index)

		result = append(result, Tuple7[A, B, C, D, E, F, G]{
			A: _a,
			B: _b,
			C: _c,
			D: _d,
			E: _e,
			F: _f,
			G: _g,
		})
	}

	return result
}

// Zip8 将多个切片按下标位置打包为 Tuple8 切片。
// 参数：a、b、c、d、e、f、g、h 为输入切片。
// 返回：长度为 max(len(a)..len(h)) 的 Tuple8 切片；当输入切片长度不一致时，缺失位置以类型零值填充。
func Zip8[A any, B any, C any, D any, E any, F any, G any, H any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H) []Tuple8[A, B, C, D, E, F, G, H] {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e), len(f), len(g), len(h)})

	result := make([]Tuple8[A, B, C, D, E, F, G, H], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)
		_f, _ := Nth(f, index)
		_g, _ := Nth(g, index)
		_h, _ := Nth(h, index)

		result = append(result, Tuple8[A, B, C, D, E, F, G, H]{
			A: _a,
			B: _b,
			C: _c,
			D: _d,
			E: _e,
			F: _f,
			G: _g,
			H: _h,
		})
	}

	return result
}

// Zip9 是一个泛型函数，它接受九个切片作为参数，并将这些切片压缩成一个元组切片
//
// 参数：
//
//	a []A：第一个切片，元素类型为A
//	b []B：第二个切片，元素类型为B
//	c []C：第三个切片，元素类型为C
//	d []D：第四个切片，元素类型为D
//	e []E：第五个切片，元素类型为E
//	f []F：第六个切片，元素类型为F
//	g []G：第七个切片，元素类型为G
//	h []H：第八个切片，元素类型为H
//	i []I：第九个切片，元素类型为I
//
// 返回值：
//
//	[]Tuple9[A, B, C, D, E, F, G, H, I]：一个元组切片，其中每个元组包含了九个切片的对应位置的元素
func Zip9[A any, B any, C any, D any, E any, F any, G any, H any, I any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, i []I) []Tuple9[A, B, C, D, E, F, G, H, I] {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e), len(f), len(g), len(h), len(i)})

	result := make([]Tuple9[A, B, C, D, E, F, G, H, I], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)
		_f, _ := Nth(f, index)
		_g, _ := Nth(g, index)
		_h, _ := Nth(h, index)
		_i, _ := Nth(i, index)

		result = append(result, Tuple9[A, B, C, D, E, F, G, H, I]{
			A: _a,
			B: _b,
			C: _c,
			D: _d,
			E: _e,
			F: _f,
			G: _g,
			H: _h,
			I: _i,
		})
	}

	return result
}

// Zip10 将十个切片a, b, c, d, e, f, g, h, i, j进行zip操作，返回一个新的切片，其中每个元素是一个Tuple10类型的元组
// 参数：
//
//	a []A：类型为A的切片
//	b []B：类型为B的切片
//	c []C：类型为C的切片
//	d []D：类型为D的切片
//	e []E：类型为E的切片
//	f []F：类型为F的切片
//	g []G：类型为G的切片
//	h []H：类型为H的切片
//	i []I：类型为I的切片
//	j []J：类型为J的切片
//
// 返回值：
//
//	[]Tuple10[A, B, C, D, E, F, G, H, I, J]：一个元素类型为Tuple10[A, B, C, D, E, F, G, H, I, J]的切片
func Zip10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, i []I, j []J) []Tuple10[A, B, C, D, E, F, G, H, I, J] {
	size := Max([]int{len(a), len(b), len(c), len(d), len(e), len(f), len(g), len(h), len(i), len(j)})

	result := make([]Tuple10[A, B, C, D, E, F, G, H, I, J], 0, size)

	for index := 0; index < size; index++ {
		_a, _ := Nth(a, index)
		_b, _ := Nth(b, index)
		_c, _ := Nth(c, index)
		_d, _ := Nth(d, index)
		_e, _ := Nth(e, index)
		_f, _ := Nth(f, index)
		_g, _ := Nth(g, index)
		_h, _ := Nth(h, index)
		_i, _ := Nth(i, index)
		_j, _ := Nth(j, index)

		result = append(result, Tuple10[A, B, C, D, E, F, G, H, I, J]{
			A: _a,
			B: _b,
			C: _c,
			D: _d,
			E: _e,
			F: _f,
			G: _g,
			H: _h,
			I: _i,
			J: _j,
		})
	}

	return result
}

// Unzip2 将 Tuple2 切片拆解为两个独立切片，恢复 Zip 之前的形式。
// 参数：tuples 为要拆解的 Tuple2[A, B] 切片。
// 返回：两个切片，分别包含所有元组的 A、B 字段；按输入顺序追加。
func Unzip2[A any, B any](tuples []Tuple2[A, B]) ([]A, []B) {
	size := len(tuples)
	r1 := make([]A, 0, size)
	r2 := make([]B, 0, size)

	for _, tuple := range tuples {
		r1 = append(r1, tuple.A)
		r2 = append(r2, tuple.B)
	}

	return r1, r2
}

// Unzip3 将 Tuple3 切片拆解为三个独立切片，恢复 Zip 之前的形式。
// 参数：tuples 为要拆解的 Tuple3[A, B, C] 切片。
// 返回：三个切片，分别包含所有元组的 A、B、C 字段；按输入顺序追加。
func Unzip3[A any, B any, C any](tuples []Tuple3[A, B, C]) ([]A, []B, []C) {
	size := len(tuples)
	r1 := make([]A, 0, size)
	r2 := make([]B, 0, size)
	r3 := make([]C, 0, size)

	for _, tuple := range tuples {
		r1 = append(r1, tuple.A)
		r2 = append(r2, tuple.B)
		r3 = append(r3, tuple.C)
	}

	return r1, r2, r3
}

// Unzip4 将 Tuple4 切片拆解为四个独立切片，恢复 Zip 之前的形式。
// 参数：tuples 为要拆解的 Tuple4[A, B, C, D] 切片。
// 返回：四个切片，分别包含所有元组的 A、B、C、D 字段；按输入顺序追加。
func Unzip4[A any, B any, C any, D any](tuples []Tuple4[A, B, C, D]) ([]A, []B, []C, []D) {
	size := len(tuples)
	r1 := make([]A, 0, size)
	r2 := make([]B, 0, size)
	r3 := make([]C, 0, size)
	r4 := make([]D, 0, size)

	for _, tuple := range tuples {
		r1 = append(r1, tuple.A)
		r2 = append(r2, tuple.B)
		r3 = append(r3, tuple.C)
		r4 = append(r4, tuple.D)
	}

	return r1, r2, r3, r4
}

// Unzip5 将 Tuple5 切片拆解为五个独立切片，恢复 Zip 之前的形式。
// 参数：tuples 为要拆解的 Tuple5[A, B, C, D, E] 切片。
// 返回：五个切片，分别包含所有元组的 A、B、C、D、E 字段；按输入顺序追加。
func Unzip5[A any, B any, C any, D any, E any](tuples []Tuple5[A, B, C, D, E]) ([]A, []B, []C, []D, []E) {
	size := len(tuples)
	r1 := make([]A, 0, size)
	r2 := make([]B, 0, size)
	r3 := make([]C, 0, size)
	r4 := make([]D, 0, size)
	r5 := make([]E, 0, size)

	for _, tuple := range tuples {
		r1 = append(r1, tuple.A)
		r2 = append(r2, tuple.B)
		r3 = append(r3, tuple.C)
		r4 = append(r4, tuple.D)
		r5 = append(r5, tuple.E)
	}

	return r1, r2, r3, r4, r5
}

// Unzip6 将 Tuple6 切片拆解为六个独立切片，恢复 Zip 之前的形式。
// 参数：tuples 为要拆解的 Tuple6[A, B, C, D, E, F] 切片。
// 返回：六个切片，分别包含所有元组的 A、B、C、D、E、F 字段；按输入顺序追加。
func Unzip6[A any, B any, C any, D any, E any, F any](tuples []Tuple6[A, B, C, D, E, F]) ([]A, []B, []C, []D, []E, []F) {
	size := len(tuples)
	r1 := make([]A, 0, size)
	r2 := make([]B, 0, size)
	r3 := make([]C, 0, size)
	r4 := make([]D, 0, size)
	r5 := make([]E, 0, size)
	r6 := make([]F, 0, size)

	for _, tuple := range tuples {
		r1 = append(r1, tuple.A)
		r2 = append(r2, tuple.B)
		r3 = append(r3, tuple.C)
		r4 = append(r4, tuple.D)
		r5 = append(r5, tuple.E)
		r6 = append(r6, tuple.F)
	}

	return r1, r2, r3, r4, r5, r6
}

// Unzip7 将 Tuple7 切片拆解为七个独立切片，恢复 Zip 之前的形式。
// 参数：tuples 为要拆解的 Tuple7[A, B, C, D, E, F, G] 切片。
// 返回：七个切片，分别包含所有元组的 A、B、C、D、E、F、G 字段；按输入顺序追加。
func Unzip7[A any, B any, C any, D any, E any, F any, G any](tuples []Tuple7[A, B, C, D, E, F, G]) ([]A, []B, []C, []D, []E, []F, []G) {
	size := len(tuples)
	r1 := make([]A, 0, size)
	r2 := make([]B, 0, size)
	r3 := make([]C, 0, size)
	r4 := make([]D, 0, size)
	r5 := make([]E, 0, size)
	r6 := make([]F, 0, size)
	r7 := make([]G, 0, size)

	for _, tuple := range tuples {
		r1 = append(r1, tuple.A)
		r2 = append(r2, tuple.B)
		r3 = append(r3, tuple.C)
		r4 = append(r4, tuple.D)
		r5 = append(r5, tuple.E)
		r6 = append(r6, tuple.F)
		r7 = append(r7, tuple.G)
	}

	return r1, r2, r3, r4, r5, r6, r7
}

// Unzip8 将 Tuple8 切片拆解为八个独立切片，恢复 Zip 之前的形式。
// 参数：tuples 为要拆解的 Tuple8[A, B, C, D, E, F, G, H] 切片。
// 返回：八个切片，分别包含所有元组的 A、B、C、D、E、F、G、H 字段；按输入顺序追加。
func Unzip8[A any, B any, C any, D any, E any, F any, G any, H any](tuples []Tuple8[A, B, C, D, E, F, G, H]) ([]A, []B, []C, []D, []E, []F, []G, []H) {
	size := len(tuples)
	r1 := make([]A, 0, size)
	r2 := make([]B, 0, size)
	r3 := make([]C, 0, size)
	r4 := make([]D, 0, size)
	r5 := make([]E, 0, size)
	r6 := make([]F, 0, size)
	r7 := make([]G, 0, size)
	r8 := make([]H, 0, size)

	for _, tuple := range tuples {
		r1 = append(r1, tuple.A)
		r2 = append(r2, tuple.B)
		r3 = append(r3, tuple.C)
		r4 = append(r4, tuple.D)
		r5 = append(r5, tuple.E)
		r6 = append(r6, tuple.F)
		r7 = append(r7, tuple.G)
		r8 = append(r8, tuple.H)
	}

	return r1, r2, r3, r4, r5, r6, r7, r8
}

// Unzip9 将 Tuple9 切片拆解为九个独立切片，恢复 Zip 之前的形式。
// 参数：tuples 为要拆解的 Tuple9[A, B, C, D, E, F, G, H, I] 切片。
// 返回：九个切片，分别包含所有元组的 A、B、C、D、E、F、G、H、I 字段；按输入顺序追加。
func Unzip9[A any, B any, C any, D any, E any, F any, G any, H any, I any](tuples []Tuple9[A, B, C, D, E, F, G, H, I]) ([]A, []B, []C, []D, []E, []F, []G, []H, []I) {
	size := len(tuples)
	r1 := make([]A, 0, size)
	r2 := make([]B, 0, size)
	r3 := make([]C, 0, size)
	r4 := make([]D, 0, size)
	r5 := make([]E, 0, size)
	r6 := make([]F, 0, size)
	r7 := make([]G, 0, size)
	r8 := make([]H, 0, size)
	r9 := make([]I, 0, size)

	for _, tuple := range tuples {
		r1 = append(r1, tuple.A)
		r2 = append(r2, tuple.B)
		r3 = append(r3, tuple.C)
		r4 = append(r4, tuple.D)
		r5 = append(r5, tuple.E)
		r6 = append(r6, tuple.F)
		r7 = append(r7, tuple.G)
		r8 = append(r8, tuple.H)
		r9 = append(r9, tuple.I)
	}

	return r1, r2, r3, r4, r5, r6, r7, r8, r9
}

// Unzip10 将一个包含Tuple10的切片解压为十个切片，分别对应Tuple10中的十个元素类型
// A, B, C, D, E, F, G, H, I, J 是Tuple10中元素的类型参数
// tuples 是包含Tuple10的切片
// 返回值为十个切片，分别对应Tuple10中的十个元素类型
func Unzip10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](tuples []Tuple10[A, B, C, D, E, F, G, H, I, J]) ([]A, []B, []C, []D, []E, []F, []G, []H, []I, []J) {
	size := len(tuples)
	r1 := make([]A, 0, size)
	r2 := make([]B, 0, size)
	r3 := make([]C, 0, size)
	r4 := make([]D, 0, size)
	r5 := make([]E, 0, size)
	r6 := make([]F, 0, size)
	r7 := make([]G, 0, size)
	r8 := make([]H, 0, size)
	r9 := make([]I, 0, size)
	r10 := make([]J, 0, size)

	for _, tuple := range tuples {
		r1 = append(r1, tuple.A)
		r2 = append(r2, tuple.B)
		r3 = append(r3, tuple.C)
		r4 = append(r4, tuple.D)
		r5 = append(r5, tuple.E)
		r6 = append(r6, tuple.F)
		r7 = append(r7, tuple.G)
		r8 = append(r8, tuple.H)
		r9 = append(r9, tuple.I)
		r10 = append(r10, tuple.J)
	}

	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}
