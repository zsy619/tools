package xstring

import (
	"fmt"
	"regexp"
)

// 密码强度等级，D为最低
/*
密码强度	说明	示例
S	密码中必须存在特殊字符、大小写字母和数字	Csdn#2020
A   对特殊字符、大写字母、小写字母和数字至少存在3种 csdn#2020
B	对特殊字符、大写字母、小写字母和数字至少存在2种	csdn2020
C	对特殊字符、大写字母、小写字母和数字至少存在1种	csdn
D	不存在特殊字符、大小写字母和数字。	/、\
*/
const (
	PasswordLevelD = iota
	PasswordLevelC
	PasswordLevelB
	PasswordLevelA
	PasswordLevelS
)

/*
 *  minLength: 指定密码的最小长度
 *  maxLength：指定密码的最大长度
 *  minLevel：指定密码最低要求的强度等级
 *  pwd：明文密码
 */
func PasswordCheck(minLength, maxLength, minLevel int, pwd string) error {
	if len(pwd) < minLength {
		return fmt.Errorf("BAD PASSWORD: The password is shorter than %d characters", minLength)
	}
	if len(pwd) > maxLength {
		return fmt.Errorf("BAD PASSWORD: The password is logner than %d characters", maxLength)
	}

	var level int = PasswordLevelD
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}
	fmt.Println("", level)
	if level < minLevel {
		return fmt.Errorf("The password does not satisfy the current policy requirements. ")
	}
	return nil
}
