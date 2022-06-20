package xstring

import "testing"

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
