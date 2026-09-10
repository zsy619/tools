package xjson

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// 字段小写json
type JsonLowerCase struct {
	Value interface{}
}

// NewJsonLowerCase 创建一个 JsonLowerCase 包装器，用于将 JSON 字段键名转换为小写。
func NewJsonLowerCase(value interface{}) JsonLowerCase {
	return JsonLowerCase{Value: value}
}

// MarshalJSON 实现 json.Marshaler：序列化时将所有 JSON 字段键名转换为小写。
func (c JsonLowerCase) MarshalJSON() ([]byte, error) {
	// Regexp definitions
	keyMatchRegex := regexp.MustCompile(`\"(\w+)\":`)
	marshalled, err := Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			matchStr := string(match)
			key := matchStr[1 : len(matchStr)-2]
			resKey := strings.ToLower(key)
			return []byte(`"` + resKey + `":`)
		},
	)
	return converted, err
}

// 字段大写json
type JsonUpperCase struct {
	Value interface{}
}

// NewJsonUpperCase 创建一个 JsonUpperCase 包装器，用于将 JSON 字段键名转换为大写。
func NewJsonUpperCase(value interface{}) JsonUpperCase {
	return JsonUpperCase{Value: value}
}

// MarshalJSON 实现 json.Marshaler：序列化时将所有 JSON 字段键名转换为大写。
func (c JsonUpperCase) MarshalJSON() ([]byte, error) {
	// Regexp definitions
	keyMatchRegex := regexp.MustCompile(`\"(\w+)\":`)
	marshalled, err := Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			matchStr := string(match)
			key := matchStr[1 : len(matchStr)-2]
			resKey := strings.ToUpper(key)
			return []byte(`"` + resKey + `":`)
		},
	)
	return converted, err
}

// 下划线json
type JsonSnakeCase struct {
	Value interface{}
}

// NewJsonSnakeCase 创建一个 JsonSnakeCase 包装器，用于将 JSON 字段键名转换为下划线风格。
func NewJsonSnakeCase(value interface{}) JsonSnakeCase {
	return JsonSnakeCase{Value: value}
}

// MarshalJSON 实现 json.Marshaler：序列化时将所有 JSON 字段键名转换为下划线风格。
func (c JsonSnakeCase) MarshalJSON() ([]byte, error) {
	// Regexp definitions
	keyMatchRegex := regexp.MustCompile(`\"(\w+)\":`)
	wordBarrierRegex := regexp.MustCompile(`(\w)([A-Z])`)
	marshalled, err := Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)
	return converted, err
}

// 驼峰json
type JsonCamelCase struct {
	Value interface{}
}

// NewJsonCamelCase 创建一个 JsonCamelCase 包装器，用于将 JSON 字段键名转换为驼峰风格。
func NewJsonCamelCase(value interface{}) JsonCamelCase {
	return JsonCamelCase{Value: value}
}

// MarshalJSON 实现 json.Marshaler：序列化时将所有 JSON 字段键名转换为驼峰风格。
func (c JsonCamelCase) MarshalJSON() ([]byte, error) {
	keyMatchRegex := regexp.MustCompile(`\"(\w+)\":`)
	marshalled, err := Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			matchStr := string(match)
			key := matchStr[1 : len(matchStr)-2]
			resKey := Lcfirst(Case2Camel(key))
			return []byte(`"` + resKey + `":`)
		},
	)
	return converted, err
}

// 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	caser := cases.Title(language.English)
	name = caser.String(name)
	return strings.Replace(name, " ", "", -1)
}

// 首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// 首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

// NewBuffer 创建一个新的字符串缓冲区。
func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

// Append 将值追加写入缓冲区（支持 int、int64、uint、uint64、string、[]byte、rune 类型），返回自身便于链式调用。
func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}
