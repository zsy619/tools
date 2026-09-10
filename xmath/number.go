package xmath

import (
	"math/big"
	"strconv"
)

// BigIntToFloat64 converts big.Int to float64 type
// BigIntToFloat64 将 *big.Int 转换为 float64，可能丢失精度。
//
// 参数：
//   - bigInt: 大整数指针；nil 时返回 0。
//
// 返回值：float64 表示的数值。
func BigIntToFloat64(bigInt *big.Int) float64 {
	f := new(big.Float).SetInt(bigInt)
	val, _ := f.Float64()
	return val
}

// Float64ToBigInt converts float64 to big.Int type
// Float64ToBigInt 将 float64 转换为 *big.Int。
//
// 参数：
//   - float: 浮点数值。
//
// 返回值：与输入对应的 *big.Int；如果格式转换失败，则返回零值 *big.Int。
func Float64ToBigInt(float float64) *big.Int {
	bigInt := new(big.Int)
	bigInt.SetString(strconv.FormatFloat(float, 'f', -1, 64), 10)
	return bigInt
}

// ListIntToListString converts a list of big.int to a list of string
// ListIntToListString 将 int 切片按顺序转换为对应的字符串切片。
//
// 参数：
//   - listInt: int 切片。
//
// 返回值：与 listInt 等长的字符串切片（每个元素使用 strconv.Itoa 转换）。
func ListIntToListString(listInt []int) []string {
	listString := make([]string, len(listInt))
	for idx, i := range listInt {
		listString[idx] = strconv.Itoa(i)
	}
	return listString
}

// IntOrEmptyInt returns an empty int if a int pointer is nil.
// Otherwise returns the int value of the pointer.
// IntOrEmptyInt 在指针为 nil 时返回 0，否则返回指针指向的值。
//
// 参数：
//   - s: 可能为 nil 的 *int。
//
// 返回值：解引用后的 int 值；nil 时返回 0。
func IntOrEmptyInt(s *int) int {
	if s != nil {
		return *s
	}
	return 0
}

// IntToPtr is a convenience func that converts a int to it's pointer
// IntToPtr 将 int 包装为指针。
//
// 参数：
//   - s: 整数值。
//
// 返回值：指向 s 的 *int。
func IntToPtr(s int) *int {
	return &s
}
