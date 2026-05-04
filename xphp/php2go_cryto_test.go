package xphp

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	out, _ := Md5("Bys@123123")
	fmt.Println("", out)
}
