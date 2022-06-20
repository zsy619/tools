package xphp

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	out, _ := Md5("12aAAAA3")
	fmt.Println("", out)
}
