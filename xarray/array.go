// Package xarray 提供数组 / 切片相关的工具函数。
//
// 该包的设计目标是补足 Go 1.21+ 标准库 slices、math/rand/v2 在
// 日常业务中常见的：
//   - 矩阵旋转、切片洗牌、切片合并、切片去重；
//   - 26 字母表、A-Z 字符串生成等小工具；
//   - 通过反射支持任意类型切片的通用操作（ReverseAny、ShuffleAny）。
//
// 本包刻意保持零依赖（除同仓库的工具与 xreflect/xgeneric 外）。
package xarray

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"reflect"
	"time"
)

// ArrayRotateInt 将 NxN 整数矩阵原地顺时针旋转 90°。
//
// 要求 matrix 是方阵（行数 = 列数），否则行为未定义。
//
// 注意：本函数是 *int* 特化版本，如果矩阵元素为其它类型请改用
// 泛型版 ArrayRotate[T any]。
func ArrayRotateInt(matrix [][]int) {
	m := len(matrix)
	tmp := 0

	for i := 0; i < m/2; i++ {
		for j := i; j < m-1-i; j++ {
			tmp = matrix[i][j]
			matrix[i][j] = matrix[m-1-j][i]
			matrix[m-1-j][i] = matrix[m-1-i][m-1-j]
			matrix[m-1-i][m-1-j] = matrix[j][m-1-i]
			matrix[j][m-1-i] = tmp
		}
	}
}

// ArrayRotate 将 NxN 泛型矩阵原地顺时针旋转 90°。
//
// 要求 matrix 是方阵（行数 = 列数），否则行为未定义。
//
// 典型用途：图像处理、坐标变换等需要把二维数组转置 + 镜像的场景。
func ArrayRotate[T any](matrix [][]T) {
	m := len(matrix)

	for i := 0; i < m/2; i++ {
		for j := i; j < m-1-i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[m-1-j][i]
			matrix[m-1-j][i] = matrix[m-1-i][m-1-j]
			matrix[m-1-i][m-1-j] = matrix[j][m-1-i]
			matrix[j][m-1-i] = tmp
		}
	}
}

// ArrayReverseAny 通过反射将传入的可索引对象原地反转。
//
// 与 slices.Reverse 不同，本函数接受任意支持 Len() 与 Index 的
// reflect.Value（即切片、数组、字符串等）。
//
// 当 s 不是可索引类型时，reflect.ValueOf / Swapper 会触发 panic，
// 请调用方保证入参合法。
func ArrayReverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// ArrayShuffleAny 使用 math/rand/v2 随机源将 []any 切片原地洗牌（Fisher-Yates）。
//
// 该函数随机源使用当前时间纳秒数播种；如需可复现的随机序列请自行
// 创建 *rand.Rand 实例并改写本函数。
func ArrayShuffleAny(slice []any) {
	seed1, seed2 := splitTimeSeed()
	r := rand.New(rand.NewPCG(seed1, seed2))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.IntN(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

// ArrayShuffle 使用 math/rand/v2 随机源将切片原地洗牌（Fisher-Yates）。
//
// 该函数随机源使用当前时间纳秒数播种；如需可复现的随机序列请自行
// 创建 *rand.Rand 实例并改写本函数。
func ArrayShuffle[T any](slice []T) {
	seed1, seed2 := splitTimeSeed()
	r := rand.New(rand.NewPCG(seed1, seed2))
	for i := len(slice) - 1; i > 0; i-- {
		j := r.IntN(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// ArrayMerge 合并若干切片为一个新切片。
//
// 入参为空时返回 ("", error)。否则按入参顺序拼接，保留每个切片
// 内部顺序，但返回新切片（不会修改任何入参）。
//
// 与 slices.Concat 等价的语义。
func ArrayMerge[T any](ary ...[]T) ([]T, error) {
	if ary == nil || len(ary) <= 0 {
		return nil, errors.New("参数错误")
	}
	var b []T
	for _, item := range ary {
		b = append(b, item...)
	}
	return b, nil
}

// ArrayUnique 对 comparable 类型的切片去重并保留首次出现顺序。
//
// 输入为 nil 时返回 nil；空切片返回空切片。
//
// 注意：相等判定走 Go 内置 == 与 map hash，对包含浮点的切片请自行
// 处理 NaN。
func ArrayUnique[T comparable](slice []T) []T {
	m := make(map[T]bool)
	var result []T
	for _, v := range slice {
		if !m[v] {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}

// ArrayChar 创建一个26元素的数组，分别放置'A'-'Z'。
//
// 常用于批量生成字母枚举、测试输入、Excel 列号转换等场景。
func ArrayChar() []string {
	myChars := [26]string{}
	for i := 0; i < 26; i++ {
		myChars[i] = fmt.Sprintf("%c", 'A'+byte(i))
	}
	return myChars[:]
}

// splitTimeSeed 将当前时间纳秒数拆成两个 uint64 用于 rand.NewPCG。
//
// PCG 随机源需要两个 64 位种子；这里把当前 UnixNano 的高低位拆开，
// 既保证不同调用之间的种子差异，又避免引入额外的 rand.Source。
func splitTimeSeed() (uint64, uint64) {
	now := uint64(time.Now().UnixNano())
	return now, now ^ 0x9E3779B97F4A7C15
}
