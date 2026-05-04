package xgeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	is := assert.New(t)

	result1 := Contains([]int{0, 1, 2, 3, 4, 5}, 5)
	result2 := Contains([]int{0, 1, 2, 3, 4, 5}, 6)

	is.Equal(result1, true)
	is.Equal(result2, false)
}

func TestContainsBy(t *testing.T) {
	is := assert.New(t)

	type a struct {
		A int
		B string
	}

	a1 := []a{{A: 1, B: "1"}, {A: 2, B: "2"}, {A: 3, B: "3"}}
	result1 := ContainsBy(a1, func(t a) bool { return t.A == 1 && t.B == "2" })
	result2 := ContainsBy(a1, func(t a) bool { return t.A == 2 && t.B == "2" })

	a2 := []string{"aaa", "bbb", "ccc"}
	result3 := ContainsBy(a2, func(t string) bool { return t == "ccc" })
	result4 := ContainsBy(a2, func(t string) bool { return t == "ddd" })

	is.Equal(result1, false)
	is.Equal(result2, true)
	is.Equal(result3, true)
	is.Equal(result4, false)
}

func TestSome(t *testing.T) {
	is := assert.New(t)

	result1 := Some([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Some([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Some([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})

	is.Equal(result1, true)
	is.Equal(result2, true)
	is.Equal(result3, false)
}

func TestEvery(t *testing.T) {
	is := assert.New(t)

	result1 := Every([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Every([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Every([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})

	is.Equal(result1, true)
	is.Equal(result2, false)
	is.Equal(result3, false)
}

func TestIntersect(t *testing.T) {
	is := assert.New(t)

	result1 := Intersect([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	result2 := Intersect([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	result3 := Intersect([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	result4 := Intersect([]int{0, 6}, []int{0, 1, 2, 3, 4, 5})
	result5 := Intersect([]int{0, 6, 0}, []int{0, 1, 2, 3, 4, 5})

	is.Equal(result1, []int{0, 2})
	is.Equal(result2, []int{0})
	is.Equal(result3, []int{})
	is.Equal(result4, []int{0})
	is.Equal(result5, []int{0})
}

func TestDifference(t *testing.T) {
	is := assert.New(t)

	left1, right1 := Difference([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 6})
	is.Equal(left1, []int{1, 3, 4, 5})
	is.Equal(right1, []int{6})

	left2, right2 := Difference([]int{1, 2, 3, 4, 5}, []int{0, 6})
	is.Equal(left2, []int{1, 2, 3, 4, 5})
	is.Equal(right2, []int{0, 6})

	left3, right3 := Difference([]int{0, 1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4, 5})
	is.Equal(left3, []int{})
	is.Equal(right3, []int{})
}

func TestUnion(t *testing.T) {
	is := assert.New(t)
	result1 := Union([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10})
	result2 := Union([]int{0, 1, 2, 3, 4, 5}, []int{6, 7})
	result3 := Union([]int{0, 1, 2, 3, 4, 5}, []int{})
	result4 := Union([]int{0, 1, 2}, []int{0, 1, 2})
	result5 := Union([]int{}, []int{})
	is.Equal(result1, []int{0, 1, 2, 3, 4, 5, 10})
	is.Equal(result2, []int{0, 1, 2, 3, 4, 5, 6, 7})
	is.Equal(result3, []int{0, 1, 2, 3, 4, 5})
	is.Equal(result4, []int{0, 1, 2})
	is.Equal(result5, []int{})
}
