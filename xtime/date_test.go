package xtime

import (
	"fmt"
	"testing"
	"time"
)

func TestGetCurrentDate(t *testing.T) {
	d := GetCurrentDate("2006-01-02 15:04:05")
	fmt.Println(d)
	dd := time.Now()
	fmt.Println(FormatDateExt(dd, YYYY_MM_DD_HH_MM_SS_SSS))
	name := `老"三'家`
	postData := fmt.Sprintf("{\"name\":\"%s\"}", name)
	fmt.Println(postData)
}
