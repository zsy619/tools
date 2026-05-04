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

func NewJsonLowerCase(value interface{}) JsonLowerCase {
	return JsonLowerCase{Value: value}
}

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

func NewJsonUpperCase(value interface{}) JsonUpperCase {
	return JsonUpperCase{Value: value}
}

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

func NewJsonSnakeCase(value interface{}) JsonSnakeCase {
	return JsonSnakeCase{Value: value}
}

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

func NewJsonCamelCase(value interface{}) JsonCamelCase {
	return JsonCamelCase{Value: value}
}

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

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

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
