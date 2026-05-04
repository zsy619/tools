package xerror

import "fmt"

func PrintError(input func() error) {
	if err := input(); err != nil {
		fmt.Println(err.Error())
	}
}
