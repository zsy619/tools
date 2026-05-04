package convs

import "github.com/zsy619/tools/xstring"

// SplitStrToSet convert a string to map set after split
func SplitStrToSet(v string, sep string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, v := range xstring.Split(v, sep) {
		m[v] = struct{}{}
	}
	return m
}
