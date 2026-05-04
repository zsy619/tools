package xmath

import (
	"fmt"
	"math"
	"strconv"
)

// Num2String
/**
 * @description: 数字转字符串
 * @param {int} n
 * @return {string}
 */
func Num2String(n int) string {
	if n >= 1000 {
		rt := math.Round(float64(n)/1000*100) / 100
		return fmt.Sprintf("%vK", rt)
	}
	return strconv.Itoa(n)
}

// Range
/**
 * @description: 生成一个范围内的整数切片
 * @return {*}
 */
func Range[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](begin, end T) (result []T) {
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
