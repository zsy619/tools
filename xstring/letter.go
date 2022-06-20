package xstring

import (
	"strings"
)

// FirstLetterUpper 首字母转大写
func FirstLetterUpper(str string) string {
	if str != "" {
		fieldBytes := []byte(str)
		fieldLength := len(str)
		str = strings.ToUpper(string(fieldBytes[:1])) + string(fieldBytes[1:fieldLength])
		fieldBytes = nil
	}
	return str
}

// FirstLetterLower 首字母转小写
func FirstLetterLower(str string) string {
	if str != "" {
		fieldBytes := []byte(str)
		fieldLength := len(str)
		str = strings.ToLower(string(fieldBytes[:1])) + string(fieldBytes[1:fieldLength])
		fieldBytes = nil
	}
	return str
}

func ToLetterArray(str string) (result []string) {
	letter := []rune(str)
	for i := 0; i < len(letter); i++ {
		result = append(result, string(letter[i]))
	}
	return
}

// GetStringWidth get the string width
func GetStringWidth(str string) int {
	w := 0
	for _, c := range []rune(str) {
		if IsHalfwidth(c) {
			w = w + 1
		} else {
			w = w + 2
		}
	}
	return w
}
