package xstring

import (
	"testing"
)

func TestPasswordCheck(t *testing.T) {
	type args struct {
		minLength int
		maxLength int
		minLevel  int
		pwd       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{8, 32, PasswordLevelS, "fb0da#07E"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PasswordCheck(tt.args.minLength, tt.args.maxLength, tt.args.minLevel, tt.args.pwd); (err != nil) != tt.wantErr {
				t.Errorf("PasswordCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	type args struct {
		minLength int
		maxLength int
		pwd       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1", args{minLength: 8, maxLength: 32, pwd: "fb0da#07E"}, true,
		},
		{
			"test2", args{minLength: 8, maxLength: 32, pwd: "12358678"}, false,
		},
		{
			"test2", args{minLength: 8, maxLength: 32, pwd: "fb0da#07"}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyPassword(tt.args.minLength, tt.args.maxLength, tt.args.pwd); got != tt.want {
				t.Errorf("VerifyPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
