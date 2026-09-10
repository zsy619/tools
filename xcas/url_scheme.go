package xcas

import (
	"net/url"
	"path"
)

// URLScheme 定义生成 CAS 协议各类接口 URL 的方法集合。
type URLScheme interface {
	// Login 返回登录接口 URL。
	Login() (*url.URL, error)
	// Logout 返回登出接口 URL。
	Logout() (*url.URL, error)
	// Validate 返回 CAS 1.0 校验接口 URL。
	Validate() (*url.URL, error)
	// ServiceValidate 返回 CAS 2.0+ serviceValidate 接口 URL。
	ServiceValidate() (*url.URL, error)
	// RestGrantingTicket 返回 REST 模式下申请 TGT 的接口 URL。
	RestGrantingTicket() (*url.URL, error)
	// RestServiceTicket 返回针对指定 TGT 申请 ST 的 REST 接口 URL。
	RestServiceTicket(tgt string) (*url.URL, error)
	// RestLogout 返回销毁指定 TGT 的 REST 接口 URL。
	RestLogout(tgt string) (*url.URL, error)
}

// NewDefaultURLScheme 创建一个使用 CAS 默认路径的 URLScheme。
func NewDefaultURLScheme(base *url.URL) *DefaultURLScheme {
	return &DefaultURLScheme{
		Base:                base,
		LoginPath:           "login",
		LogoutPath:          "logout",
		ValidatePath:        "validate",
		ServiceValidatePath: "serviceValidate",
		RestEndpoint:        path.Join("v1", "tickets"),
	}
}

// DefaultURLScheme 是一个可配置的 URLScheme。
// 推荐使用 NewDefaultURLScheme 构造以获得 CAS 默认路径。
type DefaultURLScheme struct {
	Base                *url.URL
	LoginPath           string
	LogoutPath          string
	ValidatePath        string
	ServiceValidatePath string
	RestEndpoint        string
}

// Login 返回 CAS 登录接口 URL。
func (scheme *DefaultURLScheme) Login() (*url.URL, error) {
	return scheme.createURL(scheme.LoginPath)
}

// Logout 返回 CAS 登出接口 URL。
func (scheme *DefaultURLScheme) Logout() (*url.URL, error) {
	return scheme.createURL(scheme.LogoutPath)
}

// Validate 返回 CAS 1.0 票据校验接口 URL。
func (scheme *DefaultURLScheme) Validate() (*url.URL, error) {
	return scheme.createURL(scheme.ValidatePath)
}

// ServiceValidate 返回 CAS 2.0+ serviceValidate 接口 URL。
func (scheme *DefaultURLScheme) ServiceValidate() (*url.URL, error) {
	return scheme.createURL(scheme.ServiceValidatePath)
}

// RestGrantingTicket 返回通过 REST 申请 TGT 的接口 URL。
func (scheme *DefaultURLScheme) RestGrantingTicket() (*url.URL, error) {
	return scheme.createURL(scheme.RestEndpoint)
}

// RestServiceTicket 返回针对指定 TGT 申请 ST 的 REST 接口 URL。
func (scheme *DefaultURLScheme) RestServiceTicket(tgt string) (*url.URL, error) {
	return scheme.createURL(path.Join(scheme.RestEndpoint, tgt))
}

// RestLogout 返回通过 REST 销毁指定 TGT 的接口 URL。
func (scheme *DefaultURLScheme) RestLogout(tgt string) (*url.URL, error) {
	return scheme.createURL(path.Join(scheme.RestEndpoint, tgt))
}

// createURL 基于 Base 与指定 urlPath 生成新的 *url.URL。
func (scheme *DefaultURLScheme) createURL(urlPath string) (*url.URL, error) {
	return scheme.Base.Parse(path.Join(scheme.Base.Path, urlPath))
}
