// Package xuuid 提供 UUID/GUID 与 Snowflake 唯一 ID 的生成工具。
package xuuid

import (
	"crypto/rand"
	"fmt"
	"log"
)

// Guid 生成一个随机的 GUID 字符串，格式为 8-4-4-4-12 的 32 位十六进制串。
//
// 内部使用 crypto/rand 读取 16 字节随机数进行构造，因此具备密码学强度，
// 适用于全局唯一标识符场景，例如数据库主键、请求追踪 ID 等。
//
// 副作用说明：若系统熵源不可用（例如 /dev/urandom 读取失败），
// 函数会通过 log.Fatal 直接终止进程，不会返回错误。
// 调用方应确保该行为在自身上下文（例如测试或受限沙箱）中可接受。
//
// 返回值：形如 "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" 的小写十六进制字符串。
func Guid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
