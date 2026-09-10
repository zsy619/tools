package xstring

import (
	"fmt"
	"regexp"
)

// 密码强度等级常量。等级从低到高依次为 D、C、B、A、S，值由 iota 自动递增（0..4）。
/*
密码强度等级与判定规则说明：

  等级  说明                                                                  示例
  S     密码中必须同时存在特殊字符、大写字母、小写字母和数字（共 4 种）      Csdn#2020
  A     上述 4 类字符中至少存在 3 种                                         csdn#2020
  B     上述 4 类字符中至少存在 2 种                                         csdn2020
  C     上述 4 类字符中至少存在 1 种                                         csdn
  D     不存在上述任一字符（即不包含数字、大写字母、小写字母、特殊字符）    /、\

特殊字符的判定范围为：~ ! @ # $ % ^ & * ? _ -
*/
const (
	// PasswordLevelD 密码强度等级 D：不含数字、大小写字母或特殊字符。
	PasswordLevelD = iota
	// PasswordLevelC 密码强度等级 C：至少包含数字、大写字母、小写字母、特殊字符中的 1 种。
	PasswordLevelC
	// PasswordLevelB 密码强度等级 B：至少包含上述 4 类字符中的 2 种。
	PasswordLevelB
	// PasswordLevelA 密码强度等级 A：至少包含上述 4 类字符中的 3 种。
	PasswordLevelA
	// PasswordLevelS 密码强度等级 S：必须同时包含数字、大写字母、小写字母和特殊字符。
	PasswordLevelS
)

// PasswordCheck 检查密码是否符合指定的长度区间和强度等级要求。
//
// 参数：
//   - minLength int：密码允许的最小长度（包含）。当 len(pwd) < minLength 时返回错误。
//   - maxLength int：密码允许的最大长度（包含）。当 len(pwd) > maxLength 时返回错误。
//   - minLevel int：密码必须达到的最低强度等级，取值为 PasswordLevelD..PasswordLevelS 之一。
//   - pwd string：待检查的密码字符串。
//
// 返回值：
//   - error：当密码长度不在 [minLength, maxLength] 区间内、或计算出的强度等级 < minLevel 时，
//     返回带说明的 fmt.Errorf 错误；通过检查时返回 nil。
//
// 副作用：当 len(pwd) 处于合法范围时，无论最终是否通过，函数都会向标准输出打印一行空格加 level 的值，
// 用于在调用处直观观察当前算得的强度等级。
//
// 强度等级算法：从初始值 PasswordLevelD（0）开始，分别用 [0-9]+、[a-z]+、[A-Z]+、[~!@#$%^&*?_-]+
// 四个正则匹配 pwd，每命中一个字符种类 level 加 1，最终 level ∈ [0, 4]，对应 PasswordLevelD..PasswordLevelS。
//
// 注意事项：特殊字符集合固定为 ~!@#$%^&*?_-，与 VerifyPassword 中使用的 .@$!%*#_~?&^ 不一致，
// 如需统一两者的“特殊字符”定义请在使用方自行协调。
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

// VerifyPassword 校验密码是否合法：长度需处于 [minLength, maxLength] 之间，
// 且必须包含数字、大写字母、小写字母、特殊字符（.@$!%*#_~?&^）四类中至少 3 类的组合。
//
// 参数：
//   - minLength int：密码最小长度（包含）。
//   - maxLength int：密码最大长度（包含）。注意：函数内部还会用正则 `^[a-zA-Z0-9.@$!%*#_~?&^]{8,16}$` 二次校验长度，
//     若 minLength<8 或 maxLength>16 即使通过第一段区间检查也会被该正则拒绝。
//   - pwd string：待验证的密码字符串。允许为空字符串（此时会被首段长度判定直接拒绝）。
//
// 返回值：
//   - bool：通过所有校验返回 true，否则返回 false；任何阶段失败都不会 panic。
//
// 实现细节：
//  1. 先做长度区间判定；
//  2. 再用正则 `^[a-zA-Z0-9.@$!%*#_~?&^]{8,16}$` 过滤，仅允许指定字符集中的字符且长度 8-16，
//     正则编译错误时也返回 false；
//  3. 最后用 [0-9]+、[a-z]+、[A-Z]+、[.@$!%*#_~?&^]+ 四组正则统计命中的字符种类数量，
//     至少 3 种才视为合法。
//
// 该函数无副作用（不会向标准输出打印）。
func VerifyPassword(minLength, maxLength int, pwd string) bool {
	if len(pwd) < minLength || len(pwd) > maxLength {
		return false
	}
	// 仅允许 [a-zA-Z0-9.@$!%*#_~?&^] 中的字符且长度在 8-16 之间；
	// 过滤掉这四类字符以外的密码串,直接判断不合法。
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
