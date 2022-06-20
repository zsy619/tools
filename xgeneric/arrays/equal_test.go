package arrays

import "testing"

func TestEqual(t *testing.T) {
	arr1 := []string{"wxnacy", "wen", "go"}
	arr2 := []string{"wxnacy", "wen", "go"}
	flag := Equal(arr1, arr2)
	if !flag {
		t.Errorf("%v is error", flag)
	}
	arr1 = []string{"wxnacy", "go"}
	arr2 = []string{"wxnacy", "wen", "go"}
	flag = Equal(arr1, arr2)
	if flag {
		t.Errorf("%v is error", flag)
	}
}
