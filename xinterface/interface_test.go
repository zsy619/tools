package xinterface

import (
	"fmt"
	"testing"
)

func TestIsNil(t *testing.T) {
	var p *int
	var s []string
	var m map[int]bool
	var i any = p

	fmt.Println(IsNil(p)) // true（指针）
	fmt.Println(IsNil(s)) // true（切片）
	fmt.Println(IsNil(m)) // true（映射）
	fmt.Println(IsNil(i)) // true（接口包裹 nil 指针）

	var num int
	var str string
	fmt.Println(IsNil(num)) // false（值类型）
	fmt.Println(IsNil(str)) // false（值类型）
}
