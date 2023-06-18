package xdatabase

import "testing"

func TestSafeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SafeString(tt.args.s); got != tt.want {
				t.Errorf("SafeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
