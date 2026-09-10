package xphp

import (
	"errors"
	"time"
)

// pattern 存储 PHP 日期格式与 Go layout 之间的映射关系。
//
// 通过 regexp 匹配用户输入的日期字符串，再用 layout 调用 time.Parse。
type pattern struct {
	regexp string // 用于匹配的 Go 正则表达式
	layout string // Go 的时间 layout
	format string // PHP 的日期格式
}

// patterns 是 pattern 的集合类型。
type patterns []pattern

// _defaultPatterns 内置的 PHP/Go 日期格式映射表，在 init 中初始化。
var _defaultPatterns patterns

// formatMap 是 PHP 单字符日期格式 -> Go layout 的映射表。
var formatMap map[string]string

func init() {
	_defaultPatterns = patterns{
		pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{2}$",
			layout: "2006-01-02",
			format: "Y-m-d",
		},
		pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{1}$",
			layout: "2006-01-2",
			format: "Y-m-j",
		},
		pattern{
			regexp: "^\\d{4}-\\d{1}-\\d{2}$",
			layout: "2006-1-02",
			format: "Y-n-d",
		},
		pattern{
			regexp: "^\\d{4}-\\d{1}-\\d{1,2}$",
			layout: "2006-1-2",
			format: "Y-n-j",
		},
		pattern{
			regexp: "^\\d{2}-\\d{2}-\\d{2}$",
			layout: "06-01-02",
			format: "y-m-d",
		},
		pattern{
			regexp: "^\\d{2}-\\d{2}-\\d{1}$",
			layout: "06-01-2",
			format: "y-m-j",
		},
		pattern{
			regexp: "^\\d{2}-\\d{1}-\\d{2}$",
			layout: "06-1-02",
			format: "y-n-d",
		},
		pattern{
			regexp: "^\\d{2}-\\d{1}-\\d{1,2}$",
			layout: "06-1-2",
			format: "y-n-j",
		},
		pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}$",
			layout: "2006-01-02 15:04:05",
			format: "Y-m-d H:i:s",
		},
		pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{1} \\d{2}:\\d{2}:\\d{2}$",
			layout: "2006-01-2 15:04:05",
			format: "Y-m-j H:i:s",
		},
		pattern{
			regexp: "^\\d{4}-\\d{1}-\\d{2} \\d{2}:\\d{2}:\\d{2}$",
			layout: "2006-1-02 15:04:05",
			format: "Y-n-d H:i:s",
		},
		pattern{
			regexp: "^\\d{4}-\\d{1}-\\d{1} \\d{2}:\\d{2}:\\d{2}$",
			layout: "2006-1-2 15:04:05",
			format: "Y-n-j H:i:s",
		},
		pattern{
			regexp: "^\\d{2}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}$",
			layout: "06-01-02 15:04:05",
			format: "y-m-d H:i:s",
		},
		pattern{
			regexp: "^\\d{2}-\\d{2}-\\d{1} \\d{2}:\\d{2}:\\d{2}$",
			layout: "06-01-2 15:04:05",
			format: "y-m-j H:i:s",
		},
		pattern{
			regexp: "^\\d{2}-\\d{1}-\\d{2} \\d{2}:\\d{2}:\\d{2}$",
			layout: "06-1-02 15:04:05",
			format: "y-n-d H:i:s",
		},
		pattern{
			regexp: "^\\d{2}-\\d{1}-\\d{1} \\d{2}:\\d{2}:\\d{2}$",
			layout: "06-1-2 15:04:05",
			format: "y-n-j H:i:s",
		},
		pattern{
			regexp: "^\\d{8}$",
			layout: "20060102",
			format: "Ymd",
		},
		pattern{
			regexp: "^\\d{14}$",
			layout: "20060102150405",
			format: "YmdHis",
		},
		pattern{
			regexp: "^\\d{4}$",
			layout: "2006",
			format: "Y",
		},
		pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}[+-]\\d{2}:\\d{2}$",
			layout: "2006-01-02T15:04:05-07:00",
			format: "c", // ISO 8601 date
		},
		pattern{
			regexp: "^[[:alpha:]]{3}, \\d{2} [[:alpha:]]{3} \\d{4} \\d{2}:\\d{2}:\\d{2} [+-]\\d{4}$",
			layout: time.RFC1123Z,
			format: "r", // RFC 2822 formatted date
		},
	}

	formatMap = map[string]string{
		"d": "02",
		"j": "2",
		"m": "01",
		"n": "1",
		"Y": "2006",
		"y": "06",
		"H": "15",
		"h": "03",
		"g": "3",
		"i": "04",
		"s": "05",
		"M": "Jan",
		"F": "January",
		"D": "Mon",
		"l": "Monday",
		"e": "MST",
		"T": "MST",
		"O": "-0700",
		"P": "-07:00",
	}
}

// getPattern 按 PHP 风格日期格式字符串查找对应的 pattern。
//
// 找不到匹配项时返回 errors.New("no pattern found")；调用方应自行
// 决定是回退到 convertLayout 还是直接报错。
func getPattern(format string) (pattern, error) {
	for _, p := range _defaultPatterns {
		if p.format == format {
			return p, nil
		}
	}

	return pattern{}, errors.New("no pattern found")
}

// convertLayout 把任意 PHP 日期格式字符串转换为对应的 Go layout。
//
// 与 getPattern 不同，本函数不需要输入在 _defaultPatterns 中已注册；
// 它按字符逐位查 formatMap，未识别的字符（如分隔符 / 字面量）原样保留。
// 因此可用于任意自定义 PHP 格式字符串的转换，但无法保证一定能被
// time.Parse 成功解析。
func convertLayout(format string) string {
	var layout string
	for _, s := range format {
		if v, ok := formatMap[string(s)]; ok {
			layout += v
		} else {
			layout += string(s)
		}
	}
	return layout
}
