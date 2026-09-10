package xphp

import (
	"fmt"
	"regexp"
)

// PregMatch 使用正则表达式 expr 匹配字符串 s，将整体匹配及各捕获分组存入 result，
// 并在标准输出逐条打印（[index] value）。
func PregMatch(expr string, s string) (result []string, err error) {
	r, err := regexp.Compile(expr)
	for index, match := range r.FindStringSubmatch(s) {
		fmt.Printf("[%d] %s\n", index, match)
		result = append(result, match)
	}
	return
}
