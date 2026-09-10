package xphp

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Microtime 返回当前时间对应的 Unix 时间戳（带微秒精度）。
func Microtime() float64 {
	return Round(float64(time.Now().UnixNano()) / 1000000000)
}

// IsLeapYear 判断 t 是否处于闰年（基于 UTC 年末的 YearDay 判断）。
func IsLeapYear(t time.Time) bool {
	t2 := time.Date(t.Year(), time.December, 31, 0, 0, 0, 0, time.UTC)
	return t2.YearDay() == 366
}

// LastDateOfMonth 返回 t 所在月份的最后一天的 00:00:00 UTC 时间。
func LastDateOfMonth(t time.Time) time.Time {
	t2 := FirstDateOfNextMonth(t)
	return time.Unix(t2.Unix()-86400, 0)
}

// FirstDateOfMonth 返回 t 所在月份第一天的 00:00:00 UTC 时间。
func FirstDateOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

// FirstDateOfNextMonth 返回 t 所在月份的下一个月的第一天（00:00:00 UTC）。
func FirstDateOfNextMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	if month == time.December {
		year++
		month = time.January
	} else {
		month++
	}
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

// FirstDateOfLastMonth 返回 t 所在月份上一个月的第一天（00:00:00 UTC）。
func FirstDateOfLastMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	if month == time.January {
		year--
		month = time.December
	} else {
		month--
	}
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

// recognize 根据 PHP 日期格式字符 c 返回 t 中对应字段的字符串表示。
// 未识别的字符原样返回 c 自身。
func recognize(c string, t time.Time) string {
	switch c {
	// 日
	case "d":
		return fmt.Sprintf("%02d", t.Day())
	case "D":
		return t.Format("Mon")
	case "j":
		return fmt.Sprintf("%d", t.Day())
	case "l":
		return t.Weekday().String()
	case "w":
		return fmt.Sprintf("%d", t.Weekday())
	case "z":
		return fmt.Sprintf("%v", t.YearDay()-1)

	// 周
	case "W":
		_, w := t.ISOWeek()
		return fmt.Sprintf("%d", w)

	// 月
	case "F":
		return t.Month().String()
	case "m":
		return fmt.Sprintf("%02d", t.Month())
	case "M":
		return t.Format("Jan")
	case "n":
		return fmt.Sprintf("%d", t.Month())
	case "t":
		return LastDateOfMonth(t).Format("2")

	// 年
	case "L":
		if IsLeapYear(t) {
			return "1"
		}
		return "0"
	case "o":
		fallthrough
	case "Y":
		return fmt.Sprintf("%v", t.Year())
	case "y":
		return t.Format("06")

	// 时间
	case "a":
		return t.Format("pm")
	case "A":
		return strings.ToUpper(t.Format("pm"))
	case "g":
		return t.Format("3")
	case "G":
		return fmt.Sprintf("%d", t.Hour())
	case "h":
		return t.Format("03")
	case "H":
		return fmt.Sprintf("%02d", t.Hour())
	case "i":
		return fmt.Sprintf("%02d", t.Minute())
	case "s":
		return fmt.Sprintf("%02d", t.Second())
	case "e":
		fallthrough
	case "T":
		return t.Format("MST")
	case "O":
		return t.Format("-0700")
	case "P":
		return t.Format("-07:00")
	case "U":
		return fmt.Sprintf("%v", t.Unix())

	default:
		return c
	}
}

// parse 将 PHP 风格的 format 字符串逐字符展开为 t 的文本表示。
func parse(format string, t time.Time) string {
	result := ""
	for _, s := range format {
		result += recognize(string(s), t)
	}
	return result
}

// format 优先按内置 pattern 的 layout 格式化，匹配失败则按 PHP 风格逐字符解析。
func format(f string, t time.Time) string {
	pattern, err := getPattern(f)
	if err != nil {
		return parse(f, t)
	}

	return t.Format(pattern.layout)
}

// Today 使用当前时间按 PHP 风格 format 字符串输出；时区使用本地时区。
func Today(f string) string {
	return format(f, time.Now())
}

// LocalDate 将 Unix 秒级时间戳 timestamp 按本地时区与 PHP 风格 format 字符串输出。
func LocalDate(f string, timestamp int64) string {
	return format(f, time.Unix(timestamp, 0))
}

// DateCreateFromFormat 按 PHP 风格 format 解析字符串 t 为 time.Time。
// 不识别的字符会原样保留在 layout 中，无法解析时返回 time.Parse 的错误。
func DateCreateFromFormat(f string, t string) (time.Time, error) {
	return time.Parse(convertLayout(f), t)
}

// DateCreate 将字符串 str 解析为 time.Time，支持以下输入：
//   - "now"：当前时间；
//   - 形如 "+N day" / "-2 hours" 的相对时间字符串；
//   - _defaultPatterns 中预定义的日期/时间格式。
//
// 解析失败时返回 time.Time{} 与 "Unsupported date/time string: ..." 错误。
func DateCreate(str string) (time.Time, error) {
	if strings.ToLower(str) == "now" {
		return time.Now(), nil
	}

	duration, err := DateIntervalCreateFromDateString(str)
	if err == nil {
		return time.Now().Add(duration), nil
	}

	for _, p := range _defaultPatterns {
		reg := regexp.MustCompile(p.regexp)
		if reg.MatchString(str) {
			t, err := time.Parse(p.layout, str)
			if err == nil {
				return t, nil
			}
		}
	}

	return time.Time{}, errors.New("Unsupported date/time string: " + str)
}

// DateDateSet 根据 year/month/day 设置一个 time.Time（按 2006-1-2 layout 解析）。
// month/day 不会补零，传入 0 可能导致解析失败。
func DateDateSet(year, month, day int) (time.Time, error) {
	return time.Parse("2006-1-2", fmt.Sprintf("%04d-%d-%d", year, month, day))
}

// DateDefaultTimezoneGet 获取当前本地时区的名称（如 CST）。
func DateDefaultTimezoneGet() string {
	tz, _ := time.Now().Local().Zone()
	return tz
}

// DateDefaultTimezoneSet 设置进程的默认时区为 tz（影响 time.Local）。
// tz 非法时返回错误且不修改 time.Local。
func DateDefaultTimezoneSet(tz string) error {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return err
	}
	time.Local = loc
	return nil
}

// DateTimezoneGet 返回 t 所在时区的缩写名。
func DateTimezoneGet(t time.Time) string {
	tz, _ := t.Zone()
	return tz
}

// DateTimezoneSet 返回 t 在 tz 时区下的副本；tz 非法时返回原始 t 与错误。
func DateTimezoneSet(t time.Time, tz string) (time.Time, error) {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return t, err
	}
	return t.In(loc), nil
}

// DateDiff 返回 t2 - t1 的时间间隔。
func DateDiff(t1 time.Time, t2 time.Time) time.Duration {
	return t2.Sub(t1)
}

// DateFormat 按 PHP 风格 format 字符串输出 t 的字符串表示。
func DateFormat(t time.Time, f string) string {
	return format(f, t)
}

// DateIntervalCreateFromDateString 将形如 "+1 day -2 hours" 的字符串解析为 time.Duration。
// 支持的单位：day/month/year/week/hour/minute/second，可省略数字（默认 1）。
// 无法匹配时返回 0 与 "unsupported string format" 错误。
func DateIntervalCreateFromDateString(str string) (time.Duration, error) {
	reg := regexp.MustCompile(`((\+|\-)?\s*(\d*)\s*(day|month|year|week|hour|minute|second)s?\s*)+?`)
	matches := reg.FindAllStringSubmatch(str, -1)
	if matches != nil {
		var duration int64
		for _, match := range matches {
			var diff, num int64
			if match[3] == "" {
				num = 1
			} else {
				num, _ = strconv.ParseInt(match[3], 10, 64)
			}
			switch match[4] {
			case "day":
				diff = num * 86400
			case "month":
				diff = num * 86400 * 30
			case "year":
				diff = num * 86400 * 365
			case "week":
				diff = num * 86400 * 7
			case "hour":
				diff = num * 3600
			case "minute":
				diff = num * 60
			case "second":
				diff = num
			}
			if match[2] == "-" {
				diff = -diff
			}
			duration += diff
		}
		return time.Duration(duration) * time.Second, nil
	}
	return 0, errors.New("unsupported string format")
}

// DateISODateSet 按 ISO 8601 周历返回 year 年第 week 周第 day 天的 time.Time。
// 输入超出合法范围时可能返回意料之外的时间，但不会 panic。
func DateISODateSet(year, week, day int) (time.Time, error) {
	firstDateOfYear, err := time.Parse("2006-1-2", fmt.Sprintf("%04d-%d-%d", year, 1, 1))
	if err != nil {
		return time.Time{}, err
	}

	offset := time.Duration(1-firstDateOfYear.Weekday()) * 24 * time.Hour
	firstDateOfFirstWeek := firstDateOfYear.Add(offset)

	return firstDateOfFirstWeek.Add(time.Duration(((week-1)*7+day-1)*24) * time.Hour), nil
}

// DateModify 按 PHP 风格相对时间字符串（如 "+1 day"）修改 t 并返回新值。
// 解析 modify 失败时返回原始 t 与错误。
func DateModify(t time.Time, modify string) (time.Time, error) {
	duration, err := DateIntervalCreateFromDateString(modify)
	if err != nil {
		return t, err
	}
	return t.Add(duration), nil
}

// DateOffsetGet 返回 t 所在时区相对 UTC 的偏移秒数（东时区为正）。
func DateOffsetGet(t time.Time) int {
	_, offset := t.Zone()
	return offset
}

// DateAdd 在 t 上加上时长 d，返回新的 time.Time（t 不被修改）。
func DateAdd(t time.Time, d time.Duration) time.Time {
	return t.Add(d)
}

// DateSub 在 t 上减去时长 d，返回新的 time.Time（t 不被修改）。
func DateSub(t time.Time, d time.Duration) time.Time {
	return t.Add(-d)
}

// DateTimestampGet 返回 t 对应的 Unix 秒级时间戳。
func DateTimestampGet(t time.Time) int64 {
	return t.Unix()
}

// DateTimestampSet 根据 Unix 秒级时间戳构造对应的 UTC time.Time。
func DateTimestampSet(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
