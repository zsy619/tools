package xcas

import (
	"fmt"
	"net/http"

	"github.com/golang/glog"
)

const (
	// sessionCookieName 客户端会话 Cookie 名称。
	sessionCookieName = "_cas_session"
)

// clientHandler 处理 CAS 协议相关的 HTTP 请求。
type clientHandler struct {
	c *Client
	h http.Handler
}

// ServeHTTP 处理 HTTP 请求，识别并处理 CAS 请求后将请求转发给子 http.Handler。
// 单点登出（SLO）请求会被直接处理，不会进入下游处理器。
func (ch *clientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if glog.V(2) {
		glog.Infof("cas: handling %v request for %v", r.Method, r.URL)
	}

	setClient(r, ch.c)

	if isSingleLogoutRequest(r) {
		ch.performSingleLogout(w, r)
		return
	}

	ch.c.getSession(w, r)
	ch.h.ServeHTTP(w, r)
}

// isSingleLogoutRequest 判断当前请求是否为 CAS 单点登出（SLO）请求。
//
// SLO 请求的特征：HTTP POST、Content-Type 为 application/x-www-form-urlencoded，
// 且表单中包含 logoutRequest 参数。
func isSingleLogoutRequest(r *http.Request) bool {
	if r.Method != "POST" {
		return false
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/x-www-form-urlencoded" {
		return false
	}

	if v := r.FormValue("logoutRequest"); v == "" {
		return false
	}

	return true
}

// performSingleLogout 处理单点登出请求：解析 logoutRequest XML，删除对应票据与会话。
// 解析失败或票据删除失败会以 500 响应写出错误信息。
func (ch *clientHandler) performSingleLogout(w http.ResponseWriter, r *http.Request) {
	rawXML := r.FormValue("logoutRequest")
	logoutRequest, err := parseLogoutRequest([]byte(rawXML))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ch.c.tickets.Delete(logoutRequest.SessionIndex); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ch.c.deleteSession(logoutRequest.SessionIndex)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}
