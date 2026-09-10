package xphp

import (
	"encoding/base64"
	"net/url"
	"strings"
)

// ParseURL 解析 URL 并返回其组成部分（对应 PHP parse_url()）。
// component 掩码：-1 表示全部；1：scheme(协议)；2：host(主机)；4：port(端口)；8：user(用户名)；16：pass(密码)；32：path(路径)；64：query(查询串)；128：fragment(片段)。
func ParseURL(str string, component int) (map[string]string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return nil, err
	}
	if component == -1 {
		component = 1 | 2 | 4 | 8 | 16 | 32 | 64 | 128
	}
	components := make(map[string]string)
	if (component & 1) == 1 {
		components["scheme"] = u.Scheme
	}
	if (component & 2) == 2 {
		components["host"] = u.Hostname()
	}
	if (component & 4) == 4 {
		components["port"] = u.Port()
	}
	if (component & 8) == 8 {
		components["user"] = u.User.Username()
	}
	if (component & 16) == 16 {
		components["pass"], _ = u.User.Password()
	}
	if (component & 32) == 32 {
		components["path"] = u.Path
	}
	if (component & 64) == 64 {
		components["query"] = u.RawQuery
	}
	if (component & 128) == 128 {
		components["fragment"] = u.Fragment
	}
	return components, nil
}

// URLEncode 对字符串进行 URL 编码，空格编码为 +（对应 PHP urlencode()）。
func URLEncode(str string) string {
	return url.QueryEscape(str)
}

// URLDecode 解码 URL 编码的字符串，+ 解码为空格（对应 PHP urldecode()）。
func URLDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

// URLDecodeMap 将 URL 编码的查询字符串 str 解析为键值对 map：先整体 URL 解码，再按 & 拆分键值对、按 = 拆分键与值。
func URLDecodeMap(str string) (results map[string]string, err error) {
	result, err := url.QueryUnescape(str)
	if err != nil {
		return
	}
	results = make(map[string]string)
	resultArray := strings.Split(result, "&")
	for _, val := range resultArray {
		valArray := strings.Split(val, "=")
		if len(valArray) == 2 {
			results[valArray[0]] = valArray[1]
		}
	}
	return
}

// Rawurlencode 按 RFC 3986 对字符串进行 URL 编码，空格编码为 %20（对应 PHP rawurlencode()）。
func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

// Rawurldecode 解码按 RFC 3986 编码的 URL 字符串，%20 还原为空格（对应 PHP rawurldecode()）。
func Rawurldecode(str string) (string, error) {
	return url.QueryUnescape(strings.Replace(str, "%20", "+", -1))
}

// HTTPBuildQuery 将查询数据 queryData 编码为 URL 查询字符串（对应 PHP http_build_query()）。
func HTTPBuildQuery(queryData url.Values) string {
	return queryData.Encode()
}

// Base64Encode 将字符串 str 编码为 Base64（对应 PHP base64_encode()）。
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode 将 Base64 字符串 str 解码为原始字符串，缺少的填充符 = 会自动补齐（对应 PHP base64_decode()）。
func Base64Decode(str string) (string, error) {
	switch len(str) % 4 {
	case 2:
		str += "=="
	case 3:
		str += "="
	}

	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
