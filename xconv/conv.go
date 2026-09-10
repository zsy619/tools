// Package xconv 提供一组与「类型 / 进制 / 字符串」转换相关的工具函数。
//
// 设计目标：
//   - 使用 Go 1.21+ 引入的泛型，覆盖整型、浮点、字符串等多种类型；
//   - 兼容 Java/PHP 等其它语言中常见的字符串转数字、数字转字符串、
//     进制转换、布尔识别、中文大写数字解析等模式；
//   - 不引入任何第三方依赖，仅依赖标准库。
//
// 本文件常用 API 概览：
//   - CItoa：将中文大写数字（如 一亿二千三百四十五万）转换为阿拉伯数字字符串；
//   - Itoa/FormatInt：泛型版数字转字符串；
//   - ToPtr：把任意值快速装箱为指针；
//   - IsNumber：判断字符串是否符合多种进制的数字字面量格式；
//   - ParseInt/ParseFloat：在 IsNumber 基础上增加负数与指数支持；
//   - Append/AppendString：把任意值格式化为字节或字符串并追加；
//   - ParseBool：识别字符串或整型的布尔值。
package xconv

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/zsy619/tools/xarray"
	"github.com/zsy619/tools/xgeneric"
	"github.com/zsy619/tools/xgeneric/lists"
	"github.com/zsy619/tools/xstring"
)

const (
	// Binary、Octal、Hex 分别对应二进制、八进制、十六进制。
	//
	// 注意：Octal 沿用历史定义为 10，请勿直接当作 "8" 使用；如需真正的
	// 八进制请直接传入字面量 8。
	Binary = 2
	Octal  = 10
	Hex    = 16

	// FixedUnitValue 是中文大写数字中区分万 / 亿这类进位单位的阈值。
	FixedUnitValue = 4

	// 中文数字字符常量。
	CH_ZERO   = '零'
	CH_TEN    = '十'
	CH_ONE    = '一'
	RUNE_ONE  = 1
	RUNE_ZERO = 0
)

// 中文数字 / 单位映射表，在 init() 中填充，对调用方只读。
var (
	Chinese_Uint_Number map[rune]int
	Chinese_Number      map[rune]byte

	Chinese_Unit []rune
)

// CNum 表示中文数字解析过程中的一个最小单元：当前数字、所在单位、
// 是否位于万 / 亿这种固定进位单位上。
type CNum struct {
	num       byte
	unit      int
	fixedunit int
}

func init() {
	// 中文单位 -> 进位大小。
	Chinese_Uint_Number = map[rune]int{}
	Chinese_Uint_Number['亿'] = 8
	Chinese_Uint_Number['万'] = 4
	Chinese_Uint_Number['千'] = 3
	Chinese_Uint_Number['百'] = 2
	Chinese_Uint_Number[CH_TEN] = 1

	Chinese_Number = map[rune]byte{}
	Chinese_Number['一'] = RUNE_ONE
	Chinese_Number['二'] = 2
	Chinese_Number['三'] = 3
	Chinese_Number['四'] = 4
	Chinese_Number['五'] = 5
	Chinese_Number['六'] = 6
	Chinese_Number['七'] = 7
	Chinese_Number['七'] = 7
	Chinese_Number['八'] = 8
	Chinese_Number['九'] = 9
	Chinese_Number[CH_ZERO] = RUNE_ZERO
}

// CItoa 将中文大写数字字符串（如 一亿二千三百四十五万）转换为阿拉伯数字
// 字符串（如 123450000）。
//
// 支持以下常见模式：
//   - 单字符：零、一、二 ... 九；
//   - 复合：十五、一百二十三；
//   - 高级单位：万、亿，并能处理 万亿 这种嵌套。
//
// 返回值为阿拉伯数字字符串与 nil 错误；遇到无法识别的字符时返回
// ("", error)。
//
// 注意：本函数仅识别简体中文常见字「零 一 二 三 四 五 六 七 八 九
// 十 百 千 万 亿」，不包含大写金额字符（壹贰叁等）。
func CItoa(chinesenum string) (string, error) {
	if xstring.IsBlank(chinesenum) {
		return xstring.EMPTY_STRING, fmt.Errorf("invalid chinese number. %s", chinesenum)
	}

	ls := lists.NewList[CNum]()
	nums := []rune(chinesenum)
	unit := 1
	fixedunit := 0
	biggestUnit := 0
	hasNumBeforeUnit := false
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		if v, ok := Chinese_Number[num]; ok {
			cNum := CNum{v, unit, fixedunit}
			ls.PushFront(cNum)
			hasNumBeforeUnit = true
		} else if v, ok := Chinese_Uint_Number[num]; ok {
			if v >= FixedUnitValue {
				if biggestUnit < v {
					biggestUnit = v
					fixedunit = v // just ajust to fixed unit
				} else {
					// if just like '万亿' add num unit
					fixedunit += v
				}
				unit = 1
			} else {
				unit = v + 1
				// check if miss num before unit, so we check next one
				i2 := i - 1
				if i2 >= 0 {
					if v, ok := Chinese_Uint_Number[nums[i2]]; ok {
						if v < FixedUnitValue {
							cNum := CNum{RUNE_ONE, unit, fixedunit} // add one as default unit value
							ls.PushFront(cNum)
						}
					} else if nums[i2] == CH_ZERO {
						cNum := CNum{RUNE_ONE, unit, fixedunit} // add one as default unit value
						ls.PushFront(cNum)
					}
				}

			}
			hasNumBeforeUnit = false
		} else {
			return xstring.EMPTY_STRING, fmt.Errorf("invalid chinese number. %s", chinesenum)
		}
	}

	size := unit + fixedunit
	numbers := make([]byte, size)

	// fix if has no num before unit, just like as '十' or '十五'
	if !hasNumBeforeUnit {
		ls.PushFront(CNum{RUNE_ONE, unit, fixedunit})
	}

	ls.Range(func(c CNum) bool {
		if c.num != RUNE_ZERO {
			offset := size - c.unit - c.fixedunit
			numbers[offset] = c.num + numbers[offset]
		}

		return true
	})

	// refix every bit if needs carry bit
	var maxBit byte = 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 9 {
			if i == 0 {
				maxBit = numbers[i] / 10
			} else {
				numbers[i-1] = numbers[i-1] + 1
			}
			numbers[i] = numbers[i] % 10
		}
	}
	if maxBit > 0 {
		numbers = append([]byte{maxBit}, numbers...)
	}

	return xarray.Join(numbers, xstring.EMPTY_STRING), nil
}

// Itoa 是 fmt.Sprintf("%v", i) 的泛型封装。
//
// 通过泛型约束 E 同时支持整数与浮点类型，避免调用方手写类型断言。
// 当浮点数为 NaN 或 ±Inf 时，结果遵循 fmt 的默认行为。
//
// 示例：
//
//	xconv.Itoa(42)         // "42"
//	xconv.Itoa(3.14)       // "3.14"
//	xconv.Itoa(int8(-7))   // "-7"
func Itoa[E xgeneric.Integer | xgeneric.Float](i E) string {
	return fmt.Sprintf("%v", i)
}

// FormatInt 将整数 i 转换为指定进制（2 ~ 36）的字符串。
//
// 行为参考 strconv.FormatInt 的语义：
//   - 当 i 为正数且超出 int64 范围时，自动回退到 uint64 转换，保证
//     大整数也能正确格式化；
//   - 当 i 为负数时使用 FormatInt 实现；
//   - 进制大于 36 或小于 2 时会返回 strconv 的错误行为（空字符串），
//     调用方请确保 base 在 2 ~ 36 之间。
func FormatInt[E xgeneric.Integer](i E, base int) string {
	if i > 0 && uint64(i) > math.MaxInt64 {
		return strconv.FormatUint(uint64(i), base)
	}
	return strconv.FormatInt(int64(i), base)
}

// ToPtr 将任意值装箱为指针。
//
// 主要用于以下场景：
//   - 在结构体中希望字段为 nil 时表示 "未设置"；
//   - 调用只接收指针参数的第三方 API；
//   - 模拟 Java 中的 Optional.ofNullable() 之类的语义。
//
// 每次调用都会返回新的指针，互不影响。
func ToPtr[E any](t E) *E {
	return &t
}

// IsNumber 判断字符串是否符合数字字面量格式。
//
// 支持的形式包括：
//   - 十进制整数："123"、" -12"；
//   - 十六进制："0x12"、"0X1A"；
//   - 八进制："0o17"；
//   - 二进制："0b1010"；
//   - 浮点数："19.1"、"-12.11"、"12e-9"。
//
// 与 Java 的 StringUtils.isNumeric 相比，本函数更严格：
//   - 不接受纯空串；
//   - 不接受多小数点；
//   - 不接受无数字的 "0x"。
//
// 返回 true 时表示字符串可被 ParseInt 或 ParseFloat 安全解析（指数
// 形式需调用 ParseInt/ParseFloat 而不是 strconv.ParseInt）。
//
// 示例：
//
//	IsNumber("0x12")   = true
//	IsNumber("0x")     = false
//	IsNumber("0o10")   = true
//	IsNumber("0o18")   = false
//	IsNumber("0b10")   = true
//	IsNumber("-12.11") = true
//	IsNumber("12e-9")  = true
//	IsNumber("19.1")   = true
//	IsNumber("-12.1.1") = false
func IsNumber(str string) bool {
	if xstring.IsEmpty(str) {
		return false
	}

	chars := []byte(str)
	sz := len(chars)
	hasExp := false
	hasDecPoint := false
	allowSigns := false
	foundDigit := false

	// deal with any possible sign up front
	start := 0
	if chars[0] == '-' || chars[0] == '+' {
		start = 1
	}

	if sz > start+1 && chars[start] == '0' && !strings.Contains(str, ".") { // leading 0, skip if is a decimal number
		if chars[start+1] == 'x' || chars[start+1] == 'X' { // leading 0x/0X for hex number
			i := start + 2
			if i == sz {
				return false // str == "0x"
			}

			// checking hex (it can't be anything else)
			for ; i < len(chars); i++ {
				if (chars[i] < '0' || chars[i] > '9') && (chars[i] < 'a' || chars[i] > 'f') && (chars[i] < 'A' || chars[i] > 'F') {
					return false
				}
			}
			return true
		} else if chars[start+1] == 'o' || chars[start+1] == 'O' { // leading 0o/0O for octal number
			i := start + 2
			for ; i < len(chars); i++ {
				if chars[i] < '0' || chars[i] > '7' {
					return false
				}
			}
			return true
		} else if chars[start+1] == 'b' || chars[start+1] == 'B' { // leading 0o/0O for binary number
			i := start + 2
			for ; i < len(chars); i++ {
				if chars[i] < '0' || chars[i] > '1' {
					return false
				}
			}
			return true
		}
	}

	sz-- // don't want to loop to the last char, check it afterwords
	// for type qualifiers
	i := start
	// loop to the next to last char or to the last char if we need another digit to
	// make a valid number (e.g. chars[0..5] = "1234E")
	for i < sz || i < sz+1 && allowSigns && !foundDigit {
		if chars[i] >= '0' && chars[i] <= '9' {
			foundDigit = true
			allowSigns = false

		} else if chars[i] == '.' {
			if hasDecPoint || hasExp {
				// two decimal points or dec in exponent
				return false
			}
			hasDecPoint = true
		} else if chars[i] == 'e' || chars[i] == 'E' {
			// we've already taken care of hex.
			if hasExp {
				// two E's
				return false
			}
			if !foundDigit {
				return false
			}
			hasExp = true
			allowSigns = true
		} else if chars[i] == '+' || chars[i] == '-' {
			if !allowSigns {
				return false
			}
			allowSigns = false
			foundDigit = false // we need a digit after the E
		} else {
			return false
		}
		i++
	}

	if i < len(chars) {
		if chars[i] >= '0' && chars[i] <= '9' {
			// no type qualifier, OK
			return true
		}
		if chars[i] == 'e' || chars[i] == 'E' {
			// can't have an E at the last byte
			return false
		}
		if chars[i] == '.' {
			if hasDecPoint || hasExp {
				// two decimal points or dec in exponent
				return false
			}
			// single trailing decimal point after non-exponent is ok
			return foundDigit
		}
		if !allowSigns {
			return foundDigit
		}
		// last character is illegal
		return false
	}
	// allowSigns is true iff the val ends in 'E'
	// found digit it to make sure weird stuff like '.' and '1E-' doesn't pass
	return !allowSigns && foundDigit
}

// ParseInt 将字符串解析为 int。
//
// 在 strconv.ParseInt 之上额外支持：
//   - 0x/0o/0b 前缀；
//   - 指数形式（"1e3" 视作 1000，但小数部分会被截断）。
//
// 解析失败时返回 (-1, error)；字符串不合法时返回描述性错误。
//
// 注意：当结果超出 int 范围时会进行隐式截断；如需精确无截断的解析，
// 请直接使用 strconv.ParseInt/ParseUint。
func ParseInt(str string) (int, error) {
	if !IsNumber(str) && !strings.Contains(str, ".") {
		return -1, fmt.Errorf("string '%s' is not a valid int number", str)
	}

	chars := []byte(str)
	neg := 1
	start := 0
	if chars[0] == '-' {
		neg = -1
		start = 1
	} else if chars[0] == '+' {
		start = 1
	}

	if chars[start] == '0' { // leading 0, maybe base specified int eg 0x, 0b, 0o
		base := 0
		if chars[start+1] == 'x' || chars[start+1] == 'X' {
			base = 16
		} else if chars[start+1] == 'o' || chars[start+1] == 'O' {
			base = 8
		} else if chars[start+1] == 'b' || chars[start+1] == 'B' {
			base = 2
		}
		v, err := strconv.ParseInt(string(chars[start+2:]), base, 64)
		if err != nil {
			return -1, err
		}
		return int(v) * neg, nil
	}

	// if has e or E
	s := strings.ToLower(string(chars[start:]))
	if strings.Contains(s, "e") {
		splits := strings.Split(s, "e")
		v, err := strconv.ParseInt(splits[0], 10, 64)
		if err != nil {
			return -1, err
		}
		v2, err := strconv.ParseInt(splits[1], 10, 64)
		if err != nil {
			return -1, err
		}
		pow := math.Pow10(int(v2))
		return int(float64(v) * pow), nil
	}

	v, err := strconv.ParseInt(str, 10, 64)
	return int(v), err
}

// ParseFloat 将字符串解析为 float64。
//
// 在 strconv.ParseFloat 之上额外支持：
//   - 0x/0o/0b 前缀（仅整数部分）；
//   - 显式指数形式（"1e3" 等）。
//
// 解析失败时返回 (-1, error)；带进制前缀但超过 int 范围的输入会返回
// 截断后的值，需要高精度场景请改用 strconv.ParseFloat 自行处理。
func ParseFloat(str string) (float64, error) {
	if !IsNumber(str) {
		return -1, fmt.Errorf("string '%s' is not a valid int number", str)
	}

	chars := []byte(str)
	neg := 1
	start := 0
	if chars[0] == '-' {
		neg = -1
		start = 1
	} else if chars[0] == '+' {
		start = 1
	}

	if chars[start] == '0' { // leading 0, maybe base specified int eg 0x, 0b, 0o
		base := 0
		if chars[start+1] == 'x' || chars[start+1] == 'X' {
			base = 16
		} else if chars[start+1] == 'o' || chars[start+1] == 'O' {
			base = 8
		} else if chars[start+1] == 'b' || chars[start+1] == 'B' {
			base = 2
		}
		v, err := strconv.ParseInt(string(chars[start+2:]), base, 64)
		if err != nil {
			return -1, err
		}
		return float64(int(v) * neg), nil
	}

	// if has e or E
	s := strings.ToLower(string(chars[start:]))
	if strings.Contains(s, "e") {
		splits := strings.Split(s, "e")
		v, err := strconv.ParseFloat(splits[0], 64)
		if err != nil {
			return -1, err
		}
		v2, err := strconv.ParseInt(splits[1], 10, 64)
		if err != nil {
			return -1, err
		}
		pow := math.Pow10(int(v2))
		return v * pow, nil
	}

	return strconv.ParseFloat(str, 64)
}

// Append 将任意值格式化为字符串后追加到字节切片末尾并返回。
//
// 内部使用 fmt.Sprintf("%v", e)，可正确处理字符串、整数、浮点数、
// 布尔值等基础类型；对自定义结构体也会调用其 String()/Error() 方法。
func Append[E any](dst []byte, e E) []byte {
	toAppend := fmt.Sprintf("%v", e)
	return append(dst, []byte(toAppend)...)
}

// AppendString 是 Append 在字符串上的版本。
//
// 与 strings.Builder 相比更适合在表达式中链式调用：
//
//	result := xconv.AppendString("id=", id)
func AppendString[E any](str string, e E) string {
	return string(Append([]byte(str), e))
}

// ParseBool 解析字符串或整数形式的布尔值。
//
// 接受以下输入：
//   - 字符串："1"、"t"、"T"、"TRUE" 等；
//   - 整数：0 表示 false，1 表示 true，其它整数返回错误。
//
// 对非法类型（如浮点数、结构体）会直接 panic，因为泛型约束本身
// 已经过滤了大多数非法值，这里 panic 仅作为最后的兜底。
func ParseBool[E string | xgeneric.Signed](e E) (bool, error) {
	val := reflect.ValueOf(e)
	switch val.Kind() {
	case reflect.String:
		s := val.String()
		s = strings.ToLower(s)
		return strconv.ParseBool(s)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := val.Int()
		switch i {
		case 0:
			return false, nil
		case 1:
			return true, nil
		}
		return false, fmt.Errorf("%d 不是合法的布尔整数", i)
	}
	// 泛型约束保证了能落到这里的只能是反射不支持的类型。
	panic(fmt.Sprintf("ParseBool 收到非法类型: %s", val.Kind().String()))
}
