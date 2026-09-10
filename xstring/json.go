package xstring

import (
	"fmt"
	"unsafe"

	"github.com/zsy619/tools/xjson"
)

// JsonToObject 将 JSON 字符串解析到 result 指向的目标对象。
//
// 内部使用 xjson.Unmarshal，相比 encoding/json 在大量小对象场景下
// 通常有更好的性能与更友好的错误信息。
//
// 参数：
//   - meta：JSON 字符串；
//   - result：目标对象的指针（必须非 nil，且指向的类型与 JSON 内容
//     兼容）。
//
// 返回 error：解析失败时返回错误，nil 表示成功。
func JsonToObject(meta string, result any) error {
	return xjson.Unmarshal(StringToBytes(meta), result)
}

// StringToBytes 在不复制底层字节的情况下将字符串转换为 []byte。
//
// 实现基于 unsafe.Pointer，对调用方有一个重要约定：**返回的 []byte
// 绝不能被修改**。一旦修改，会破坏字符串的不可变性，可能在其它
// goroutine 读取同一字符串时触发未定义行为。
//
// 推荐仅在以下场景使用：
//   - 把字符串原样传给期望 []byte 的 API（如 xjson.Unmarshal、io.Writer）；
//   - 临时的 hash / 比较 / 切分操作。
//
// 不适用场景：
//   - 需要修改返回值；
//   - 需要长期持有返回值（因为 Go 的 GC 可能回收原始 string）。
//
// 安全替代方案是 []byte(s)，会复制一份内存。
func StringToBytes(value string) []byte {
	return *(*[]byte)(unsafe.Pointer(&value))
}

// MetaToJsonContent 用 value 构造一个 JSON 对象字面量字符串。
//
// 简单地把 value 用大括号包起来，等价于 fmt.Sprintf("{%s}", value)，
// 主要用于将一段已经拼接好的 JSON 片段包裹成对象。
//
// 注意：本函数不做任何转义与校验，如果 value 内本身含有未转义的
// 引号或逗号，调用方需自行保证 JSON 语法正确。
func MetaToJsonContent(value string) string {
	return fmt.Sprintf("{%s}", value)
}
