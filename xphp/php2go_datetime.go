package xphp

import "time"

// Checkdate 实现 PHP checkdate()：校验是否为合法公历日期。
// 规则：1<=month<=12、day 落在当月合法范围、1<=year<=32767。
func Checkdate(month, day, year int) bool {
	// 检查月份
	if month > 12 || month < 1 {
		return false
	}
	// 检查日期
	if day < 1 {
		return false
	}
	switch month {
	case 1, 3, 5, 7, 8, 10, 12: // 31 天
		if day > 31 {
			return false
		}
	case 4, 6, 9, 11: // 30 天
		if day > 30 {
			return false
		}
	default: // 二月
		if CheckIfLeapYear(year) {
			if day > 29 {
				return false
			}
		} else {
			if day > 28 {
				return false
			}
		}
	}
	// 检查年份
	if year < 1 || year > 32767 {
		return false
	}
	return true
}

// CheckIfLeapYear 按格里高利历规则判断 year 是否为闰年。
func CheckIfLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}
	if year%4 == 0 {
		return true
	}
	return false
}

// Time 返回当前时间的 Unix 秒级时间戳（PHP time()）。
func Time() int64 {
	return time.Now().Unix()
}

// Strtotime 使用给定 format 解析 strtime，并返回其 Unix 秒级时间戳。
// 与 PHP strtotime() 不同：本函数需要显式传入 Go 的 layout。
// 示例：Strtotime("02/01/2006 15:04:05", "02/01/2016 15:04:05") == 1451747045。
func Strtotime(format, strtime string) (int64, error) {
	t, err := time.Parse(format, strtime)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// StrToTime 按 "2006-01-02 15:04:05" 解析 str 并返回 Unix 秒级时间戳。
func StrToTime(str string) (int64, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// Date 使用 Go 的 layout 字符串 format 格式化 Unix 秒级时间戳 timestamp。
// 注意：此处 format 直接传给 time.Format，需使用 Go 时间常量。
// 示例：Date("02/01/2006 15:04:05 PM", 1524799394)。
func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}

// Sleep 暂停当前 goroutine t 秒。
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// Usleep 暂停当前 goroutine t 微秒。
func Usleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}
