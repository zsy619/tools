package xmath

// Min
/**
 * @description:  获取最小值
 * @return {*}
 */
func Min[T NumberAll](vs ...T) T {
	min := vs[0]
	for i := 1; i < len(vs); i++ {
		if min > vs[i] {
			min = vs[i]
		}
	}
	return min
}
