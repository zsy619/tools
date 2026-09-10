package sets

// Union 返回多个集合的并集。
// 即出现在任一传入集合中的所有元素（去重）组成的 Set[T]。
// 参数：sets 为可变数量的 Set[T]。
// 返回：包含所有集合元素（去重）的新 Set[T]。
func Union[T comparable](sets ...Set[T]) Set[T] {
	r := NewSet[T]()
	for _, s := range sets {
		r.AddSlice(s.ToSlice()...)
	}
	return r
}
