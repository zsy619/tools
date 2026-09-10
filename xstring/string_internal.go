package xstring

import (
	"unicode"
)

// splitIntoStrings 根据Unicode字符类型将字符串分割成多个子串，并返回字符串切片。
// 如果upperCase为true，则所有返回的字符串为大写，否则为小写。
// 参数：
//   s: 待分割的字符串。
//   upperCase: 是否将返回字符串转为大写，true表示转为大写，false表示转为小写。
// 返回值：
//   []string: 分割后的字符串切片。
func splitIntoStrings(s string, upperCase bool) []string {
	var runes [][]rune
	lastCharType := 0
	charType := 0

	// 根据 Unicode 字符类型将字符串切分为多个 rune 切片
	for _, r := range s {
		switch true {
		case isLower(r):
			charType = 1
		case isUpper(r):
			charType = 2
		case isDigit(r):
			charType = 3
		default:
			charType = 4
		}

		if charType == lastCharType {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastCharType = charType
	}

	for i := 0; i < len(runes)-1; i++ {
		if isUpper(runes[i][0]) && isLower(runes[i+1][0]) {
			length := len(runes[i]) - 1
			temp := runes[i][length]
			runes[i+1] = append([]rune{temp}, runes[i+1]...)
			runes[i] = runes[i][:length]
		}
	}

	// 过滤掉既非字母也非数字开头的分组，并按 upperCase 决定统一大小写
	var result []string
	for _, rs := range runes {
		if len(rs) > 0 && (unicode.IsLetter(rs[0]) || isDigit(rs[0])) {
			if upperCase {
				result = append(result, string(toUpperAll(rs)))
			} else {
				result = append(result, string(toLowerAll(rs)))
			}
		}
	}

	return result
}

// isDigit 判断给定的 rune 是否为 ASCII 十进制数字（'0' 到 '9'）。
// 注意：本函数仅覆盖 ASCII 数字，未包含全角数字或 Unicode 其他数字字符。
// 参数：
//   r: 待检测的字符。
// 返回值：
//   bool: 当且仅当 r 处于 '0'..'9' 区间时返回 true，否则返回 false。
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// isLower 判断给定的 rune 是否为 ASCII 小写字母（'a' 到 'z'）。
// 注意：本函数仅覆盖 ASCII 小写字母，不包含其他 Unicode 字符。
// 参数：
//   r: 待检测的字符。
// 返回值：
//   bool: 当且仅当 r 处于 'a'..'z' 区间时返回 true，否则返回 false。
func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

// isUpper 判断给定的 rune 是否为 ASCII 大写字母（'A' 到 'Z'）。
// 注意：本函数仅覆盖 ASCII 大写字母，不包含其他 Unicode 字符。
// 参数：
//   r: 待检测的字符。
// 返回值：
//   bool: 当且仅当 r 处于 'A'..'Z' 区间时返回 true，否则返回 false。
func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// toLower 将 ASCII 大写字母（'A' 到 'Z'）转换为对应的小写字母。
// 通过偏移量 +32 实现大写到小写的转换；非 ASCII 大写字母将原样返回。
// 参数：
//   r: 待转换的字符。
// 返回值：
//   rune: 转换后的字符；若 r 不是 ASCII 大写字母，则原样返回。
func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

// toLowerAll 将 rune 切片中的每个 ASCII 大写字母（'A' 到 'Z'）转换为对应的小写字母。
// 会对传入切片进行原地修改并返回同一引用。
// 参数：
//   rs: 待转换的 rune 切片（会被修改）。
// 返回值：
//   []rune: 转换后的 rune 切片（即入参数 rs 本身）。
func toLowerAll(rs []rune) []rune {
	for i := range rs {
		rs[i] = toLower(rs[i])
	}
	return rs
}

// toUpper 将 ASCII 小写字母（'a' 到 'z'）转换为对应的大写字母。
// 通过偏移量 -32 实现小写到大写的转换；非 ASCII 小写字母将原样返回。
// 参数：
//   r: 待转换的字符。
// 返回值：
//   rune: 转换后的字符；若 r 不是 ASCII 小写字母，则原样返回。
func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}

// toUpperAll 将 rune 切片中的每个 ASCII 小写字母（'a' 到 'z'）转换为对应的大写字母。
// 会对传入切片进行原地修改并返回同一引用。
// 参数：
//   rs: 待转换的 rune 切片（会被修改）。
// 返回值：
//   []rune: 转换后的 rune 切片（即入参数 rs 本身）。
func toUpperAll(rs []rune) []rune {
	for i := range rs {
		rs[i] = toUpper(rs[i])
	}
	return rs
}

// padAtPosition 将字符串 str 按指定填充方式与位置扩展到目标长度 length。
// 当 length 不大于 str 长度时，直接返回原字符串；否则按 position 在左侧、
// 两侧或右侧填充 padStr 字符（若 padStr 为空则默认使用单个空格）。
// 参数：
//   str: 待填充的原始字符串。
//   length: 目标长度（按字节计算，使用 len(str)）。
//   padStr: 用于填充的字符串；若为空则使用单个空格 " " 作为填充字符。
//   position: 填充位置。0 表示两侧居中填充；1 表示左侧填充；其他值表示右侧填充。
// 返回值：
//   string: 填充后的字符串，长度等于 length（按字节）。
func padAtPosition(str string, length int, padStr string, position int) string {
	if len(str) >= length {
		return str
	}

	if padStr == "" {
		padStr = " "
	}

	length = length - len(str)
	startPadLen := 0
	if position == 0 {
		startPadLen = length / 2
	} else if position == 1 {
		startPadLen = length
	}
	endPadLen := length - startPadLen

	charLen := len(padStr)
	leftPad := ""
	cur := 0
	for cur < startPadLen {
		leftPad += string(padStr[cur%charLen])
		cur++
	}

	cur = 0
	rightPad := ""
	for cur < endPadLen {
		rightPad += string(padStr[cur%charLen])
		cur++
	}

	return leftPad + str + rightPad
}

// isLetter 判断给定的 rune 是否为字母但排除 CJK 字符。
// 即：是 Unicode 字母（unicode.IsLetter），并且不在以下 CJK 区块内：
//   - 平假名与片假名（日文）
//   - CJK 统一表意符号扩展 A
//   - CJK 统一表意符号
//   - CJK 兼容表意字符
//   - 半角片假名（日文）
// 参数：
//   r: 待检测的字符。
// 返回值：
//   bool: 当 r 是非 CJK 字母时返回 true，否则返回 false。
func isLetter(r rune) bool {
	if !unicode.IsLetter(r) {
		return false
	}

	switch {
	// CJK 字符范围：/[\u3040-\u30ff\u3400-\u4dbf\u4e00-\u9fff\uf900-\ufaff\uff66-\uff9f]/

	// 平假名与片假名（仅日文）
	case r >= '\u3034' && r < '\u30ff':
		return false

	// CJK 统一表意符号扩展 A（中、日、韩）
	case r >= '\u3400' && r < '\u4dbf':
		return false

	// CJK 统一表意符号（中、日、韩）
	case r >= '\u4e00' && r < '\u9fff':
		return false

	// CJK 兼容表意字符（中、日、韩）
	case r >= '\uf900' && r < '\ufaff':
		return false

	// 半角片假名（仅日文）
	case r >= '\uff66' && r < '\uff9f':
		return false
	}

	return true
}
