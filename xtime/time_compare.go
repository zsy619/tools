package xtime

import "time"

// Equal 等于
func Equal(t1, t2 time.Time) bool {
	return t1.Equal(t2)
}

// Eq 判断两个时间是否相等。
func Eq(t1, t2 time.Time) bool { return Equal(t1, t2) }

// NotEqual 不等于
func NotEqual(t1, t2 time.Time) bool {
	return !t1.Equal(t2)
}

// Ne 判断两个时间是否不相等。
func Ne(t1, t2 time.Time) bool { return NotEqual(t1, t2) }

// LessThan 小于
func LessThan(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

// Lt 判断第一个时间是否早于第二个时间。
func Lt(t1, t2 time.Time) bool { return LessThan(t1, t2) }

// LessEqual 小于等于
func LessEqual(t1, t2 time.Time) bool {
	return t1.Equal(t2) || t1.Before(t2)
}

// Le 判断第一个时间是否早于或等于第二个时间。
func Le(t1, t2 time.Time) bool { return LessEqual(t1, t2) }

// GreaterThan 大于
func GreaterThan(t1, t2 time.Time) bool {
	return t1.After(t2)
}

// Gt 判断第一个时间是否晚于第二个时间。
func Gt(t1, t2 time.Time) bool { return GreaterThan(t1, t2) }

// GreaterEqual 大于等于
func GreaterEqual(t1, t2 time.Time) bool {
	return t1.Equal(t2) || t1.After(t2)
}

// Ge 判断第一个时间是否晚于或等于第二个时间。
func Ge(t1, t2 time.Time) bool { return GreaterEqual(t1, t2) }
