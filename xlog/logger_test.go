package xlog

import (
	"reflect"
	"testing"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	lfshook "github.com/rifflock/lfshook"
)

func Test_newRotateHook(t *testing.T) {
	type args struct {
		logPath      string
		maxAge       time.Duration
		rotationTime time.Duration
	}
	tests := []struct {
		name string
		args args
		want *lfshook.LfsHook
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRotateHook(tt.args.logPath, tt.args.maxAge, tt.args.rotationTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRotateHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newLogLevel(t *testing.T) {
	type args struct {
		logPath      string
		logFileName  string
		maxAge       time.Duration
		rotationTime time.Duration
	}
	tests := []struct {
		name string
		args args
		want *rotatelogs.RotateLogs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newLogLevel(tt.args.logPath, tt.args.logFileName, tt.args.maxAge, tt.args.rotationTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
