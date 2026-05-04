package xphp

import (
	"fmt"
	"testing"
)

func TestArrayFill(t *testing.T) {
	fmt.Println(ArrayFill(0, 30, 10))
	fmt.Println(ArrayFill(20, 30, 10))
}
