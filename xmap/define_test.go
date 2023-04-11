package xmap

import "testing"

func TestH_GetInt(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		h    H
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test1", h: H{"a": "1"}, args: args{key: "a"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.h.GetInt(tt.args.key); got != tt.want {
				t.Errorf("H.GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
