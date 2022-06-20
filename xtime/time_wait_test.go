// https://github.com/baidu/go-lib/blob/master/time/time_wait/time_wait.go

package xtime

import (
	"testing"
	"time"
)

func TestWaitTill_case1(t *testing.T) {
	waitSecs := int64(2)
	start := time.Now().Unix()

	toTime := start + waitSecs

	WaitTill(toTime)

	passSecs := time.Now().Unix() - start

	if passSecs != waitSecs {
		t.Errorf("err in WaitTill(): wait=%d, pass=%d", waitSecs, passSecs)
	}
}

func TestWaitTill_case2(t *testing.T) {
	start := time.Now().Unix()

	toTime := start - 2

	WaitTill(toTime)

	passSecs := time.Now().Unix() - start

	if passSecs != 0 {
		t.Errorf("err in WaitTill(): wait=0, pass=%d", passSecs)
	}
}
