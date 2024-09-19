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

// PasswordCheck 检查密码是否符合要求
//
// 参数：
// minLength int - 密码最小长度
// maxLength int - 密码最大长度
// minLevel int - 密码最低安全等级
// pwd string - 待检查的密码
//
// 返回值：
// error - 如果密码不符合要求，返回错误信息；否则返回nil
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
		return fmt.Errorf("the password does not satisfy the current policy requirements. ")
	}
	return nil
}

// VerifyPassword 函数用于验证密码是否合法: 必须包含数字、大写字母、小写字母、特殊字符(如.@$!%*#_~?&^)至少3种的组合且长度在8-16之间
// minLength：密码最小长度
// maxLength：密码最大长度
// pwd：待验证的密码
// 返回值：如果密码合法返回true，否则返回false
func VerifyPassword(minLength, maxLength int, pwd string) bool {
	if len(pwd) < minLength || len(pwd) > maxLength {
		return false
	}
	// 过滤掉这四类字符以外的密码串,直接判断不合法
	re, err := regexp.Compile(`^[a-zA-Z0-9.@$!%*#_~?&^]{8,16}$`)
	if err != nil {
		return false
	}
	match := re.MatchString(pwd)
	if !match {
		return false
	}

	level := 0
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[.@$!%*#_~?&^]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}
	return level >= 3
}
