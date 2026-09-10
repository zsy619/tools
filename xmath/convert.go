package xmath

import (
	"fmt"
	"reflect"
)

// IntToBool 将整数转换为 bool（仅在 value == 1 时返回 true）。
//
// 参数：
//   - value: 整数输入。
//
// 返回值：value == 1 时返回 true，其他值返回 false。
func IntToBool(value int) bool {
	if value == 1 {
		return true
	} else {
		return false
	}
}

// convert any numeric value to int64
// ToInt64 将任意数值类型转换为 int64。
//
// 参数：
//   - value: 任意 int*/uint* 类型的值。
//
// 返回值：
//   - d: 转换后的 int64 数值。
//   - err: 当 value 不是数值类型时返回错误。
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}
