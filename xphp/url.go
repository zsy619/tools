package xphp

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// GetHeaders 获取服务器对 HTTP 请求响应时发送的全部头部信息。
func GetHeaders(url string) (http.Header, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	resp.Header.Set("status", resp.Status)

	return resp.Header, nil
}

// GetMetaTags 从指定 URL 的 HTTP 内容中提取所有 meta 标签的 content 属性，并返回映射。
func GetMetaTags(url string) (map[string]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ParseDocument(resp.Body), nil
}

// ParseDocument 解析 HTML 文档，提取其中的 meta 数据（包括页面标题 title）。
func ParseDocument(doc io.Reader) map[string]string {
	data := make(map[string]string)
	z := html.NewTokenizer(doc)

	var titleFound bool
	for {
		tt := z.Next()
		t := z.Token()
		switch tt {
		case html.ErrorToken:
			return data
		case html.EndTagToken:
			if t.Data == "head" {
				return data
			}
		case html.StartTagToken, html.SelfClosingTagToken:
			if t.Data == "title" {
				titleFound = true
			}
			if t.Data == "meta" {
				var property, content string
				for _, attr := range t.Attr {
					switch attr.Key {
					case "property", "name":
						property = strings.ToLower(attr.Val)
					case "content":
						content = attr.Val
					}
				}

				if property != "" {
					data[strings.TrimSpace(property)] = content
				}
			}
		case html.TextToken:
			if titleFound {
				data["title"] = t.Data
				titleFound = false
			}
		}
	}
}

// RawURLDecode 解码 URL 编码的字符串。
//
// RawURLDecode 与 URLDecode 基本相同，区别在于它不会把 '+' 解码为 ' '（空格）。
func RawURLDecode(str string) string {
	res, _ := url.PathUnescape(str)
	return res
}

// RawURLEncode 按照 RFC 3986 对字符串进行 URL 编码。
//
// RawURLEncode 与 URLEncode 基本相同，区别在于它不会把空格转义为 +。
func RawURLEncode(str string) string {
	return url.PathEscape(str)
}
