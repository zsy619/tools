package convs

import "github.com/zsy619/tools/xstring"

// SplitStrToSet 将字符串按分隔符拆分后转换为 map[string]struct{} 集合。
// 参数：v 为输入字符串，sep 为分隔符。
// 返回：以拆分后的子串为键的集合（值为空结构体）；相同子串只会保留一份。
func SplitStrToSet(v string, sep string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, v := range xstring.Split(v, sep) {
		m[v] = struct{}{}
	}
	return m
}
