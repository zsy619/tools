package xcas

import (
	"context"
	"net/http"
	"time"
)

// key 是为 context.Value 定义的私有类型，避免与其他包的 key 冲突。
type key int

const (
	// ClientKey 在 http.Request Context 中存储 *Client 使用的 key。
	ClientKey key = iota
	// AuthenticationResponseKey 在 http.Request Context 中存储 *AuthenticationResponse 使用的 key。
	AuthenticationResponseKey
)

// setClient 将 Client 关联到 http.Request 的 Context 中。
func setClient(r *http.Request, c *Client) {
	ctx := context.WithValue(r.Context(), ClientKey, c)
	r2 := r.WithContext(ctx)
	*r = *r2
}

// getClient 从 http.Request Context 中取出已关联的 Client；不存在时返回 nil。
func getClient(r *http.Request) *Client {
	if c := r.Context().Value(ClientKey); c != nil {
		return c.(*Client)
	}

	return nil // 显式返回 nil，便于调用方统一处理
}

// RedirectToLogin 供受 CAS 保护的处理器调用，将请求重定向到 CAS 登录页。
// 若 Context 中未关联 Client，则返回 500 错误。
func RedirectToLogin(w http.ResponseWriter, r *http.Request) {
	c := getClient(r)
	if c == nil {
		err := "cas: redirect to cas failed as no client associated with request"
		http.Error(w, err, http.StatusInternalServerError)
		return
	}

	c.RedirectToLogin(w, r)
}

// RedirectToLogout 供受 CAS 保护的处理器调用，将请求重定向到 CAS 登出页。
// 若 Context 中未关联 Client，则返回 500 错误。
func RedirectToLogout(w http.ResponseWriter, r *http.Request) {
	c := getClient(r)
	if c == nil {
		err := "cas: redirect to cas failed as no client associated with request"
		http.Error(w, err, http.StatusInternalServerError)
		return
	}

	c.RedirectToLogout(w, r)
}

// setAuthenticationResponse 将 AuthenticationResponse 关联到 http.Request 的 Context 中。
func setAuthenticationResponse(r *http.Request, a *AuthenticationResponse) {
	ctx := context.WithValue(r.Context(), AuthenticationResponseKey, a)
	r2 := r.WithContext(ctx)
	*r = *r2
}

// getAuthenticationResponse 从 http.Request Context 中取出已关联的 AuthenticationResponse；不存在时返回 nil。
func getAuthenticationResponse(r *http.Request) *AuthenticationResponse {
	if a := r.Context().Value(AuthenticationResponseKey); a != nil {
		return a.(*AuthenticationResponse)
	}

	return nil // 显式返回 nil，便于调用方统一处理
}

// IsAuthenticated 判断当前请求是否已经通过 CAS 认证。
func IsAuthenticated(r *http.Request) bool {
	if a := getAuthenticationResponse(r); a != nil {
		return true
	}

	return false
}

// Username 返回当前请求中已认证用户的用户名；未认证时返回空字符串。
func Username(r *http.Request) string {
	if a := getAuthenticationResponse(r); a != nil {
		return a.User
	}

	return ""
}

// Attributes 返回当前请求中已认证用户的属性；未认证时返回 nil。
func Attributes(r *http.Request) UserAttributes {
	if a := getAuthenticationResponse(r); a != nil {
		return a.Attributes
	}

	return nil
}

// AuthenticationDate 返回认证发生的时间。
//
// 若 CAS 服务端响应中未包含 AuthenticationDate（例如 CAS 2.0），
// 则可能返回 time.IsZero。
func AuthenticationDate(r *http.Request) time.Time {
	var t time.Time
	if a := getAuthenticationResponse(r); a != nil {
		t = a.AuthenticationDate
	}

	return t
}

// IsNewLogin 表示当前服务票据是否由一次全新的认证流程签发。
//
// 若 CAS 服务端响应中未包含 isNewLogin 字段（例如 CAS 2.0），
// 可能错误地返回 false。
func IsNewLogin(r *http.Request) bool {
	if a := getAuthenticationResponse(r); a != nil {
		return a.IsNewLogin
	}

	return false
}

// IsRememberedLogin 表示当前服务票据是否由长期认证票据（如 Remember-Me）签发。
//
// 若 CAS 服务端响应中未包含 Remembered Login 信息（例如 CAS 2.0），
// 可能错误地返回 false。
func IsRememberedLogin(r *http.Request) bool {
	if a := getAuthenticationResponse(r); a != nil {
		return a.IsRememberedLogin
	}

	return false
}

// MemberOf 返回当前已认证用户所属的用户组列表；未认证时返回 nil。
func MemberOf(r *http.Request) []string {
	if a := getAuthenticationResponse(r); a != nil {
		return a.MemberOf
	}

	return nil
}
