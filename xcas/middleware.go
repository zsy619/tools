package xcas

import (
	"net/http"

	"github.com/golang/glog"
)

// Handler 返回一个标准 http.HandlerFunc：检查当前请求是否已通过 CAS 认证。
// 未认证则重定向到 CAS 登录页；路径为 "/logout" 时重定向到登出页；
// 其余情况调用 h.ServeHTTP 处理请求。
func (c *Client) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if glog.V(2) {
			glog.Infof("cas: handling %v request for %v", r.Method, r.URL)
		}

		setClient(r, c)

		if !IsAuthenticated(r) {
			RedirectToLogin(w, r)
			return
		}

		if r.URL.Path == "/logout" {
			RedirectToLogout(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}
