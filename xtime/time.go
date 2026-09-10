package xtime

import (
	"context"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// Time 用于表示 MySQL Unix 时间戳并进行数据库转换。
type Time int64

// Scan 从数据库驱动读取时间值。src 为 time.Time 时使用其 Unix 秒数；为字符串时按十进制解析，非法类型不修改当前值。
func (jt *Time) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case time.Time:
		*jt = Time(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		*jt = Time(i)
	}
	return
}

// Value 将当前 Unix 秒数转换为 time.Time 供数据库驱动使用，不返回错误。
func (jt Time) Value() (driver.Value, error) {
	return time.Unix(int64(jt), 0), nil
}

// Time 将当前 Unix 秒数转换为 time.Time。
func (jt Time) Time() time.Time {
	return time.Unix(int64(jt), 0)
}

// Duration 用于从 TOML 文本（如 1s、500ms）解析时间间隔。
type Duration time.Duration

// UnmarshalText 使用 time.ParseDuration 解析文本；解析失败时返回错误且不修改当前值。
func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := time.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}

// Shrink 将时长限制为不超过上下文截止时间；上下文无截止时间时按原时长创建超时上下文。
// 返回值依次为调整后的时长、带截止时间的上下文及取消函数。
func (d Duration) Shrink(c context.Context) (Duration, context.Context, context.CancelFunc) {
	if deadline, ok := c.Deadline(); ok {
		if ctimeout := time.Until(deadline); ctimeout < time.Duration(d) {
			// deliver small timeout
			return Duration(ctimeout), c, func() {}
		}
	}
	ctx, cancel := context.WithTimeout(c, time.Duration(d))
	return d, ctx, cancel
}

// CurrentEpochSecsInFloat 返回当前 Unix 纪元时间，单位为秒的 float64。
// 纪元按 1970-01-01 00:00:00 UTC 计算。
func CurrentEpochSecsInFloat() float64 {
	now := time.Now()
	ts := float64(now.UnixNano()) / float64(1000*1000*1000)
	return ts
}

// CurrentEpochSecsInInt64 返回当前 Unix 纪元时间，单位为秒的 int64。
// 纪元按 1970-01-01 00:00:00 UTC 计算。
func CurrentEpochSecsInInt64() int64 {
	return time.Now().Unix()
}

// CurrentEpochSecsInInt 返回当前 Unix 纪元时间，单位为秒的 int。
// 纪元按 1970-01-01 00:00:00 UTC 计算。
func CurrentEpochSecsInInt() int {
	return int(CurrentEpochSecsInInt64())
}

// CurrentEpochNanoSecsInInt64 返回当前 Unix 纪元时间，单位为纳秒的 int64。
// 纪元按 1970-01-01 00:00:00 UTC 计算。
func CurrentEpochNanoSecsInInt64() int64 {
	return time.Now().UnixNano()
}

// SecsToNanoSecsInInt64 将秒数转换为纳秒数。
func SecsToNanoSecsInInt64(secs int64) int64 {
	return secs * int64(1000000000)
}

// SecsFromEpochToTime 将 Unix 纪元秒数转换为 time.Time。
func SecsFromEpochToTime(ts int64) time.Time {
	return time.Unix(ts, 0)
}

// NanoSecsFromEpochToTime 将 Unix 纪元纳秒数转换为 time.Time。
func NanoSecsFromEpochToTime(ts int64) time.Time {
	return time.Unix(0, ts)
}

// ToSecsFromEpoch 将 time.Time 转换为 Unix 纪元秒数；nil 指针会导致 panic。
func ToSecsFromEpoch(t *time.Time) int64 {
	return t.Unix()
}

// ToNanoSecsFromEpoch 将 time.Time 转换为 Unix 纪元纳秒数；nil 指针会导致 panic。
func ToNanoSecsFromEpoch(t *time.Time) int64 {
	return t.UnixNano()
}

// TimestampToString 将 Unix 时间戳转换为十进制字符串。
func TimestampToString(timestamp int64) string {
	return strconv.FormatInt(timestamp, 10)
}

// StringToTimestamp 将十进制字符串解析为 int64；解析失败时返回错误。
func StringToTimestamp(timestamp string) (int64, error) {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return i, fmt.Errorf("could not convert timestamp from string to int64: %v", err)
	}
	return i, nil
}

// IsActive 判断 now 是否位于 start（含）之后且早于 stop；边界相等时仍按包含规则处理。
func IsActive(now, start, stop time.Time) bool {
	return (start.Before(now) || start.Equal(now)) && now.Before(stop)
}
