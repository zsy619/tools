package xmath

// Avg
/**
 * @description: 计算平均值
 * @return {*}
 */
func Avg[T NumberAll](vs ...T) T {
	return Sum(vs...) / T(len(vs))
}
