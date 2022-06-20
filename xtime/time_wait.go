// https://github.com/baidu/go-lib/blob/master/time/time_wait/time_wait.go

package xtime

import "time"

// WaitTill waits until toTime
//
// Params:
//     - toTime: time to wait until. the number of seconds elapsed since January 1, 1970 UTC.
func WaitTill(toTime int64) {
	waitSecs := toTime - time.Now().Unix()
	if waitSecs > 0 {
		time.Sleep(time.Second * time.Duration(waitSecs))
	}
}

// CalcNextTime calculates the nearest time from now, given cycle and offset
//
// Params:
//  - cycle: cycle in seconds
//  - offset: offset of the next time; in seconds
//
// Return:
//  - timestamp of next time
func CalcNextTime(cycle int64, offset int64) int64 {
	current := time.Now().Unix()

	if current%cycle == 0 {
		return current + offset
	} else {
		return current - current%cycle + cycle + offset
	}
}
