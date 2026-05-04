package arrays

import "github.com/zsy619/tools/xgeneric"

// Unique 数组或slice边去重
func Unique[T xgeneric.Ordered](arr []T) []T {
	tmp := make(map[T]struct{})
	l := len(arr)
	if l == 0 {
		return arr
	}

	rel := make([]T, 0, l)
	for _, item := range arr {
		_, ok := tmp[item]
		if ok {
			continue
		}
		tmp[item] = struct{}{}
		rel = append(rel, item)
	}

	return rel[:len(tmp)]
}
