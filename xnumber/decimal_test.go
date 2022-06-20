package xnumber

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	fmt.Println("", Range(5, 5))
	fmt.Println("", Range(6, 5))
	fmt.Println("", Range(-5, 6))
	fmt.Println("", Range(8, 6))
}

func TestDecimal0(t *testing.T) {
	fmt.Println(Decimal0(888.0009))
}
