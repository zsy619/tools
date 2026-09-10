package xjson

import (
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

// 参考文档链接：
// https://mp.weixin.qq.com/s?__biz=MzkyMzIyNjIxMQ==&mid=2247485083&idx=1&sn=e828d5b3d30727bb13a16ca00d63324d&chksm=c1e91f97f69e9681ef2a6b79674c7202869bebc478713caf3fc8e34c4021e6ee6da37a95949f&scene=21#wechat_redirect
// https://www.jianshu.com/p/040f8b735a47

// 全局 JSON 编解码器实例。
// JSON 与标准库 encoding/json API 兼容，性能更高。
// JSONDefault 为 json-iterator 默认配置。
// JSONFastest 为 json-iterator 最快配置（牺牲部分兼容性换取速度）。
var (
	JSON        = jsoniter.ConfigCompatibleWithStandardLibrary
	JSONDefault = jsoniter.ConfigDefault
	JSONFastest = jsoniter.ConfigFastest
)

// init 包初始化函数。注册 PHP 模糊解码器，使解码器对 PHP 风格的 JSON（如带前导零的数字等）更加宽容。
func init() {
	extra.RegisterFuzzyDecoders() // 开启PHP兼容模式
}

// Unmarshal 将 JSON 字节数据反序列化到 v 指向的值中。
// data 为要解析的 JSON 字节切片，v 为目标对象指针。
// 当 JSON 不合法或与 v 类型不匹配时返回错误。
func Unmarshal(data []byte, v interface{}) error {
	return JSON.Unmarshal(data, v)
}

// UnmarshalUseNumber 将 JSON 字节数据反序列化到 v 中，数字类型解码为 json.Number 而非 float64。
// 适用于需要保留原始数字精度的场景。
// data 为 JSON 字节切片，v 为目标对象指针；解析失败时返回错误。
func UnmarshalUseNumber(data []byte, v interface{}) error {
	s := string(data)
	decoder := jsoniter.NewDecoder(strings.NewReader(s))
	decoder.UseNumber()
	err := decoder.Decode(&v)
	return err
}

// UnmarshalFromString 将 JSON 字符串反序列化到 v 指向的值中。
// str 为要解析的 JSON 字符串，v 为目标对象指针；解析失败时返回错误。
func UnmarshalFromString(str string, v interface{}) error {
	return JSON.UnmarshalFromString(str, v)
}

// Marshal 将 v 序列化为 JSON 字节切片。
// v 为任意可序列化类型；序列化失败时返回错误。
func Marshal(v interface{}) ([]byte, error) {
	return JSON.Marshal(v)
}

// MarshalToString 将 v 序列化为 JSON 字符串。
// v 为任意可序列化类型；序列化失败时返回错误。
func MarshalToString(v interface{}) (string, error) {
	return JSON.MarshalToString(v)
}
