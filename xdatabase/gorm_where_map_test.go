package xdatabase

import (
	"testing"
)

func TestWhereBuild(t *testing.T) {
	cond, vals, err := WhereBuild(map[string]interface{}{
		"name":   "jinzhu",
		"age in": []int{20, 19, 18},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cond)
	t.Log(vals)
}

func TestWhereBuildExtension(t *testing.T) {
	wh := WhereBuildExtension(map[string]interface{}{
		"name": "jin'zhu",
		"c":    []string{"aa'1", "2"},
		"d":    []int{1},
		"or":   "a=b",
		"abc":  "like '%444%'",
	})
	t.Log(wh)
}
