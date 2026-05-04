package xstring

import (
	"fmt"
	"unsafe"

	"github.com/zsy619/tools/xjson"
)

// JsonToObject 将JSON字符串转换为Go对象
// meta：待转换的JSON字符串
// result：用于接收转换结果的Go对象指针
// 返回值：若转换成功返回nil，否则返回错误信息
func JsonToObject(meta string, result any) error {
	return xjson.Unmarshal(StringToBytes(meta), result)
}

// StringToBytes 将字符串转换为字节数组
// 参数:
//
//	value: 待转换的字符串
//
// 返回值:
//
//	转换后的字节数组
func StringToBytes(value string) []byte {
	return *(*[]byte)(unsafe.Pointer(&value))
}

// MetaToJsonContent 将字符串转换为JSON格式并返回
// 参数value：需要转换的字符串
// 返回值：转换后的JSON字符串
func MetaToJsonContent(value string) string {
	return fmt.Sprintf("{%s}", value)
}
