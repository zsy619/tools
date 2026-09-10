// https://github.com/baidu/go-lib/blob/master/time/time_wait/time_wait.go

package xtime

import "time"

// WaitTill waits until toTime。
//

// Params:。
// - toTime: time to wait until. the number of seconds elapsed since January 1, 1970 UTC.。
// WaitTill 阻塞当前 goroutine，直到 Unix 秒时间戳 toTime 到来；目标时间已过时不等待。
func WaitTill(toTime int64) {
	waitSecs := toTime - time.Now().Unix()
	if waitSecs > 0 {
		time.Sleep(time.Second * time.Duration(waitSecs))
	}
}

// CalcNextTime calculates the nearest time from now, given cycle and offset。
//

// Params:。
// - cycle: cycle in seconds。
// - offset: offset of the next time; in seconds。
//

// Return:。
// - timestamp of next time。
// CalcNextTime 根据周期和偏移计算从当前时间起的最近下一次 Unix 秒时间戳。
func CalcNextTime(cycle int64, offset int64) int64 {
	current := time.Now().Unix()

	if current%cycle == 0 {
		return current + offset
	} else {
		return current - current%cycle + cycle + offset
	}
}
