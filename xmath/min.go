package xmath

// Min
/**
 * @description:  获取最小值
 * @return {*}
 */
// Min 返回一组数值中的最小值（要求至少传入一个参数）。
//
// 参数：
//   - vs: 任意长度的数值切片，元素必须实现 NumberAll 接口。
//
// 返回值：vs 中的最小元素（与元素同类型 T）。
func Min[T NumberAll](vs ...T) T {
	min := vs[0]
	for i := 1; i < len(vs); i++ {
		if min > vs[i] {
			min = vs[i]
		}
	}
	return min
}
