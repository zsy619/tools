package xtime

import "time"

// Equal 等于
func Equal(t1, t2 time.Time) bool {
	return t1.Equal(t2)
}

func Eq(t1, t2 time.Time) bool { return Equal(t1, t2) }

// NotEqual 不等于
func NotEqual(t1, t2 time.Time) bool {
	return !t1.Equal(t2)
}

func Ne(t1, t2 time.Time) bool { return NotEqual(t1, t2) }

// LessThan 小于
func LessThan(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

func Lt(t1, t2 time.Time) bool { return LessThan(t1, t2) }

// LessEqual 小于等于
func LessEqual(t1, t2 time.Time) bool {
	return t1.Equal(t2) || t1.Before(t2)
}

func Le(t1, t2 time.Time) bool { return LessEqual(t1, t2) }

// GreaterThan 大于
func GreaterThan(t1, t2 time.Time) bool {
	return t1.After(t2)
}

func Gt(t1, t2 time.Time) bool { return GreaterThan(t1, t2) }

// GreaterEqual 大于等于
func GreaterEqual(t1, t2 time.Time) bool {
	return t1.Equal(t2) || t1.After(t2)
}

func Ge(t1, t2 time.Time) bool { return GreaterEqual(t1, t2) }
