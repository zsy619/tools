package xtest

import (
	"fmt"
	"strings"
	"testing"
)

// Test_fmt 验证 fmt.Sprintf 使用 %02d 格式化整数时，会将不足两位的数字左侧补零输出。
func Test_fmt(t *testing.T) {
	fmt.Println(fmt.Sprintf("%02d", 10))
}

// checkSpiltRune 是 strings.FieldsFunc 使用的分隔符判定函数。
// 参数 r 为当前扫描到的 rune；码点大于 97 即视为分隔符。
// 当前实现中两个分支的返回值均为 true，等价于始终返回 true。
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
