package xmath

// Max 返回一组数值中的最大值（要求至少传入一个参数）。
//
// 参数：
//   - vs: 任意长度的数值切片，元素必须实现 NumberAll 接口。
//
// 返回值：vs 中的最大元素（与元素同类型 T）。
func Max[T NumberAll](vs ...T) T {
	max := vs[0]
	for i := 1; i < len(vs); i++ {
		if max < vs[i] {
			max = vs[i]
		}
	}
	return max
}

// MaxInt 获取最大整数
// MaxInt 返回一组 int 中的最大值（要求至少传入一个参数）。
//
// 参数：
//   - vs: 任意长度的 int 切片。
//
// 返回值：vs 中的最大元素。
func MaxInt(vs ...int) int {
	max := vs[0]
	for i := 1; i < len(vs); i++ {
		if max < vs[i] {
			max = vs[i]
		}
	}
	return max
}

// MinInt 获取最小整数
// MinInt 返回一组 int 中的最小值（要求至少传入一个参数）。
//
// 参数：
//   - vs: 任意长度的 int 切片。
//
// 返回值：vs 中的最小元素。
func MinInt(vs ...int) int {
	min := vs[0]
	for i := 1; i < len(vs); i++ {
		if min > vs[i] {
			min = vs[i]
		}
	}
	return min
}

// SumInt 计算整数总和
// SumInt 对一组 int 求和（空参数返回 0）。
//
// 参数：
//   - args: 任意长度的 int 列表。
//
// 返回值：args 中所有元素的累加和。
func SumInt(args ...int) (v int) {
	for _, arg := range args {
		v += arg
	}
	return
}

// Sum 对一组数值求和，元素类型必须实现 NumberAll。
//
// 参数：
//   - args: 任意长度的数值列表。
//
// 返回值：累加和，与元素同类型 T；入参为空时返回 T 的零值。
func Sum[T NumberAll](args ...T) (v T) {
	for _, arg := range args {
		v += arg
	}
	return
}

// BitSplitInt 按位分解整数
// BitSplitInt 将 v 的二进制中所有为 1 的位对应的 2 的幂收集为切片返回。
//
// 参数：
//   - v: 待分解的非负整数。
//
// 返回值：v 的所有置位对应的 2 的幂切片（升序）。
func BitSplitInt(v int) (vs []int) {
	vs = make([]int, 0)
	for i := 1; i <= v; i *= 2 {
		if (i & v) == i {
			vs = append(vs, i)
		}
	}
	return vs
}

// InSliceInt 判断某值是否在切片中
// InSliceInt 判断 v 是否出现在 is 中（线性查找）。
//
// 参数：
//   - v: 待查找的值。
//   - is: 候选切片。
//
// 返回值：v 存在时返回 true，否则 false。
func InSliceInt(v int, is []int) bool {
	for _, i := range is {
		if i == v {
			return true
		}
	}
	return false
}

// InSlice
/**
 * @description:  判断某值是否在切片中
 * @return {*}
 */
func InSlice[T NumberAll](v T, is []T) bool {
	for _, i := range is {
		if i == v {
			return true
		}
	}
	return false
}

// MaxInt64 获取最大整数
// MaxInt64 返回一组 int64 中的最大值（要求至少传入一个参数）。
//
// 参数：
//   - vs: 任意长度的 int64 切片。
//
// 返回值：vs 中的最大元素。
func MaxInt64(vs ...int64) int64 {
	max := vs[0]
	for i := 1; i < len(vs); i++ {
		if max < vs[i] {
			max = vs[i]
		}
	}
	return max
}

// MinInt64 获取最小整数
// MinInt64 返回一组 int64 中的最小值（要求至少传入一个参数）。
//
// 参数：
//   - vs: 任意长度的 int64 切片。
//
// 返回值：vs 中的最小元素。
func MinInt64(vs ...int64) int64 {
	min := vs[0]
	for i := 1; i < len(vs); i++ {
		if min > vs[i] {
			min = vs[i]
		}
	}
	return min
}

// SumInt64 计算64位整数总和
// SumInt64 对一组 int64 求和（空参数返回 0）。
//
// 参数：
//   - args: 任意长度的 int64 列表。
//
// 返回值：args 中所有元素的累加和。
func SumInt64(args ...int64) (v int64) {
	for _, arg := range args {
		v += arg
	}
	return
}

// BitSplitInt64 按位分解64位整数
// BitSplitInt64 将 v 的二进制中所有为 1 的位对应的 2 的幂收集为 int64 切片返回。
//
// 参数：
//   - v: 待分解的非负 int64。
//
// 返回值：v 的所有置位对应的 2 的幂切片（升序）。
func BitSplitInt64(v int64) (vs []int64) {
	vs = make([]int64, 0)
	for i := int64(1); i <= v; i *= 2 {
		if (i & v) == i {
			vs = append(vs, i)
		}
	}
	return vs
}

// InSliceInt64 判断某值是否在切片中
// InSliceInt64 判断 v 是否出现在 is 中（线性查找）。
//
// 参数：
//   - v: 待查找的值。
//   - is: 候选切片。
//
// 返回值：v 存在时返回 true，否则 false。
func InSliceInt64(v int64, is []int64) bool {
	for _, i := range is {
		if i == v {
			return true
		}
	}
	return false
}
