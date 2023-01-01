package xtest

import "testing"

/*
1584 ± : go test -bench=. -cpu=1,2,4 -run=^BenchmarkLen$                                                                                                                                                                 [21m] ✭
goos: darwin
goarch: amd64
pkg: haedu.gov.cn/tools/xtest
cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
BenchmarkLen            1000000000               0.6546 ns/op
BenchmarkLen-2          1000000000               0.6408 ns/op
BenchmarkLen-4          1000000000               0.6377 ns/op
PASS
ok      haedu.gov.cn/tools/xtest        359.207s
*/
func BenchmarkLen(b *testing.B) {
	l := []int{}
	for i := 0; i < b.N; i++ {
		l = append(l, i)
	}
	b.ResetTimer()
	for j := 0; j < len(l); j++ {
	}
}

/*
1585 ± : go test -bench=. -cpu=1,2,4 -run=^BenchmarkLen2$                                                                                                                                                                [27m] ✭
goos: darwin
goarch: amd64
pkg: haedu.gov.cn/tools/xtest
cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
BenchmarkLen            1000000000               0.6353 ns/op
BenchmarkLen-2          1000000000               0.6510 ns/op
BenchmarkLen-4          1000000000               0.6362 ns/op
BenchmarkLen2           1000000000               0.3236 ns/op
BenchmarkLen2-2         1000000000               0.3295 ns/op
BenchmarkLen2-4         1000000000               0.3302 ns/op
PASS
ok      haedu.gov.cn/tools/xtest        568.737ss
*/
func BenchmarkLen2(b *testing.B) {
	l := []int{}
	for i := 0; i < b.N; i++ {
		l = append(l, i)
	}
	b.ResetTimer()
	ll := len(l)
	for j := 0; j < ll; j++ {
	}
}
