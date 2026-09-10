package xphp

import (
	"errors"
	"fmt"
	"html"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// PadTypes 为 StrPad 提供可选的填充位置类型。
const (
	// StrPadRight 表示在字符串右侧填充。
	StrPadRight = iota
	// StrPadLeft 表示在字符串左侧填充。
	StrPadLeft
	// StrPadBoth 表示在字符串两侧同时填充。
	StrPadBoth
)

const (
	// PHP_EOL 定义 PHP 风格的换行符常量。
	PHP_EOL = "\n"
)

// RuneMatchFunc 是用于判断某个 rune 是否满足条件的函数类型。
type RuneMatchFunc func(rune) bool

// buildReplaceSlice 是 Replace 与 Ireplace 的辅助函数，负责将 search 与 replace 参数归一化为字符串切片。
func buildReplaceSlice(search, replace interface{}) ([]string, []string, error) {
	var aSearch, aReplace []string

	switch s := search.(type) {
	case int:
		aSearch = append(aSearch, strconv.Itoa(s))
	case rune:
		aSearch = append(aSearch, string(s))
	case string:
		aSearch = append(aSearch, s)
	case []string:
		aSearch = s
	default:
		return aSearch, aReplace, errors.New("unsupported type of search")
	}

	switch r := replace.(type) {
	case int:
		aReplace = append(aReplace, strconv.Itoa(r))
	case rune:
		aReplace = append(aReplace, string(r))
	case string:
		aReplace = append(aReplace, r)
	case []string:
		aReplace = r
	default:
		return aSearch, aReplace, errors.New("unsupported type of replace")
	}

	return aSearch, aReplace, nil
}

// Replace 将 subject 中所有出现的 search 字符串替换为 replace 字符串。
//
// 该函数是 PHP str_replace() 的 Go 实现。
//
// 参见 http://php.net/manual/en/function.str-replace.php
func Replace(search, replace interface{}, subject string) string {
	var aSearch, aReplace []string

	aSearch, aReplace, err := buildReplaceSlice(search, replace)
	if err != nil {
		return subject
	}

	r := ""
	for index, s := range aSearch {
		if index < len(aReplace) {
			r = aReplace[index]
		}
		subject = strings.Replace(subject, s, r, -1)
	}

	return subject
}

// Ireplace 是 Replace() 的不区分大小写版本。
func Ireplace(search, replace interface{}, subject string) string {
	var aSearch, aReplace []string

	aSearch, aReplace, err := buildReplaceSlice(search, replace)
	if err != nil {
		return subject
	}

	r := ""
	var reg *regexp.Regexp
	for index, s := range aSearch {
		if index < len(aReplace) {
			r = aReplace[index]
		}
		reg = regexp.MustCompile("(?i:" + s + ")")
		subject = reg.ReplaceAllString(subject, r)
	}

	return subject
}

// Stristr 是不区分大小写的 Strstr()。
func Stristr(haystack, needle string) string {
	haystackx := strings.ToLower(haystack)
	needlex := strings.ToLower(needle)
	pos := Stripos(haystackx, needlex, 0)
	if pos < 0 {
		return ""
	}
	return Substr(haystack, uint(pos), 0)
}

// HTMLSpecialchars 将字符串中的特殊字符转换为 HTML 实体。
func HTMLSpecialchars(str string) string {
	return html.EscapeString(str)
}

// HTMLSpecialcharsDecode 将特殊的 HTML 实体还原为对应的字符。
func HTMLSpecialcharsDecode(str string) string {
	return html.UnescapeString(str)
}

// DefaultNumberFormat 以英文千分位记法调用 NumberFormat：小数点为 "."，千位分隔符为 ","。
func DefaultNumberFormat(number float64, decimals int) string {
	return NumberFormat(number, uint(decimals), ".", ",")
}

// StrPad 使用 padString 将 input 填充到指定长度 padLength，padType 指定填充位置。
func StrPad(input string, padLength int, padString string, padType int) string {
	// if the value of padLength is less than or equal to the length of the input string,
	// no padding takes place, and input will be returned.
	if padLength <= len(input) {
		return input
	}
	// default padType is StrPadRight
	if padType > StrPadBoth {
		padType = StrPadRight
	}
	// default padString is space
	if padString == "" {
		padString = " "
	}

	var s string
	l := padLength - len(input)
	switch padType {
	case StrPadRight:
		s = input
		i := 0
		for i < l {
			for j := 0; j < len(padString); j++ {
				s += padString[j : j+1]
				i++
				if i >= l {
					break
				}
			}
		}
	case StrPadLeft:
		i := 0
		for i < l {
			for j := 0; j < len(padString); j++ {
				s += padString[j : j+1]
				i++
				if i >= l {
					break
				}
			}
		}
		s += input
	case StrPadBoth:
		l1 := l / 2
		s = StrPad(StrPad(input, len(input)+l1, padString, StrPadLeft), padLength, padString, StrPadRight)
	}

	return s
}

// Num2String 将数字 num 转为可读字符串：num >= 1000 时格式化为保留两位小数的 "x.xxK" 形式，否则返回其十进制表示。
func Num2String(num int) string {
	if num >= 1000 {
		rt := fmt.Sprintf("%0.2f", math.Round(float64(num)/1000*100)/100)
		rt = strings.TrimRight(rt, "0")
		rt = strings.TrimRight(rt, ".")
		return rt + "K"
	}
	return strconv.Itoa(num)
}
