package xmath

import "fmt"

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
