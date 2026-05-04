package xtest

import (
	"regexp"
	"testing"
)

func Test_regex(t *testing.T) {
	// match := "http(s?)://(.*).hnzhjypt.com/index/login"
	match := "^http(s?)://(.*).hnzhjypt.com/Index/login.html"
	// match := "http(s?)://(.*).hncaiyun.com/Index/login.html"
	// match := "http(s?)://(.*).hncaiyun.com/index/login"
	// url := "http://www.hnzhjypt.com/index/login"
	url := "https://www.hnzhjypt.com/Index/login.html"
	// url := "http://jytest.hncaiyun.com/Index/login.html"
	// url := "http://jytest.hncaiyun.com/index/login"
	if ok, _ := regexp.Match(match, []byte(url)); !ok {
		t.Error("RegexMatch failed")
	}
}
