// Package xstring 提供了一系列常用的字符串处理工具函数。
//
// 本包以标准库 strings 为基础，针对常见业务场景补充了大量扩展工具，主要涵盖：
//   - 整型/无符号整型/浮点型与布尔类型等基本类型与字符串之间的相互转换（含默认值兜底）；
//   - 切片（int、字符串、common.Address 等）与字符串之间的相互转换（拼接、拆分、类型互转）；
//   - 字符串常用操作：截取、填充、翻转、命名风格转换（驼峰、蛇形、短横线等）；
//   - 字符串查找与判断：包含、前缀/后缀、子串定位、URL 与合约地址合法性校验等；
//   - 字符串清洗与格式化：HTML 标签去除、空格处理、不可打印字符过滤、敏感信息掩码等；
//   - 文本相似度等高级工具：汉明距离、UTF-8 合法性校验等。
//
// 所有函数均为无状态纯函数（除复用 sync.Pool 提升性能的小段缓冲外），可在并发场景下安全使用。
// 涉及以太坊相关能力时复用 go-ethereum 的 common 包；JSON 序列化复用 xjson 包以保证行为一致。
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
	// validEthAddressExp 用于校验以太坊节点/合约调用 URL 的正则：
	// 要求协议必须是 http/https/ws/wss 之一，主机部分可以是：
	//   - 形如 "a.bcdef" 的常规域名（顶级域长度 2~5）；
	//   - 或特殊主机名 localhost、ethereum、127.0.0.1。
	// 端口部分可选，端口号至少为两位数字，后面允许任意路径/参数。
	validEthAddressExp = regexp.MustCompile(`^(http|https|ws|wss):\/\/((.+?)\.(.{2,5})|localhost|ethereum|127\.0\.0\.1)(\:[0-9]{2,})*.*$`)
	// bfPool 是一个 *bytes.Buffer 的 sync.Pool，用于复用字符串拼接时的中间缓冲，
	// 减少高频调用（如 JoinInts）场景下的内存分配开销。
	bfPool = sync.Pool{
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

// IsValidEthAPIURL 判断传入的字符串是否为合法的以太坊节点/合约调用 URL。
//
// 参数：
//   - url string：待校验的 URL 字符串。
//
// 返回值：
//   - bool：合法返回 true，不合法返回 false。该函数只判断格式，不访问远程地址。
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

// ToVal 把任意类型的变量转换为字符串表示。
//
// 支持的常见类型映射如下：
//   - float32 / float64：使用 strconv.FormatFloat 转为十进制字符串，例如 float64(3.0) 会变成 "3"；
//   - int / int8 / int16 / int32 / int64：使用 strconv.Itoa 或 strconv.FormatInt 转为十进制字符串；
//   - uint / uint8 / uint16 / uint32 / uint64：使用 strconv.Itoa 或 strconv.FormatUint 转为十进制字符串；
//   - string：原样返回；
//   - []byte：使用 string(bytes) 转为字符串；
//   - 其他类型：尝试使用 xjson.Marshal 序列化为 JSON 字符串，序列化失败时得到空字节数组并返回 ""。
//
// 参数：
//   - value any：任意输入值；nil 时返回空字符串 ""。
//
// 返回值：
//   - string：转换后的字符串结果。
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

// Reverse 把字符串 s 按 rune（Unicode 码点）顺序翻转，并返回翻转结果。
//
// 函数首先校验输入是否合法 UTF-8；遇到非法 UTF-8 时直接返回原字符串与 error，
// 不会尝试做任何修复或替换。
//
// 参数：
//   - s string：待翻转的字符串。
//
// 返回值：
//   - string：翻转后的字符串；输入不是合法 UTF-8 时返回原始 s。
//   - error：输入不是合法 UTF-8 时返回 errors.New("input is not valid UTF-8")，否则返回 nil。
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

// KebabCase 把字符串 s 转换为短横线连接的小写命名（kebab-case）。
//
// 内部通过 splitIntoStrings 先将字符串拆分为若干子串，再以 "-" 连接，
// 每个子串保持原样（不做大小写转换）。空字符串调用返回 ""。
//
// 参数：
//   - s string：待转换的字符串。
//
// 返回值：
//   - string：kebab-case 形式的小写命名字符串。
func KebabCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "-")
}

// UpperKebabCase 把字符串 s 转换为短横线连接的大写命名（UPPER-KEBAB-CASE）。
//
// 内部通过 splitIntoStrings 将字符串拆分为若干子串后再以 "-" 连接，
// 子串按 upperCase=true 的策略统一转大写。空字符串调用返回 ""。
//
// 参数：
//   - s string：待转换的字符串。
//
// 返回值：
//   - string：UPPER-KEBAB-CASE 形式的大写命名字符串。
func UpperKebabCase(s string) string {
	result := splitIntoStrings(s, true)
	return strings.Join(result, "-")
}

// SnakeCase 把字符串 s 转换为下划线连接的小写命名（snake_case）。
//
// 内部通过 splitIntoStrings 先将字符串拆分为若干子串，再以 "_" 连接，
// 每个子串保持原样（不做大小写转换）。空字符串调用返回 ""。
//
// 参数：
//   - s string：待转换的字符串。
//
// 返回值：
//   - string：snake_case 形式的小写命名字符串。
func SnakeCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "_")
}

// UpperSnakeCase 把字符串 s 转换为下划线连接的大写命名（UPPER_SNAKE_CASE）。
//
// 内部通过 splitIntoStrings 将字符串拆分为若干子串后再以 "_" 连接，
// 子串按 upperCase=true 的策略统一转大写。空字符串调用返回 ""。
//
// 参数：
//   - s string：待转换的字符串。
//
// 返回值：
//   - string：UPPER_SNAKE_CASE 形式的大写命名字符串。
func UpperSnakeCase(s string) string {
	result := splitIntoStrings(s, true)
	return strings.Join(result, "_")
}

// Before 返回字符串 s 中第一次出现 char 之前的子串（不含 char 本身）。
//
// 边界行为：
//   - s 为空、char 为空、或 s 中不包含 char 时，返回原始 s；
//   - char 出现在首位置时返回 ""。
//
// 参数：
//   - s string：源字符串。
//   - char string：分隔标记字符串，按完整子串匹配。
//
// 返回值：
//   - string：char 首次出现之前的子串。
func Before(s, char string) string {
	i := strings.Index(s, char)

	if s == "" || char == "" || i == -1 {
		return s
	}

	return s[0:i]
}

// BeforeLast 返回字符串 s 中最后一次出现 char 之前的子串（不含 char 本身）。
//
// 边界行为：
//   - s 为空、char 为空、或 s 中不包含 char 时，返回原始 s；
//   - char 出现在首位置时返回 ""。
//
// 参数：
//   - s string：源字符串。
//   - char string：分隔标记字符串，按完整子串匹配。
//
// 返回值：
//   - string：char 最后一次出现之前的子串。
func BeforeLast(s, char string) string {
	i := strings.LastIndex(s, char)

	if s == "" || char == "" || i == -1 {
		return s
	}

	return s[0:i]
}

// After 返回字符串 s 中第一次出现 char 之后的子串（不含 char 本身）。
//
// 边界行为：
//   - s 为空、char 为空、或 s 中不包含 char 时，返回原始 s；
//   - char 出现在末位置时返回 ""。
//
// 参数：
//   - s string：源字符串。
//   - char string：分隔标记字符串，按完整子串匹配。
//
// 返回值：
//   - string：char 首次出现之后的子串。
func After(s, char string) string {
	i := strings.Index(s, char)

	if s == "" || char == "" || i == -1 {
		return s
	}

	return s[i+len(char):]
}

// AfterLast 返回字符串 s 中最后一次出现 char 之后的子串（不含 char 本身）。
//
// 边界行为：
//   - s 为空、char 为空、或 s 中不包含 char 时，返回原始 s；
//   - char 出现在末位置时返回 ""。
//
// 参数：
//   - s string：源字符串。
//   - char string：分隔标记字符串，按完整子串匹配。
//
// 返回值：
//   - string：char 最后一次出现之后的子串。
func AfterLast(s, char string) string {
	i := strings.LastIndex(s, char)

	if s == "" || char == "" || i == -1 {
		return s
	}

	return s[i+len(char):]
}

// IsString 判断任意类型 v 是否为字符串类型。
//
// 参数：
//   - v any：任意输入值，nil 时返回 false。
//
// 返回值：
//   - bool：v 的动态类型为 string 时返回 true，否则返回 false（包含 nil）。
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

// Unwrap 把字符串 str 开头和结尾处的 wrapToken 包装层去掉。
//
// 仅当 wrapToken 同时出现在字符串首部和尾部时才会被剥除，其它情况下原样返回。
//
// 边界行为：
//   - str 或 wrapToken 为空时直接返回 str；
//   - wrapToken 仅在尾部出现、或仅在首部出现、或不在首尾位置出现时，str 不会被修改；
//   - 内部还要求 wrapToken 的长度不超过 lastIndex，否则也不剥除（保持现状）。
//
// 参数：
//   - str string：原始字符串。
//   - wrapToken string：包装标记字符串，按完整子串匹配。
//
// 返回值：
//   - string：去掉首尾包装后的字符串；若条件不满足则返回原始 str。
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

// SplitEx 按分隔符 sep 拆分字符串 s，并可通过 removeEmptyString 控制是否跳过空串。
//
// 行为说明：
//   - sep 为空时直接返回空切片 []string{}；
//   - removeEmptyString 为 true 时，位于分隔符之间的空子串将被丢弃；
//   - removeEmptyString 为 false 时，保留拆分产生的所有元素，包括末尾可能出现的空串。
//
// 该函数与 strings.Split 类似，但提供了“是否丢弃空串”的可配置能力。
//
// 参数：
//   - s string：待拆分的字符串。
//   - sep string：分隔符，sep 为空时直接返回空切片。
//   - removeEmptyString bool：是否丢弃拆分产生的空子串。
//
// 返回值：
//   - []string：拆分后的字符串切片。
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

// Substring 在字符串 s 中按 rune（Unicode 码点）位置截取子串，并去掉 NUL 字符。
//
// offset 与 length 的语义：
//   - offset 为负值时按从字符串末尾倒数的方式处理（offset = size + offset），
//     仍小于 0 时会被夹紧到 0；
//   - offset 超过字符串长度时返回 ""；
//   - length 超出可取范围时会被夹紧到剩余长度；
//   - 最终结果会移除所有 '\x00' 字符。
//
// 参数：
//   - s string：源字符串。
//   - offset int：起始下标，可为负值。
//   - length uint：要截取的长度（按 rune 计数）。
//
// 返回值：
//   - string：截取后的子串；越界时返回 ""。
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

// SplitWords 把字符串 s 按“字母序列”切分为单词切片。
//
// 规则：连续字母（unicode.IsLetter 判定为 true）被视为一个单词；
// 单词内部的 '\” 与 '-' 被视作单词的一部分，不会触发拆分；
// 遇到其它字符时结束当前单词。最终可能返回 0 个或多个单词。
//
// 参数：
//   - s string：待拆分的字符串。
//
// 返回值：
//   - []string：拆分得到的单词切片。
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

// WordCount 统计字符串 s 中按字母序列计数的单词个数。
//
// 规则与 SplitWords 保持一致：连续字母序列计为 1 个单词；
// 单词内部的 '\” 与 '-' 不影响计数。
//
// 参数：
//   - s string：待统计的字符串。
//
// 返回值：
//   - int：单词数量，无单词时返回 0。
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

// RemoveNonPrintable 移除字符串 s 中所有不可打印的 rune（unicode.IsPrint 返回 false 的字符）。
//
// 使用 strings.Map 遍历，对每个 rune 调用 unicode.IsPrint：
//   - 可打印字符原样保留；
//   - 不可打印字符返回 -1，由 strings.Map 删除。
//
// 参数：
//   - s string：源字符串。
//
// 返回值：
//   - string：过滤掉不可打印字符后的字符串。
func RemoveNonPrintable(s string) string {
	result := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, s)

	return result
}

// BytesToString 通过 unsafe.Pointer 把 []byte 零拷贝转换为 string。
//
// 警告：该函数不会复制底层字节数组，因此返回的 string 与入参 bytes 共享同一段内存；
// 在 bytes 被修改或回收后，string 内容也会随之改变。请仅在 bytes 的生命周期
// 明确长于 string 的使用周期，或 bytes 之后不会再被修改的场景下使用。
//
// 参数：
//   - bytes []byte：待转换的字节切片。
//
// 返回值：
//   - string：与 bytes 共享底层内存的字符串。
func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// IsNotBlank 是 IsBlank 的取反便捷封装。
//
// 当字符串 s 至少包含一个非空白字符时返回 true；
// 全部由空白字符组成或为空字符串时返回 false。
//
// 参数：
//   - s string：待检测的字符串。
//
// 返回值：
//   - bool：s 非空白返回 true，否则返回 false。
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// HasPrefixAny 判断字符串 s 是否以 prefixes 切片中的任意一个前缀开头。
//
// 边界行为：
//   - s 为空、或 prefixes 为空时返回 false；
//   - 一旦命中任一前缀即返回 true，剩余前缀不再继续判断。
//
// 参数：
//   - s string：待检测的字符串。
//   - prefixes []string：候选前缀集合。
//
// 返回值：
//   - bool：命中任一前缀返回 true；否则返回 false。
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

// HasSuffixAny 判断字符串 s 是否以 suffixes 切片中的任意一个后缀结尾。
//
// 边界行为：
//   - s 为空、或 suffixes 为空时返回 false；
//   - 一旦命中任一后缀即返回 true，剩余后缀不再继续判断。
//
// 参数：
//   - s string：待检测的字符串。
//   - suffixes []string：候选后缀集合。
//
// 返回值：
//   - bool：命中任一后缀返回 true；否则返回 false。
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

// IndexOffset 在字符串 s 中从下标 idxFrom 开始向后搜索 substr，并返回其在 s 中的原始下标。
//
// 边界行为：
//   - idxFrom < 0 或 idxFrom > len(s)-1 时直接返回 -1；
//   - 找不到时返回 -1（来自 strings.Index 的返回值）；
//   - 命中时返回的是 substr 在 s 中的绝对下标，不是相对于 idxFrom 的偏移。
//
// 参数：
//   - s string：源字符串。
//   - substr string：要查找的子串。
//   - idxFrom int：搜索起始位置（绝对下标）。
//
// 返回值：
//   - int：substr 在 s 中的绝对下标；未命中或越界时返回 -1。
func IndexOffset(s string, substr string, idxFrom int) int {
	if idxFrom > len(s)-1 || idxFrom < 0 {
		return -1
	}

	return strings.Index(s[idxFrom:], substr) + idxFrom
}

// ReplaceWithMap 按 replaces 映射批量替换字符串 s 中的子串。
//
// 按 map 的迭代顺序依次对 s 执行 strings.ReplaceAll，每次替换作用于上一次的结果；
// 因此当不同键之间存在包含关系时，替换结果依赖于 Go map 的迭代顺序，结果不确定。
//
// 参数：
//   - s string：源字符串。
//   - replaces map[string]string：key 为待替换子串，value 为替换后的内容。
//
// 返回值：
//   - string：批量替换完成后的字符串。
func ReplaceWithMap(s string, replaces map[string]string) string {
	for k, v := range replaces {
		s = strings.ReplaceAll(s, k, v)
	}

	return s
}

// DefaultTrimChars 是 Trim 类函数默认要去除的字符集合。
//
// 包含：制表符 ('\t')、垂直制表符 ('\v')、换行符 ('\n')、
// 回车符 ('\r')、换页符 ('\f')、普通空格 (' ')、NUL 字节 (0x00)、
// 删除控制符 (0x85) 以及不间断空格 (0xA0)。
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
	// whitespaceRegexMatcher 用于匹配任意一种空白字符（\\s：空格、制表符、换行、回车等）。
	// 在 RemoveWhiteSpace 中用于将单个非连续空白归一化为单空格。
	whitespaceRegexMatcher = regexp.MustCompile(`\s`)
	// mutiWhitespaceRegexMatcher 用于匹配连续两个及以上空白字符，
	// 包括 POSIX 空白类（[:space:]）和 Unicode 空格分隔符（\\p{Zs}）。
	// 在 RemoveWhiteSpace 中用于将连续空白折叠为单空格。
	mutiWhitespaceRegexMatcher = regexp.MustCompile(`[[:space:]]{2,}|[\s\p{Zs}]{2,}`)
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
