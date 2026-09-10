package xphp

import (
	"math"
	"math/cmplx"
	"math/rand"
	"strconv"
	"time"
)

// MaxPrecision 是数学相关函数使用的默认最大精度。
var MaxPrecision = 14

// MtRand 是 Rand() 的别名（对应 PHP mt_rand()）。
func MtRand(min, max int) int {
	return Rand(min, max)
}

// Acos 计算复数 val 的反余弦。
func Acos(val complex128) complex128 {
	return cmplx.Acos(val)
}

// Acosh 计算复数 val 的反双曲余弦。
func Acosh(val complex128) complex128 {
	return cmplx.Acosh(val)
}

// Asin 计算复数 val 的反正弦。
func Asin(val complex128) complex128 {
	return cmplx.Asin(val)
}

// Asinh 计算复数 val 的反双曲正弦。
func Asinh(val complex128) complex128 {
	return cmplx.Asinh(val)
}

// Atan 计算复数 val 的反正切。
func Atan(val complex128) complex128 {
	return cmplx.Atan(val)
}

// Atanh 计算复数 val 的反双曲正切。
func Atanh(val complex128) complex128 {
	return cmplx.Atanh(val)
}

// Base_convert 将数字字符串 num 从 frombase 进制转换为 tobase 进制（对应 PHP base_convert()）。
func Base_convert(num string, frombase, tobase int) (string, error) {
	i, err := strconv.ParseInt(num, frombase, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, tobase), nil
}

// Abs 返回 number 的绝对值（对应 PHP abs()）。
func Abs(number float64) float64 {
	return math.Abs(number)
}

// Rand 返回 [min, max] 闭区间内的随机整数（对应 PHP rand()）。
// 取值范围限制在 [0, 2147483647]。
func Rand(min, max int) int {
	if min > max {
		panic("min: min cannot be greater than max")
	}
	// PHP: getrandmax()
	if int31 := 1<<31 - 1; max > int31 {
		panic("max: max can not be greater than " + strconv.Itoa(int31))
	}
	if min == max {
		return min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max+1-min) + min
}

// Round 返回 value 四舍五入后的值（对应 PHP round()）。
func Round(value float64) float64 {
	return math.Floor(value + 0.5)
}

// Floor 返回不大于 value 的最大整数，即向下取整（对应 PHP floor()）。
func Floor(value float64) float64 {
	return math.Floor(value)
}

// Ceil 返回不小于 value 的最小整数，即向上取整（对应 PHP ceil()）。
func Ceil(value float64) float64 {
	return math.Ceil(value)
}

// Pi 返回圆周率 π 的近似值（对应 PHP pi()）。
func Pi() float64 {
	return math.Pi
}

// Max 返回 nums 中的最大值，参数个数不得少于 2（对应 PHP max()）。
func Max(nums ...float64) float64 {
	if len(nums) < 2 {
		panic("nums: the nums length is less than 2")
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		max = math.Max(max, nums[i])
	}
	return max
}

// Min 返回 nums 中的最小值，参数个数不得少于 2（对应 PHP min()）。
func Min(nums ...float64) float64 {
	if len(nums) < 2 {
		panic("nums: the nums length is less than 2")
	}
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		min = math.Min(min, nums[i])
	}
	return min
}

// Decbin 将十进制整数 number 转换为二进制字符串（对应 PHP decbin()）。
func Decbin(number int64) string {
	return strconv.FormatInt(number, 2)
}

// Bindec 将二进制字符串 str 解析为十进制数字字符串（对应 PHP bindec()）。
func Bindec(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 10), nil
}

// Hex2bin 将十六进制字符串 data 解析并输出为二进制字符串（对应 PHP hex2bin()）。
func Hex2bin(data string) (string, error) {
	i, err := strconv.ParseInt(data, 16, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 2), nil
}

// Bin2hex 将二进制字符串 str 解析并输出为十六进制字符串（对应 PHP bin2hex()）。
func Bin2hex(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}

// Dechex 将十进制整数 number 转换为十六进制字符串（对应 PHP dechex()）。
func Dechex(number int64) string {
	return strconv.FormatInt(number, 16)
}

// Hexdec 将十六进制字符串 str 转换为十进制整数（对应 PHP hexdec()）。
func Hexdec(str string) (int64, error) {
	return strconv.ParseInt(str, 16, 0)
}

// Decoct 将十进制整数 number 转换为八进制字符串（对应 PHP decoct()）。
func Decoct(number int64) string {
	return strconv.FormatInt(number, 8)
}

// Octdec 将八进制字符串 str 转换为十进制整数（对应 PHP octdec()）。
func Octdec(str string) (int64, error) {
	return strconv.ParseInt(str, 8, 0)
}

// BaseConvert 将数字字符串 number 从 frombase 进制转换为 tobase 进制（对应 PHP base_convert()）。
func BaseConvert(number string, frombase, tobase int) (string, error) {
	i, err := strconv.ParseInt(number, frombase, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, tobase), nil
}

// IsNan 判断 val 是否为非数值 NaN（对应 PHP is_nan()）。
func IsNan(val float64) bool {
	return math.IsNaN(val)
}
