package convs

import (
	"fmt"
	"strconv"

	"github.com/zsy619/tools/xstring"
)

// SplitStrToSlice 按指定分隔符将字符串拆分为 T 类型切片（转换失败时返回 nil）。
// 参数：s 为输入字符串，sep 为分隔符。
// 返回：拆分后元素类型为 T 的切片；底层调用 SplitStrToSliceE[T]，转换失败时返回 nil（不返回错误）。
func SplitStrToSlice[T any](s, sep string) []T {
	v, _ := SplitStrToSliceE[T](s, sep)
	return v
}

// SplitStrToSliceE 按指定分隔符将字符串拆分为 T 类型切片，转换失败时返回错误。
// 注意：该函数基于 Go 1.18 泛型实现，调用时需显式指定元素类型，例如 SplitStrToSliceE[int]("1,2,3", ",")。
// 参数：s 为输入字符串，sep 为分隔符。
// 返回：拆分得到的 []T 与 nil 错误；若 T 不在支持的类型列表中或解析失败则返回 error。
func SplitStrToSliceE[T any](s, sep string) ([]T, error) {
	ss := xstring.Split(s, sep)
	dst := make([]T, len(ss))
	var t T
	for i, v := range ss {
		switch any(t).(type) {
		case string:
			dst[i] = any(v).(T)
		case int:
			v, err := strconv.ParseInt(v, 0, 0)
			if err != nil {
				return nil, err
			}
			dst[i] = any(int(v)).(T)
		case int8:
			v, err := strconv.ParseInt(v, 0, 8)
			if err != nil {
				return nil, err
			}
			dst[i] = any(int8(v)).(T)
		case int16:
			v, err := strconv.ParseInt(v, 0, 16)
			if err != nil {
				return nil, err
			}
			dst[i] = any(int16(v)).(T)
		case int32:
			v, err := strconv.ParseInt(v, 0, 32)
			if err != nil {
				return nil, err
			}
			dst[i] = any(int32(v)).(T)
		case int64:
			v, err := strconv.ParseInt(v, 0, 32)
			if err != nil {
				return nil, err
			}
			dst[i] = any(v).(T)
		case uint:
			v, err := strconv.ParseUint(v, 0, 0)
			if err != nil {
				return nil, err
			}
			dst[i] = any(uint(v)).(T)
		case uint8:
			v, err := strconv.ParseUint(v, 0, 8)
			if err != nil {
				return nil, err
			}
			dst[i] = any(uint8(v)).(T)
		case uint16:
			v, err := strconv.ParseUint(v, 0, 16)
			if err != nil {
				return nil, err
			}
			dst[i] = any(uint16(v)).(T)
		case uint32:
			v, err := strconv.ParseUint(v, 0, 32)
			if err != nil {
				return nil, err
			}
			dst[i] = any(uint32(v)).(T)
		case uint64:
			v, err := strconv.ParseUint(v, 0, 64)
			if err != nil {
				return nil, err
			}
			dst[i] = any(v).(T)
		case float32:
			v, err := strconv.ParseFloat(v, 32)
			if err != nil {
				return nil, err
			}
			dst[i] = any(float32(v)).(T)
		case float64:
			v, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, err
			}
			dst[i] = any(v).(T)
		case bool:
			v, err := strconv.ParseBool(v)
			if err != nil {
				return nil, err
			}
			dst[i] = any(v).(T)
		default:
			return nil, fmt.Errorf("the type %T is not supported", t)
		}
	}
	return dst, nil
}
