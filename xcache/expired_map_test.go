package xcache

import (
	"testing"
)

func Benchmark_expired_map(b *testing.B) {
	m := NewExpiredMap()
	for i := 0; i < b.N; i++ {
		m.Set(i, i, 10)
	}
}
