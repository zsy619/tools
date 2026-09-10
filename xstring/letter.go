package xstring

import "strings"

// FirstLetterUpper 将字符串首字符转换为大写。
//
// 仅作用于首字符，其余字符保持原样；空字符串返回空字符串。
//
// 注意：本函数按字节处理首字符，对纯 ASCII 字符串行为符合直觉；
// 对多字节 UTF-8 字符串可能切出非法 rune，请仅在已知为 ASCII 的场景
// 使用，否则请改用 utf8.DecodeRuneInString。
func FirstLetterUpper(s string) string {
	if s != "" {
		fieldBytes := []byte(s)
		fieldLength := len(s)
		s = strings.ToUpper(string(fieldBytes[:1])) + string(fieldBytes[1:fieldLength])
		fieldBytes = nil
	}
	return s
}

// FirstLetterLower 将字符串首字符转换为小写。
//
// 与 FirstLetterUpper 对称，仅处理首字符；空字符串返回空字符串。
func FirstLetterLower(s string) string {
	if s != "" {
		fieldBytes := []byte(s)
		fieldLength := len(s)
		s = strings.ToLower(string(fieldBytes[:1])) + string(fieldBytes[1:fieldLength])
		fieldBytes = nil
	}
	return s
}

// ToLetterArray 把字符串拆成单字符切片。
//
// 使用 rune 解码，因此中文等多字节字符会被正确切分为单个 rune。
// 返回的每个元素都是一个长度为 1 的字符（rune 转字符串）。
func ToLetterArray(s string) (result []string) {
	letter := []rune(s)
	for i := 0; i < len(letter); i++ {
		result = append(result, string(letter[i]))
	}
	return
}

// GetStringWidth 按终端等宽字体的显示宽度估算字符串宽度。
//
// 实现遵循 East Asian Width 规范：
//   - 半角字符（ASCII 等）宽度为 1；
//   - 全角字符（中文、日文、韩文等）宽度为 2。
//
// 常用于对齐日志、表格、终端输出等场景。注意这是估算值，遇到组合
// 字符、控制字符或 emoji 时可能与真实终端表现略有差异。
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
