package convs

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

//
// Desc: 将任意类型的值转换为指定类型的切片。
//

// ToStrSlice 将任意类型转换为 []string（转换失败时返回 nil）。
// 例如：[]int{1, 2, 3} 会转换为 []string{"1", "2", "3"}。
func ToStrSlice(i any) []string {
	v, _ := ToStrSliceE(i)
	return v
}

// ToStrSliceE 将任意类型转换为 []string，并在转换失败时返回错误。
func ToStrSliceE(i any) ([]string, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []string", i, i)
	}

	switch v := i.(type) {
	case []string:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		s := make([]string, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToStringE(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			s[j] = v
		}
		return s, nil
	}

	// If i is a single value.
	switch v := i.(type) {
	case string:
		return strings.Fields(v), nil
	default:
		s, err := ToStringE(v)
		if err != nil {
			return nil, err
		}
		return []string{s}, nil
	}
}

// ToIntSlice 将任意类型转换为 []int（转换失败时返回 nil）。
// 例如：[]string{"1", "2", "3"} 会转换为 []int{1, 2, 3}。
func ToIntSlice(i any) []int {
	v, _ := ToIntSliceE(i)
	return v
}

// ToIntSliceE 将任意类型转换为 []int，并在转换失败时返回错误。
func ToIntSliceE(i any) ([]int, error) {
	if i == nil {
		return []int{}, nil
	}

	switch v := i.(type) {
	case []int:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToIntE(s.Index(j).Interface())
			if err != nil {
				return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
	}
}

// ToInt8Slice 将任意类型转换为 []int8（转换失败时返回 nil）。
// 例如：[]string{"1", "2", "3"} 会转换为 []int8{1, 2, 3}。
func ToInt8Slice(i any) []int8 {
	v, _ := ToInt8SliceE(i)
	return v
}

// ToInt8SliceE 将任意类型转换为 []int8，并在转换失败时返回错误。
func ToInt8SliceE(i any) ([]int8, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int8", i, i)
	}

	switch v := i.(type) {
	case []int8:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		i8s := make([]int8, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToInt8E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			i8s[j] = v
		}
		return i8s, nil
	}

	// If i is a single value.
	v, err := ToInt8E(i)
	if err != nil {
		return nil, err
	}
	return []int8{v}, nil
}

// ToInt16Slice 将任意类型转换为 []int16（转换失败时返回 nil）。
// 例如：[]string{"1", "2", "3"} 会转换为 []int16{1, 2, 3}。
func ToInt16Slice(i any) []int16 {
	v, _ := ToInt16SliceE(i)
	return v
}

// ToInt16SliceE 将任意类型转换为 []int16，并在转换失败时返回错误。
func ToInt16SliceE(i any) ([]int16, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int16", i, i)
	}

	switch v := i.(type) {
	case []int16:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		i16s := make([]int16, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToInt16E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			i16s[j] = v
		}
		return i16s, nil
	}

	// If i is a single value.
	v, err := ToInt16E(i)
	if err != nil {
		return nil, err
	}
	return []int16{v}, nil
}

// ToInt32Slice 将任意类型转换为 []int32（转换失败时返回 nil）。
// 例如：[]string{"1", "2", "3"} 会转换为 []int32{1, 2, 3}。
func ToInt32Slice(i any) []int32 {
	v, _ := ToInt32SliceE(i)
	return v
}

// ToInt32SliceE 将任意类型转换为 []int32，并在转换失败时返回错误。
func ToInt32SliceE(i any) ([]int32, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int32", i, i)
	}

	switch v := i.(type) {
	case []int32:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		i32s := make([]int32, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToInt32E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			i32s[j] = v
		}
		return i32s, nil
	}

	// If i is a single value.
	v, err := ToInt32E(i)
	if err != nil {
		return nil, err
	}
	return []int32{v}, nil
}

// ToInt64Slice 将任意类型转换为 []int64（转换失败时返回 nil）。
// 例如：[]string{"1", "2", "3"} 会转换为 []int64{1, 2, 3}。
func ToInt64Slice(i any) []int64 {
	v, _ := ToInt64SliceE(i)
	return v
}

// ToInt64SliceE 将任意类型转换为 []int64，并在转换失败时返回错误。
func ToInt64SliceE(i any) ([]int64, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
	}

	switch v := i.(type) {
	case []int64:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		i64s := make([]int64, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToInt64E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			i64s[j] = v
		}
		return i64s, nil
	}

	// If i is a single value.
	v, err := ToInt64E(i)
	if err != nil {
		return nil, err
	}
	return []int64{v}, nil
}

// ToUintSlice 将任意类型转换为 []uint（转换失败时返回 nil）。
// 例如：[]string{"1", "2", "3"} 会转换为 []uint{1, 2, 3}。
func ToUintSlice(i any) []uint {
	v, _ := ToUintSliceE(i)
	return v
}

// ToUintSliceE 将任意类型转换为 []uint，并在转换失败时返回错误。
func ToUintSliceE(i any) ([]uint, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint", i, i)
	}

	switch v := i.(type) {
	case []uint:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToUintE(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// If i is a single value.
	v, err := ToUintE(i)
	if err != nil {
		return nil, err
	}
	return []uint{v}, nil
}

// ToUint8Slice 将任意类型转换为 []uint8（转换失败时返回 nil）。
// 例如：[]string{"1", "2", "3"} 会转换为 []uint8{1, 2, 3}。
func ToUint8Slice(i any) []uint8 {
	v, _ := ToUint8SliceE(i)
	return v
}

// ToUint8SliceE 将任意类型转换为 []uint8，并在转换失败时返回错误。
func ToUint8SliceE(i any) ([]uint8, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint8", i, i)
	}

	switch v := i.(type) {
	case []uint8:
		return v, nil
	}

	// if i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint8, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToUint8E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// If i is a single value.
	u, err := ToUint8E(i)
	if err != nil {
		return nil, err
	}
	return []uint8{u}, nil
}

// ToUint16Slice 将任意类型转换为 []uint16（转换失败时返回 nil）。
// 例如：[]string{"1", "2", "3"} 会转换为 []uint16{1, 2, 3}。
func ToUint16Slice(i any) []uint16 {
	v, _ := ToUint16SliceE(i)
	return v
}

// ToUint16SliceE 将任意类型转换为 []uint16，并在转换失败时返回错误。
func ToUint16SliceE(i any) ([]uint16, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint16", i, i)
	}

	switch v := i.(type) {
	case []uint16:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint16, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToUint16E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// If i is a single value.
	u, err := ToUint16E(i)
	if err != nil {
		return nil, err
	}
	return []uint16{u}, nil
}

// ToUint32Slice 将任意类型转换为 []uint32（转换失败时返回 nil）。
// 例如：[]string{"1", "2", "3"} 会转换为 []uint32{1, 2, 3}。
func ToUint32Slice(i any) []uint32 {
	v, _ := ToUint32SliceE(i)
	return v
}

// ToUint32SliceE 将任意类型转换为 []uint32，并在转换失败时返回错误。
func ToUint32SliceE(i any) ([]uint32, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint32", i, i)
	}

	switch v := i.(type) {
	case []uint32:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint32, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToUint32E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// If i is a single value.
	u, err := ToUint32E(i)
	if err != nil {
		return nil, err
	}
	return []uint32{u}, nil
}

// ToUint64Slice 将任意类型转换为 []uint64（转换失败时返回 nil）。
// 例如：[]string{"1", "2", "3"} 会转换为 []uint64{1, 2, 3}。
func ToUint64Slice(i any) []uint64 {
	v, _ := ToUint64SliceE(i)
	return v
}

// ToUint64SliceE 将任意类型转换为 []uint64，并在转换失败时返回错误。
func ToUint64SliceE(i any) ([]uint64, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []uint64", i, i)
	}

	switch v := i.(type) {
	case []uint64:
		return v, nil
	}

	// If i is a slice or array.
	kind := reflect.TypeOf(i).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		sl := reflect.ValueOf(i)
		u := make([]uint64, sl.Len())
		for j := 0; j < sl.Len(); j++ {
			v, err := ToUint64E(sl.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			u[j] = v
		}
		return u, nil
	}

	// If i is a single value.
	u, err := ToUint64E(i)
	if err != nil {
		return nil, err
	}
	return []uint64{u}, nil
}

// ToByteSlice 将任意类型转换为 []byte（转换失败时返回 nil）。
// 例如：[]string{"1", "2", "3"} 会转换为 []byte{1, 2, 3}。
func ToByteSlice(i any) []byte {
	return ToUint8Slice(i)
}

// ToByteSliceE 将任意类型转换为 []byte，并在转换失败时返回错误。
func ToByteSliceE(i any) ([]byte, error) {
	return ToUint8SliceE(i)
}

// ToBoolSlice 将任意类型转换为 []bool（转换失败时返回 nil）。
func ToBoolSlice(a any) []bool {
	v, _ := ToBoolSliceE(a)
	return v
}

// ToBoolSliceE 将任意类型转换为 []bool，并在转换失败时返回错误。
func ToBoolSliceE(i any) ([]bool, error) {
	if i == nil {
		return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
	}

	switch v := i.(type) {
	case []bool:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]bool, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToBoolE(s.Index(j).Interface())
			if err != nil {
				return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
	}
}

// ToDurationSlice 将任意类型转换为 []time.Duration（转换失败时返回 nil）。
func ToDurationSlice(i any) []time.Duration {
	v, _ := ToDurationSliceE(i)
	return v
}

// ToDurationSliceE 将任意类型转换为 []time.Duration，并在转换失败时返回错误。
func ToDurationSliceE(i any) ([]time.Duration, error) {
	if i == nil {
		return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
	}

	switch v := i.(type) {
	case []time.Duration:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]time.Duration, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToDurationE(s.Index(j).Interface())
			if err != nil {
				return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
	}
}
