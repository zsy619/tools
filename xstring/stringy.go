// Package xstring 提供一系列对标准库 strings 包的便捷增强与封装，
// 用于补齐标准库在常见字符串场景下缺失或容易踩坑的行为。
package xstring

import (
	"bytes"
	"regexp"
	"strings"
)

// Split 是对 strings.Split 的封装。
//
// 该函数修复了 strings.Split 的一个常见陷阱：
// 当 s 为空字符串时，strings.Split("", sep) 会返回一个长度为 1 且元素为空字符串的切片，
// 而本函数在 s 为空字符串时直接返回长度为 0 的空切片，更符合调用者的直觉。
//
// 参数:
//   - s: 待分割的字符串。若为空字符串 ""，则返回空切片（不包含任何元素）。
//   - sep: 用作分隔符的字符串（语义与 strings.Split 相同）。
//
// 返回值:
//   - 当 s 为空字符串时，返回长度为 0 的 []string。
//   - 否则返回 strings.Split(s, sep) 的结果，行为与标准库完全一致。
func Split(s, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}

// JoinStrSkipEmpty 将多个字符串用指定分隔符合并为一个字符串，并自动跳过空字符串元素。
//
// 与 strings.Join 的关键区别在于：本函数会跳过所有空字符串 ""，
// 不会在结果中产生多余的连续分隔符，也不会把空串作为首个或末尾元素写入。
//
// 参数:
//   - sep: 相邻两个非空字符串之间的分隔符。允许为空字符串 ""，表示无分隔符直接拼接。
//   - s:   可变参数，待合并的字符串列表。可包含任意数量的空字符串，它们会被静默忽略。
//
// 返回值:
//   - 已合并的字符串。若所有输入均为空字符串或未传入任何参数，则返回空字符串 ""。
//   - 不会在结果的开头或结尾附加分隔符。
func JoinStrSkipEmpty(sep string, s ...string) string {
	var buf bytes.Buffer
	for _, v := range s {
		if v == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(v)
	}
	return buf.String()
}

// JoinStr 将多个字符串用指定分隔符合并为一个字符串。
//
// 注意：本函数不会跳过空字符串元素，空串会被原样写入结果，
// 因此可能会出现连续分隔符或首尾为空串的情况。
// 如需跳过空串，请使用 JoinStrSkipEmpty。
//
// 参数:
//   - sep: 相邻字符串之间的分隔符。允许为空字符串 ""，表示无分隔符直接拼接。
//   - s:   可变参数，待合并的字符串列表（包含空字符串在内的所有元素都会被保留）。
//
// 返回值:
//   - 已合并的字符串。若未传入任何字符串，则返回空字符串 ""。
func JoinStr(sep string, s ...string) string {
	var buf bytes.Buffer
	for i, v := range s {
		if i != 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(v)
	}
	return buf.String()
}

// ReverseStr 将指定字符串反转，不会修改原字符串（Go 中字符串本身不可变，因此也不会有副作用）。
//
// 该函数按 rune（即 Unicode 码点）进行反转，能够正确处理多字节字符（如中文、emoji 等），
// 不会出现按字节反转时的乱码问题。
//
// 参数:
//   - s: 待反转的字符串。可为空字符串 ""，此时返回空字符串。
//
// 返回值:
//   - 反转后的新字符串。原字符串不受影响。
func ReverseStr(s string) string {
	rs := []rune(s)
	var r []rune
	for i := len(rs) - 1; i >= 0; i-- {
		r = append(r, rs[i])
	}
	return string(r)
}

// GetAlphanumericNumByASCII 通过比对字符的 ASCII 码值，统计字符串中字母与数字的总数。
//
// 计数范围（按 ASCII 码）：
//   - '0'..'9'（48..57）数字
//   - 'A'..'Z'（65..90）大写字母
//   - 'a'..'z'（97..122）小写字母
//
// 注意：该函数以字节（byte）为单位遍历，对于非 ASCII（如中文、emoji）不会被计入，
// 因为这些字符的 UTF-8 首字节均落在 128..255 区间。性能优于基于正则表达式的版本，
// 因此在仅关心英文字母与数字的场景下推荐使用本函数。
//
// 参数:
//   - s: 待统计的输入字符串。可为空字符串 ""，此时返回 0。
//
// 返回值:
//   - 字符串中属于上述字母 / 数字字符集合的字节总数。
func GetAlphanumericNumByASCII(s string) int {
	num := int(0)
	for i := 0; i < len(s); i++ {
		switch {
		case 48 <= s[i] && s[i] <= 57: // digits
			fallthrough
		case 65 <= s[i] && s[i] <= 90: // uppercase letters
			fallthrough
		case 97 <= s[i] && s[i] <= 122: // lowercase letters
			num++
		default:
		}
	}
	return num
}

// GetAlphanumericNumByASCIIV2 通过比对字符的 ASCII 码值，统计字符串中字母与数字的总数。
//
// 与 GetAlphanumericNumByASCII 等价：计数范围同样覆盖 '0'..'9'、'A'..'Z'、'a'..'z'，
// 但本函数使用 range s 逐 rune 迭代，因此能够正确解码 UTF-8 多字节序列，
// 同时每次迭代都要做 UTF-8 解码开销，性能比按字节遍历的 GetAlphanumericNumByASCII 差。
//
// 参数:
//   - s: 待统计的输入字符串。可为空字符串 ""，此时返回 0。
//
// 返回值:
//   - 字符串中属于上述字母 / 数字字符集合的字符（rune）总数。
func GetAlphanumericNumByASCIIV2(s string) int {
	num := int(0)
	for _, c := range s {
		switch {
		case '0' <= c && c <= '9':
			fallthrough
		case 'a' <= c && c <= 'z':
			fallthrough
		case 'A' <= c && c <= 'Z':
			num++
		default:
		}
	}
	return num
}

// GetAlphanumericNumByRegExp 使用正则表达式统计字符串中字母与数字的总数。
//
// 该函数通过分别匹配 \d（数字）与 [a-zA-Z]（大小写英文字母），
// 再累加两者的匹配结果得到总计数。
//
// 注意：本函数每次调用都会通过 regexp.MustCompile 重新编译正则表达式，
// 性能远低于基于 ASCII 码直接比较的 GetAlphanumericNumByASCII，
// 因此在性能敏感的场景中推荐改用 GetAlphanumericNumByASCII。
//
// 副作用：regexp.MustCompile 在模式非法时会 panic，
// 本函数使用的两个模式（\d 与 [a-zA-Z]）均为合法模式，因此不会触发 panic。
//
// 参数:
//   - s: 待统计的输入字符串。可为空字符串 ""，此时返回 0。
//
// 返回值:
//   - 字符串中数字字符与英文字母字符的总数（数字与字母不会重复计数）。
func GetAlphanumericNumByRegExp(s string) int {
	rNum := regexp.MustCompile(`\d`)
	rLetter := regexp.MustCompile("[a-zA-Z]")
	return len(rNum.FindAllString(s, -1)) + len(rLetter.FindAllString(s, -1))
}
