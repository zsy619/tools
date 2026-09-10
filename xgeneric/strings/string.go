package strings

import (
	"math/rand"
	"unicode/utf8"

	"github.com/zsy619/tools/xgeneric"
)

var (
	LowerCaseLettersCharset = []rune("abcdefghijklmnopqrstuvwxyz")
	UpperCaseLettersCharset = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LettersCharset          = append(LowerCaseLettersCharset, UpperCaseLettersCharset...)
	NumbersCharset          = []rune("0123456789")
	AlphanumericCharset     = append(LettersCharset, NumbersCharset...)
	SpecialCharset          = []rune("!@#$%^&*()_+-=[]{}|;':\",./<>?")
	AllCharset              = append(AlphanumericCharset, SpecialCharset...)
)

// RandomString 返回一个由指定字符集随机生成的字符串。
// 示例：https://go.dev/play/p/rRseOQVVum4
// 参数：size 为结果长度，charset 为可用字符集。
// 返回：长度为 size 的随机字符串；当 size <= 0 或 charset 为空时返回空串。
func RandomString(size int, charset []rune) string {
	if size <= 0 {
		return ""
	}
	if len(charset) <= 0 {
		return ""
	}

	b := make([]rune, size)
	possibleCharactersCount := len(charset)
	for i := range b {
		b[i] = charset[rand.Intn(possibleCharactersCount)]
	}
	return string(b)
}

// Substring 返回字符串的子串。
// 示例：https://go.dev/play/p/TQlxQi82Lu1
// 参数：str 为输入字符串，offset 为起始偏移（负数表示从尾部倒数），length 为长度。
// 返回：以 offset 开始、长度为 length 的子串；当 offset 越界时返回空串，length 超过剩余长度会被截断。
func Substring[T ~string](str T, offset int, length uint) T {
	size := len(str)

	if offset < 0 {
		offset = size + offset
		if offset < 0 {
			offset = 0
		}
	}

	if offset > size {
		return xgeneric.Empty[T]()
	}

	if length > uint(size)-uint(offset) {
		length = uint(size - offset)
	}

	return str[offset : offset+int(length)]
}

// ChunkString 将字符串按 size 长度切分为多个子串切片。
// 当无法均匀切分时，最后一个块包含剩余字符。
// 示例：https://go.dev/play/p/__FLTuJVz54
// 参数：str 为输入字符串，size 为每个块的长度。
// 返回：包含切分后所有块的 []T。
// 当 size <= 0 时 panic("lo.ChunkString: Size parameter must be greater than 0")；空串返回 [""]。
func ChunkString[T ~string](str T, size int) []T {
	if size <= 0 {
		panic("lo.ChunkString: Size parameter must be greater than 0")
	}

	if len(str) == 0 {
		return []T{""}
	}

	if size >= len(str) {
		return []T{str}
	}

	var chunks []T = make([]T, 0, ((len(str)-1)/size)+1)
	currentLen := 0
	currentStart := 0
	for i := range str {
		if currentLen == size {
			chunks = append(chunks, str[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, str[currentStart:])
	return chunks
}

// RuneLength 是 utf8.RuneCountInString 的别名，返回字符串中 rune 的数量。
// 示例：https://go.dev/play/p/tuhgW_lWY8l
// 参数：str 为输入字符串。
// 返回：str 中的 rune 个数（按 UTF-8 解码）。
func RuneLength(str string) int {
	return utf8.RuneCountInString(str)
}
