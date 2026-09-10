package xtime

import "time"

// TimeSpan 用于记录一段耗时，以纳秒为单位保存开始与结束时刻。
// 字段未导出，使用方通过 Start、End 与 GetTimeSpanMS 控制与读取。
type TimeSpan struct {
	startNS int64
	endNS   int64
}

// Start 标记当前时刻为计时的起点，覆盖之前记录的 startNS。
// 若尚未调用 End，结果值仅包含部分跨度。
func (ts *TimeSpan) Start() {
	ts.startNS = time.Now().UnixNano()
}

// End 标记当前时刻为计时的终点，覆盖之前记录的 endNS。
// 若尚未调用 Start，endNS 与未初始化的 startNS 计算差值，结果无意义。
func (ts *TimeSpan) End() {
	ts.endNS = time.Now().UnixNano()
}

// GetTimeSpanMS 返回从 Start 到 End 所经过的时间（毫秒，小数形式）。
// 若调用顺序错误（例如 End 在 Start 之前）将返回负值。
func (ts *TimeSpan) GetTimeSpanMS() float64 {
	return float64(ts.endNS-ts.startNS) / 1000000
}
