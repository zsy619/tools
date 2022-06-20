package xphp

import (
	"fmt"
	"regexp"
)

func PregMatch(expr string, s string) (result []string, err error) {
	r, err := regexp.Compile(expr)
	for index, match := range r.FindStringSubmatch(s) {
		fmt.Printf("[%d] %s\n", index, match)
		result = append(result, match)
	}
	return
}
