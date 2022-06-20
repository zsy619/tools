package xregexp

import "regexp"

var (
	emailRegexp = regexp.MustCompile(
		"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
	)

	mobileRegexp = regexp.MustCompile(
		`^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4[579]\d{2})\d{6}$`,
	)
)

// IsEmail 是否为邮箱
func IsEmail(v string) bool {
	return emailRegexp.MatchString(v)
}

// IsMobile 是否为手机
func IsMobile(v string) bool {
	return mobileRegexp.MatchString(v)
}
