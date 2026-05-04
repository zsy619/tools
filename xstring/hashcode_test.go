package xstring

import (
	"testing"
)

// go test -run=none -bench=BenchmarkString -count=10 -cpu=2,4,6,8
func BenchmarkString(b *testing.B) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"888777777"}, 3495473196},
		{"test2", args{"www.haedu.gov.cn"}, 4059702535},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			if got := ToHashcode(tt.args.s); got != tt.want {
				b.Errorf("String() = %v, want %v", got, tt.want)
			}
		}
	}
}

func TestString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"888777777"}, 3495473196},
		{"test2", args{"www.haedu.gov.cn"}, 4059702535},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToHashcode(tt.args.s); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
