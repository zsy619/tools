package xhttp

import (
	"net/http"
	"strings"
)

// IsGzipEncode 判断 HTTP 响应头中是否声明使用 gzip 编码。
// 当 Content-Encoding 头存在且（不区分大小写）等于 "gzip" 时返回 true。
func IsGzipEncode(header http.Header) bool {
	value := header["Content-Encoding"]

	return value != nil && strings.EqualFold(value[0], "gzip")
}
