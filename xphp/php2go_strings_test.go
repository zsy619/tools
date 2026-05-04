package xphp

import (
	"fmt"
	"testing"
)

func TestSubstr(t *testing.T) {
	str := "河南省教育网有限公司"
	fmt.Println(Substr(str, 0, len(str)-1))
}

func TestJson_decode(t *testing.T) {
	jsonobj := `{"Peter":35,"Ben":37,"Joe":43}`
	rt, _ := Json_decode(jsonobj)
	fmt.Println("", rt)
}
