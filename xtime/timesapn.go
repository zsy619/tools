package xtime

import "time"

type TimeSpan struct {
	startNS int64
	endNS   int64
}

func (ts *TimeSpan) Start() {
	ts.startNS = time.Now().UnixNano()
}

func (ts *TimeSpan) End() {
	ts.endNS = time.Now().UnixNano()
}

func (ts *TimeSpan) GetTimeSpanMS() float64 {
	return float64(ts.endNS-ts.startNS) / 1000000
}
