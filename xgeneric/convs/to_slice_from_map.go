package convs

import (
	"fmt"
	"reflect"
)

//
// 将 map 的键和值转换为切片（顺序随机）。
// 例如：map[string]int{"a":1,"b":2, "c":3} 会转换为 []string{"a", "c", "b"} 与 []int{1, 3, 2}。
//

// MapKeys 返回包含 map 中所有键的切片。
// 返回的键顺序随机（取决于 map 的迭代顺序）。
// 也可以使用标准库 golang.org/x/exp/maps#Keys。
// 参数：m 为输入 map。
// 返回：以任意顺序排列的所有键组成的切片。
func MapKeys[K comparable, V any](m map[K]V) []K {
	s := make([]K, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

// MapVals 返回包含 map 中所有值的切片。
// 返回的值顺序随机（取决于 map 的迭代顺序）。
// 也可以使用标准库 golang.org/x/exp/maps#Values。
// 参数：m 为输入 map。
// 返回：以任意顺序排列的所有值组成的切片。
func MapVals[K comparable, V any](m map[K]V) []V {
	s := make([]V, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

// MapKeysVals 返回 map 中所有键和值的两个切片。
// 返回的键与值均按随机顺序排列。
// 参数：m 为输入 map。
// 返回：包含所有键的 []K 与包含所有值的 []V，二者按各自顺序对齐。
func MapKeysVals[K comparable, V any](m map[K]V) ([]K, []V) {
	ks, vs := make([]K, 0, len(m)), make([]V, 0, len(m))
	for k, v := range m {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	return ks, vs
}

// Map2Slice 将 map 的键和值转换为两个切片（顺序随机，失败时返回 nil, nil）。
// 参数：a 为要转换的 map（任意具体类型）。
// 返回：键切片与值切片（均为 any 类型）；底层调用 Map2SliceE，失败时返回 nil, nil（不返回错误）。
func Map2Slice(a any) (ks any, vs any) {
	ks, vs, _ = Map2SliceE(a)
	return
}

// Map2SliceE 将 map 的键和值转换为两个切片（顺序随机），失败时返回错误。
// 参数：a 为要转换的 map（任意具体类型，使用反射处理）。
// 返回：键切片、值切片以及错误信息；a 为 nil 或非 map 类型时返回 error。
func Map2SliceE(a any) (ks any, vs any, err error) {
	// Check param.
	if a == nil {
		return nil, nil, fmt.Errorf("unable to converts %#v of type %T to slice", a, a)
	}
	t := reflect.TypeOf(a)
	if t.Kind() != reflect.Map {
		err = fmt.Errorf("the input %#v of type %T isn't a map", a, a)
		return
	}

	// Execute the conversion.
	m := reflect.ValueOf(a)
	l := m.Len()
	keys := m.MapKeys()
	ksT, vsT := reflect.SliceOf(t.Key()), reflect.SliceOf(t.Elem())
	ksV, vsV := reflect.MakeSlice(ksT, 0, l), reflect.MakeSlice(vsT, 0, l)
	for _, k := range keys {
		ksV = reflect.Append(ksV, k)
		vsV = reflect.Append(vsV, m.MapIndex(k))
	}
	return ksV.Interface(), vsV.Interface(), nil
}
