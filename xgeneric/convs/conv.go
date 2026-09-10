// Package conv 提供与类型转换相关的工具函数，
// 例如使用 ToAny[int](any) 转为 int、将切片或数组转为集合类型（map[bool]struct{}）、
// 将 map 的键或值转为切片等。
package convs

import (
	"fmt"
)

// ToAny 将任意值 a 转换为类型 T（转换失败时静默返回 T 的零值）。
// 参数：a 为要转换的任意类型值。
// 返回：转换后的 T 值；若转换失败返回 T 的零值（不返回 error，请使用 ToAnyE 接收错误）。
func ToAny[T any](a any) T {
	v, _ := ToAnyE[T](a)
	return v
}

// ToAnyE 将任意值 a 转换为类型 T，并在转换失败时返回错误。
// 参数：a 为要转换的任意类型值。
// 返回：转换后的 T 值与 nil 错误；若 T 不在支持的类型列表中则返回 error（信息为 "the type %T isn't supported"）。
func ToAnyE[T any](a any) (T, error) {
	var t T
	switch any(t).(type) {
	case bool:
		v, err := ToBoolE(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int:
		v, err := ToIntE(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int8:
		v, err := ToInt8E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int16:
		v, err := ToInt16E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int32:
		v, err := ToInt32E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case int64:
		v, err := ToInt64E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint:
		v, err := ToUintE(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint8:
		v, err := ToUint8E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint16:
		v, err := ToUint16E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint32:
		v, err := ToUint32E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case uint64:
		v, err := ToUint64E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case float32:
		v, err := ToFloat32E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case float64:
		v, err := ToFloat64E(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	case string:
		v, err := ToStringE(a)
		if err != nil {
			return t, err
		}
		t = any(v).(T)
	default:
		return t, fmt.Errorf("the type %T isn't supported", t)
	}
	return t, nil
}
