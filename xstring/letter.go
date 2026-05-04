package xstring

import (
	"strings"
)

// FirstLetterUpper 将字符串的首字母转换为大写并返回转换后的字符串
// 如果传入的字符串为空，则返回空字符串
// 参数：
//
//	s string - 待转换的字符串
//
// 返回值：
//
//	string - 转换后的字符串
func FirstLetterUpper(s string) string {
	if s != "" {
		fieldBytes := []byte(s)
		fieldLength := len(s)
		s = strings.ToUpper(string(fieldBytes[:1])) + string(fieldBytes[1:fieldLength])
		fieldBytes = nil
	}
	return s
}

// FirstLetterLower 将字符串s的首字母转换为小写并返回转换后的字符串
// 如果s为空字符串，则返回空字符串
func FirstLetterLower(s string) string {
	if s != "" {
		fieldBytes := []byte(s)
		fieldLength := len(s)
		s = strings.ToLower(string(fieldBytes[:1])) + string(fieldBytes[1:fieldLength])
		fieldBytes = nil
	}
	return s
}

// ToLetterArray 将字符串s转换为字符数组并返回
// 参数s：待转换的字符串
// 返回值result：转换后的字符数组
func ToLetterArray(s string) (result []string) {
	letter := []rune(s)
	for i := 0; i < len(letter); i++ {
		result = append(result, string(letter[i]))
	}
	return
}

// GetStringWidth 函数用于计算字符串s的显示宽度
// 参数s是要计算宽度的字符串
// 返回值是字符串s的显示宽度
func GetStringWidth(s string) int {
	w := 0
	rstr := []rune(s)
	for _, c := range rstr {
		if IsHalfwidth(c) {
			w = w + 1
		} else {
			w = w + 2
		}
	}
	return w
}
