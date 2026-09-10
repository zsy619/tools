// Package xstring 提供字符串与字符串切片相关的常用辅助函数集合。
//
// 该包聚焦两类能力：
//  1. 字符串切片的集合/函数式操作（类似 lodash 中针对 []string 的子集）。
//  2. 字符串长度校验、字符集校验等与字符/字符集相关的校验函数，
//     并通过回调函数（errFunc）将错误返回交由调用方控制，
//     从而支持本地化或自定义错误格式。
//
// 所有函数都不修改入参，无副作用，且不会 panic（除依赖调用方提供的 errFunc 的行为外）。
package xstring

// Len 检查字符串 s 的字节长度是否落在 [min, max] 区间内。
//
// 参数说明：
//   - s：待检查的字符串。注意：长度计算使用的是 len(s)，即字节数，而非 rune 数，
//     因此对多字节（中文字符等）的字符串需自行考虑是否合适。
//   - min：允许的最小字节长度（含）。
//   - max：允许的最大字节长度（含）。当 min > max 时，区间天然为空，任何 s 都会触发 errFunc。
//
// 返回值：
//   - 当 s 的长度在 [min, max] 区间内时返回 nil。
//   - 否则调用 errFunc(min, max) 并把其返回值原样返回。
//
// 不会 panic；但 errFunc 不应返回 nil，否则调用方需自行决定如何处理。
//
// 示例：
//
//	if err := xstring.Len(name, 1, 32, func(min, max int) error {
//	    return fmt.Errorf("名称长度必须在 %d-%d 之间", min, max)
//	}); err != nil {
//	    return err
//	}
func Len(s string, min, max int, errFunc func(min, max int) error) error {
	if len(s) < min || len(s) > max {
		return errFunc(min, max)
	}
	return nil
}

// StringToCharset 将字符串 s 转换为一个字符集合（去重后的 rune 集合）。
//
// 以 map[rune]struct{} 的形式返回，键为 s 中出现的每个 rune，值为空结构体。
//
// 参数说明：
//   - s：输入字符串。空字符串会返回非 nil 的空 map。
//
// 返回值：
//   - 一个包含 s 中所有不重复 rune 的集合；当 s 为空时集合为空，但 map 本身非 nil（已 make）。
//
// 用途：配合 InCharset / IncludeCharset 等函数做字符白名单/黑名单检查。
//
// 示例：
//
//	set := xstring.StringToCharset("abc") // { 'a': {}, 'b': {}, 'c': {} }
func StringToCharset(s string) map[rune]struct{} {
	m := map[rune]struct{}{}
	for _, c := range s {
		m[c] = struct{}{}
	}
	return m
}

// InCharset 检查字符串 s 中的每一个 rune 是否都出现在 charset 集合内。
//
// 适用于"字符串必须全部由 charset 内的字符组成"的场景，例如密码强度或字符白名单。
//
// 参数说明：
//   - s：待检查的字符串。空字符串视为满足条件（没有任何越界字符）。
//   - charset：允许出现的 rune 集合，通常由 StringToCharset 生成。
//   - errFunc：发现第一个不在 charset 中的 rune 时被调用，参数为该非法字符，
//     返回值会作为 InCharset 的返回值原样上抛。
//
// 返回值：
//   - 全部 rune 均在 charset 中时返回 nil。
//   - 否则返回 errFunc(invalidChar)。
//
// 不会 panic；如果 errFunc 为 nil 且触发了非法字符，会因调用 nil 产生 panic，
// 调用方需保证 errFunc 非空。
//
// 示例：
//
//	digits := xstring.StringToCharset("0123456789")
//	err := xstring.InCharset(input, digits, func(c rune) error {
//	    return fmt.Errorf("非法字符: %q", c)
//	})
func InCharset(s string, charset map[rune]struct{}, errFunc func(invalidChar rune) error) error {
	for _, c := range s {
		if _, ok := charset[c]; !ok {
			return errFunc(c)
		}
	}
	return nil
}

// IncludeCharset 检查字符串 s 中是否至少包含 charset 集合里的一个字符。
//
// 与 InCharset 相反，这里只要求"包含即可"，不做完整字符校验。
//
// 参数说明：
//   - s：待检查的字符串。空字符串永远不包含任何字符，必然触发 errFunc。
//   - charset：字符集合。
//   - errFunc：当 s 不包含 charset 中任何字符时被调用（无参数），返回值作为错误上抛。
//
// 返回值：
//   - 命中 charset 中至少一个 rune 时返回 nil。
//   - 否则返回 errFunc()。
//
// 不会 panic；但若 errFunc 为 nil 且未命中，会因调用 nil 产生 panic。
//
// 示例：
//
//	vowels := xstring.StringToCharset("aeiouAEIOU")
//	err := xstring.IncludeCharset(pwd, vowels, func() error {
//	    return errors.New("密码必须至少包含一个元音字母")
//	})
func IncludeCharset(s string, charset map[rune]struct{}, errFunc func() error) error {
	if !includeCharset(s, charset) {
		return errFunc()
	}
	return nil
}

// includeCharset 内部辅助函数，判断字符串 s 中是否包含 charset 集合里的任意一个字符。
//
// 返回 true 表示 s 至少包含 charset 中的一个 rune；空字符串或空 charset 始终返回 false。
//
// 参数：
//   - s：待检查的字符串。
//   - charset：rune 集合。
//
// 返回值：bool，true 表示命中至少一个字符。
func includeCharset(s string, charset map[rune]struct{}) bool {
	for _, c := range s {
		if _, ok := charset[c]; ok {
			return true
		}
	}
	return false
}

// IncludeCharsets 检查字符串 s 是否同时包含 charsets 列表中每一个字符集里的至少一个字符。
//
// 即对每个 charset 调用 includeCharset，一旦发现某个字符集未被命中，立即调用 errFunc
// 并把未命中的 charset 作为参数传入。
//
// 参数说明：
//   - s：待检查的字符串。
//   - charsets：字符集列表，按顺序检查；空列表会立即返回 nil。
//   - errFunc：当某个 charset 未被命中时被调用，参数为该 charset，返回值作为错误上抛。
//
// 返回值：
//   - 全部 charset 都命中时返回 nil。
//   - 否则返回 errFunc(notIncludedCharset)。
//
// 不会 panic；但若 errFunc 为 nil 且未命中，调用 nil 会 panic。
//
// 示例：
//
//	lower := xstring.StringToCharset("abcdefghijklmnopqrstuvwxyz")
//	upper := xstring.StringToCharset("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
//	digit := xstring.StringToCharset("0123456789")
//	err := xstring.IncludeCharsets(pwd, []map[rune]struct{}{lower, upper, digit},
//	    func(cs map[rune]struct{}) error {
//	        return errors.New("密码必须同时包含大小写字母与数字")
//	    })
func IncludeCharsets(s string, charsets []map[rune]struct{}, errFunc func(notIncludedCharset map[rune]struct{}) error) error {
	for _, charset := range charsets {
		if !includeCharset(s, charset) {
			return errFunc(charset)
		}
	}
	return nil
}

// IncludeCharsetsCount 统计字符串 s 中命中的字符集数量，并校验是否落在 [min, max] 区间内。
//
// 常用于"密码必须命中 N 类字符集（数字、小写、大写、符号）"这类按种类数量校验的场景。
//
// 参数说明：
//   - s：待检查的字符串。
//   - min：要求至少命中的字符集数量（含）。
//   - max：要求至多命中的字符集数量（含）。
//   - charsets：字符集列表，按顺序逐个判定；空列表会得到 cnt=0，落在 [min, max] 时返回 nil。
//   - errFunc：当 cnt 落在区间外时被调用，参数为 (min, max, cnt)，
//     其返回值会原样作为错误返回。
//
// 返回值：
//   - 命中字符集数量处于 [min, max] 时返回 nil。
//   - 否则返回 errFunc(min, max, cnt)。
//
// 不会 panic；但若 errFunc 为 nil 且需要触发错误，调用 nil 会 panic。
//
// 示例：
//
//	lower := xstring.StringToCharset("abcdefghijklmnopqrstuvwxyz")
//	upper := xstring.StringToCharset("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
//	digit := xstring.StringToCharset("0123456789")
//	symbol := xstring.StringToCharset("!@#$%^&*")
//	err := xstring.IncludeCharsetsCount(pwd, 2, 4,
//	    []map[rune]struct{}{lower, upper, digit, symbol},
//	    func(min, max, cnt int) error {
//	        return fmt.Errorf("密码需命中 %d-%d 类字符，实际 %d 类", min, max, cnt)
//	    })
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
