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
// Num2String 将整数格式化为易读的字符串；>=1000 时换算为 K 单位（保留两位小数）。
//
// 参数：
//   - n: 待格式化的整数。
//
// 返回值：转换后的字符串，例如 999 -> "999"，1500 -> "1.5K"。
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
// Range 生成一个 [begin, end] 闭区间序列切片，步长恒为 1。
//
// 参数：
//   - begin: 起始值（含）。
//   - end: 终止值（含）。
//
// 返回值：从 begin 到 end 的切片；当 begin == end 时返回仅包含 begin 的单元素切片。
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
