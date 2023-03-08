package xtest

import (
	"fmt"
	"strings"
	"testing"
)

func Test_fmt(t *testing.T) {
	fmt.Println(fmt.Sprintf("%02d", 10))
}

func checkSpiltRune(r rune) bool {
	if r > 97 {
		return true
	}
	return true
}

func Test_ToValidUTF8(t *testing.T) {
	strings.ToValidUTF8("街角\xF3魔族是最好的动漫\x80", "替换字符")

	strHaiCoder := "嗨客网(www.haicoder.net)Hello,HaiCoder,Hello,World"
	strArr := strings.FieldsFunc(strHaiCoder, checkSpiltRune)
	fmt.Println("strArr =", strArr)

	letter := []rune(strHaiCoder)
	fmt.Println("", letter)
	for i := 0; i < len(letter); i++ {
		fmt.Println("", string(letter[i]))
	}
}

func Test_Douhao(t *testing.T) {
	fmt.Println(fmt.Sprintf(`{"job_state":0,"natural_check":%v,"msg":"批量下架"}`, 120))
}
