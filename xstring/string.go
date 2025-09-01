// Package strings contains various common utils for strings
package xstring

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"

	"github.com/zsy619/tools/xjson"
)

var (
	validEthAddressExp = regexp.MustCompile(`^(http|https|ws|wss):\/\/((.+?)\.(.{2,5})|localhost|ethereum|127\.0\.0\.1)(\:[0-9]{2,})*.*$`)
	bfPool             = sync.Pool{
		New: func() any {
			return bytes.NewBuffer([]byte{})
		},
	}
)

// JoinInts 将一个int64类型的切片转换成以逗号分隔的字符串
//
// 参数:
//
//	is: 一个int64类型的切片
//
// 返回值:
//
//	返回一个以逗号分隔的字符串
func JoinInts(is []int64) string {
	if len(is) == 0 {
		return ""
	}
	if len(is) == 1 {
		return strconv.FormatInt(is[0], 10)
	}
	buf := bfPool.Get().(*bytes.Buffer)
	for _, i := range is {
		buf.WriteString(strconv.FormatInt(i, 10))
		buf.WriteByte(',')
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	s := buf.String()
	buf.Reset()
	bfPool.Put(buf)
	return s
}

// SplitInts 函数将字符串s按照逗号分隔，并将每个分隔后的子串转换成int64类型的整数切片返回
//
// 如果字符串s为空，则返回nil和一个nil错误
//
// 参数s为待处理的字符串
//
// 返回值res为转换后的整数切片，error为可能出现的错误
func SplitInts(s string) ([]int64, error) {
	if s == "" {
		return nil, nil
	}
	sArr := strings.Split(s, ",")
	res := make([]int64, 0, len(sArr))
	for _, sc := range sArr {
		i, err := strconv.ParseInt(sc, 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}

// StrOrEmptyStr 函数接收一个字符串指针s作为参数，返回字符串类型的值。
//
// 如果s不为nil，则返回s指向的字符串；否则返回空字符串。
func StrOrEmptyStr(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// StrToPtr 函数将传入的字符串转换为指向该字符串的指针
//
// 参数s: 需要转换为指针的字符串
//
// 返回值: 指向传入的字符串的指针
func StrToPtr(s string) *string {
	return &s
}

// IsValidEthAPIURL returns true if the given string matches a valid
// eth endpoint URL
func IsValidEthAPIURL(url string) bool {
	return validEthAddressExp.MatchString(url)
}

// IsValidContractAddress 判断传入的字符串是否为有效的合约地址
//
// 参数：
// address string - 待验证的字符串地址
//
// 返回值：
// bool - 若字符串为有效的合约地址则返回true，否则返回false
func IsValidContractAddress(address string) bool {
	return common.IsHexAddress(address)
}

// RandomHexStr 生成指定长度的随机十六进制字符串
//
// 参数：
//
//	n int - 生成字符串的长度
//
// 返回值：
//
//	string - 随机十六进制字符串
//	error - 如果生成过程中出错，则返回错误信息，否则为nil
func RandomHexStr(n int) (string, error) {
	bys := make([]byte, n)
	if _, err := rand.Read(bys); err != nil {
		return "", err
	}
	return hex.EncodeToString(bys), nil
}

// RandomFileName 生成指定长度的随机文件名
//
// 参数：
//
//	n int - 文件名长度
//
// 返回值：
//
//	string - 随机文件名，如果生成失败则返回空字符串
func RandomFileName(n int) string {
	fn, err := RandomHexStr(n)
	if err != nil {
		return ""
	}
	return fn
}

// ListCommonAddressToListString 将一组 common.Address 类型的地址转换成字符串类型的切片
//
// 参数：
//   - addresses: 待转换的 common.Address 类型地址切片
//
// 返回值：
//   - []string: 转换后的字符串类型地址切片
func ListCommonAddressToListString(addresses []common.Address) []string {
	addressesString := make([]string, len(addresses))
	for i, address := range addresses {
		addressesString[i] = address.Hex()
	}
	return addressesString
}

// ListStringToListCommonAddress 将字符串类型的地址列表转换为common.Address类型的地址列表
//
// 参数：
//
//	addresses []string - 字符串类型的地址列表
//
// 返回值：
//
//	[]common.Address - common.Address类型的地址列表
func ListStringToListCommonAddress(addresses []string) []common.Address {
	addressesCommon := make([]common.Address, len(addresses))
	for i, address := range addresses {
		addressesCommon[i] = common.HexToAddress(address)
	}
	return addressesCommon
}

// ListCommonAddressesToString 将 common.Address 类型的切片转换成字符串，并返回
//
// 参数：
//
//	addresses: 要转换的 common.Address 类型的切片
//
// 返回值：
//
//	转换后的字符串，以逗号分隔
func ListCommonAddressesToString(addresses []common.Address) string {
	addressesString := ListCommonAddressToListString(addresses)
	return strings.Join(addressesString, ",")
}

// ListIntToListString 将整型切片转换为字符串切片
//
// 参数：
//
//	listInt []int 待转换的整型切片
//
// 返回值：
//
//	[]string 转换后的字符串切片
func ListIntToListString(listInt []int) []string {
	listString := make([]string, len(listInt))
	for idx, i := range listInt {
		listString[idx] = strconv.Itoa(i)
	}
	return listString
}

// StringToCommonAddressesList 将逗号分隔的字符串转换为通用地址列表
//
// 参数：
//
//	addresses string - 逗号分隔的字符串
//
// 返回值：
//
//	[]common.Address - 通用地址列表
func StringToCommonAddressesList(addresses string) []common.Address {
	addressesString := strings.Split(addresses, ",")
	return ListStringToListCommonAddress(addressesString)
}

// SubString 返回字符串s中从start到end的子字符串
//
// 如果start小于0或end大于字符串长度或start大于end，则返回空字符串
//
// 如果start等于0且end等于字符串长度，则返回原字符串s
//
// 参数：
//   - s：待截取子字符串的源字符串
//   - start：子字符串的起始位置（包含）
//   - end：子字符串的结束位置（不包含）
//
// 返回值：
//   - 截取到的子字符串
func SubString(s string, start int, end int) string {
	r := []rune(s)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return s
	}

	return string(r[start:end])
}

// ToInt 字符串转换为int
func ToInt(s string, def ...int) int {
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			if len(def) > 0 {
				return def[0]
			}
			return 0
		}
		return i
	}
	return 0
}

// ToInt64 字符串转换为int数组
func ToIntArray(s, sep string) ([]int, bool) {
	if len(s) > 0 {
		var intArr []int
		strArr := strings.Split(s, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToInt(str))
		}
		return intArr, true
	}
	return nil, false
}

// ToInt8 字符串转换为int8
func ToInt8(s string, def ...int8) int8 {
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			if len(def) > 0 {
				return def[0]
			}
			return 0
		}
		return int8(i)
	}
	return 0
}

// ToInt8Array 字符串转换为int8数组
func ToInt8Array(s, sep string) ([]int8, bool) {
	if len(s) > 0 {
		var intArr []int8
		strArr := strings.Split(s, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToInt8(str))
		}
		return intArr, true
	}
	return nil, false
}

// ToInt16 字符串转换为int16
func ToInt16(s string, def ...int16) int16 {
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			if len(def) > 0 {
				return def[0]
			}
			return 0
		}
		return int16(i)
	}
	return 0
}

// ToInt16Array 字符串转换为int16数组
func ToInt16Array(s, sep string) ([]int16, bool) {
	if len(s) > 0 {
		var intArr []int16
		strArr := strings.Split(s, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToInt16(str))
		}
		return intArr, true
	}
	return nil, false
}

// ToInt32 字符串转换为int32
func ToInt32(s string, def ...int32) int32 {
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			if len(def) > 0 {
				return def[0]
			}
			return 0
		}
		return int32(i)
	}
	return 0
}

// ToInt32Array 字符串转换为int32数组
func ToInt32Array(s string, sep string) ([]int32, bool) {
	if len(s) > 0 {
		var intArr []int32
		strArr := strings.Split(s, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToInt32(str))
		}
		return intArr, true
	}
	return nil, false
}

// ToInt64 字符串转换为int64
func ToInt64(s string, def ...int64) int64 {
	if len(s) > 0 {
		value, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			if len(def) > 0 {
				return def[0]
			}
			return 0
		}
		return value
	}
	return 0
}

// ToInt64Array 字符串转换为int64数组
func ToInt64Array(s, sep string) ([]int64, bool) {
	if len(s) > 0 {
		var intArr []int64
		strArr := strings.Split(s, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToInt64(str))
		}
		return intArr, true
	}
	return nil, false
}

// ToUint 字符串转换为uint
func ToUnit(s string, def ...uint) uint {
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			if len(def) > 0 {
				return def[0]
			}
			return 0
		}
		return uint(i)
	}
	return 0
}

// ToUintArray 字符串转换为uint数组
func ToUintArray(s, sep string) ([]uint, bool) {
	if len(s) > 0 {
		var intArr []uint
		strArr := strings.Split(s, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToUnit(str))
		}
		return intArr, true
	}
	return nil, false
}

// ToUint8 字符串转换为uint8
func ToUint8(s string) uint8 {
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return uint8(i)
	}
	return 0
}

// ToUint8Array 字符串转换为uint8数组
func ToUint8Array(s, sep string) ([]uint8, bool) {
	if len(s) > 0 {
		var intArr []uint8
		strArr := strings.Split(s, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToUint8(str))
		}
		return intArr, true
	}
	return nil, false
}

// ToUint16 字符串转换为uint16
func ToUint16(s string) uint16 {
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return uint16(i)
	}
	return 0
}

// ToUint16Array 字符串转换为uint16数组
func ToUint16Array(s, sep string) ([]uint16, bool) {
	if len(s) > 0 {
		var intArr []uint16
		strArr := strings.Split(s, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToUint16(str))
		}
		return intArr, true
	}
	return nil, false
}

// ToUint32 字符串转换为uint32
func ToUint32(s string) uint32 {
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return uint32(i)
	}
	return 0
}

// ToUint32Array 字符串转换为uint32数组
func ToUint32Array(s, sep string) ([]uint32, bool) {
	if len(s) > 0 {
		var intArr []uint32
		strArr := strings.Split(s, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToUint32(str))
		}
		return intArr, true
	}
	return nil, false
}

// ToUint64 将字符串转换为uint64类型
//
// 参数：
//
//	s 待转换的字符串
//
// 返回值：
// 转换后的uint64类型数值，如果转换失败则返回0
func ToUint64(s string) uint64 {
	if len(s) > 0 {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return uint64(i)
	}
	return 0
}

// ToUint64Array 将字符串按指定分隔符拆分成多个字符串，并将每个字符串转换成uint64类型存入切片中返回。
//
// 如果输入字符串为空，则返回nil和false。
//
//	s：待转换的字符串
//
//	sep：字符串中的分隔符
//
// 返回值1：转换后的uint64类型切片
//
// 返回值2：转换是否成功，成功为true，失败为false
func ToUint64Array(s, sep string) ([]uint64, bool) {
	if len(s) > 0 {
		var intArr []uint64
		strArr := strings.Split(s, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToUint64(str))
		}
		return intArr, true
	}
	return nil, false
}

// Without 函数从字符串切片中移除指定的字符串元素
//
// arr: 待处理的字符串切片
//
// remove: 需要移除的字符串元素
//
// 返回移除指定元素后的新字符串切片
func Without(arr []string, remove string) (chopped []string) {
	for _, s := range arr {
		if s != remove {
			chopped = append(chopped, s)
		}
	}
	return
}

// GenerateSlug 将输入字符串 s 转换为 slug 格式，并返回转换后的 slug 字符串
//
// 参数：
// s：需要转换的字符串
//
// 返回值：
// slug：转换后的 slug 字符串
func GenerateSlug(s string) (slug string) {
	return strings.Map(func(r rune) rune {
		switch {
		case r == ' ', r == '-':
			return '-'
		case r == '_', unicode.IsLetter(r), unicode.IsDigit(r):
			return r
		default:
			return -1
		}
	}, strings.ToLower(strings.TrimSpace(s)))
}

// InChain 判断字符串 s 是否在字符串切片 haystack 中
// 参数：
//
//	s string - 待查找的字符串
//	haystack []string - 字符串切片
//
// 返回值：
//
//	bool - 如果 s 在 haystack 中，则返回 true；否则返回 false
func InChain(s string, haystack []string) bool {
	if haystack == nil {
		return false
	}
	for _, straw := range haystack {
		if s == straw {
			return true
		}
	}
	return false
}

// TrimHtml 去除字符串中的HTML标签，并返回处理后的字符串。
// 参数：
//
//	s string - 待处理的字符串
//
// 返回值：
//
//	string - 处理后的字符串
func TrimHtml(s string) string {
	// 将HTML标签全转换成小写
	re, _ := regexp.Compile(`\\<[\\S\\s]+?\\>`)
	s = re.ReplaceAllStringFunc(s, strings.ToLower)
	// 去除STYLE
	re, _ = regexp.Compile(`\\<style[\\S\\s]+?\\</style\\>`)
	s = re.ReplaceAllString(s, "")
	// 去除SCRIPT
	re, _ = regexp.Compile(`\\<script[\\S\\s]+?\\</script\\>`)
	s = re.ReplaceAllString(s, "")
	// 去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile(`\\<[\\S\\s]+?\\>`)
	s = re.ReplaceAllString(s, "\n")
	// 去除连续的换行符
	re, _ = regexp.Compile(`\\s{2,}`)
	s = re.ReplaceAllString(s, "\n")
	return strings.TrimSpace(s)
}

// CutstrHtml 去除字符串中的HTML标签，换行符，制表符，空格和&nbsp;并返回处理后的字符串
// 参数：
//   - s: 需要处理的字符串
//
// 返回值：
//   - string: 处理后的字符串
func CutstrHtml(s string) string {
	strx := TrimHtml(s)
	strx = strings.ReplaceAll(strx, "\r\n", "")
	strx = strings.ReplaceAll(strx, "\r", "")
	strx = strings.ReplaceAll(strx, " ", "")
	strx = strings.ReplaceAll(strx, "\t", "")
	strx = strings.ReplaceAll(strx, "&nbsp;", "")
	strx = strings.ReplaceAll(strx, "\n", "")
	return strx
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func ToVal(value any) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value := value.(type) {
	case float64:
		ft := value
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value
		key = strconv.Itoa(it)
	case uint:
		it := value
		key = strconv.Itoa(int(it))
	case int8:
		it := value
		key = strconv.Itoa(int(it))
	case uint8:
		it := value
		key = strconv.Itoa(int(it))
	case int16:
		it := value
		key = strconv.Itoa(int(it))
	case uint16:
		it := value
		key = strconv.Itoa(int(it))
	case int32:
		it := value
		key = strconv.Itoa(int(it))
	case uint32:
		it := value
		key = strconv.Itoa(int(it))
	case int64:
		it := value
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value
		key = strconv.FormatUint(it, 10)
	case string:
		key = value
	case []byte:
		key = string(value)
	default:
		newValue, _ := xjson.Marshal(value)
		key = string(newValue)
	}

	return key
}

// Reverse 函数接收一个字符串参数s，返回翻转后的字符串和可能发生的错误
func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

// CamelCase 将给定的字符串s转换为驼峰命名法（CamelCase）并返回结果字符串
//
// 如果字符串s为空，则返回空字符串
//
// 参数：
//
//	s string - 待转换的字符串
//
// 返回值：
//
//	string - 转换后的驼峰命名法字符串
func CamelCase(s string) string {
	var builder strings.Builder

	strs := splitIntoStrings(s, false)
	for i, str := range strs {
		if i == 0 {
			builder.WriteString(strings.ToLower(str))
		} else {
			builder.WriteString(Capitalize(str))
		}
	}

	return builder.String()
}

// UpperFirst 将字符串首字母转换为大写，并返回转换后的字符串
//
// 如果传入的字符串为空，则返回空字符串
//
// 参数：
//
//	s string - 待转换的字符串
//
// 返回值：
//
//	string - 转换后的字符串
func UpperFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToUpper(r)

	return string(r) + s[size:]
}

// LowerFirst 将字符串的第一个字符转换为小写并返回新的字符串
//
// 如果字符串为空，则返回空字符串
//
// 参数：
//
//	s：待处理的字符串
//
// 返回值：
//
//	转换后的字符串
func LowerFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToLower(r)

	return string(r) + s[size:]
}

// Pad 函数用于在字符串s的右侧填充指定数量的padStr字符串，使总长度达到size个字符。
//
// 如果s的长度已经大于或等于size，则直接返回s。
//
// 参数：
//
//	s：待填充的字符串
//	size：目标字符串长度
//	padStr：用于填充的字符串
//
// 返回值：
//
//	填充后的字符串
func Pad(s string, size int, padStr string) string {
	return padAtPosition(s, size, padStr, 0)
}

// PadStart 在字符串s的开头填充指定字符串padStr，直到长度达到size。
//
// 如果s的长度已经大于或等于size，则直接返回s。
//
// 参数：
//   - s：需要填充的字符串
//   - size：填充后的字符串长度
//   - padStr：用于填充的字符串
//
// 返回值：
//   - 填充后的字符串
func PadStart(s string, size int, padStr string) string {
	return padAtPosition(s, size, padStr, 1)
}

// PadEnd 函数用于在字符串末尾填充指定字符，使字符串长度达到指定大小
//
// 参数：
//   - source：待填充的源字符串
//   - size：目标字符串长度
//   - padStr：用于填充的字符串
//
// 返回值：
//   - 返回填充后的字符串
func PadEnd(source string, size int, padStr string) string {
	return padAtPosition(source, size, padStr, 2)
}

func KebabCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "-")
}

func UpperKebabCase(s string) string {
	result := splitIntoStrings(s, true)
	return strings.Join(result, "-")
}

func SnakeCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "_")
}

func UpperSnakeCase(s string) string {
	result := splitIntoStrings(s, true)
	return strings.Join(result, "_")
}

func Before(s, char string) string {
	i := strings.Index(s, char)

	if s == "" || char == "" || i == -1 {
		return s
	}

	return s[0:i]
}

func BeforeLast(s, char string) string {
	i := strings.LastIndex(s, char)

	if s == "" || char == "" || i == -1 {
		return s
	}

	return s[0:i]
}

func After(s, char string) string {
	i := strings.Index(s, char)

	if s == "" || char == "" || i == -1 {
		return s
	}

	return s[i+len(char):]
}

func AfterLast(s, char string) string {
	i := strings.LastIndex(s, char)

	if s == "" || char == "" || i == -1 {
		return s
	}

	return s[i+len(char):]
}

func IsString(v any) bool {
	if v == nil {
		return false
	}
	switch v.(type) {
	case string:
		return true
	default:
		return false
	}
}

func Unwrap(str string, wrapToken string) string {
	if str == "" || wrapToken == "" {
		return str
	}

	firstIndex := strings.Index(str, wrapToken)
	lastIndex := strings.LastIndex(str, wrapToken)

	if firstIndex == 0 && lastIndex > 0 && lastIndex <= len(str)-1 {
		if len(wrapToken) <= lastIndex {
			str = str[len(wrapToken):lastIndex]
		}
	}

	return str
}

func SplitEx(s, sep string, removeEmptyString bool) []string {
	if sep == "" {
		return []string{}
	}

	n := strings.Count(s, sep) + 1
	a := make([]string, n)
	n--
	i := 0
	sepSave := 0
	ignore := false

	for i < n {
		m := strings.Index(s, sep)
		if m < 0 {
			break
		}
		ignore = false
		if removeEmptyString {
			if s[:m+sepSave] == "" {
				ignore = true
			}
		}
		if !ignore {
			a[i] = s[:m+sepSave]
			s = s[m+len(sep):]
			i++
		} else {
			s = s[m+len(sep):]
		}
	}

	var ret []string
	if removeEmptyString {
		if s != "" {
			a[i] = s
			ret = a[:i+1]
		} else {
			ret = a[:i]
		}
	} else {
		a[i] = s
		ret = a[:i+1]
	}

	return ret
}

func Substring(s string, offset int, length uint) string {
	rs := []rune(s)
	size := len(rs)

	if offset < 0 {
		offset = size + offset
		if offset < 0 {
			offset = 0
		}
	}
	if offset > size {
		return ""
	}

	if length > uint(size)-uint(offset) {
		length = uint(size - offset)
	}

	str := string(rs[offset : offset+int(length)])

	return strings.Replace(str, "\x00", "", -1)
}

func SplitWords(s string) []string {
	var word string
	var words []string
	var r rune
	var size, pos int

	isWord := false

	for len(s) > 0 {
		r, size = utf8.DecodeRuneInString(s)

		switch {
		case isLetter(r):
			if !isWord {
				isWord = true
				word = s
				pos = 0
			}

		case isWord && (r == '\'' || r == '-'):
			// is word

		default:
			if isWord {
				isWord = false
				words = append(words, word[:pos])
			}
		}

		pos += size
		s = s[size:]
	}

	if isWord {
		words = append(words, word[:pos])
	}

	return words
}

func WordCount(s string) int {
	var r rune
	var size, count int

	isWord := false

	for len(s) > 0 {
		r, size = utf8.DecodeRuneInString(s)

		switch {
		case isLetter(r):
			if !isWord {
				isWord = true
				count++
			}

		case isWord && (r == '\'' || r == '-'):
			// is word

		default:
			isWord = false
		}

		s = s[size:]
	}

	return count
}

func RemoveNonPrintable(s string) string {
	result := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, s)

	return result
}

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func HasPrefixAny(s string, prefixes []string) bool {
	if len(s) == 0 || len(prefixes) == 0 {
		return false
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

func HasSuffixAny(s string, suffixes []string) bool {
	if len(s) == 0 || len(suffixes) == 0 {
		return false
	}
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

func IndexOffset(s string, substr string, idxFrom int) int {
	if idxFrom > len(s)-1 || idxFrom < 0 {
		return -1
	}

	return strings.Index(s[idxFrom:], substr) + idxFrom
}

func ReplaceWithMap(s string, replaces map[string]string) string {
	for k, v := range replaces {
		s = strings.ReplaceAll(s, k, v)
	}

	return s
}

// DefaultTrimChars are the characters which are stripped by Trim* functions in default.
var DefaultTrimChars = string([]byte{
	'\t', // Tab.
	'\v', // Vertical tab.
	'\n', // New line (line feed).
	'\r', // Carriage return.
	'\f', // New page.
	' ',  // Ordinary space.
	0x00, // NUL-byte.
	0x85, // Delete.
	0xA0, // Non-breaking space.
})

// SplitAndTrim 函数将输入的字符串s按照指定的分隔符delimiter进行拆分，并对拆分后的每个子串去除指定字符集characterMask中的字符。
//
// 返回拆分并去除字符后的字符串切片。
//
// 如果不传入characterMask，则默认去除字符串两端的空格。
//
// 参数：
//   - s：待拆分的字符串
//   - delimiter：分隔符
//   - characterMask：可选参数，要去除的字符集，可传入多个字符集
//
// 返回值：
//   - []string：拆分并去除字符后的字符串切片
func SplitAndTrim(s, delimiter string, characterMask ...string) []string {
	result := make([]string, 0)

	for _, v := range strings.Split(s, delimiter) {
		v = Trim(v, characterMask...)
		if v != "" {
			result = append(result, v)
		}
	}

	return result
}

// Trim 函数接收一个字符串s和可选的字符掩码characterMask作为参数，并返回一个新的字符串。
//
// 该函数会删除字符串s开头和结尾处与trimChars中任一字符匹配的所有字符。
//
// 如果characterMask参数不为空，则将其首个元素追加到trimChars中。
//
// 如果没有提供characterMask参数，则使用默认的字符掩码DefaultTrimChars。
//
// 参数：
//   - s: 需要处理的字符串
//   - characterMask: 可选的字符掩码，用于指定需要删除的字符集
//
// 返回值：
//   - 返回处理后的字符串
func Trim(s string, characterMask ...string) string {
	trimChars := DefaultTrimChars

	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}

	return strings.Trim(s, trimChars)
}

// HideString 函数将给定字符串中的指定部分替换为指定的字符。
//
// 参数：
//   - s：待处理的字符串
//   - start：需要替换部分的起始位置
//   - end：需要替换部分的结束位置
//   - replaceChar：用于替换的字符
//
// 返回值：
//   - 替换后的字符串
func HideString(s string, start, end int, replaceChar string) string {
	size := len(s)

	if replaceChar == "" {
		return s
	}

	if start > size-1 || start < 0 || end < 0 || start > end {
		return s
	}

	if end > size {
		end = size
	}

	startStr := s[0:start]
	endStr := s[end:size]

	replaceSize := end - start
	replaceStr := strings.Repeat(replaceChar, replaceSize)

	return startStr + replaceStr + endStr
}

// ContainsAll 判断字符串s中是否包含所有子字符串substrs中的元素
//
// 参数：
//
//	s: 待判断的字符串
//	substrs: 子字符串数组
//
// 返回值：
//
//	如果s中包含substrs中所有元素，则返回true；否则返回false
func ContainsAll(s string, substrs []string) bool {
	for _, v := range substrs {
		if !strings.Contains(s, v) {
			return false
		}
	}
	return true
}

// ContainsAny 判断字符串str中是否包含substrs中的任意一个子串
//
// 参数：
//   - str string - 待判断的字符串
//   - substrs []string - 子串列表
//
// 返回值：
//   - bool - 如果str中包含substrs中的任意一个子串，则返回true；否则返回false
func ContainsAny(str string, substrs []string) bool {
	for _, v := range substrs {
		if strings.Contains(str, v) {
			return true
		}
	}

	return false
}

var (
	whitespaceRegexMatcher     *regexp.Regexp = regexp.MustCompile(`\s`)
	mutiWhitespaceRegexMatcher *regexp.Regexp = regexp.MustCompile(`[[:space:]]{2,}|[\s\p{Zs}]{2,}`)
)

// RemoveWhiteSpace 函数用于去除字符串中的空格
//
// 参数：
//   - s: 待处理的字符串
//   - repalceAll: 是否去除所有空格，若为真，则去除所有空格，否则仅去除连续的空格
//
// 返回值：
//   - 处理后的字符串
func RemoveWhiteSpace(s string, repalceAll bool) string {
	if repalceAll && s != "" {
		return strings.Join(strings.Fields(s), "")
	} else if s != "" {
		s = mutiWhitespaceRegexMatcher.ReplaceAllString(s, " ")
		s = whitespaceRegexMatcher.ReplaceAllString(s, " ")
	}

	return strings.TrimSpace(s)
}

// SubInBetween 函数返回在字符串 s 中位于 start 和 end 之间（不包含 start 和 end）的子串
//
// 如果字符串 s 中没有包含 start，或者 start 后没有包含 end，则返回空字符串 ""
//
// 参数：
//   - s: 待处理的字符串
//   - start: 子串开始位置的标记字符串
//   - end: 子串结束位置的标记字符串
//
// 返回值：
//
//   - 返回位于 start 和 end 之间（不包含 start 和 end）的子串，如果未找到则返回空字符串 ""
func SubInBetween(s string, start string, end string) string {
	if _, after, ok := strings.Cut(s, start); ok {
		if before, _, ok := strings.Cut(after, end); ok {
			return before
		}
	}

	return ""
}

// HammingDistance 计算两个等长字符串a和b之间的汉明距离（即对应位置上不同字符的个数）
// 如果a和b的长度不相等，则返回-1和错误信息
// 参数：
//
//   - a string - 字符串a
//   - b string - 字符串b
//
// 返回值：
//
//   - int - a和b之间的汉明距离
//   - error - 如果a和b的长度不相等，则返回错误信息；否则返回nil
func HammingDistance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("a length and b length are unequal")
	}

	ar := []rune(a)
	br := []rune(b)

	var distance int
	for i, codepoint := range ar {
		if codepoint != br[i] {
			distance++
		}
	}

	return distance, nil
}

// Concat 函数将传入的字符串参数拼接起来并返回拼接后的结果
//
// 参数：
//   - length int - 拼接后字符串的最大长度，如果为0或负数则不进行长度限制
//   - str ...string - 要拼接的字符串参数，可变参数
//
// 返回值：
//   - string - 拼接后的字符串结果
func Concat(length int, str ...string) string {
	if len(str) == 0 {
		return ""
	}

	sb := strings.Builder{}
	if length <= 0 {
		sb.Grow(len(str[0]) * len(str))
	} else {
		sb.Grow(length)
	}

	for _, s := range str {
		sb.WriteString(s)
	}
	return sb.String()
}
