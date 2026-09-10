package xcas

import (
	"net/http"

	"github.com/golang/glog"
)

// restClientHandler 通过 HTTP Basic 认证处理 CAS REST 协议的请求。
type restClientHandler struct {
	c *RestClient
	h http.Handler
}

// ServeHTTP 处理 HTTP 请求：先做 HTTP Basic 认证，再通过 CAS REST API 完成登录，
// 最后将请求转发给下游 http.Handler。
// 任何认证失败都会返回 401 并附带 WWW-Authenticate 响应头。
func (ch *restClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if glog.V(2) {
		glog.Infof("cas: handling %v request for %v", r.Method, r.URL)
	}

	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="CAS Protected Area"`)
		w.WriteHeader(401)
		return
	}

	// TODO: 可加入短期缓存，避免每次请求都访问 CAS 服务端。
	// 缓存 key 可以是 Authorization 头，value 为 authenticationResponse。

	success, err := ch.authenticate(username, password)
	if err != nil {
		if glog.V(1) {
			glog.Infof("cas: rest authentication failed %v", err)
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="CAS Protected Area"`)
		w.WriteHeader(401)
		return
	}

	setAuthenticationResponse(r, success)
	ch.h.ServeHTTP(w, r)
}

// authenticate 通过 TGT -> ST -> 校验 ST 的流程完成 CAS REST 认证。
// 任意步骤失败都会返回相应错误；成功时返回 AuthenticationResponse。
func (ch *restClientHandler) authenticate(username string, password string) (*AuthenticationResponse, error) {
	tgt, err := ch.c.RequestGrantingTicket(username, password)
	if err != nil {
		return nil, err
	}

	st, err := ch.c.RequestServiceTicket(tgt)
	if err != nil {
		return nil, err
	}

	return ch.c.ValidateServiceTicket(st)
}
