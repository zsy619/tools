package xtest

import (
	"fmt"
	"testing"
)

func Test_Interface(t *testing.T) {
	options := make([]interface{}, 0)
	options = append(options, []string{"xm", "xb"})
	options = append(options, []string{"xm", "xb"})
	printOption(options)
}

func printOption(options ...interface{}) {
	for _, v := range options {
		fmt.Println(v)
	}
}
