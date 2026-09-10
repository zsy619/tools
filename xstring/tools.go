package xstring

import (
	"fmt"
	"regexp"
	"strings"
	"unsafe"

	"github.com/zsy619/tools"
	"github.com/zsy619/tools/xarray"
	"github.com/zsy619/tools/xgeneric/sets"
	"github.com/zsy619/tools/xunsafe"
)

const (
	// INDEX_NOT_FOUND 表示「未找到」的位置，与 strings.Index/LastIndex
	// 的 -1 返回值保持一致。
	INDEX_NOT_FOUND = -1
	// EMPTY_STRING 是空字符串的语义常量，便于在表达式中复用并表达
	// 「此处意图为空」的语义。
	EMPTY_STRING = ""
)

// Abbreviate 在固定宽度内截断字符串，超出部分用 abbrevMarker 替换。
//
// 行为细节：
//   - 当 str 或 abbrevMarker 为空时，函数不会真正截断；
//   - offset 表示从 str 哪个位置开始考虑截断；
//   - 当 maxWidth 过小（小于 abbrevMarker 长度 + 1）会返回错误；
//   - 当 maxWidth 已经能装下整个 str，原样返回。
//
// 该函数常用于日志、表格、UI 等需要严格限制字符串宽度的场景。
func Abbreviate(str, abbrevMarker string, offset, maxWidth int) (string, error) {
	if IsEmpty(str) && IsEmpty(abbrevMarker) {
		return str, nil
	} else if !IsEmpty(str) && abbrevMarker == "" && maxWidth > 0 {
		return SubString(str, 0, maxWidth), nil
	} else if IsEmpty(str) || IsEmpty(abbrevMarker) {
		return str, nil
	}

	abbrevMarkerLength := len(abbrevMarker)
	minAbbrevWidth := abbrevMarkerLength + 1
	minAbbrevWidthOffset := abbrevMarkerLength + abbrevMarkerLength + 1

	if maxWidth < minAbbrevWidth {
		return str, fmt.Errorf("minimum abbreviation width is %d", minAbbrevWidth)
	}
	l := len(str)
	if l <= maxWidth {
		return str, nil
	}
	if offset > l {
		offset = l
	}
	if l-offset < maxWidth-abbrevMarkerLength {
		offset = l - (maxWidth - abbrevMarkerLength)
	}
	if offset <= abbrevMarkerLength+1 {
		return SubString(str, 0, maxWidth-abbrevMarkerLength) + abbrevMarker, nil
	}
	if maxWidth < minAbbrevWidthOffset {
		return str, fmt.Errorf("minimum abbreviation width with offset is %d", minAbbrevWidthOffset)
	}
	if offset+maxWidth-abbrevMarkerLength < l {
		ns, err := Abbreviate(SubString(str, offset, -1), abbrevMarker, 0, maxWidth-abbrevMarkerLength)
		if err != nil {
			return str, err
		}
		return abbrevMarker + ns, nil
	}
	return abbrevMarker + SubString(str, l-(maxWidth-abbrevMarkerLength), -1), nil
}

// AbbreviateMiddle 将字符串截断到指定长度，并用 middle 替换中间字符。
//
// 行为细节：
//   - 当 str 或 middle 为空时直接返回原字符串，不做截断；
//   - 当 length >= len(str) 时原样返回；
//   - 当 length < len(middle)+2 时无法容纳替换标记，同样原样返回；
//   - 其余情况下按 (length-len(middle))/2 对称地从前/后各取一段，中间用 middle 拼接。
//
// 典型场景：脱敏显示长字符串（如手机号、身份证号、文件名等）。
func AbbreviateMiddle(str, middle string, length int) string {
	if IsEmpty(str) || IsEmpty(middle) {
		return str
	}

	if length >= len(str) || length < len(middle)+2 {
		return str
	}

	targetSting := length - len(middle)
	startOffset := targetSting/2 + targetSting%2
	endOffset := len(str) - targetSting/2

	return SubString(str, 0, startOffset) +
		middle +
		SubString(str, endOffset, -1)
}

// Capitalize 将字符串首字符转换为大写形式（非 ASCII 字符保持原样）。
//
// 示例：
//   - Capitalize("hello") == "Hello"
//   - Capitalize("HEllo") == "HEllo"（首字符已经大写则不变）
//   - Capitalize("")      == ""
//   - Capitalize("12h")   == "12h"（数字开头则不变）
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	b := []byte(s)
	b[0] = []byte(strings.ToUpper(string(b[0])))[0]
	return string(b)
}

// Uncapitalize 将字符串首字符转换为小写形式（非 ASCII 字符保持原样）。
//
// 示例：
//   - Uncapitalize("hello") == "hello"
//   - Uncapitalize("HEllo") == "hEllo"（仅首个字符变小写）
//   - Uncapitalize("")      == ""
//   - Uncapitalize("12h")   == "12h"（数字开头则不变）
func Uncapitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	b := []byte(s)
	b[0] = []byte(strings.ToLower(string(b[0])))[0]
	return string(b)
}

// SubstringAfter 返回 s 中第一次出现 separator 之后的子串。
//
// 行为细节：
//   - 当 s 为空时返回空串；
//   - 当 separator 为空时返回 EMPTY_STRING（无明确分隔点）；
//   - 当 separator 不在 s 中出现时返回 EMPTY_STRING。
func SubstringAfter(s string, separator string) string {
	if len(s) == 0 {
		return s
	}

	if len(separator) == 0 {
		return EMPTY_STRING
	}

	pos := strings.Index(s, separator)
	if pos == INDEX_NOT_FOUND {
		return EMPTY_STRING
	}

	return string(s[pos+len(separator):])
}

// SubstringAfterLast 返回 s 中最后一次出现 separator 之后的子串。
//
// 行为细节：
//   - 当 s 为空时返回空串；
//   - 当 separator 为空或未在 s 中出现时返回 EMPTY_STRING；
//   - 当 separator 出现在 s 末尾（即后无内容）时也返回 EMPTY_STRING。
func SubstringAfterLast(s string, separator string) string {
	if len(s) == 0 {
		return s
	}

	if len(separator) == 0 {
		return EMPTY_STRING
	}

	pos := strings.LastIndex(s, separator)
	if pos == INDEX_NOT_FOUND || pos == len(s)-len(separator) {
		return EMPTY_STRING
	}

	return string(s[pos+len(separator):])
}

// SubstringBefore 返回 s 中第一次出现 separator 之前的子串。
//
// 行为细节：
//   - 当 s 为空时返回空串；
//   - 当 separator 为空或未在 s 中出现时返回 EMPTY_STRING。
func SubstringBefore(s string, separator string) string {
	if len(s) == 0 {
		return s
	}

	if len(separator) == 0 {
		return EMPTY_STRING
	}

	pos := strings.Index(s, separator)
	if pos == INDEX_NOT_FOUND {
		return EMPTY_STRING
	}

	return string(s[:pos])
}

// SubstringBeforeLast 返回 s 中最后一次出现 separator 之前的子串。
//
// 行为细节：
//   - 当 s 为空时返回空串；
//   - 当 separator 为空时返回 EMPTY_STRING；
//   - 当 separator 未在 s 中出现时直接返回 s 本身（与 SubstringBefore 不同）。
func SubstringBeforeLast(s string, separator string) string {
	if len(s) == 0 {
		return s
	}

	if len(separator) == 0 {
		return EMPTY_STRING
	}

	pos := strings.LastIndex(s, separator)
	if pos == INDEX_NOT_FOUND {
		return s
	}
	return string(s[:pos])
}

// SubstringMatch 判断 s 从 index 位置开始的字节序列是否与 sub 完全相等。
//
// 行为细节：
//   - 当 index+len(sub) 超出 s 的长度时直接返回 false（不会越界访问）；
//   - 通过逐字节比较确认匹配，匹配返回 true，否则 false；
//   - 不调用正则，纯字节比较，性能可控。
func SubstringMatch(s string, index int, sub string) bool {
	if index+len(sub) > len(s) {
		return false
	}

	for i := 0; i < len(sub); i++ {
		if s[index+i] != sub[i] {
			return false
		}
	}
	return true
}

// Repeat 把单字节 s 重复 count 次拼接成字符串返回。
//
// 行为细节：
//   - count <= 0 时由 xarray.CreateAndFill 的语义决定结果（一般返回空串）；
//   - 返回值等价于 strings.Repeat(string(s), count)，但只接受单字节，节省分配。
func Repeat(s byte, count int) string {
	ret := xarray.CreateAndFill(count, s)
	return string(ret)
}

// StringToSlice 基于 unsafe 把 string 零拷贝转换为 []byte。
//
// 关键点：将 string 的 (Data, Len) 字段直接写入切片头，再额外设置 Cap。
//
// 注意：
//   - 返回的 []byte 与原 string 共享底层内存，修改 []byte 会影响 string；
//   - 仅在确实需要避免拷贝时才使用，且应避免在 string 仍被其他地方引用时写入。
func StringToSlice(value string) []byte {
	// create a new []byte
	var ret []byte

	// 把string的引用指向 ret的空间
	*(*string)(unsafe.Pointer(&ret)) = value

	// 设置slice的Cap值 ，用unsafe.Add操作，执行偏移操作 16个字节
	offset := uintptr(8) * 2
	*(*int)(unsafe.Add(unsafe.Pointer(&ret), offset)) = len(value)

	return ret
}

// SliceToString 基于 unsafe 把 []byte 零拷贝转换为 string。
//
// 关键点：将切片的 (Data, Len) 字段直接写入 string 头。
//
// 注意：
//   - 当 b 为 nil 时返回 EMPTY_STRING，避免解引用空切片；
//   - 返回的 string 与原切片共享底层内存，修改 b 会改变 string 内容（违反 Go 的不可变约定，慎用）。
func SliceToString(b []byte) string {
	if b == nil {
		return EMPTY_STRING
	}

	// just share Slice's Data and Len content
	return *(*string)(unsafe.Pointer(&b))
}

// IsEmpty return if s is empty
func IsEmpty(s string) bool {
	if s == EMPTY_STRING || len(s) == 0 {
		return true
	}
	return false
}

// IsBlank return if s is empty or blank string
// IsBlank("")  == true
// IsBlank(" ")  == true
// IsBlank(" a ")  == false
func IsBlank(s string) bool {
	if IsEmpty(s) || strings.TrimSpace(s) == EMPTY_STRING {
		return true
	}
	return false
}

// IndexFromOffset to return index of sub from Index
func IndexFromOffset(s, sub string, fromIndex int) int {
	if fromIndex < 0 || fromIndex+len(sub) >= len(s) {
		return -1
	}
	left := SubString(s, fromIndex, len(s))
	index := strings.Index(left, sub)
	if index != -1 {
		index += fromIndex
	}
	return index
}

// Wrap Wraps a String with another String.
func Wrap(s string, wrap string) string {
	if IsEmpty(s) || IsEmpty(wrap) {
		return s
	}
	b := make([]byte, len(s)+2*len(wrap))
	b1 := xunsafe.StringToSlice(s)
	b2 := xunsafe.StringToSlice(wrap)
	copy(b, b2)
	copy(b[len(wrap):], b1)
	copy(b[len(s)+len(wrap):], b2)

	return xunsafe.SliceToString(b)
}

// Expand to solve place holder value with prefix and suffix and replace by 'fn' call back function
func Expand(s, prefix, suffix string, fn tools.Func[string, string]) (string, error) {
	if IsBlank(prefix) || IsBlank(suffix) {
		return EMPTY_STRING, fmt.Errorf("invalid prefix or suffix")
	}

	startIndex := strings.Index(s, prefix)
	if startIndex == -1 {
		return s, nil
	}

	visitedPlaceholders := sets.NewSet[string]()

	for startIndex != -1 {
		endIndex := findPlaceholderEndIndex(s, startIndex, prefix, suffix)
		if endIndex != -1 {
			placeholder := SubString(s, startIndex+len(prefix), endIndex)
			originalPlaceholder := placeholder

			if !visitedPlaceholders.Add(originalPlaceholder) {
				return EMPTY_STRING, fmt.Errorf("circular placeholder reference %s in property definitions", originalPlaceholder)
			}

			// Recursive invocation, parsing placeholders contained in the placeholder key.
			placeholder, err := Expand(placeholder, prefix, suffix, fn)
			if err != nil {
				return EMPTY_STRING, err
			}

			// Now obtain the value for the fully resolved key...
			propVal := fn(placeholder)
			if !IsBlank(propVal) {
				propVal, err := Expand(propVal, prefix, suffix, fn)
				if err != nil {
					return EMPTY_STRING, err
				}
				s = ReplaceByOffset(s, startIndex, endIndex+len(suffix), propVal)
				offset := startIndex + len(propVal)
				startIndex = IndexFromOffset(s, prefix, offset)
			} else {
				offset := endIndex + len(prefix)
				startIndex = IndexFromOffset(s, prefix, offset)
			}
			visitedPlaceholders.Remove(originalPlaceholder)
		} else {
			startIndex = -1
		}

	}

	return s, nil
}

// ReplaceByOffset to replace sub string by offset begin and end index
func ReplaceByOffset(s string, begin, end int, replace string) string {
	sz := len(s)
	if begin > sz || end > sz {
		return s
	}
	if begin > end {
		return s
	}

	bb := []byte(s)
	b1 := bb[:begin]
	ret := string(b1) + replace
	b2 := bb[end:]
	return ret + string(b2)
}

func findPlaceholderEndIndex(buf string, startIndex int, prefix, suffix string) int {
	index := startIndex + len(prefix)
	withinNestedPlaceholder := 0
	for index < len(buf) {
		if SubstringMatch(buf, index, suffix) {
			if withinNestedPlaceholder > 0 {
				withinNestedPlaceholder--
				index = index + len(suffix)
			} else {
				return index
			}
		} else if SubstringMatch(buf, index, prefix) {
			withinNestedPlaceholder++
			index = index + len(prefix)
		} else {
			index++
		}
	}

	return -1
}

// HasSuffix to check if s has suffix
func HasSuffix(s string, suffix ...string) bool {
	if IsEmpty(s) || len(suffix) == 0 {
		return false
	}
	for _, v := range suffix {
		if strings.HasSuffix(s, v) {
			return true
		}
	}
	return false
}

// HasPrefix to check if s has prefix
func HasPrefix(s string, prefix ...string) bool {
	if IsEmpty(s) || len(prefix) == 0 {
		return false
	}
	for _, v := range prefix {
		if strings.HasPrefix(s, v) {
			return true
		}
	}
	return false
}

var (
	sensitiveStrings []string
	sensitiveRegex   *regexp.Regexp
)

func init() {
	// 定义敏感字符串列表
	sensitiveStrings = []string{
		"<script>", "</script>", "<style>", "</style>", "<iframe>", "</iframe>", "<img>", "</img>",
		"<video>", "</video>", "<audio>", "</audio>", "<source>", "</source>", "<track>", "</track>",
		"<object>", "</object>", "<embed>", "</embed>", "<form>", "</form>", "<button>", "</button>",
		"<input>", "</input>", "<textarea>", "</textarea>", "<select>", "</select>", "<option>", "</option>",
		"<frame>", "</frame>", "<frameset>", "</frameset>", "<noframes>", "</noframes>", "<marquee>", "</marquee>",
		"<blink>", "</blink>", "<link>", "</link>", "<meta>", "</meta>", "<base>", "</base>", "<basefont>", "</basefont>",
		"<applet>", "</applet>", "<param>", "</param>", "<object>", "</object>", "<embed>", "</embed>",
		"select", "delete", "update", "where", "drop", "truncate", "exec", "execute", "declare", "master", "create",
	}

	// 构建正则表达式，忽略大小写，并对敏感字符串进行转义处理
	pattern := "(?i)" + strings.Join(sensitiveStrings, "|")
	sensitiveRegex = regexp.MustCompile(pattern)
}

// GetSafeString 过滤敏感字符串
func GetSafeString(data string) string {
	if data == "" {
		return ""
	}
	return sensitiveRegex.ReplaceAllString(data, "")
}
