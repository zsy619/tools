package xreflect

import (
	"bytes"
	"reflect"
)

// byteSliceType 缓存 []byte 类型的 reflect.Type，
// 用于在 DeepEquals.Matches 中快速判定目标值是否为字节切片。
var byteSliceType reflect.Type = reflect.TypeOf([]byte{})

// NewDeepEquals 创建一个基于"深度相等"的匹配器。
// 深度相等的判定遵循 reflect 包的 reflect.DeepEqual 语义。
// 该匹配器要求被比较的值 c 与 x 具有完全相同的类型，否则一定不匹配。
// 对于 []byte 类型，会特殊处理以走更高效的 bytes.Equal 路径。
func NewDeepEquals(x interface{}) *DeepEquals {
	return &DeepEquals{x}
}

// DeepEquals 是 NewDeepEquals 返回的匹配器类型，封装了用于比较的基准值 x。
type DeepEquals struct {
	x interface{}
}

// Matches 判断 c 是否与 m.x 深度相等。
// 当 c 与 m.x 的反射类型不一致时直接返回 false；
// 当两者均为非 nil 的 []byte 时使用 bytes.Equal 比较以提升性能；
// 其他情形委托给 reflect.DeepEqual；
// 当 c 是 nil 切片时，reflect.DeepEqual 通常会判为不等，函数会返回 false。
func (m *DeepEquals) Matches(c interface{}) bool {
	// Make sure the types match.
	ct := reflect.TypeOf(c)
	xt := reflect.TypeOf(m.x)

	if ct != xt {
		return false
	}

	// Special case: handle byte slices more efficiently.
	cValue := reflect.ValueOf(c)
	xValue := reflect.ValueOf(m.x)

	if ct == byteSliceType && !cValue.IsNil() && !xValue.IsNil() {
		xBytes := m.x.([]byte)
		cBytes := c.([]byte)
		return bytes.Equal(cBytes, xBytes)
	}

	// Defer to the reflect package.
	if reflect.DeepEqual(m.x, c) {
		return true
	}

	// Special case: if the comparison failed because c is the nil slice, given
	// an indication of this (since its value is printed as "[]").
	if cValue.Kind() == reflect.Slice && cValue.IsNil() {
		return false
	}

	return false
}
