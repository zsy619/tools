package xbyte

import "testing"

func TestByteSize_String(t *testing.T) {
	tests := []struct {
		name string
		b    ByteSize
		want string
	}{
		{"test1", ByteSize(0), "0.00B"},
		{"test2", ByteSize(512), "512.00B"},
		{"test3", ByteSize(1024), "1.00KB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("ByteSize.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
