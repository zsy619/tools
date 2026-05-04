package xgeneric

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	src := []int{-2, -1, -0, 1, 2}
	dst := Filter(func(v int) bool { return v >= 0 }, src)
	fmt.Println(dst)
}

func TestIFF(t *testing.T) {
	a := -1
	fmt.Println(IFF(a > 0, a, 0))
}

func TestIFN(t *testing.T) {
	a := -1
	fmt.Println(IFN(a > 0, func() int { return a }, func() int { return 0 }))
}
