package xcache

import (
	"fmt"
	"testing"
	"time"
)

func TestSetCahce(t *testing.T) {
	SetCahce("page", 1, time.Minute*60)
	rt, ok := GetCache("page")
	if ok {
		fmt.Println(rt)
	}
}

func TestGetCache(t *testing.T) {
	rt, ok := GetCache("page")
	if ok {
		fmt.Println(rt)
	} else {
		fmt.Println("No cache")
	}
}
