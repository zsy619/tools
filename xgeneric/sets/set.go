package sets

import "github.com/zsy619/tools"

// Set 表示一个元素集合，内部基于 map[T]struct{} 实现以节省内存。
type Set[T comparable] map[T]struct{}

// NewSet 创建一个包含指定元素的集合。
// 参数：es 为要加入集合的可变元素列表。
// 返回：包含 es（去重后）所有元素的 Set[T]。
func NewSet[T comparable](es ...T) Set[T] {
	s := Set[T]{}
	for _, e := range es {
		s.Add(e)
	}
	return s
}

// Len 返回集合 s 中的元素数量。
func (s *Set[T]) Len() int {
	return len(*s)
}

// IsEmpty 判断集合 s 是否为空。
// 返回：当 s 中元素数量为 0 时返回 true，否则返回 false。
func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

// Add 将 target 加入集合。
// 参数：target 为要添加的元素。
// 返回：成功添加返回 true；若 target 已存在则返回 false 且集合保持不变。
func (s *Set[T]) Add(target T) bool {
	_, ok := (*s)[target]
	if ok {
		return false
	}

	(*s)[target] = tools.Empty
	return true
}

// AddSlice 将一组元素加入集合。
// 若元素已存在则不产生任何效果（无幂幂）。
func (s *Set[T]) AddSlice(es ...T) {
	for _, e := range es {
		(*s)[e] = struct{}{}
	}
}

// Remove 从集合中移除指定元素。
// 若元素不在集合中则不产生任何效果。
func (s *Set[T]) Remove(es ...T) {
	for _, e := range es {
		delete(*s, e)
	}
}

// Contains 判断元素 v 是否在集合 s 中。
// 返回：找到则返回 true，否则返回 false。
func (s *Set[T]) Contains(v T) bool {
	_, ok := (*s)[v]
	return ok
}

// Clone 创建包含相同元素的新集合（深拷贝元素）。
// 返回：与 s 包含相同元素的全新 Set[T]，修改新集合不会影响原集合。
func (s *Set[T]) Clone() Set[T] {
	r := Set[T]{}
	r.AddSlice(s.ToSlice()...)
	return r
}

// ToSlice 将集合转换为切片。
// 返回：包含集合中所有元素的 []T；空集合时返回空切片。元素顺序随机。
func (s *Set[T]) ToSlice() []T {
	r := make([]T, 0, s.Len())

	for e := range *s {
		r = append(r, e)
	}

	return r
}
