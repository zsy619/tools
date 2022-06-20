package xstring

import (
	"fmt"
	"strings"
	"testing"
)

func TestFirstLetterUpper(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{""}, ""},
		{"test2", args{"zhushuyan"}, "Zhushuyan"},
		{"test3", args{"Zhushuyan"}, "Zhushuyan"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstLetterUpper(tt.args.str); got != tt.want {
				t.Errorf("FirstLetterUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstLetterLower(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{""}, ""},
		{"test2", args{"zhushuyan"}, "zhushuyan"},
		{"test3", args{"Zhushuyan"}, "zhushuyan"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstLetterLower(tt.args.str); got != tt.want {
				t.Errorf("FirstLetterLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStringWidth(t *testing.T) {
	str := `ssadaf
    asfafd`
	fmt.Println("-->", str)
	fmt.Println("-->", strings.ReplaceAll(strings.ReplaceAll(str, "\n", ""), "\t", ""))
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"中国"}, 4},
		{"t1", args{"中国\t"}, 5},
		{"t1", args{"ILove中国"}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStringWidth(tt.args.str); got != tt.want {
				t.Errorf("GetStringWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}
