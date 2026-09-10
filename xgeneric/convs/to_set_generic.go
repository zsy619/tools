package convs

import (
	"fmt"
	"reflect"
)

// ToSet 将切片或数组转换为 map[T]struct{}，转换失败时返回 nil。
// 参数：i 为任意类型的切片或数组。
// 返回：以 T 为键、struct{}{} 为值的 map；底层调用 ToSetE[T]，转换失败时返回 nil（不返回错误）。
func ToSet[T comparable](i any) map[T]struct{} {
	m, _ := ToSetE[T](i)
	return m
}

// ToSetE 将切片或数组转换为 map[T]struct{}，并在转换失败时返回错误。
// 注意：输入元素的类型不必与映射键类型完全相同，例如 []uint64{1, 2, 3} 既可以转换为
// map[uint64]struct{}{1:struct{}{}, 2:struct{}{}, 3:struct{}{}}，也可以转换为
// map[string]struct{}{"1":struct{}{}, "2":struct{}{}, "3":struct{}{}}。
// 注意：该函数基于 Go 1.18 泛型实现，调用时需显式指定元素类型，例如 ToSetE[int]([]int{1,2,3})。
func ToSetE[T comparable](i any) (map[T]struct{}, error) {
	// Check param.
	if i == nil {
		return nil, fmt.Errorf("the input i is nil")
	}
	t := reflect.TypeOf(i)
	kind := t.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, fmt.Errorf("the type %T of input %#v isn't a slice or array", i, i)
	}

	// Execute the conversion.
	v := reflect.ValueOf(i)
	mapT := reflect.MapOf(t.Elem(), reflect.TypeOf(struct{}{}))
	mapV := reflect.MakeMapWithSize(mapT, v.Len())
	for j := 0; j < v.Len(); j++ {
		mapV.SetMapIndex(v.Index(j), reflect.ValueOf(struct{}{}))
	}
	if v, ok := mapV.Interface().(map[T]struct{}); ok {
		return v, nil
	}
	// Convert the element to the type T.
	set := make(map[T]struct{}, v.Len())
	for _, k := range mapV.MapKeys() {
		v, err := ToAnyE[T](k.Interface())
		if err != nil {
			return nil, err
		}
		set[v] = struct{}{}
	}
	return set, nil
}
