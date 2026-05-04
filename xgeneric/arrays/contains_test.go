package arrays

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	arr := []string{"wxnacy", "winn"}
	s := "wxnacy"
	i := Contains(arr, s)
	fmt.Println("", i)
	if i != 0 {
		t.Error(i)
	}
}
