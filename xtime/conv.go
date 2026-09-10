package xtime

import "time"

// DateTime2UTS 将 YYYY-MM-DD hh:mm:ss 格式的日期时间字符串解析为 Unix 时间戳（秒）。
// 使用本地时区解析；解析失败时返回 0。空字符串同样会导致解析失败并返回 0。
func DateTime2UTS(dt string) int64 {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(DateTimeFormat, dt, loc)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// UTS2DateTime 将 Unix 时间戳（秒）格式化为 YYYY-MM-DD hh:mm:ss 形式的字符串。
// 使用本地时区格式化；uts 取负值亦可正确处理。
func UTS2DateTime(uts int64) string {
	return time.Unix(uts, 0).Format(DateTimeFormat)
}
