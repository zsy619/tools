package xnumber

import (
	"fmt"
	"math"
	"strconv"
)

func Num2String(n int) string {
	if n >= 1000 {
		rt := math.Round(float64(n)/1000*100) / 100
		return fmt.Sprintf("%vK", rt)
	}
	return strconv.Itoa(n)
}

func Range(begin, end int) (result []int) {
	step := end - begin
	if step == 0 {
		result = append(result, begin)
		return
	}
	current := begin
	for i := 0; i <= int(math.Abs(float64(step))); i++ {
		result = append(result, current)
		if step > 0 {
			current += 1
		} else {
			current -= 1
		}
	}
	return
}
