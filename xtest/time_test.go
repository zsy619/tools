package xtest

import (
	"fmt"
	"testing"
	"time"
)

func Test_Time(t *testing.T) {
	format := "2006-01-02 15:04:05"
	now := time.Now()
	// now, _ := time.Parse(format, time.Now().Format(format))
	a, _ := time.Parse(format, "2041-03-10 11:00:00")
	b, _ := time.Parse(format, "2015-03-10 16:00:00")

	fmt.Println("now:", now.Format(format), "\na:", a.Format(format), "\nb:", b.Format(format))
	fmt.Println("---    a > now  >  b   -----------")
	fmt.Println("now  a   After: ", now.After(a))
	fmt.Println("now  a   Before:", now.Before(a))
	fmt.Println("now  b   After:", now.After(b))
	fmt.Println("now  b   Before:", now.Before(b))
	fmt.Println("a  b   After:", a.After(b))
	fmt.Println("a  b   Before:", a.Before(b))
	fmt.Println("a  a   Before:", a.Before(a))
	fmt.Println("a  a   After:", a.After(a))
	fmt.Println(now.Format(format), a.Format(format), b.Format(format))
	fmt.Println(now.Unix(), a.Unix(), b.Unix())
}

func Test_Now_Format(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 00:00:00"))
	fmt.Println(now.Format("2006-01-02 23:59:59"))
}
