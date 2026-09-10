package xreflect

import "reflect"

// IsByteType 判断 v 的类型是否为 byte（即 reflect.Kind == Uint8）。
// 当 v 为 nil 时，reflect.TypeOf 会返回 nil，调用 Kind 会 panic，调用方需自行保证 v 非 nil。
func IsByteType(v any) bool {
	typ := reflect.TypeOf(v)
	return typ.Kind() == reflect.Uint8
}

// TypeOfName 返回 v 的类型名称（reflect.Type.Name）。
// 对于指针、切片、Map、通道等非命名类型，可能返回空字符串。
// 当 v 为 nil 接口值时，reflect.TypeOf 返回 nil，调用 Name 会 panic。
func TypeOfName[E any](v E) string {
	t := reflect.TypeOf(v)
	return t.Name()
}

// ConvertibleTo 判断 v 的类型是否可以转换为 checked 的类型。
// 通过 reflect.Value.Convert 的可转换性规则判定：例如数值之间可转换、实现某接口的类型可转换为该接口等。
// v 或 checked 为 nil 接口时，reflect.TypeOf 返回 nil，会导致 panic，需保证两者均非 nil 接口。
func ConvertibleTo[E any](v E, checked any) bool {
	t := reflect.TypeOf(v)
	return t.ConvertibleTo(reflect.TypeOf(checked))
}

// AssignIfConvertibleTo 当 checked 可以转换为 v 的类型时，将 checked 转换后返回，并附带 true。
// 不可转换时原样返回 v 和 false。
// 注意：内部对转换结果做了 .(E) 的类型断言，若 checked 的运行时类型与泛型参数 E 的实例类型不一致会 panic。
func AssignIfConvertibleTo[E any](v E, checked any) (E, bool) {
	if ConvertibleTo(v, checked) {
		va := reflect.ValueOf(checked)
		newVa := va.Convert(reflect.TypeOf(v))
		r := newVa.Interface().(E)
		return r, true
	}

	return v, false
}
