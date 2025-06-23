package xinterface

import (
	"reflect"

	"haedu.gov.cn/tools/xreflect"
)

// Similar to "extend" in JS, only updates fields that are specified and not empty in newData
//
// Both newData and mainObj must be pointers to struct objects
func Update(mainObj any, newData any) bool {
	newDataVal, mainObjVal := reflect.ValueOf(newData).Elem(), reflect.ValueOf(mainObj).Elem()
	fieldCount := newDataVal.NumField()
	changed := false
	for i := 0; i < fieldCount; i++ {
		newField := newDataVal.Field(i)
		// They passed in a value for this field, update our DB user
		if newField.IsValid() && !xreflect.IsEmpty(newField) {
			dbField := mainObjVal.Field(i)
			dbField.Set(newField)
			changed = true
		}
	}
	return changed
}

// IsNil 函数用于判断传入的参数是否为 nil
func IsNil(i any) bool {
	// 使用 defer 和 recover 来捕获 panic，防止程序崩溃
	defer func() {
		recover()
	}()
	// 如果传入的参数为 nil，则直接返回 true
	if i == nil {
		return true
	}
	// 使用 reflect 包获取传入参数的值
	vi := reflect.ValueOf(i)
	// 检查是否为可赋 nil 值的类型
	switch vi.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		// 如果是可赋 nil 值的类型，则调用 IsNil 方法判断是否为 nil
		return vi.IsNil()
	default:
		return false // 基础类型、结构体等不可能为 nil
	}
}
