// Package xreflect 提供反射相关的增强工具，主要包括：
//   - DeepEqual：对 reflect.DeepEqual 的增强，先比较再回退到 JSON 比较；
//   - 其它子文件提供类型提取、MatchPattern 等常见反射操作。
package xreflect

import (
	"bytes"
	"reflect"

	"github.com/zsy619/tools/xjson"
)

// DeepEqual 是对 reflect.DeepEqual 的增强版。
//
// 先尝试 reflect.DeepEqual（速度快、零分配），如果失败则把两侧值
// 都用 xjson.Marshal 序列化为 JSON 后逐字节比较。
//
// 该增强逻辑能解决 reflect.DeepEqual 在某些语义不同但 JSON 表示
// 相同的场景下的「误判为不等」问题，例如：
//   - map 字段顺序不同；
//   - float 与 int 的可表达相等；
//   - 自定义 UnmarshalJSON 实现的不同结构体。
//
// 注意：JSON 比较会忽略 time.Time 的 Location 等反射能区分的细节，
// 因此本函数等价性比 reflect.DeepEqual 略弱，请按需选用。
//
// 参考文献：
//   - https://mp.weixin.qq.com/s?__biz=MzkyMzIyNjIxMQ==&mid=2247485356&idx=1&sn=5c8f73b895fa284de887fe05f19e4012&chksm=c1e91ea0f69e97b699f59468b78a44dc660293d95771510dbc4466acb246c525ce075bd8a554&scene=132#wechat_redirect
func DeepEqual(v1, v2 interface{}) bool {
	if reflect.DeepEqual(v1, v2) {
		return true
	}
	bytesA, _ := xjson.Marshal(v1)
	bytesB, _ := xjson.Marshal(v2)
	return bytes.Equal(bytesA, bytesB)
}
