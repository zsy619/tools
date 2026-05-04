package xio

import "testing"

func TestDiv(t *testing.T) {
	type args struct {
		Num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"0", args{0}, "A"},
		{"1", args{1}, "A"},
		{"2", args{2}, "B"},
		{"20", args{20}, "T"},
		{"33", args{27}, "AA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Div(tt.args.Num); got != tt.want {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}
