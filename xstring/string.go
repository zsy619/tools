// Package strings contains various common utils for strings
package xstring

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/ethereum/go-ethereum/common"

	"haedu.gov.cn/tools/xjson"
)

var (
	validEthAddressExp = regexp.MustCompile(`^(http|https|ws|wss):\/\/((.+?)\.(.{2,5})|localhost|ethereum|127\.0\.0\.1)(\:[0-9]{2,})*.*$`)
	bfPool             = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

// JoinInts format int64 slice like:n1,n2,n3.
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

// SplitInts split string into int64 slice.
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

// StrOrEmptyStr returns an empty string if a string pointer is nil.
// Otherwise returns the string value of the pointer.
func StrOrEmptyStr(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// StrToPtr is a convenience func that converts a string to it's pointer
func StrToPtr(s string) *string {
	return &s
}

// IsValidEthAPIURL returns true if the given string matches a valid
// eth endpoint URL
func IsValidEthAPIURL(url string) bool {
	return validEthAddressExp.MatchString(url)
}

// IsValidContractAddress returns true is the given string matches a valid
// smart contract address
func IsValidContractAddress(address string) bool {
	return common.IsHexAddress(address)
}

// RandomHexStr generates a hex string from a byte slice of n random numbers.
func RandomHexStr(n int) (string, error) {
	bys := make([]byte, n)
	if _, err := rand.Read(bys); err != nil {
		return "", err
	}
	return hex.EncodeToString(bys), nil
}

// ListCommonAddressToListString converts a list of common.address to list of string
func ListCommonAddressToListString(addresses []common.Address) []string {
	addressesString := make([]string, len(addresses))
	for i, address := range addresses {
		addressesString[i] = address.Hex()
	}
	return addressesString
}

// ListStringToListCommonAddress converts a list of strings to list of common.address
func ListStringToListCommonAddress(addresses []string) []common.Address {
	addressesCommon := make([]common.Address, len(addresses))
	for i, address := range addresses {
		addressesCommon[i] = common.HexToAddress(address)
	}
	return addressesCommon
}

// ListCommonAddressesToString converts a list of common.address to a comma delimited string
func ListCommonAddressesToString(addresses []common.Address) string {
	addressesString := ListCommonAddressToListString(addresses)
	return strings.Join(addressesString, ",")
}

// ListIntToListString converts a list of big.int to a list of string
func ListIntToListString(listInt []int) []string {
	listString := make([]string, len(listInt))
	for idx, i := range listInt {
		listString[idx] = strconv.Itoa(i)
	}
	return listString
}

// StringToCommonAddressesList converts a comma delimited string to a list of common.address
func StringToCommonAddressesList(addresses string) []common.Address {
	addressesString := strings.Split(addresses, ",")
	return ListStringToListCommonAddress(addressesString)
}

// SubString 截取字符串
// 获取source的子串,如果start小于0或者end大于source长度则返回""
// start:开始index，从0开始，包括0
// end:结束index，以end结束，但不包括end
func SubString(source string, start int, end int) string {
	r := []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

// ToInt 字符串转换为int
func ToInt(input string) int {
	if len(input) > 0 {
		i, err := strconv.Atoi(input)
		if err != nil {
			return 0
		}
		return i
	}
	return 0
}

// ToInt64 字符串转换为int数组
func ToIntArray(input, sep string) []int {
	if len(input) > 0 {
		var intArr []int
		strArr := strings.Split(input, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToInt(str))
		}
		return intArr
	}
	return nil
}

// ToInt8 字符串转换为int8
func ToInt8(input string) int8 {
	if len(input) > 0 {
		i, err := strconv.Atoi(input)
		if err != nil {
			return 0
		}
		return int8(i)
	}
	return 0
}

// ToInt8Array 字符串转换为int8数组
func ToInt8Array(input, sep string) []int8 {
	if len(input) > 0 {
		var intArr []int8
		strArr := strings.Split(input, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToInt8(str))
		}
		return intArr
	}
	return nil
}

// ToInt16 字符串转换为int16
func ToInt16(input string) int16 {
	if len(input) > 0 {
		i, err := strconv.Atoi(input)
		if err != nil {
			return 0
		}
		return int16(i)
	}
	return 0
}

func ToInt16Array(input, sep string) []int16 {
	if len(input) > 0 {
		var intArr []int16
		strArr := strings.Split(input, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToInt16(str))
		}
		return intArr
	}
	return nil
}

// ToInt32 字符串转换为int32
func ToInt32(input string) int32 {
	if len(input) > 0 {
		i, err := strconv.Atoi(input)
		if err != nil {
			return 0
		}
		return int32(i)
	}
	return 0
}

// ToInt32Array 字符串转换为int32数组
func ToInt32Array(input string, sep string) []int32 {
	intArray := []int32{}
	// Split string into array
	arr := strings.Split(input, sep)

	for _, n := range arr {
		// Convert each string element to an int
		intVal, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		// Add the value to the final int array
		intArray = append(intArray, int32(intVal))
	}

	// intArray now contains [1,2,3]
	return intArray
}

// ToInt64 字符串转换为int64
func ToInt64(input string) int64 {
	if len(input) > 0 {
		i, err := strconv.Atoi(input)
		if err != nil {
			return 0
		}
		return int64(i)
	}
	return 0
}

// ToInt64Array 字符串转换为int64数组
func ToInt64Array(input, sep string) []int64 {
	if len(input) > 0 {
		var intArr []int64
		strArr := strings.Split(input, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToInt64(str))
		}
		return intArr
	}
	return nil
}

// ToUint 字符串转换为uint
func ToUnit(input string) uint {
	if len(input) > 0 {
		i, err := strconv.Atoi(input)
		if err != nil {
			return 0
		}
		return uint(i)
	}
	return 0
}

// ToUintArray 字符串转换为uint数组
func ToUintArray(input, sep string) []uint {
	if len(input) > 0 {
		var intArr []uint
		strArr := strings.Split(input, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToUint(str))
		}
		return intArr
	}
	return nil
}

// ToUint8 字符串转换为uint8
func ToUint8(input string) uint8 {
	if len(input) > 0 {
		i, err := strconv.Atoi(input)
		if err != nil {
			return 0
		}
		return uint8(i)
	}
	return 0
}

// ToUint8Array 字符串转换为uint8数组
func ToUint8Array(input, sep string) []uint8 {
	if len(input) > 0 {
		var intArr []uint8
		strArr := strings.Split(input, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToUint8(str))
		}
		return intArr
	}
	return nil
}

// ToUint16 字符串转换为uint16
func ToUint16(input string) uint16 {
	if len(input) > 0 {
		i, err := strconv.Atoi(input)
		if err != nil {
			return 0
		}
		return uint16(i)
	}
	return 0
}

func ToUint16Array(input, sep string) []uint16 {
	if len(input) > 0 {
		var intArr []uint16
		strArr := strings.Split(input, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToUint16(str))
		}
		return intArr
	}
	return nil
}

// ToUint32 字符串转换为uint32
func ToUint32(input string) uint32 {
	if len(input) > 0 {
		i, err := strconv.Atoi(input)
		if err != nil {
			return 0
		}
		return uint32(i)
	}
	return 0
}

// ToUint32Array 字符串转换为uint32数组
func ToUint32Array(input, sep string) []uint32 {
	if len(input) > 0 {
		var intArr []uint32
		strArr := strings.Split(input, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToUint32(str))
		}
		return intArr
	}
	return nil
}

// ToUint64 字符串转换为uint64
func ToUint64(input string) uint64 {
	if len(input) > 0 {
		i, err := strconv.Atoi(input)
		if err != nil {
			return 0
		}
		return uint64(i)
	}
	return 0
}

// ToUint64Array 字符串转换为uint64数组
func ToUint64Array(input, sep string) []uint64 {
	if len(input) > 0 {
		var intArr []uint64
		strArr := strings.Split(input, sep)
		for _, str := range strArr {
			intArr = append(intArr, ToUint64(str))
		}
		return intArr
	}
	return nil
}

// Without returns a copy of the slice without the remove parameter
//
// [TODO] Extend to work with all standard types
func Without(arr []string, remove string) (chopped []string) {
	for _, s := range arr {
		if s != remove {
			chopped = append(chopped, s)
		}
	}
	return
}

// GenerateSlug converts a string into a lowercase dasherized slug
//
// For example: GenerateSlug("My cool object") returns "my-cool-object"
func GenerateSlug(str string) (slug string) {
	return strings.Map(func(r rune) rune {
		switch {
		case r == ' ', r == '-':
			return '-'
		case r == '_', unicode.IsLetter(r), unicode.IsDigit(r):
			return r
		default:
			return -1
		}
		return -1
	}, strings.ToLower(strings.TrimSpace(str)))
}

// InChain returns a boolean if a string is already in a slice of strings
//
// [TODO] Extend this to work for all standard types
func InChain(needle string, haystack []string) bool {
	if haystack == nil {
		return false
	}
	for _, straw := range haystack {
		if needle == straw {
			return true
		}
	}
	return false
}

func TrimHtml(src string) string {
	// 将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	// 去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	// 去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	// 去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	// 去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func CutstrHtml(src string) string {
	strx := TrimHtml(src)
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
func StrVal(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := xjson.Marshal(value)
		key = string(newValue)
	}

	return key
}

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
