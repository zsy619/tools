package xmath

import "fmt"

// NewSlice 生成一个等差数值切片 [start, end]，步长为 step（包含两端）。
//
// 参数：
//   - start: 起始值（包含）。
//   - end: 终止值（包含）。
//   - step: 步长，必须 > 0；否则或 end < start 时返回空切片。
//
// 返回值：等差数值切片。
func NewSlice[T NumberAll](start, end, step T) []T {
	if step <= 0 || end < start {
		return []T{}
	}
	s := make([]T, 0, 1+int((end-start)/step))
	for start <= end {
		s = append(s, start)
		start += step
	}
	return s
}

// XRange 以 channel 形式返回一个序列，类似 Python 的 xrange/PHP 的 range。
//
// 参数（1 ~ 3 个）：
//   - 1 个参数 stop：序列从 0 到 stop（不含 stop），步长 1。
//   - 2 个参数 start, stop：序列从 start 到 stop（不含 stop），步长 1。
//   - 3 个参数 start, stop, step：序列从 start 到 stop（不含 stop），步长为 step（可正可负）。
//
// 返回值：仅写 chan T；调用方需持续读取直至 channel 关闭。
//
// 副作用：会启动一个 goroutine 负责写入；该 goroutine 写入完成后会关闭 channel。
//
// 注意：参数个数不合法（小于 1 或大于 3）时会向 stderr 打印错误并继续执行默认行为。
func XRange[T NumberAll](args ...T) chan T {
	if l := len(args); l < 1 || l > 3 {
		fmt.Println("error args length, xRangeInt requires 1-3 int arguments")
	}
	var start, stop T
	var step T = 1
	switch len(args) {
	case 1:
		stop = args[0]
		start = 0
	case 2:
		start, stop = args[0], args[1]
	case 3:
		start, stop, step = args[0], args[1], args[2]
	}

	ch := make(chan T)
	go func() {
		if step > 0 {
			for start < stop {
				ch <- start
				start = start + step
			}
		} else {
			for start > stop {
				ch <- start
				start = start + step
			}
		}
		close(ch)
	}()

	return ch
}
