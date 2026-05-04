package xtime

import (
	"fmt"
	"testing"
	"time"
)

func TestLessThan(t *testing.T) {
	time1 := "2015-03-20 08:50:29"
	time2 := "2015-03-21 09:04:25"
	// 先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	if err != nil {
		t.Fatal(err)
	}
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && LessThan(t1, t2) { // t1 < t2
		// 处理逻辑
		fmt.Println("true")
	}
}

func TestEqual(t *testing.T) {
	time1 := "2015-03-20 08:50:29"
	time2 := "2015-03-20 08:50:29"
	// 先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	if err != nil {
		t.Fatal(err)
	}
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && Equal(t2, t1) { // t1 == t2
		// 处理逻辑
		fmt.Println("true")
	}
}

func TestNotEqual(t *testing.T) {
	time1 := "2015-03-20 08:50:29"
	time2 := "2015-03-21 09:04:25"
	// 先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	if err != nil {
		t.Fatal(err)
	}
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && NotEqual(t1, t2) { // t1 != t2
		// 处理逻辑
		fmt.Println("true")
	}
}

func TestGreaterThan(t *testing.T) {
	time1 := "2015-03-20 08:50:29"
	time2 := "2015-03-21 09:04:25"
	// 先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	if err != nil {
		t.Fatal(err)
	}
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && GreaterThan(t2, t1) { //  t2 > t1
		// 处理逻辑
		fmt.Println("true")
	}
}

func TestGreaterEqual(t *testing.T) {
	time1 := "2015-03-20 08:50:29"
	time2 := "2015-03-21 09:04:25"
	// 先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	if err != nil {
		t.Fatal(err)
	}
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && GreaterEqual(t2, t1) { //  t2 >= t1
		// 处理逻辑
		fmt.Println("true")
	}
}
