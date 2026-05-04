package xcrypto

import (
	"testing"
)

func TestBase64StdEncode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{"sssf###@$%sdf"}, "c3NzZiMjI0AkJXNkZg=="},
		{"test2", args{"河南省教育网有限公司"}, "5rKz5Y2X55yB5pWZ6IKy572R5pyJ6ZmQ5YWs5Y+4"},
		{"test3", args{"河南省教育网有限公司Abc"}, "5rKz5Y2X55yB5pWZ6IKy572R5pyJ6ZmQ5YWs5Y+4QWJj"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64StdEncode(tt.args.s); got != tt.want {
				t.Errorf("Base64StdEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64StdDecode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{"c3NzZiMjI0AkJXNkZg=="}, "sssf###@$%sdf"},
		{"test2", args{"5rKz5Y2X55yB5pWZ6IKy572R5pyJ6ZmQ5YWs5Y+4"}, "河南省教育网有限公司"},
		{"test3", args{"5rKz5Y2X55yB5pWZ6IKy572R5pyJ6ZmQ5YWs5Y+4QWJj"}, "河南省教育网有限公司Abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64StdDecode(tt.args.s); got != tt.want {
				t.Errorf("Base64StdDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
