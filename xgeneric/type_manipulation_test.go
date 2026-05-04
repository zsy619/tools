package xgeneric

import (
	"fmt"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	o1 := IsEmpty("")
	fmt.Println(o1)
	o2 := IsEmpty(0)
	fmt.Println(o2)
}

func TestFromAnySlice(t *testing.T) {
	var in []any
	in = append(in, 10)
	in = append(in, 1)
	in = append(in, 9)
	in = append(in, 6)
	in = append(in, 2)
	fmt.Println(in)

	var out []int
	var ok bool
	out, ok = FromAnySlice[int](in)
	fmt.Println(out, ok)
}
