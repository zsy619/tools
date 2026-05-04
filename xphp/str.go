package xphp

import (
	"errors"
	"fmt"
	"html"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// PadTypes for StrPad
const (
	StrPadRight = iota
	StrPadLeft
	StrPadBoth
)

const (
	PHP_EOL = "\n"
)

// RuneMatchFunc is function to check if a rune match some condition
type RuneMatchFunc func(rune) bool

// buildReplaceSlice is a helper function for Replace and Ireplace
func buildReplaceSlice(search, replace interface{}) ([]string, []string, error) {
	var aSearch, aReplace []string

	switch s := search.(type) {
	case int:
		aSearch = append(aSearch, strconv.Itoa(s))
	case rune:
		aSearch = append(aSearch, string(s))
	case string:
		aSearch = append(aSearch, s)
	case []string:
		aSearch = s
	default:
		return aSearch, aReplace, errors.New("unsupported type of search")
	}

	switch r := replace.(type) {
	case int:
		aReplace = append(aReplace, strconv.Itoa(r))
	case rune:
		aReplace = append(aReplace, string(r))
	case string:
		aReplace = append(aReplace, r)
	case []string:
		aReplace = r
	default:
		return aSearch, aReplace, errors.New("unsupported type of replace")
	}

	return aSearch, aReplace, nil
}

// Replace all occurrences of the search string with the replacement string
//
// This function is an implement of PHP's str_replace
//
// see http://php.net/manual/en/function.str-replace.php
func Replace(search, replace interface{}, subject string) string {
	var aSearch, aReplace []string

	aSearch, aReplace, err := buildReplaceSlice(search, replace)
	if err != nil {
		return subject
	}

	r := ""
	for index, s := range aSearch {
		if index < len(aReplace) {
			r = aReplace[index]
		}
		subject = strings.Replace(subject, s, r, -1)
	}

	return subject
}

// Ireplace is case-insensitive version of Replace()
func Ireplace(search, replace interface{}, subject string) string {
	var aSearch, aReplace []string

	aSearch, aReplace, err := buildReplaceSlice(search, replace)
	if err != nil {
		return subject
	}

	r := ""
	var reg *regexp.Regexp
	for index, s := range aSearch {
		if index < len(aReplace) {
			r = aReplace[index]
		}
		reg = regexp.MustCompile("(?i:" + s + ")")
		subject = reg.ReplaceAllString(subject, r)
	}

	return subject
}

// Stristr is case-insensitive Strstr()
func Stristr(haystack, needle string) string {
	haystackx := strings.ToLower(haystack)
	needlex := strings.ToLower(needle)
	pos := Stripos(haystackx, needlex, 0)
	if pos < 0 {
		return ""
	}
	return Substr(haystack, uint(pos), 0)
}

// HTMLSpecialchars converts special characters to HTML entities
func HTMLSpecialchars(str string) string {
	return html.EscapeString(str)
}

// HTMLSpecialcharsDecode converts special HTML entities back to characters
func HTMLSpecialcharsDecode(str string) string {
	return html.UnescapeString(str)
}

// DefaultNumberFormat is default NumberFormat for english notation with thousands separator
func DefaultNumberFormat(number float64, decimals int) string {
	return NumberFormat(number, uint(decimals), ".", ",")
}

// StrPad pads a string to a certain length with another string
func StrPad(input string, padLength int, padString string, padType int) string {
	// if the value of padLength is less than or equal to the length of the input string,
	// no padding takes place, and input will be returned.
	if padLength <= len(input) {
		return input
	}
	// default padType is StrPadRight
	if padType > StrPadBoth {
		padType = StrPadRight
	}
	// default padString is space
	if padString == "" {
		padString = " "
	}

	var s string
	l := padLength - len(input)
	switch padType {
	case StrPadRight:
		s = input
		i := 0
		for i < l {
			for j := 0; j < len(padString); j++ {
				s += padString[j : j+1]
				i++
				if i >= l {
					break
				}
			}
		}
	case StrPadLeft:
		i := 0
		for i < l {
			for j := 0; j < len(padString); j++ {
				s += padString[j : j+1]
				i++
				if i >= l {
					break
				}
			}
		}
		s += input
	case StrPadBoth:
		l1 := l / 2
		s = StrPad(StrPad(input, len(input)+l1, padString, StrPadLeft), padLength, padString, StrPadRight)
	}

	return s
}

func Num2String(num int) string {
	if num >= 1000 {
		rt := fmt.Sprintf("%0.2f", math.Round(float64(num)/1000*100)/100)
		rt = strings.TrimRight(rt, "0")
		rt = strings.TrimRight(rt, ".")
		return rt + "K"
	}
	return strconv.Itoa(num)
}
