package xjson

import (
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

// https://mp.weixin.qq.com/s?__biz=MzkyMzIyNjIxMQ==&mid=2247485083&idx=1&sn=e828d5b3d30727bb13a16ca00d63324d&chksm=c1e91f97f69e9681ef2a6b79674c7202869bebc478713caf3fc8e34c4021e6ee6da37a95949f&scene=21#wechat_redirect
// https://www.jianshu.com/p/040f8b735a47
var (
	JSON        = jsoniter.ConfigCompatibleWithStandardLibrary
	JSONDefault = jsoniter.ConfigDefault
	JSONFastest = jsoniter.ConfigFastest
)

func init() {
	extra.RegisterFuzzyDecoders() // 开启PHP兼容模式
}

func Unmarshal(data []byte, v interface{}) error {
	return JSON.Unmarshal(data, v)
}

func UnmarshalUseNumber(data []byte, v interface{}) error {
	s := string(data)
	decoder := jsoniter.NewDecoder(strings.NewReader(s))
	decoder.UseNumber()
	err := decoder.Decode(&v)
	return err
}

func UnmarshalFromString(str string, v interface{}) error {
	return JSON.UnmarshalFromString(str, v)
}

func Marshal(v interface{}) ([]byte, error) {
	return JSON.Marshal(v)
}

func MarshalToString(v interface{}) (string, error) {
	return JSON.MarshalToString(v)
}
