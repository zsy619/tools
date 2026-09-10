// Package xunsafe 提供基于 unsafe 包的低层类型转换工具。
//
// 这些函数把 string/byte/任意结构体之间的零拷贝转换、指针强转、
// 指针偏移等操作集中到一个包中，方便上层业务在确知数据生命周期
// 安全的前提下，避开深拷贝与反射带来的开销。
//
// **警告**：本包所有函数均依赖 Go 内存模型与运行时对指针语义的
// 隐式约束，请仅在你完全理解返回值的生命周期后再使用。
package xunsafe

import (
	"reflect"
	"unsafe"
)

const (
	// Max_Size 用于把任意指针转成超大定长数组时的最大长度。
	// 设为 2<<31 (即 4 GiB)，足以覆盖任何 Go 对象的 sizeof。
	Max_Size = 2 << 31
)

// MappingToArray 把任意类型的值零拷贝转成字节切片。
//
// 实现原理：把 obj 的内存布局直接视为 [Max_Size]byte 并取前
// unsafe.Sizeof(obj) 个字节。
//
// **警告**：
//   - 返回切片与 obj 的内存重叠，仅当 obj 在调用方控制下不被修改
//     且生命周期内有效时才安全；
//   - 不要用于含指针或可变状态的字段，否则会出现意外别名。
func MappingToArray[T any](obj T) []byte {
	arr := (*[Max_Size]byte)(unsafe.Pointer(&obj))
	size := unsafe.Sizeof(obj)
	return arr[:size]
}

// ArrayMapping 是 MappingToArray 的逆操作。
//
// 把字节切片按 T 类型的内存布局重新解释为指向 T 的指针；常见用法是
// 与 MappingToArray 配合实现结构体与字节的零拷贝互转。
//
// **警告**：
//   - bytes 必须确实对应 T 类型的内存表示（大小、对齐一致），否则
//     解引用会导致崩溃或读到错误数据；
//   - 返回指针与 bytes 共享内存，调用方需自行保证生命周期安全。
func ArrayMapping[T any](bytes []byte) *T {
	ret := (*T)(unsafe.Pointer(&bytes[0]))
	return ret
}

// Slice 在 ptr 指向的内存上构造一个长度为 size 的 []R 切片。
//
// 与 MappingToArray 类似，本函数不会拷贝内存，调用方需保证 ptr 指向
// 的空间至少 size*sizeof(R) 字节有效。
//
// 注意：T 与 R 仅用于推导元素大小，必须在内存布局上严格一致，否则
// 会读到错误数据。
func Slice[T, R any](ptr *T, size int) []R {
	ret := (*[Max_Size]R)(unsafe.Pointer(ptr))[:size]
	return ret
}

// As 将 E* 强转成 R*。
//
// 等价于 (*R)(unsafe.Pointer(ptr))，但类型推导更直观。
//
// **警告**：调用方必须确保 E 与 R 内存布局兼容（如不同结构体的零拷贝
// 重解释），否则解引用结果无意义。
func As[E, R any](ptr *E) *R {
	return (*R)(unsafe.Pointer(ptr))
}

// OffsetValue 把 ptr 视为 T 类型数组，按 sizeof(R) * offsetN 的偏移
// 返回 R*，相当于数组按下标的指针访问。
//
// 用途：在结构体中按字段偏移访问子结构。
//
// **警告**：调用方需保证 ptr 之后至少 (offsetN+1) 个 R 大小的空间
// 可用，否则会越界读取到非法内存。
func OffsetValue[T, R any](ptr *T, offsetN int) *R {
	var r R
	return (*R)(unsafeIndex(unsafe.Pointer(ptr), 0, ValueSizeof(r), offsetN))
}

// unsafeIndex 是 OffsetValue 的低层助手：
// 在 base + offset 之后，按 elemsz 步长跳过 n 个元素，返回新指针。
//
// 该函数不参与类型检查，仅供同包内部 OffsetValue 使用。
func unsafeIndex(base unsafe.Pointer, offset uintptr, elemsz uintptr, n int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(base) + offset + uintptr(n)*elemsz)
}

// ValueSizeof 返回值的内存大小。
//
// 与 unsafe.Sizeof 不同，本函数对指针 v 返回其指向类型的大小
// （即 *E 的 sizeof 等价于 E 的 sizeof），便于在不确定入参是值还是
// 指针的场景下统一行为。
func ValueSizeof(v any) uintptr {
	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Pointer {
		return typ.Elem().Size()
	}

	return typ.Size()
}

// StringToSlice 把 string 零拷贝转成 []byte。
//
// 实现原理：string 与 []byte 在运行时都表示为 (Data, Len) 的形式，
// 通过 unsafe 直接复用底层字节数组并把切片头部第二字段设置为 len(value)
// 表示 Cap（从而与 Len 对齐，避免 append 误以为有额外空间）。
//
// **警告**：返回的 []byte 与原 string 共享同一段只读内存，
//   - 不可对返回切片做任何写操作，否则会触发运行时的 fatal 错误
//     "slice bounds out of range" 或者写入只读段崩溃；
//   - 仅适用于 string 在整个使用期间不会被修改的场景。
func StringToSlice(value string) []byte {
	// create a new []byte
	var ret []byte

	// 把string的引用指向 ret的空间
	*(*string)(unsafe.Pointer(&ret)) = value

	// 设置slice的Cap值 ，用unsafe.Add操作，执行偏移操作 16个字节
	*(*int)(unsafe.Add(unsafe.Pointer(&ret), uintptr(8)*2)) = len(value)

	return ret
}

// SliceToString 把 []byte 零拷贝转成 string。
//
// b 为 nil 时返回空字符串，否则直接复用 b 的底层字节构造 string。
//
// **警告**：返回 string 与 b 共享同一段内存；如果 b 在使用期间被
// 修改，string 内容也会随之变化（这是合法的但很容易出 bug）。
// 仅用于短生命周期的 b 或确保不会被修改的场景。
func SliceToString(b []byte) string {
	if b == nil {
		return ""
	}

	// just share Slice's Data and Len content
	return *(*string)(unsafe.Pointer(&b))
}

// Copy 返回 v 所指值的深拷贝（Go 默认赋值语义）。
//
// 对于含指针字段的类型，本函数只复制指针引用（浅拷贝）；如果需要
// 完全独立的新对象，请自行实现序列化/反序列化。
//
// 注意：如果 E 包含 sync.Mutex 等「不可复制」字段，go vet 会报警
// "copies lock"，应避免对含锁的结构体使用本函数。
func Copy[E any](v *E) *E {
	n := *v
	return &n
}
