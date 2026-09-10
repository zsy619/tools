// Package xregexp 提供常用的正则校验函数。
//
// 该包在标准库 regexp 之上预编译了几个高频使用的正则表达式，
// 避免调用方每次自行 MustCompile，提高运行时性能。
package xregexp

import "regexp"

var (
	// emailRegexp 匹配符合 RFC 5321 / RFC 5322 简化语法的邮箱地址。
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	// mobileRegexp 匹配中国大陆 11 位手机号（含可选 +86 / 86 前缀）。
	mobileRegexp = regexp.MustCompile(`^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4[579]\d{2})\d{6}$`)
)

// IsEmail 判断字符串是否为合法邮箱地址。
//
// 实现依据为 RFC 5322 的简化语法：本地部分支持常见的特殊字符，
// 域名部分按 RFC 1035 的标签规则匹配。
//
// 注意：返回 true 仅表示格式上合规，无法验证该邮箱是否真实可达。
func IsEmail(v string) bool {
	return emailRegexp.MatchString(v)
}

// IsMobile 判断字符串是否为中国大陆 11 位手机号。
//
// 支持的可选前缀包括 +86、86；号段覆盖三大运营商（含虚拟运营商）
// 历史上分配过的号段，例如 13x、14[579]、16[567]、17[01356789]、
// 19[189] 等。
//
// 注意：返回 true 仅表示号段格式合规，不能验证该号码是否已启用。
func IsMobile(v string) bool {
	return mobileRegexp.MatchString(v)
}
