package xmath

// Avg
/**
 * @description: 计算平均值
 * @return {*}
 */
// Avg 计算一组数值的算术平均值。
//
// 参数：
//   - vs: 任意长度的数值切片，元素必须实现 NumberAll 接口。
//
// 返回值：vs 的算术平均值（与元素同类型 T）。
//
// 注意：入参为空时返回 T 的零值（0）；调用方需自行处理该情况。
func Avg[T NumberAll](vs ...T) T {
	return Sum(vs...) / T(len(vs))
}
