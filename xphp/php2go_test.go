package xphp

import "testing"

func TestUcwords(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{StrToLower("test")}, "Test"},
		{"test2", args{StrToLower("UCWORDS")}, "Ucwords"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ucwords(tt.args.str); got != tt.want {
				t.Errorf("Ucwords() = %v, want %v", got, tt.want)
			}
		})
	}
}
