package xstring

import (
	"fmt"

	"haedu.gov.cn/tools/xjson"
)

func JsonToObject(meta string, result interface{}) error {
	return xjson.Unmarshal(StringToBytes(meta), result)
}

func StringToBytes(value string) []byte {
	return []byte(value)
}

func MetaToJsonContent(value string) string {
	return fmt.Sprintf("{%s}", value)
}
