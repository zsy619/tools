package xdatabase

import "strings"

// SafeString 转义字符串
func SafeString(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")  // 转义 \
	s = strings.ReplaceAll(s, "\n", "\\n")   // 转义换行符
	s = strings.ReplaceAll(s, "\r", "\\r")   // 转义回车符
	s = strings.ReplaceAll(s, "\t", "\\t")   // 转义制表符
	s = strings.ReplaceAll(s, "\a", "\\a")   // 转义响铃符
	s = strings.ReplaceAll(s, "\b", "\\b")   // 转义退格符
	s = strings.ReplaceAll(s, "\f", "\\f")   // 转义换页符
	s = strings.ReplaceAll(s, "'", "\\'")    // 转义 '
	s = strings.ReplaceAll(s, "\"", "\\\"")  // 转义 "
	s = strings.ReplaceAll(s, "\x1a", "\\Z") // 转义 EOF
	return s
}
