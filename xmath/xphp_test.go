package xmath

import (
	"reflect"
	"testing"
)

func TestNewSlice(t *testing.T) {
	{
		start := 1
		end := 10
		step := 2
		s := NewSlice(start, end, step)
		if !reflect.DeepEqual(s, []int{1, 3, 5, 7, 9}) {
			t.Errorf("NewSlice failed")
		}
	}

	{
		start := 1.0
		end := 10.0
		step := 100.0
		s := NewSlice(start, end, step)
		if !reflect.DeepEqual(s, []int{1, 3, 5, 7, 9}) {
			t.Errorf("NewSlice failed")
		}
	}
}

func Test_XRange(t *testing.T) {
	{
		ch := XRange(1, 10, 2)
		s := []int{}
		for i := range ch {
			s = append(s, i)
		}
		if !reflect.DeepEqual(s, []int{1, 3, 5, 7, 9}) {
			t.Errorf("c failed")
		}
	}

	{
		ch := XRange(1.0, 10.0, 2.0)
		s := []float64{}
		for i := range ch {
			s = append(s, i)
		}
		t.Log(s)
	}
}
