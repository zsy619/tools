package convs

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Struct2Map 将结构体转换为 map[string]any。
// 例如 struct{I int, S string}{I: 1, S: "a"} 会转为 map["I":1 "S":"a"]。
// 注意：结构体中未导出的字段会被忽略，无法被转换。
// 参数：a 为任意结构体值。
// 返回：以字段名为键、字段值为值的 map[string]any；当 a 不是结构体类型时返回 nil。
func Struct2Map(a any) map[string]any {
	// Check param.
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := reflect.TypeOf(a)
	var m = make(map[string]any)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			m[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	return m
}

// Struct2MapString 将结构体转换为 map[string]string（字段值使用 ToAny[string] 转为字符串）。
// 例如 struct{I int, S string}{I: 1, S: "a"} 会转为 map["I":"1" "S":"a"]。
// 注意：结构体中未导出的字段会被忽略，无法被转换。
// 参数：obj 为任意结构体值。
// 返回：以字段名为键、字段值字符串为值的 map[string]string；当 obj 不是结构体类型时返回 nil。
func Struct2MapString(obj any) map[string]string {
	// Check param.
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := reflect.TypeOf(obj)
	var m = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			m[t.Field(i).Name] = ToAny[string](v.Field(i).Interface())
		}
	}
	return m
}

// ToMapStrStr 将任意类型转换为 map[string]string（转换失败时返回空 map）。
// 参数：i 为要转换的任意类型值。
// 返回：转换得到的 map[string]string；底层调用 ToMapStrStrE，转换失败时返回空 map（不返回错误）。
func ToMapStrStr(i any) map[string]string {
	v, _ := ToMapStrStrE(i)
	return v
}

// ToMapStrStrE 将任意类型转换为 map[string]string 并在失败时返回错误。
// 参数：i 为要转换的任意类型值（支持 map[string]string / map[string]any / map[any]string / map[any]any / JSON 字符串）。
// 返回：转换得到的 map[string]string 与 nil 错误；当 i 为字符串时会尝试 JSON 反序列化；不支持的类型返回 error。
func ToMapStrStrE(i any) (map[string]string, error) {
	var m = map[string]string{}

	switch v := i.(type) {
	case map[string]string:
		return v, nil
	case map[string]any:
		for k, val := range v {
			m[k] = ToAny[string](val)
		}
		return m, nil
	case map[any]string:
		for k, val := range v {
			m[ToAny[string](k)] = val
		}
		return m, nil
	case map[any]any:
		for k, val := range v {
			m[ToAny[string](k)] = ToAny[string](val)
		}
		return m, nil
	case string:
		err := jsonStringToObject(v, &m)
		return m, err
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]string", i, i)
	}
}

// jsonStringToObject 尝试将字符串 s 作为 JSON 反序列化到 v 指向的对象中。
// 参数：s 为 JSON 字符串，v 为目标对象的指针。
// 返回：json.Unmarshal 的错误（成功时为 nil）。
func jsonStringToObject(s string, v any) error {
	data := []byte(s)
	return json.Unmarshal(data, v)
}
