package xtime

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// https://github.com/pinpt/go-common/blob/master/datetime/time.go
const (
	// SignalTimeUnitNOW is now
	SignalTimeUnitNOW int32 = 0
	// SignalTimeUnitMONTH is 30 days
	SignalTimeUnitMONTH int32 = 30
	// SignalTimeUnitQUARTER is 90 days
	SignalTimeUnitQUARTER int32 = 90
	// SignalTimeUnitHALFYEAR is 180 days
	SignalTimeUnitHALFYEAR int32 = 180
	// SignalTimeUnitYEAR is 365 days
	SignalTimeUnitYEAR int32 = 365
	// SignalTimeUnitTHIRDQUARTER is 270 days
	SignalTimeUnitTHIRDQUARTER int32 = 270
	// SignalTimeUnitALLTIME is all time
	SignalTimeUnitALLTIME int32 = -1
	// SignalTimeUnitBIMONTH is two month
	SignalTimeUnitBIMONTH int32 = 60
	// DaysInMilliseconds is one day in milliseconds
	DaysInMilliseconds int64 = 86400000
)

// RFC3339 是始终保留时区偏移的 RFC3339 格式。
// 未设置时区时，格式化结果可能类似
// 2019-09-03T20:48:57.073Z。
// 使用此自定义格式可始终获得偏移量。
// 格式使用 2006-01-02T15:04:05.999999999-07:00。
const RFC3339 = "2006-01-02T15:04:05.999999999-07:00"

// GetTimeUnitString 将时间单位编号转换为可读字符串；未知编号返回空字符串。
func GetTimeUnitString(timeUnit int32) string {
	switch timeUnit {
	case SignalTimeUnitNOW:
		{
			return "now"
		}
	case SignalTimeUnitMONTH:
		{
			return "month"
		}
	case SignalTimeUnitQUARTER:
		{
			return "quarter"
		}
	case SignalTimeUnitBIMONTH:
		{
			return "bimonth"
		}
	case SignalTimeUnitHALFYEAR:
		{
			return "halfyear"
		}
	case SignalTimeUnitTHIRDQUARTER:
		{
			return "thirdquarter"
		}
	case SignalTimeUnitYEAR:
		{
			return "year"
		}
	}
	return "alltime"
}

// ISODate 返回当前时间的 RFC3339 格式字符串。
func ISODate() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// ISODateFromTime 将给定 time.Time 格式化为 RFC3339 字符串。
func ISODateFromTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

// ISODateToTime 解析 RFC3339 日期字符串；格式错误或空字符串返回错误。
func ISODateToTime(date string) (time.Time, error) {
	if strings.HasSuffix(date, "Z") {
		return time.Parse("2006-01-02T15:04:05Z", date)
	}
	return ISODateOffsetToTime(date)
}

// ISODateOffsetToTime 解析带时区偏移的 RFC3339 日期字符串；错误时返回错误。
func ISODateOffsetToTime(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, nil
	}
	if strings.Contains(date, "Z") {
		// 2017-01-20T15:56:23.000000Z-08:00
		tv, err := time.Parse("2006-01-02T15:04:05.999999999Z-07:00", date)
		if err == nil {
			return tv, nil
		}
	}
	if strings.Contains(date, ".") {
		tv, err := time.Parse("2006-01-02T15:04:05.999999999-07:00", date)
		if err != nil {
			return time.Parse("2006-01-02T15:04:05.999999999-0700", date)
		}
		return tv, nil
	}

	match, _ := regexp.MatchString("([+-]\\d{2}:\\d{2})", date)
	if match {
		return time.Parse("2006-01-02T15:04:05-07:00", date)
	}

	return time.Parse("2006-01-02T15:04:05-0700", date)
}

// ISODateToEpoch 将 ISO 日期解析为 Unix 纪元秒数；无效或空字符串返回 0 和错误。
func ISODateToEpoch(date string) (int64, error) {
	if date == "" {
		return 0, nil
	}
	ts, err := ISODateToTime(date)
	if err != nil {
		return 0, err
	}
	return TimeToEpoch(ts), nil
}

// TimeToEpoch 将 time.Time 转换为 UTC Unix 纪元毫秒数。
func TimeToEpoch(tv time.Time) int64 {
	if tv.IsZero() {
		return 0
	}
	tv = tv.UTC()
	// we want to round down to microsecond precision from nano second before we return as milliseconds
	// so we can get the microseconds in the value of epoch
	return (tv.UnixNano() + 500000) / 1000000
}

// EpochNow 返回当前时间的 UTC Unix 纪元毫秒数。
func EpochNow() int64 {
	return TimeToEpoch(time.Now())
}

// DateFromEpoch 将 Unix 纪元毫秒数转换为 time.Time。
func DateFromEpoch(t int64) time.Time {
	return time.Unix(0, t*1000000)
}

// ShortDateFromEpoch 将 Unix 纪元毫秒数格式化为短日期。
func ShortDateFromEpoch(t int64) string {
	tv := DateFromEpoch(t)
	return tv.UTC().Format("2006-01-02")
}

// ShortDateFromTime 将 time.Time 格式化为短日期。
func ShortDateFromTime(tv time.Time) string {
	return tv.UTC().Format("2006-01-02")
}

// ShortDate 从 RFC3339 字符串提取日期部分；解析失败返回空字符串。
func ShortDate(date string) string {
	if strings.Contains(date, "T") {
		t, err := time.Parse("2006-01-02T15:04:05Z", date)
		if err != nil {
			return fmt.Sprintf("<error parsing date: %s. %v>", date, err)
		}
		return t.UTC().Format("2006-01-02")
	}
	return date
}

// DateRange 返回指定时间单位的日期范围起止值；timeunit 为 -1 时返回两个 0。
// timeunit 的单位由调用方按 GetTimeUnitString 的定义解释。
func DateRange(ref time.Time, timeunit int64) (int64, int64) {
	end := EndofDay(TimeToEpoch(ref))
	begin := 1000 + (end - DaysInMilliseconds*timeunit)
	if timeunit == int64(-1) {
		begin = 0
	}
	return StartofDay(begin), end
}

// DateRangePrevious 根据参考时间和时间单位返回前一时间范围。
func DateRangePrevious(ref int64, timeunit int64) (int64, int64) {
	end := EndofDay(ref)
	begin := (end - DaysInMilliseconds*timeunit)
	priorstart, priorend := DateRange(DateFromEpoch(begin), timeunit)
	return priorstart, priorend // we go to the next day
}

// EndofDay 返回指定 Unix 秒时间当天结束时刻的 Unix 毫秒时间。
func EndofDay(tv int64) int64 {
	t := DateFromEpoch(tv).UTC()
	return TimeToEpoch(time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 9999, time.UTC))
}

// StartofDay 返回指定 Unix 秒时间当天开始时刻的 Unix 毫秒时间。
func StartofDay(tv int64) int64 {
	t := DateFromEpoch(tv).UTC()
	return TimeToEpoch(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC))
}

// ToTimeRange 以 tv 所在日期为基准加 days 天，返回起止 Unix 毫秒时间；days 可为负数。
func ToTimeRange(tv time.Time, days int) (int64, int64) {
	tv = tv.UTC()
	end := time.Date(tv.Year(), tv.Month(), tv.Day(), 23, 59, 59, 9999, time.UTC)
	startend := time.Date(tv.Year(), tv.Month(), tv.Day(), 0, 0, 0, 0, time.UTC)
	start := startend.AddDate(0, 0, days)
	return TimeToEpoch(start), TimeToEpoch(end)
}

// GetSignalDate 根据参考日期和信号时间单位返回短日期字符串。
func GetSignalDate(timeUnit int32, refDate time.Time) string {
	switch timeUnit {
	case SignalTimeUnitNOW:
		{
			return ShortDateFromTime(refDate)
		}
	case SignalTimeUnitMONTH:
		{
			return ShortDateFromTime(refDate.AddDate(0, 0, -30))
		}
	case SignalTimeUnitBIMONTH:
		{
			return ShortDateFromTime(refDate.AddDate(0, 0, -60))
		}
	case SignalTimeUnitQUARTER:
		{
			return ShortDateFromTime(refDate.AddDate(0, 0, -90))
		}
	case SignalTimeUnitHALFYEAR:
		{
			return ShortDateFromTime(refDate.AddDate(0, 0, -180))
		}
	case SignalTimeUnitTHIRDQUARTER:
		{
			return ShortDateFromTime(refDate.AddDate(0, 0, -270))
		}
	case SignalTimeUnitYEAR:
		{
			return ShortDateFromTime(refDate.AddDate(0, 0, -365))
		}
	}
	return ""
}

// GetSignalTime 根据参考日期和信号时间单位返回对齐后的 time.Time。
// 计算时会因日期边界和截断规则产生偏移。
// 例如从 2017-02-26 23:59:59.9999 取一天可能落到 2017-02-25。
// 截断后为 2017-02-25 00:00:00，因此语义上可能跨越两天。
func GetSignalTime(timeUnit int32, refDate time.Time) time.Time {
	var t time.Time
	switch timeUnit {
	case SignalTimeUnitNOW:
		{
			return refDate.UTC().Truncate(time.Hour * 24)
		}
	case SignalTimeUnitMONTH:
		{
			t = refDate.UTC().AddDate(0, 0, -30)
		}
	case SignalTimeUnitBIMONTH:
		{
			t = refDate.UTC().AddDate(0, 0, -60)
		}
	case SignalTimeUnitQUARTER:
		{
			t = refDate.UTC().AddDate(0, 0, -90)
		}
	case SignalTimeUnitHALFYEAR:
		{
			t = refDate.UTC().AddDate(0, 0, -180)
		}
	case SignalTimeUnitTHIRDQUARTER:
		{
			t = refDate.UTC().AddDate(0, 0, -270)
		}
	case SignalTimeUnitYEAR:
		{
			t = refDate.UTC().AddDate(0, 0, -365)
		}
	}

	return t.Truncate(time.Hour * 24)
}

// ToMilliSec 将 time.Time 转换为 Unix 纪元毫秒数。
func ToMilliSec(date time.Time) int64 {
	return date.UnixNano() / 1000000
}

// AddDaysToStrDate 将指定天数加到日期字符串并返回新字符串；解析失败时返回错误。
func AddDaysToStrDate(date string, days int) (string, error) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}
	return ShortDateFromTime(d.AddDate(0, 0, days)), nil
}

// Date 表示可格式化输出的日期结构。
type Date struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

// NewDateNow 返回当前时间的 Date。
func NewDateNow() Date {
	epoch := EpochNow()
	val := DateFromEpoch(epoch).Format(RFC3339)
	tv, _ := ISODateToTime(val)
	_, timezone := tv.Zone()
	return Date{
		Epoch:   epoch,
		Rfc3339: val,
		Offset:  int64(timezone) / 60,
	}
}

// NewDate 从日期字符串创建 Date；格式错误或日期无效时返回错误。
func NewDate(val string) (*Date, error) {
	tv, err := ISODateToTime(val)
	if err != nil {
		return nil, err
	}
	_, timezone := tv.Zone()
	return &Date{
		Epoch:   TimeToEpoch(tv),
		Rfc3339: tv.Round(time.Millisecond).Format(RFC3339),
		Offset:  int64(timezone) / 60,
	}, nil
}

// NewDateWithTime 从 time.Time 创建 Date。
func NewDateWithTime(tv time.Time) *Date {
	_, timezone := tv.Zone()
	return &Date{
		Epoch:   TimeToEpoch(tv),
		Rfc3339: tv.Round(time.Millisecond).Format(RFC3339),
		Offset:  int64(timezone) / 60,
	}
}

// NewDateFromEpoch 从 Unix 纪元值创建 Date。
func NewDateFromEpoch(epoch int64) Date {
	val := DateFromEpoch(epoch).Format(RFC3339)
	tv, _ := ISODateToTime(val)
	_, timezone := tv.Zone()
	return Date{
		Epoch:   epoch,
		Rfc3339: val,
		Offset:  int64(timezone) / 60,
	}
}

// TimeFromDate 将 Date 转换为 time.Time。
func TimeFromDate(date Date) time.Time {
	ts := DateFromEpoch(date.Epoch)
	if ts.IsZero() {
		return ts
	}
	// apply timezone
	loc := time.FixedZone("", int(date.Offset*60))
	return ts.In(loc)
}

// EpochMinuteApart 判断两个 Unix 纪元值是否相隔不超过一分钟。
// 绝对差值小于或等于 60000 毫秒时返回 true。
func EpochMinuteApart(epoch1, epoch2 int64) bool {
	big := epoch1
	small := epoch2
	// return tv1.Truncate(time.Minute).Equal(tv2.Truncate(time.Minute))
	// get the names right
	if small > big {
		big = epoch2
		small = epoch1
	}
	return big-small <= 1000*60
}

// ConvertToModel 按时间填充 dateModel；不匹配的结构字段保持不变。
func ConvertToModel(ts time.Time, dateModel interface{}) {
	if ts.IsZero() {
		return
	}

	date := NewDateWithTime(ts)

	t := reflect.ValueOf(dateModel).Elem()
	t.FieldByName("Rfc3339").Set(reflect.ValueOf(date.Rfc3339))
	t.FieldByName("Epoch").Set(reflect.ValueOf(date.Epoch))
	t.FieldByName("Offset").Set(reflect.ValueOf(date.Offset))
}
