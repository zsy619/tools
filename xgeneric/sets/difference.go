package sets

// Difference 返回两个集合的差集（即 l 中存在但不在 r 中的元素集合）。
// 参数：l、r 为两个 Set[T]。
// 返回：包含 l 中去除 r 共有元素后的新 Set[T]。
func Difference[T comparable](l, r Set[T]) Set[T] {
	re := l.Clone()
	re.Remove(r.ToSlice()...)
	return re
}
