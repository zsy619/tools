package convs

import (
	"fmt"
	"reflect"
)

//
// 将任意元素类型的切片或数组转为指定类型的映射集合（即 map[K]struct{}）。
// 注意：输入元素的类型不必与映射键类型完全相同，例如 []uint64{1, 2, 3} 可以转换为
// map[uint64]struct{}{1:struct{}{}, 2:struct{}{}, 3:struct{}{}}，也可以按需转为
// map[string]struct{}{"1":struct{}{}, "2":struct{}{}, "3":struct{}{}}。
//

// ToBoolSet 将切片或数组转换为 map[bool]struct{}（转换失败时返回 nil）。
func ToBoolSet(i any) map[bool]struct{} {
	m, _ := ToBoolSetE(i)
	return m
}

// ToBoolSetE 将切片或数组转换为 map[bool]struct{} 并在转换失败时返回错误。
func ToBoolSetE(i any) (map[bool]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[bool]struct{}); ok {
		return v, nil
	}
	dst := make(map[bool]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToBoolE(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToIntSet 将切片或数组转换为 map[int]struct{}（转换失败时返回 nil）。
func ToIntSet(i any) map[int]struct{} {
	m, _ := ToIntSetE(i)
	return m
}

// ToIntSetE 将切片或数组转换为 map[int]struct{} 并在转换失败时返回错误。
func ToIntSetE(i any) (map[int]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int]struct{}); ok {
		return v, nil
	}
	dst := make(map[int]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToIntE(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToInt8Set 将切片或数组转换为 map[int8]struct{}（转换失败时返回 nil）。
func ToInt8Set(i any) map[int8]struct{} {
	m, _ := ToInt8SetE(i)
	return m
}

// ToInt8SetE 将切片或数组转换为 map[int8]struct{} 并在转换失败时返回错误。
func ToInt8SetE(i any) (map[int8]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int8]struct{}); ok {
		return v, nil
	}
	dst := make(map[int8]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToInt8E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToInt16Set 将切片或数组转换为 map[int16]struct{}（转换失败时返回 nil）。
func ToInt16Set(i any) map[int16]struct{} {
	m, _ := ToInt16SetE(i)
	return m
}

// ToInt16SetE 将切片或数组转换为 map[int16]struct{} 并在转换失败时返回错误。
func ToInt16SetE(i any) (map[int16]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int16]struct{}); ok {
		return v, nil
	}
	dst := make(map[int16]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToInt16E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToInt32Set 将切片或数组转换为 map[int32]struct{}（转换失败时返回 nil）。
func ToInt32Set(i any) map[int32]struct{} {
	m, _ := ToInt32SetE(i)
	return m
}

// ToInt32SetE 将切片或数组转换为 map[int32]struct{} 并在转换失败时返回错误。
func ToInt32SetE(i any) (map[int32]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int32]struct{}); ok {
		return v, nil
	}
	dst := make(map[int32]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToInt32E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToInt64Set 将切片或数组转换为 map[int64]struct{}（转换失败时返回 nil）。
func ToInt64Set(i any) map[int64]struct{} {
	m, _ := ToInt64SetE(i)
	return m
}

// ToInt64SetE 将切片或数组转换为 map[int64]struct{} 并在转换失败时返回错误。
func ToInt64SetE(i any) (map[int64]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[int64]struct{}); ok {
		return v, nil
	}
	dst := make(map[int64]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToInt64E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToUintSet 将切片或数组转换为 map[uint]struct{}（转换失败时返回 nil）。
func ToUintSet(i any) map[uint]struct{} {
	m, _ := ToUintSetE(i)
	return m
}

// ToUintSetE 将切片或数组转换为 map[uint]struct{} 并在转换失败时返回错误。
func ToUintSetE(i any) (map[uint]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint]struct{}); ok {
		return v, nil
	}
	dst := make(map[uint]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToUintE(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToUint8Set 将切片或数组转换为 map[uint8]struct{}（转换失败时返回 nil）。
func ToUint8Set(i any) map[uint8]struct{} {
	m, _ := ToUint8SetE(i)
	return m
}

// ToUint8SetE 将切片或数组转换为 map[uint8]struct{} 并在转换失败时返回错误。
func ToUint8SetE(i any) (map[uint8]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint8]struct{}); ok {
		return v, nil
	}
	dst := make(map[uint8]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToUint8E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToUint16Set 将切片或数组转换为 map[uint16]struct{}（转换失败时返回 nil）。
func ToUint16Set(i any) map[uint16]struct{} {
	m, _ := ToUint16SetE(i)
	return m
}

// ToUint16SetE 将切片或数组转换为 map[uint16]struct{} 并在转换失败时返回错误。
func ToUint16SetE(i any) (map[uint16]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint16]struct{}); ok {
		return v, nil
	}
	dst := make(map[uint16]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToUint16E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToUint32Set 将切片或数组转换为 map[uint32]struct{}（转换失败时返回 nil）。
func ToUint32Set(i any) map[uint32]struct{} {
	m, _ := ToUint32SetE(i)
	return m
}

// ToUint32SetE 将切片或数组转换为 map[uint32]struct{} 并在转换失败时返回错误。
func ToUint32SetE(i any) (map[uint32]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint32]struct{}); ok {
		return v, nil
	}
	dst := make(map[uint32]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToUint32E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToUint64Set 将切片或数组转换为 map[uint64]struct{}（转换失败时返回 nil）。
func ToUint64Set(i any) map[uint64]struct{} {
	m, _ := ToUint64SetE(i)
	return m
}

// ToUint64SetE 将切片或数组转换为 map[uint64]struct{} 并在转换失败时返回错误。
func ToUint64SetE(i any) (map[uint64]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[uint64]struct{}); ok {
		return v, nil
	}
	dst := make(map[uint64]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToUint64E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToFloat32Set 将切片或数组转换为 map[float32]struct{}（转换失败时返回 nil）。
func ToFloat32Set(i any) map[float32]struct{} {
	m, _ := ToFloat32SetE(i)
	return m
}

// ToFloat32SetE 将切片或数组转换为 map[float32]struct{}，并在转换失败时返回错误。
func ToFloat32SetE(i any) (map[float32]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[float32]struct{}); ok {
		return v, nil
	}
	dst := make(map[float32]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToFloat32E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToFloat64Set 将切片或数组转换为 map[float64]struct{}（转换失败时返回 nil）。
func ToFloat64Set(i any) map[float64]struct{} {
	m, _ := ToFloat64SetE(i)
	return m
}

// ToFloat64SetE 将切片或数组转换为 map[float64]struct{}，并在转换失败时返回错误。
func ToFloat64SetE(i any) (map[float64]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[float64]struct{}); ok {
		return v, nil
	}
	dst := make(map[float64]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToFloat64E(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// ToStrSet 将切片或数组转换为 map[string]struct{}（转换失败时返回 nil）。
func ToStrSet(i any) map[string]struct{} {
	m, _ := ToStrSetE(i)
	return m
}

// ToStrSetE 将切片或数组转换为 map[string]struct{}，并在转换失败时返回错误。
func ToStrSetE(i any) (map[string]struct{}, error) {
	m, err := toSetE(i)
	if err != nil {
		return nil, err
	}
	if v, ok := m.(map[string]struct{}); ok {
		return v, nil
	}
	dst := make(map[string]struct{}, reflect.ValueOf(m).Len())
	for _, k := range reflect.ValueOf(m).MapKeys() {
		v, err := ToStringE(k.Interface())
		if err != nil {
			return nil, err
		}
		dst[v] = struct{}{}
	}
	return dst, nil
}

// toSetE 将切片或数组转换为 map[any]struct{}，并在校验失败时返回错误。
func toSetE(a any) (any, error) {
	// Check param.
	if a == nil {
		return nil, fmt.Errorf("the input argument is nil")
	}
	t := reflect.TypeOf(a)
	if t.Kind() != reflect.Slice && t.Kind() != reflect.Array {
		return nil, fmt.Errorf("the input %#v of type %T isn't a slice or array", a, a)
	}

	// Execute the conversion.
	v := reflect.ValueOf(a)
	mT := reflect.MapOf(t.Elem(), reflect.TypeOf(struct{}{}))
	mV := reflect.MakeMapWithSize(mT, v.Len())
	for i := 0; i < v.Len(); i++ {
		mV.SetMapIndex(v.Index(i), reflect.ValueOf(struct{}{}))
	}
	return mV.Interface(), nil
}
