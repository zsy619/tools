package xreflect

import (
	"bytes"
	"reflect"

	"haedu.gov.cn/tools/xjson"
)

// 增强型 DeepEqual 函数
// https://mp.weixin.qq.com/s?__biz=MzkyMzIyNjIxMQ==&mid=2247485356&idx=1&sn=5c8f73b895fa284de887fe05f19e4012&chksm=c1e91ea0f69e97b699f59468b78a44dc660293d95771510dbc4466acb246c525ce075bd8a554&scene=132#wechat_redirect
func DeepEqual(v1, v2 interface{}) bool {
	if reflect.DeepEqual(v1, v2) {
		return true
	}
	bytesA, _ := xjson.Marshal(v1)
	bytesB, _ := xjson.Marshal(v2)
	return bytes.Equal(bytesA, bytesB)
}
