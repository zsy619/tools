package xinterface

import (
	"testing"
	"time"
)

func TestToString(t *testing.T) {
	type args struct {
		i interface{}
	}
	var tint8 int8 = 1
	var ttime time.Time = time.Now()
	tests := []struct {
		name    string
		args    args
		wantStr string
	}{
		// TODO: Add test cases.
		{"int", args{1}, "1"},
		{"int16", args{tint8}, "1"},
		{"ttime", args{ttime}, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStr := ToString(tt.args.i); gotStr != tt.wantStr {
				t.Errorf("ToString() = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}

func TestToStringExt(t *testing.T) {
}
