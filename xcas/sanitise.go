package xcas

import (
	"net/url"
)

// urlCleanParameters 需要从 URL 中移除的 CAS 协议相关查询参数。
var urlCleanParameters = []string{"gateway", "renew", "service", "ticket"}

// sanitisedURL 移除 URL 中 CAS 协议相关的查询参数（gateway/renew/service/ticket）。
// 原始 *url.URL 重新解析不应出错，错误被忽略。
func sanitisedURL(unclean *url.URL) *url.URL {
	// 解析已存在的 *url.URL 通常不会出错
	u, _ := url.Parse(unclean.String())
	q := u.Query()

	for _, param := range urlCleanParameters {
		q.Del(param)
	}

	u.RawQuery = q.Encode()
	return u
}

// sanitisedURLString 清理 URL 中的 CAS 相关参数并返回其字符串表示。
func sanitisedURLString(unclean *url.URL) string {
	return sanitisedURL(unclean).String()
}
