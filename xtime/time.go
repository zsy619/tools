package xtime

import (
	"context"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// Time be used to MySql timestamp converting.
type Time int64

// Scan scan time.
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

// Value get time value.
func (jt Time) Value() (driver.Value, error) {
	return time.Unix(int64(jt), 0), nil
}

// Time get time.
func (jt Time) Time() time.Time {
	return time.Unix(int64(jt), 0)
}

// Duration be used toml unmarshal string time, like 1s, 500ms.
type Duration time.Duration

// UnmarshalText unmarshal text to duration.
func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := time.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}

// Shrink will decrease the duration by comparing with context's timeout duration
// and return new timeout\context\CancelFunc.
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

// CurrentEpochSecsInFloat returns the current time as a timestamp
// from epoch as type float64 in seconds.
func CurrentEpochSecsInFloat() float64 {
	now := time.Now()
	ts := float64(now.UnixNano()) / float64(1000*1000*1000)
	return ts
}

// CurrentEpochSecsInInt64 returns the current time as a timestamp
// from epoch as type int64 in seconds.
func CurrentEpochSecsInInt64() int64 {
	return time.Now().Unix()
}

// CurrentEpochSecsInInt returns the current time as a timestamp
// from epoch as type int in seconds.
func CurrentEpochSecsInInt() int {
	return int(CurrentEpochSecsInInt64())
}

// CurrentEpochNanoSecsInInt64 returns the current time as a timestamp
// from epoch as type int64 in nanoseconds.
func CurrentEpochNanoSecsInInt64() int64 {
	return time.Now().UnixNano()
}

// SecsToNanoSecsInInt64 converts a value from secs to nanoseconds.
func SecsToNanoSecsInInt64(secs int64) int64 {
	return secs * int64(1000000000)
}

// SecsFromEpochToTime converts an int64 of seconds from epoch to Time struct
func SecsFromEpochToTime(ts int64) time.Time {
	return time.Unix(ts, 0)
}

// NanoSecsFromEpochToTime converts an int64 of nanoseconds from epoch to Time struct
func NanoSecsFromEpochToTime(ts int64) time.Time {
	return time.Unix(0, ts)
}

// ToSecsFromEpoch converts a time.Time struct to nanoseconds from epoch.
func ToSecsFromEpoch(t *time.Time) int64 {
	return t.Unix()
}

// ToNanoSecsFromEpoch converts a time.Time struct to nanoseconds from epoch.
func ToNanoSecsFromEpoch(t *time.Time) int64 {
	return t.UnixNano()
}

// TimestampToString converts an int64 timestamp to string
func TimestampToString(timestamp int64) string {
	return strconv.FormatInt(timestamp, 10)
}

// StringToTimestamp converts a string timestamp to int64
func StringToTimestamp(timestamp string) (int64, error) {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return i, fmt.Errorf("could not convert timestamp from string to int64: %v", err)
	}
	return i, nil
}

func IsActive(now, start, stop time.Time) bool {
	return (start.Before(now) || start.Equal(now)) && now.Before(stop)
}
