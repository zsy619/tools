package xtime

import "time"

// TimeCounter 用于记录从设定时刻起经过的时间间隔。
type TimeCounter struct {
	time.Time
	int64
}

// NewTimeCounter 创建并初始化时间计数器。
func NewTimeCounter() (t *TimeCounter) {
	t = new(TimeCounter)
	t.Set()
	return t
}

// Set 将开始时间设为当前时刻。
func (t *TimeCounter) Set() {
	t.Time = time.Now()
	t.int64 = t.Time.UnixNano()
}

// GetD 返回从开始时间到当前的 time.Duration。
func (t *TimeCounter) GetD() time.Duration {
	return time.Since(t.Time)
}

// GetS 返回从开始时间到当前的秒数。
func (t *TimeCounter) GetS() int64 {
	return (time.Now().UnixNano() - t.int64) / int64(time.Second)
}

// GetMs 返回从开始时间到当前的毫秒数。
func (t *TimeCounter) GetMs() int64 {
	return (time.Now().UnixNano() - t.int64) / int64(time.Millisecond)
}

// GetUs 返回从开始时间到当前的微秒数。
func (t *TimeCounter) GetUs() int64 {
	return (time.Now().UnixNano() - t.int64) / int64(time.Microsecond)
}

// GetNs 返回从开始时间到当前的纳秒数。
func (t *TimeCounter) GetNs() int64 {
	return time.Now().UnixNano() - t.int64
}

// TimeCost 返回一个闭包，调用闭包时返回自创建以来经过的 time.Duration。
func TimeCost() func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		return time.Since(start)
	}
}
