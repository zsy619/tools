package arrays

import (
	"testing"
)

func TestUnique(t *testing.T) {
	arrayInts := []int{1, 2, 2, 3, 4}
	arrayStrs := []string{"A", "B", "C", "D", "C"}

	// 测试整数
	t.Log(Unique(arrayInts))

	// 测试字符串
	t.Log(Unique(arrayStrs))
}

func BenchmarkUniqueGeneric(b *testing.B) {
	arrayStrs := []string{"A", "B", "C", "D", "C"}
	for i := 0; i < b.N; i++ {
		_ = Unique(arrayStrs)
	}
}
