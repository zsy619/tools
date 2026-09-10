package xphp

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"math"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/zsy619/tools/xjson"
)

// StringSaveTo 对输入字符串做保存前处理：先将双引号替换为 HTML 实体 &#34;，再调用 Stripslashes 去除转义反斜杠。
func StringSaveTo(intput string) string {
	rt := strings.ReplaceAll(intput, "\"", "&#34;")
	rt = Stripslashes(rt)
	return rt
}

// Htmlspecialchars_decode 将 &lt;、&gt;、&amp;、&quot;、&#039; 等 HTML 实体还原为对应的 <、>、&、"、' 字符（对应 PHP htmlspecialchars_decode()）。
func Htmlspecialchars_decode(s string) string {
	s = strings.Replace(s, "&lt;", "<", -1)
	s = strings.Replace(s, "&gt;", ">", -1)
	s = strings.Replace(s, "&amp;", "&", -1)
	s = strings.Replace(s, `&quot;`, `"`, -1)
	s = strings.Replace(s, `&#039;`, `'`, -1)
	return s
}

// Htmlspecialchars 将 <、>、&、"、单引号分别替换为 &lt;、&gt;、&amp;、&quot;、&#039;（对应 PHP htmlspecialchars()）。
func Htmlspecialchars(s string) string {
	s = strings.Replace(s, "<", "&lt;", -1)
	s = strings.Replace(s, ">", "&gt;", -1)
	s = strings.Replace(s, "&", "&amp;", -1)
	s = strings.Replace(s, `"`, "&quot;", -1)
	s = strings.Replace(s, `'`, "&#039;", -1)
	return s
}

// Strrev 将字符串 s 中的字符顺序反转（对应 PHP strrev()，按 rune 处理）。
func Strrev(s string) string {
	if s == "" {
		return s
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	length := len([]rune(s))
	var re []rune = make([]rune, length)
	i := 0
	for _, v := range s {
		re[length-i-1] = v
		i++
	}
	return string(re)
}

// Chr 返回 ASCII 码 ascii 对应的单字符字符串（对应 PHP chr()）；ascii 超出 32-127 范围时返回错误。
func Chr(ascii int) (string, error) {
	if ascii > 127 || ascii < 32 {
		return "", errors.New("invalid ascii code")
	}
	var buf bytes.Buffer
	buf.Write([]byte{byte(ascii)})
	return buf.String(), nil
}

/**
 *使用此函数将字符串分割成小块非常有用。
 *例如将 base64_encode() 的输出转换成符合 RFC 2045 语义的字符串。
 *它会在每 chunklen 个字符后边插入 end。(按byte计数，不是rune！)
 */
func Chunk_split(body string, chunklen int, end string) (str string, err error) {
	count := int(math.Ceil(float64(len(body) / chunklen)))
	var buf bytes.Buffer
	for i := 0; i < count-1; i++ { // 防止访问越界
		start := i * chunklen
		_, err = buf.Write([]byte(body)[start : start+chunklen])
		if err != nil {
			break
		}
		_, err = buf.WriteString(end)
		if err != nil {
			break
		}
	}
	_, err = buf.Write([]byte(body)[chunklen*(count-1):])
	if err != nil {
		return "", err
	}
	_, err = buf.WriteString(end)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Strpos 在 haystack 中从 offset 起查找 needle 首次出现的位置（字节下标，offset 可为负数；对应 PHP strpos()）。找不到时返回 -1。
func Strpos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// Stripos 是不区分大小写的 Strpos()（对应 PHP stripos()）。
func Stripos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	haystack = haystack[offset:]
	if offset < 0 {
		offset += length
	}
	pos := strings.Index(strings.ToLower(haystack), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// Strrpos 在 haystack 中查找 needle 最后一次出现的位置（字节下标，支持负数 offset；对应 PHP strrpos()）。找不到时返回 -1。
func Strrpos(haystack, needle string, offset int) int {
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}
	pos = strings.LastIndex(haystack, needle)
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// Strripos 是不区分大小写的 Strrpos()（对应 PHP strripos()）。
func Strripos(haystack, needle string, offset int) int {
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}
	pos = strings.LastIndex(strings.ToLower(haystack), strings.ToLower(needle))
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// StrReplace 将 subject 中前 count 次出现的 search 替换为 replace（对应 PHP str_replace()）。
func StrReplace(search, replace, subject string, count int) string {
	return strings.Replace(subject, search, replace, count)
}

// StrReplaceNoLimit 将 subject 中所有出现的 search 替换为 replace（等价于替换次数无限制的 StrReplace）。
func StrReplaceNoLimit(search, replace, subject string) string {
	return strings.Replace(subject, search, replace, -1)
}

// StrReplaceEnter 移除 subject 中的所有换行符（\r\n、\r、\n），用于将多行文本压缩为单行。
func StrReplaceEnter(subject string) (result string) {
	result = subject
	result = strings.ReplaceAll(result, "\r\n", "")
	result = strings.ReplaceAll(result, "\r", "")
	result = strings.ReplaceAll(result, "\n", "")
	return
}

// Strtoupper 将 str 中的字母全部转为大写（对应 PHP strtoupper()）。
func Strtoupper(str string) string {
	return strings.ToUpper(str)
}

// Strtolower 将 str 中的字母全部转为小写（对应 PHP strtolower()）。
func Strtolower(str string) string {
	return strings.ToLower(str)
}

// Ucfirst 将 str 的首字母转为大写（对应 PHP ucfirst()）。
func Ucfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToUpper(v))
		return u + str[len(u):]
	}
	return ""
}

// Lcfirst 将 str 的首字母转为小写（对应 PHP lcfirst()）。
func Lcfirst(str string) string {
	for _, v := range str {
		u := string(unicode.ToLower(v))
		return u + str[len(u):]
	}
	return ""
}

// Ucwords 将 str 中每个单词的首字母转为大写（对应 PHP ucwords()）。
func Ucwords(str string) string {
	caser := cases.Title(language.English)
	return caser.String(str)
}

// Substr 返回 str 从 start 起、长度为 length 的子串（对应 PHP substr()；length 为 -1 表示截到末尾，为 0 返回空串，按字节处理）。
func Substr(str string, start uint, length int) string {
	if length < -1 {
		return str
	}
	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}
	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}

// ParseStr 将 URL 编码的查询字符串解析到 result 中，支持嵌套 key（对应 PHP parse_str()）。
// f1=m&f2=n -> map[f1:m f2:n]
// f[a]=m&f[b]=n -> map[f:map[a:m b:n]]
// f[a][a]=m&f[a][b]=n -> map[f:map[a:map[a:m b:n]]]
// f[]=m&f[]=n -> map[f:[m n]]
// f[a][]=m&f[a][]=n -> map[f:map[a:[m n]]]
// f[][]=m&f[][]=n -> map[f:[map[]]] // 目前不支持嵌套切片。
// f=m&f[a]=n -> error // 这与 PHP 的行为不同。
// a .[[b=c -> map[a___[b:c]
func ParseStr(encodedString string, result map[string]interface{}) error {
	// build nested map.
	var build func(map[string]interface{}, []string, interface{}) error

	build = func(result map[string]interface{}, keys []string, value interface{}) error {
		length := len(keys)
		// trim ',"
		key := strings.Trim(keys[0], "'\"")
		if length == 1 {
			result[key] = value
			return nil
		}

		// The end is slice. like f[], f[a][]
		if keys[1] == "" && length == 2 {
			// todo nested slice
			if key == "" {
				return nil
			}
			val, ok := result[key]
			if !ok {
				result[key] = []interface{}{value}
				return nil
			}
			children, ok := val.([]interface{})
			if !ok {
				return fmt.Errorf("expected type '[]interface{}' for key '%s', but got '%T'", key, val)
			}
			result[key] = append(children, value)
			return nil
		}

		// The end is slice + map. like f[][a]
		if keys[1] == "" && length > 2 && keys[2] != "" {
			val, ok := result[key]
			if !ok {
				result[key] = []interface{}{}
				val = result[key]
			}
			children, ok := val.([]interface{})
			if !ok {
				return fmt.Errorf("expected type '[]interface{}' for key '%s', but got '%T'", key, val)
			}
			if l := len(children); l > 0 {
				if child, ok := children[l-1].(map[string]interface{}); ok {
					if _, ok := child[keys[2]]; !ok {
						_ = build(child, keys[2:], value)
						return nil
					}
				}
			}
			child := map[string]interface{}{}
			_ = build(child, keys[2:], value)
			result[key] = append(children, child)

			return nil
		}

		// map. like f[a], f[a][b]
		val, ok := result[key]
		if !ok {
			result[key] = map[string]interface{}{}
			val = result[key]
		}
		children, ok := val.(map[string]interface{})
		if !ok {
			return fmt.Errorf("expected type 'map[string]interface{}' for key '%s', but got '%T'", key, val)
		}

		return build(children, keys[1:], value)
	}

	// split encodedString.
	parts := strings.Split(encodedString, "&")
	for _, part := range parts {
		pos := strings.Index(part, "=")
		if pos <= 0 {
			continue
		}
		key, err := url.QueryUnescape(part[:pos])
		if err != nil {
			return err
		}
		for key[0] == ' ' {
			key = key[1:]
		}
		if key == "" || key[0] == '[' {
			continue
		}
		value, err := url.QueryUnescape(part[pos+1:])
		if err != nil {
			return err
		}

		// split into multiple keys
		var keys []string
		left := 0
		for i, k := range key {
			if k == '[' && left == 0 {
				left = i
			} else if k == ']' {
				if left > 0 {
					if len(keys) == 0 {
						keys = append(keys, key[:left])
					}
					keys = append(keys, key[left+1:i])
					left = 0
					if i+1 < len(key) && key[i+1] != '[' {
						break
					}
				}
			}
		}
		if len(keys) == 0 {
			keys = append(keys, key)
		}
		// first key
		first := ""
		for i, chr := range keys[0] {
			if chr == ' ' || chr == '.' || chr == '[' {
				first += "_"
			} else {
				first += string(chr)
			}
			if chr == '[' {
				first += keys[0][i+1:]
				break
			}
		}
		keys[0] = first

		// build nested map
		if err := build(result, keys, value); err != nil {
			return err
		}
	}

	return nil
}

// NumberFormat 对 number 进行千分位格式化（对应 PHP number_format()）。
// decimals：设置小数位数（四舍五入）。
// decPoint：设置小数点的分隔符。
// thousandsSep：设置千位分隔符。
func NumberFormat(number float64, decimals uint, decPoint, thousandsSep string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimals)
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousandsSep)
	n, l1, l2 := 0, len(prefix), len(sep)
	// thousands sep num
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	s := string(tmp)
	if dec > 0 {
		s += decPoint + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

// ChunkSplit 将 body 按 chunklen 个字符（rune）切分成小块，并用 end 连接每一块（对应 PHP chunk_split()）。
func ChunkSplit(body string, chunklen uint, end string) string {
	if end == "" {
		end = "\r\n"
	}
	runes, erunes := []rune(body), []rune(end)
	l := uint(len(runes))
	if l <= 1 || l < chunklen {
		return body + end
	}
	ns := make([]rune, 0, len(runes)+len(erunes))
	var i uint
	for i = 0; i < l; i += chunklen {
		if i+chunklen > l {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+chunklen]...)
		}
		ns = append(ns, erunes...)
	}
	return string(ns)
}

// StrWordCount 返回 str 中的单词列表，单词按空白字符分隔（对应 PHP str_word_count()）。
func StrWordCount(str string) []string {
	return strings.Fields(str)
}

// Wordwrap 按 width 宽度对 str 进行换行，br 为换行符；cut 为 true 时允许在超宽单词中间强制断开（对应 PHP wordwrap()）。
func Wordwrap(str string, width uint, br string, cut bool) string {
	strlen := len(str)
	brlen := len(br)
	linelen := int(width)

	if strlen == 0 {
		return ""
	}
	if brlen == 0 {
		panic("break string cannot be empty")
	}
	if linelen == 0 && cut {
		panic("can't force cut when width is zero")
	}

	current, laststart, lastspace := 0, 0, 0
	var ns []byte
	for current = 0; current < strlen; current++ {
		if str[current] == br[0] && current+brlen < strlen && str[current:current+brlen] == br {
			ns = append(ns, str[laststart:current+brlen]...)
			current += brlen - 1
			lastspace = current + 1
			laststart = lastspace
		} else if str[current] == ' ' {
			if current-laststart >= linelen {
				ns = append(ns, str[laststart:current]...)
				ns = append(ns, br[:]...)
				laststart = current + 1
			}
			lastspace = current
		} else if current-laststart >= linelen && cut && laststart >= lastspace {
			ns = append(ns, str[laststart:current]...)
			ns = append(ns, br[:]...)
			laststart = current
			lastspace = current
		} else if current-laststart >= linelen && laststart < lastspace {
			ns = append(ns, str[laststart:lastspace]...)
			ns = append(ns, br[:]...)
			lastspace++
			laststart = lastspace
		}
	}

	if laststart != current {
		ns = append(ns, str[laststart:current]...)
	}
	return string(ns)
}

// Strlen 返回 str 的字节长度（对应 PHP strlen()）。
func Strlen(str string) int {
	return len(str)
}

// MbStrlen 返回 str 的字符数，按 rune（Unicode 码点）计数（对应 PHP mb_strlen()）。
func MbStrlen(str string) int {
	return utf8.RuneCountInString(str)
}

// MbSubstr 返回 str 从 start 起、长度为 length 的子串（按 rune 处理，start 与 length 可为负；对应 PHP mb_substr()）。
func MbSubstr(str string, start, length int) string {
	runes := []rune(str)
	if start < 0 {
		start = 0
	}
	if start > len(runes) {
		start = len(runes)
	}
	end := start + length
	if end < 0 {
		end = 0
	}
	if end > len(runes) {
		end = len(runes)
	}
	return string(runes[start:end])
}

// StrRepeat 将 input 重复 multiplier 次后拼接返回（对应 PHP str_repeat()）。
func StrRepeat(input string, multiplier int) string {
	return strings.Repeat(input, multiplier)
}

// Strstr 在 haystack 中查找 needle 首次出现的位置，并返回该匹配之后的剩余子串（对应 PHP strstr()）。找不到或 needle 为空时返回空串。
func Strstr(haystack string, needle string) string {
	if needle == "" {
		return ""
	}
	idx := strings.Index(haystack, needle)
	if idx == -1 {
		return ""
	}
	return haystack[idx+len([]byte(needle))-1:]
}

// Strtr 按替换规则转换 haystack 中的字符（对应 PHP strtr()）。
//
// 若参数数量为 1，参数类型为 map[string]string：
// Strtr("baab", map[string]string{"ab": "01"}) 将返回 "ba01"
// 若参数数量为 2，参数类型为 string, string：
// Strtr("baab", "ab", "01") 将返回 "1001"，即 a => 0；b => 1。
func Strtr(haystack string, params ...interface{}) string {
	ac := len(params)
	if ac == 1 {
		pairs := params[0].(map[string]string)
		length := len(pairs)
		if length == 0 {
			return haystack
		}
		oldnew := make([]string, length*2)
		for o, n := range pairs {
			if o == "" {
				return haystack
			}
			oldnew = append(oldnew, o, n)
		}
		return strings.NewReplacer(oldnew...).Replace(haystack)
	} else if ac == 2 {
		from := params[0].(string)
		to := params[1].(string)
		trlen, lt := len(from), len(to)
		if trlen > lt {
			trlen = lt
		}
		if trlen == 0 {
			return haystack
		}

		str := make([]uint8, len(haystack))
		var xlat [256]uint8
		var i int
		var j uint8
		if trlen == 1 {
			for i = 0; i < len(haystack); i++ {
				if haystack[i] == from[0] {
					str[i] = to[0]
				} else {
					str[i] = haystack[i]
				}
			}
			return string(str)
		}
		// trlen != 1
		for {
			xlat[j] = j
			if j++; j == 0 {
				break
			}
		}
		for i = 0; i < trlen; i++ {
			xlat[from[i]] = to[i]
		}
		for i = 0; i < len(haystack); i++ {
			str[i] = xlat[haystack[i]]
		}
		return string(str)
	}

	return haystack
}

// StrShuffle 随机打乱 str 中字符的顺序（对应 PHP str_shuffle()）。
func StrShuffle(str string) string {
	runes := []rune(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]rune, len(runes))
	for i, v := range r.Perm(len(runes)) {
		s[i] = runes[v]
	}
	return string(s)
}

// Trim 去除 str 首尾的空白字符；若提供了 characterMask，则去除其中指定的字符（对应 PHP trim()）。
func Trim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimSpace(str)
	}
	return strings.Trim(str, characterMask[0])
}

// Ltrim 去除 str 开头的空白字符；若提供了 characterMask，则去除其中指定的字符（对应 PHP ltrim()）。
func Ltrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimLeftFunc(str, unicode.IsSpace)
	}
	return strings.TrimLeft(str, characterMask[0])
}

// Rtrim 去除 str 末尾的空白字符；若提供了 characterMask，则去除其中指定的字符（对应 PHP rtrim()）。
func Rtrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimRightFunc(str, unicode.IsSpace)
	}
	return strings.TrimRight(str, characterMask[0])
}

// Explode 以 delimiter 为分隔符将 str 拆分为字符串切片（对应 PHP explode()）。
func Explode(delimiter, str string) []string {
	return strings.Split(str, delimiter)
}

// ExplodeComma 以逗号 "," 为分隔符将字符串 str 拆分。
func ExplodeComma(str string) []string {
	return strings.Split(str, ",")
}

// Ord 返回 char 中首个字符的 ASCII/Unicode 码点（对应 PHP ord()）。
func Ord(char string) int {
	r, _ := utf8.DecodeRune([]byte(char))
	return int(r)
}

// Nl2br 将字符串中的换行符替换为 <br />（isXhtml 为 true）或 <br> 标签（对应 PHP nl2br()）。
// 可识别的换行形式：\n\r、\r\n、\r、\n。
func Nl2br(str string, isXhtml bool) string {
	r, n, runes := '\r', '\n', []rune(str)
	var br []byte
	if isXhtml {
		br = []byte("<br />")
	} else {
		br = []byte("<br>")
	}
	skip := false
	length := len(runes)
	var buf bytes.Buffer
	for i, v := range runes {
		if skip {
			skip = false
			continue
		}
		switch v {
		case n, r:
			if (i+1 < length) && (v == r && runes[i+1] == n) || (v == n && runes[i+1] == r) {
				buf.Write(br)
				skip = true
				continue
			}
			buf.Write(br)
		default:
			buf.WriteRune(v)
		}
	}
	return buf.String()
}

// Json_decode 将 JSON 字符串 data 解码为 map[string]interface{}。
func Json_decode(data string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	err := xjson.Unmarshal([]byte(data), &dat)
	return dat, err
}

// JSONDecode 将 JSON 数据 data 解码到 val 指向的变量中（对应 PHP json_decode()）。
func JSONDecode(data []byte, val interface{}) error {
	return xjson.Unmarshal(data, val)
}

// JSONEncode 将 val 编码为 JSON 数据（对应 PHP json_encode()）。
func JSONEncode(val interface{}) ([]byte, error) {
	return xjson.Marshal(val)
}

// Addslashes 在单引号、双引号与反斜杠前添加转义反斜杠（对应 PHP addslashes()）。
func Addslashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// Stripslashes 去除字符串中的转义反斜杠（对应 PHP stripslashes()）。
func Stripslashes(str string) string {
	var buf bytes.Buffer
	l, skip := len(str), false
	for i, char := range str {
		if skip {
			skip = false
		} else if char == '\\' {
			if i+1 < l && str[i+1] == '\\' {
				skip = true
			}
			continue
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// Quotemeta 在正则表达式元字符（. + \ ( $ ) [ ^ ] * ? 等）前添加转义反斜杠（对应 PHP quotemeta()）。
func Quotemeta(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '.', '+', '\\', '(', '$', ')', '[', '^', ']', '*', '?':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// Htmlentities 将字符串转换为对应的 HTML 实体（对应 PHP htmlentities()）。
func Htmlentities(str string) string {
	return html.EscapeString(str)
}

// HTMLEntityDecode 将 HTML 实体解码为对应字符（对应 PHP html_entity_decode()）。
func HTMLEntityDecode(str string) string {
	return html.UnescapeString(str)
}

// Levenshtein 计算 str1 与 str2 之间的编辑距离（对应 PHP levenshtein()）。
// costIns：定义插入操作的代价。
// costRep：定义替换操作的代价。
// costDel：定义删除操作的代价。
// 任一字符串长度超过 255 字节时返回 -1。
func Levenshtein(str1, str2 string, costIns, costRep, costDel int) int {
	maxLen := 255
	l1 := len(str1)
	l2 := len(str2)
	if l1 == 0 {
		return l2 * costIns
	}
	if l2 == 0 {
		return l1 * costDel
	}
	if l1 > maxLen || l2 > maxLen {
		return -1
	}

	p1 := make([]int, l2+1)
	p2 := make([]int, l2+1)
	var c0, c1, c2 int
	var i1, i2 int
	for i2 := 0; i2 <= l2; i2++ {
		p1[i2] = i2 * costIns
	}
	for i1 = 0; i1 < l1; i1++ {
		p2[0] = p1[0] + costDel
		for i2 = 0; i2 < l2; i2++ {
			if str1[i1] == str2[i2] {
				c0 = p1[i2]
			} else {
				c0 = p1[i2] + costRep
			}
			c1 = p1[i2+1] + costDel
			if c1 < c0 {
				c0 = c1
			}
			c2 = p2[i2] + costIns
			if c2 < c0 {
				c0 = c2
			}
			p2[i2+1] = c0
		}
		tmp := p1
		p1 = p2
		p2 = tmp
	}
	c0 = p1[l2]

	return c0
}

// SimilarText 计算 first 与 second 中相同字符的数量；若 percent 非空，则把相似度百分比写入其中（对应 PHP similar_text()）。
func SimilarText(first, second string, percent *float64) int {
	var similarText func(string, string, int, int) int
	similarText = func(str1, str2 string, len1, len2 int) int {
		var sum, max int
		pos1, pos2 := 0, 0

		// Find the longest segment of the same section in two strings
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
					if l+1 > max {
						max = l + 1
						pos1 = i
						pos2 = j
					}
				}
			}
		}

		if sum = max; sum > 0 {
			if pos1 > 0 && pos2 > 0 {
				sum += similarText(str1, str2, pos1, pos2)
			}
			if (pos1+max < len1) && (pos2+max < len2) {
				s1 := []byte(str1)
				s2 := []byte(str2)
				sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
			}
		}

		return sum
	}

	l1, l2 := len(first), len(second)
	if l1+l2 == 0 {
		return 0
	}
	sim := similarText(first, second, l1, l2)
	if percent != nil {
		*percent = float64(sim*200) / float64(l1+l2)
	}
	return sim
}

// Soundex 计算字符串的 soundex 键，用于按英语发音近似匹配单词（对应 PHP soundex()）。
// 传入空字符串会触发 panic。
func Soundex(str string) string {
	if str == "" {
		panic("str: cannot be an empty string")
	}
	table := [26]rune{
		// A, B, C, D
		'0', '1', '2', '3',
		// E, F, G
		'0', '1', '2',
		// H
		'0',
		// I, J, K, L, M, N
		'0', '2', '2', '4', '5', '5',
		// O, P, Q, R, S, T
		'0', '1', '2', '6', '2', '3',
		// U, V
		'0', '1',
		// W, X
		'0', '2',
		// Y, Z
		'0', '2',
	}
	last, code, small := -1, 0, 0
	sd := make([]rune, 4)
	// build soundex string
	for i := 0; i < len(str) && small < 4; i++ {
		// ToUpper
		char := str[i]
		if char < '\u007F' && 'a' <= char && char <= 'z' {
			code = int(char - 'a' + 'A')
		} else {
			code = int(char)
		}
		if code >= 'A' && code <= 'Z' {
			if small == 0 {
				sd[small] = rune(code)
				small++
				last = int(table[code-'A'])
			} else {
				code = int(table[code-'A'])
				if code != last {
					if code != 0 {
						sd[small] = rune(code)
						small++
					}
					last = code
				}
			}
		}
	}
	// pad with "0"
	for ; small < 4; small++ {
		sd[small] = '0'
	}
	return string(sd)
}

// StrToLower 返回将字符串中所有字母字符转换为小写后的结果。
func StrToLower(s string) string {
	return strings.ToLower(s)
}
