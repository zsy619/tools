package xarray

import (
	"fmt"
	"reflect"
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

func TestArrayRotate(t *testing.T) {
	// Test case 1: rotate a 3x3 integer matrix clockwise
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected1 := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	ArrayRotate(matrix1)
	if !reflect.DeepEqual(matrix1, expected1) {
		t.Errorf("Test case 1 failed: expected %v, but got %v", expected1, matrix1)
	}

	// Test case 2: rotate a 2x2 string matrix clockwise
	matrix2 := [][]string{
		{"a", "b"},
		{"c", "d"},
	}
	expected2 := [][]string{
		{"c", "a"},
		{"d", "b"},
	}
	ArrayRotate(matrix2)
	if !reflect.DeepEqual(matrix2, expected2) {
		t.Errorf("Test case 2 failed: expected %v, but got %v", expected2, matrix2)
	}

	// Test case 3: rotate a 4x4 integer matrix counter-clockwise
	matrix3 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	expected3 := [][]int{
		{4, 8, 12, 16},
		{3, 7, 11, 15},
		{2, 6, 10, 14},
		{1, 5, 9, 13},
	}
	ArrayRotate(matrix3)
	if !reflect.DeepEqual(matrix3, expected3) {
		t.Errorf("Test case 3 failed: expected %v, but got %v", expected3, matrix3)
	}

	// Test case 4: rotate an empty matrix
	matrix4 := [][]float64{}
	expected4 := [][]float64{}
	ArrayRotate(matrix4)
	if !reflect.DeepEqual(matrix4, expected4) {
		t.Errorf("Test case 4 failed: expected %v, but got %v", expected4, matrix4)
	}
}

func TestArrayShuffle(t *testing.T) {
	// Test with an integer slice.
	ints := []int{1, 2, 3, 4, 5}
	ArrayShuffle(ints)
	t.Log(ints)

	// Test with a string slice.
	strings := []string{"apple", "banana", "cherry", "date", "elderberry"}
	ArrayShuffle(strings)
	t.Log(strings)

	// Test with an empty slice.
	empty := []int{}
	ArrayShuffle(empty)
	t.Log(empty)

	// Test with a slice of custom structs.
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 23},
		{"Bob", 34},
		{"Charlie", 45},
		{"David", 56},
		{"Eve", 67},
	}
	ArrayShuffle(people)
	t.Log(people)
}

func TestArrayUnique(t *testing.T) {
	tInput1 := []int{1, 2, 3, 1, 2, 4, 5}
	tOutput1 := []int{1, 2, 3, 4, 5}
	tInput2 := []string{"foo", "bar", "baz", "foo"}
	tOutput2 := []string{"foo", "bar", "baz"}
	tInput3 := []float64{1.1, 2.2, 3.3, 1.1, 2.2, 4.4, 5.5}
	tOutput3 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	{
		result := ArrayUnique(tInput1)
		if !reflect.DeepEqual(result, tOutput1) {
			t.Errorf("Unique(%v) = %v, expected %v", tInput1, result, tOutput1)
		}
	}

	{
		result := ArrayUnique(tInput1)
		if !reflect.DeepEqual(result, tOutput1) {
			t.Errorf("Unique(%v) = %v, expected %v", tInput1, result, tOutput1)
		}
	}

	{
		result := ArrayUnique(tInput2)
		if !reflect.DeepEqual(result, tOutput2) {
			t.Errorf("Unique(%v) = %v, expected %v", tInput2, result, tOutput2)
		}
	}

	{
		result := ArrayUnique(tInput3)
		if !reflect.DeepEqual(result, tOutput3) {
			t.Errorf("Unique(%v) = %v, expected %v", tInput3, result, tOutput3)
		}
	}
}
