package xtime

import (
	"fmt"
	"time"
)

// DateTime 日期时间类型
type DateTime time.Time

// AddDays 向当前日期增加days天
func (d DateTime) AddDays(days int) DateTime {
	return DateTime(time.Time(d).AddDate(0, 0, days))
}

// SubDays 从当前日期减去days天
func (d DateTime) SubDays(days int) DateTime {
	return DateTime(time.Time(d).AddDate(0, 0, -days))
}

// Compare 比较两个日期先后
func (d DateTime) Compare(other DateTime) int {
	return time.Time(d).Compare(time.Time(other))
}

// String 转换日期时间类型到string类型
func (d DateTime) String() string {
	return fmt.Sprintf("%v", time.Time(d))
}
