// Package xinterface 提供一组与 interface{} 动态值、指针判空、字段合并、
// JSON 序列化相关的辅助函数。
//
// 该包的设计目标是补足 Go 在运行时类型判断上的常见模式，例如：
//   - 把结构体的零值字段判定为空，并在更新时只覆盖非空字段；
//   - 安全地判断任意类型是否为 nil（包括 chan / func / map 等）；
//   - 将任意对象转 JSON，或批量转换一组对象。
//
// 由于操作对象多为 interface{}，函数内部大量依赖 reflect，请避免在
// 热路径高频调用。
package xinterface

import (
	"reflect"

	"github.com/zsy619/tools/xreflect"
)

// Update 将 newData 中的非空字段合并到 mainObj 中。
//
// 行为类似 JavaScript 中的 Object.assign / spread：
//   - 只覆盖 newData 中 "非空"（通过 xreflect.IsEmpty 判定）的字段；
//   - 返回值表示 mainObj 是否被实际修改。
//
// 注意：
//   - mainObj 与 newData 必须是指向同结构体（或可赋值兼容结构体）的指针；
//   - 函数内部按字段顺序遍历，字段名不匹配时通过 reflect 同样会失败，
//     建议保持字段名 / 类型完全一致以获得最直观的行为。
func Update(mainObj any, newData any) bool {
	newDataVal, mainObjVal := reflect.ValueOf(newData).Elem(), reflect.ValueOf(mainObj).Elem()
	fieldCount := newDataVal.NumField()
	changed := false
	for i := 0; i < fieldCount; i++ {
		newField := newDataVal.Field(i)
		// 仅当 newData 中该字段合法、非零值时才覆盖。
		if newField.IsValid() && !xreflect.IsEmpty(newField) {
			dbField := mainObjVal.Field(i)
			dbField.Set(newField)
			changed = true
		}
	}
	return changed
}

// IsNil 安全地判断任意类型是否为 nil。
//
// 与直接与 nil 比较不同，本函数可以处理：
//   - 真正的 nil（var x any）;
//   - 指针、切片、map、chan、func、interface 类型的零值 nil。
//
// 对其它类型（数值、布尔、结构体等），本函数始终返回 false，因为它们
// 不可能为 nil。
//
// 函数内部使用 defer recover 来兜底 reflect 在 nil 接口上调用
// reflect.ValueOf(...).Kind() 时可能抛出的 panic，保证调用者总能拿到
// 一个布尔结果。
func IsNil(i any) bool {
	// 使用 defer 和 recover 来捕获 panic，防止程序崩溃。
	defer func() {
		_ = recover()
	}()
	// 真正的 nil 直接返回 true。
	if i == nil {
		return true
	}
	// 使用 reflect 包获取传入参数的值。
	vi := reflect.ValueOf(i)
	// 检查是否为可赋 nil 值的类型。
	switch vi.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		// 可赋 nil 值的类型走 IsNil。
		return vi.IsNil()
	default:
		// 基础类型、结构体等不可能为 nil。
		return false
	}
}
