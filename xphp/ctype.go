package xphp

import "regexp"

// CtypeAlnum 当 text 中所有字符均为字母或数字时返回 true，否则返回 false。
// 空字符串返回 false。
func CtypeAlnum(text string) bool {
	reg := regexp.MustCompile("^[[:alnum:]]+$") // 等价模式 ^[a-zA-Z0-9]+$
	return reg.MatchString(text)
}

// CtypeAlpha 当 text 中所有字符均为字母时返回 true，否则返回 false。
// 空字符串返回 false。
func CtypeAlpha(text string) bool {
	reg := regexp.MustCompile("^[[:alpha:]]+$") // 等价模式 ^[A-Za-z]+$
	return reg.MatchString(text)
}

// CtypeCntrl 当 text 中所有字符均为控制字符时返回 true，否则返回 false。
// 控制字符包括 \x00-\x1f 与 \x7f；空字符串返回 false。
func CtypeCntrl(text string) bool {
	reg := regexp.MustCompile("^[[:cntrl:]]+$") // 等价模式 ^[\x00-\x1f\x7f]+$
	return reg.MatchString(text)
}

// CtypeDigit 当 text 中所有字符均为十进制数字时返回 true，否则返回 false。
// 空字符串返回 false。
func CtypeDigit(text string) bool {
	reg := regexp.MustCompile("^[[:digit:]]+$") // 等价模式 ^[0-9]+$
	return reg.MatchString(text)
}

// CtypeGraph 当 text 中所有字符均为可打印且会产生可见输出（不含空白字符）时返回 true。
// 等价匹配范围为 [!-~]；空字符串返回 false。
func CtypeGraph(text string) bool {
	reg := regexp.MustCompile("^[[:graph:]]+$") // 等价模式 ^[!-~]+$
	return reg.MatchString(text)
}

// CtypeLower 当 text 中所有字符均为小写字母时返回 true。
// 空字符串返回 false。
func CtypeLower(text string) bool {
	reg := regexp.MustCompile("^[[:lower:]]+$") // 等价模式 ^[a-z]+$
	return reg.MatchString(text)
}

// CtypePrint 当 text 中所有字符均会产生输出（含空格）时返回 true。
// 若 text 包含控制字符或不产生任何输出 / 控制功能的字符则返回 false。
// 匹配范围为 [ -~]；空字符串返回 false。
func CtypePrint(text string) bool {
	reg := regexp.MustCompile("^[[:print:]]+$") // 等价模式 ^[ -~]+$
	return reg.MatchString(text)
}

// CtypePunct 当 text 中所有字符均为可打印字符、且不是字母、数字或空白时返回 true。
// 匹配范围为 [!-/:-@[-`{-~]；空字符串返回 false。
func CtypePunct(text string) bool {
	reg := regexp.MustCompile("^[[:punct:]]+$") // 等价模式 ^[!-\/\:-@\[-`\{-~]+$
	return reg.MatchString(text)
}

// CtypeSpace 当 text 中所有字符均为空白字符时返回 true，否则返回 false。
// 除空格外还包括制表符、垂直制表符、换行、回车与换页符。
// 匹配范围为 s；空字符串返回 false。
func CtypeSpace(text string) bool {
	reg := regexp.MustCompile("^[[:space:]]+$") // 等价模式 ^[\s]+$
	return reg.MatchString(text)
}

// CtypeUpper 当 text 中所有字符均为大写字母时返回 true。
// 空字符串返回 false。
func CtypeUpper(text string) bool {
	reg := regexp.MustCompile("^[[:upper:]]+$") // 等价模式 ^[A-Z]+$
	return reg.MatchString(text)
}

// CtypeXdigit 当 text 中所有字符均为十六进制“数字”时返回 true。
// 即十进制数字或 A-F / a-f 范围内的字符；空字符串返回 false。
func CtypeXdigit(text string) bool {
	reg := regexp.MustCompile("^[[:xdigit:]]+$") // 等价模式 ^[A-Fa-f0-9]+$
	return reg.MatchString(text)
}
