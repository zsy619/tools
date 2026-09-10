package arrays

import "github.com/zsy619/tools/xgeneric"

// Equal 判断两个切片是否相等（元素顺序与值完全一致）。
// 参数：a、b 为待比较的两个切片，元素必须满足 xgeneric.Ordered 约束。
// 返回：当长度相同、nil 状态一致（要么都为 nil 要么都非 nil）且逐元素相等时返回 true，否则返回 false。
func Equal[T xgeneric.Ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
