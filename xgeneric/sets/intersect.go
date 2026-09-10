package sets

// Intersect 返回多个集合的交集。
// 即同时存在于所有传入集合中的元素组成的 Set[T]。
// 参数：sets 为可变数量的 Set[T]。
// 返回：包含所有集合共有元素的新 Set[T]。
func Intersect[T comparable](sets ...Set[T]) Set[T] {
	records := map[T]int{}

	for _, s := range sets {
		for _, e := range s.ToSlice() {
			records[e] = records[e] + 1
		}
	}

	r := NewSet[T]()

	for e, num := range records {
		if num == len(sets) {
			r.Add(e)
		}
	}

	return r
}
