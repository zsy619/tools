package xstring

// Len 字符长度在范围内
func Len(s string, min, max int, errFunc func(min, max int) error) error {
	if len(s) < min || len(s) > max {
		return errFunc(min, max)
	}
	return nil
}

// StringToCharset 字符串转字符集
func StringToCharset(s string) map[rune]struct{} {
	m := map[rune]struct{}{}
	for _, c := range s {
		m[c] = struct{}{}
	}
	return m
}

// InCharset 字符串在字符集里面
func InCharset(s string, charset map[rune]struct{}, errFunc func(invalidChar rune) error) error {
	for _, c := range s {
		if _, ok := charset[c]; !ok {
			return errFunc(c)
		}
	}
	return nil
}

// IncludeCharset 字符串包含字符集
func IncludeCharset(s string, charset map[rune]struct{}, errFunc func() error) error {
	if !includeCharset(s, charset) {
		return errFunc()
	}
	return nil
}

// includeCharset 判断字符串s中是否包含charset中的字符
//
// 如果包含，则返回true；否则返回false
//
// 参数：
//   - s：待检查的字符串
//   - charset：包含字符的集合，使用map[rune]struct{}表示
//
// 返回值：
//   - bool类型，表示字符串s中是否包含charset中的字符
func includeCharset(s string, charset map[rune]struct{}) bool {
	for _, c := range s {
		if _, ok := charset[c]; ok {
			return true
		}
	}
	return false
}

// IncludeCharsets 检查字符串 s 是否包含字符集列表 charsets 中的所有字符集，
// 如果 s 不包含某个字符集中的字符，则调用 errFunc 函数，并将该字符集作为参数传入。
// 参数：
//   - s: 待检查的字符串
//   - charsets: 字符集列表，每个字符集为一个 rune 类型的映射
//   - errFunc: 回调函数，当 s 不包含某个字符集中的字符时调用，参数为未包含的字符集
//
// 返回值：
//   - error: 如果 s 不包含某个字符集中的字符，则返回 errFunc 函数的返回值，否则返回 nil
func IncludeCharsets(s string, charsets []map[rune]struct{}, errFunc func(notIncludedCharset map[rune]struct{}) error) error {
	for _, charset := range charsets {
		if !includeCharset(s, charset) {
			return errFunc(charset)
		}
	}
	return nil
}

// IncludeCharsetsCount 计算字符串s中符合charsets字符集要求的字符集数量，并判断是否满足min和max的限制。
//
// 如果数量不满足限制，则返回errFunc的返回值，否则返回nil。
//
// 参数：
//
//	s：待检查的字符串。
//	min：最小字符集数量限制。
//	max：最大字符集数量限制。
//	charsets：待检查的字符集列表，每个字符集为map[rune]struct{}类型。
//	errFunc：当字符集数量不满足限制时调用的错误处理函数，函数签名func(min, max, count int) error。
//
// 返回值：
//
//	若字符集数量满足限制，则返回nil；否则返回errFunc的返回值。
func IncludeCharsetsCount(s string, min, max int, charsets []map[rune]struct{}, errFunc func(min, max, count int) error) error {
	cnt := 0
	for _, charset := range charsets {
		if includeCharset(s, charset) {
			cnt++
		}
	}
	if cnt < min || cnt > max {
		return errFunc(min, max, cnt)
	}
	return nil
}
