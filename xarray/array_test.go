package xarray

import (
	"fmt"
	"testing"
)

func TestArrayRotateInt(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix)
	ArrayRotateInt(matrix)
	fmt.Println(matrix)
}

func TestArrayReverseAny(t *testing.T) {
	s := []interface{}{1, "2", "中文", uint(3), byte(4), "a", float64(5)}
	fmt.Println(s)
	ArrayReverseAny(s)
	fmt.Println(s)
}

func TestArrayShuffleAny(t *testing.T) {
	slice := []interface{}{"a", "b", "c", "d", "e", "f"}
	ArrayShuffleAny(slice)
	fmt.Println(slice)
	for _, v := range slice {
		fmt.Println(v.(string))
	}
	ArrayShuffleAny(slice)
	fmt.Println(slice)
	ArrayShuffleAny(slice)
	fmt.Println(slice)
	ArrayShuffleAny(slice)
	fmt.Println(slice)
	ArrayShuffleAny(slice)
	fmt.Println(slice)
	ArrayShuffleAny(slice)
	fmt.Println(slice)
	ArrayShuffleAny(slice)
	fmt.Println(slice)
}
