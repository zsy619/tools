package xcas

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"net/url"

	"github.com/golang/glog"
)

// Options CAS 客户端的可配置选项。
type Options struct {
	URL          *url.URL     // CAS 服务的根 URL
	Store        TicketStore  // 自定义票据存储，若为 nil 则使用 MemoryStore
	Client       *http.Client // 自定义 HTTP 客户端，便于配置连接参数
	SendService  bool         // 是否在请求中附带 service 参数
	URLScheme    URLScheme    // 自定义 URL Scheme，用于改写客户端的请求地址
	Cookie       *http.Cookie // 会话 Cookie 选项，使用 Path、Domain、MaxAge、HttpOnly 与 Secure
	SessionStore SessionStore // 自定义会话存储，若为 nil 则使用内存存储
}

// Client 实现 CAS 主协议。
type Client struct {
	tickets   TicketStore
	client    *http.Client
	urlScheme URLScheme
	cookie    *http.Cookie

	sessions    SessionStore
	sendService bool

	stValidator *ServiceTicketValidator
}

// NewClient 根据传入的 Options 创建一个 Client。
// 若 options 为 nil 会触发空指针解引用，请勿传入 nil。
func NewClient(options *Options) *Client {
	if glog.V(2) {
		glog.Infof("cas: new client with options %v", options)
	}

	var tickets TicketStore
	if options.Store != nil {
		tickets = options.Store
	} else {
		tickets = &MemoryStore{}
	}

	var sessions SessionStore
	if options.SessionStore != nil {
		sessions = options.SessionStore
	} else {
		sessions = NewMemorySessionStore()
	}

	var urlScheme URLScheme
	if options.URLScheme != nil {
		urlScheme = options.URLScheme
	} else {
		urlScheme = NewDefaultURLScheme(options.URL)
	}

	var client *http.Client
	if options.Client != nil {
		client = options.Client
	} else {
		client = &http.Client{}
	}

	var cookie *http.Cookie
	if options.Cookie != nil {
		cookie = options.Cookie
	} else {
		cookie = &http.Cookie{
			MaxAge:   86400,
			HttpOnly: false,
			Secure:   false,
		}
	}

	return &Client{
		tickets:     tickets,
		client:      client,
		urlScheme:   urlScheme,
		cookie:      cookie,
		sessions:    sessions,
		sendService: options.SendService,
		stValidator: NewServiceTicketValidator(client, options.URL),
	}
}

// Handle 将一个 http.Handler 包装为带有 CAS 认证能力的处理器。
func (c *Client) Handle(h http.Handler) http.Handler {
	return &clientHandler{
		c: c,
		h: h,
	}
}

// HandleFunc 将一个处理函数包装为带有 CAS 认证能力的处理器。
func (c *Client) HandleFunc(h func(http.ResponseWriter, *http.Request)) http.Handler {
	return c.Handle(http.HandlerFunc(h))
}

// requestURL 根据 http.Request 计算出绝对 URL。
// 优先使用 X-Forwarded-Host / X-Forwarded-Proto 头判断反向代理后的真实主机与协议，
// 否则回退到 r.Host 与 TLS 状态。
func requestURL(r *http.Request) (*url.URL, error) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		return nil, err
	}

	u.Host = r.Host
	if host := r.Header.Get("X-Forwarded-Host"); host != "" {
		u.Host = host
	}

	u.Scheme = "http"
	if scheme := r.Header.Get("X-Forwarded-Proto"); scheme != "" {
		u.Scheme = scheme
	} else if r.TLS != nil {
		u.Scheme = "https"
	}

	return u, nil
}

// LoginUrlForRequest 构造当前请求对应的 CAS 登录 URL（含 service 参数）。
// 出错时返回底层错误；URL 解析失败会返回相应错误。
func (c *Client) LoginUrlForRequest(r *http.Request) (string, error) {
	u, err := c.urlScheme.Login()
	if err != nil {
		return "", err
	}

	service, err := requestURL(r)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Add("service", sanitisedURLString(service))
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// LogoutUrlForRequest 构造当前请求对应的 CAS 登出 URL；
// 当 Options.SendService 为 true 时会附带 service 参数。
func (c *Client) LogoutUrlForRequest(r *http.Request) (string, error) {
	u, err := c.urlScheme.Logout()
	if err != nil {
		return "", err
	}

	if c.sendService {
		service, err := requestURL(r)
		if err != nil {
			return "", err
		}

		q := u.Query()
		q.Add("service", sanitisedURLString(service))
		u.RawQuery = q.Encode()
	}

	return u.String(), nil
}

// ServiceValidateUrlForRequest 构造给定 ticket 与请求对应的 CAS serviceValidate URL（CAS 2.0+）。
func (c *Client) ServiceValidateUrlForRequest(ticket string, r *http.Request) (string, error) {
	service, err := requestURL(r)
	if err != nil {
		return "", err
	}
	return c.stValidator.ServiceValidateUrl(service, ticket)
}

// ValidateUrlForRequest 构造给定 ticket 与请求对应的 CAS validate URL（CAS 1.0）。
func (c *Client) ValidateUrlForRequest(ticket string, r *http.Request) (string, error) {
	service, err := requestURL(r)
	if err != nil {
		return "", err
	}
	return c.stValidator.ValidateUrl(service, ticket)
}

// RedirectToLogout 清理本地会话并以 302 重定向到 CAS 登出地址。
// 登出 URL 解析失败时会以 500 响应写出错误信息。
func (c *Client) RedirectToLogout(w http.ResponseWriter, r *http.Request) {
	u, err := c.LogoutUrlForRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if glog.V(2) {
		glog.Infof("Logging out, redirecting client to %v with status %v",
			u, http.StatusFound)
	}

	c.clearSession(w, r)
	http.Redirect(w, r, u, http.StatusFound)
}

// RedirectToLogin 以 302 重定向到 CAS 登录地址；登录 URL 解析失败会以 500 响应。
func (c *Client) RedirectToLogin(w http.ResponseWriter, r *http.Request) {
	u, err := c.LoginUrlForRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if glog.V(2) {
		glog.Infof("Redirecting client to %v with status %v", u, http.StatusFound)
	}

	http.Redirect(w, r, u, http.StatusFound)
}

// validateTicket 使用给定的 ticket 与 service 请求远程校验票据，并将成功结果写入票据存储。
// 校验或写入失败会返回相应错误。
func (c *Client) validateTicket(ticket string, service *http.Request) error {
	serviceURL, err := requestURL(service)
	if err != nil {
		return err
	}

	success, err := c.stValidator.ValidateTicket(serviceURL, ticket)
	if err != nil {
		return err
	}

	if err := c.tickets.Write(ticket, success); err != nil {
		return err
	}

	return nil
}

// getSession 查找或创建当前请求对应的会话。
//
// 若请求未携带会话 Cookie，会在响应中下发新的 Cookie；
// 若 URL 中包含 ticket 参数，则先进行票据校验。
func (c *Client) getSession(w http.ResponseWriter, r *http.Request) {
	cookie := c.getCookie(w, r)

	if s, ok := c.sessions.Get(cookie.Value); ok {
		if t, err := c.tickets.Read(s); err == nil {
			if glog.V(1) {
				glog.Infof("Re-used ticket %s for %s", s, t.User)
			}

			setAuthenticationResponse(r, t)
			return
		} else {
			if glog.V(2) {
				glog.Infof("Ticket %v not in %T: %v", s, c.tickets, err)
			}

			if glog.V(1) {
				glog.Infof("Clearing ticket %s, no longer exists in ticket store", s)
			}

			clearCookie(w, cookie)
		}
	}

	if ticket := r.URL.Query().Get("ticket"); ticket != "" {
		if err := c.validateTicket(ticket, r); err != nil {
			if glog.V(2) {
				glog.Infof("Error validating ticket: %v", err)
			}
			return // allow ServeHTTP()
		}

		c.setSession(cookie.Value, ticket)

		if t, err := c.tickets.Read(ticket); err == nil {
			if glog.V(1) {
				glog.Infof("Validated ticket %s for %s", ticket, t.User)
			}

			setAuthenticationResponse(r, t)
			return
		} else {
			if glog.V(2) {
				glog.Infof("Ticket %v not in %T: %v", ticket, c.tickets, err)
			}

			if glog.V(1) {
				glog.Infof("Clearing ticket %s, no longer exists in ticket store", ticket)
			}

			clearCookie(w, cookie)
		}
	}
}

// getCookie 查找请求中的会话 Cookie；不存在时创建新的会话 Cookie 并写入响应。
// 注意：默认未启用 HttpOnly，以便 Ajax 请求也能携带该 Cookie。
func (c *Client) getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		// NOTE: 默认未启用 HttpOnly，便于 Ajax 请求使用。
		cookie = &http.Cookie{
			Name:     sessionCookieName,
			Value:    newSessionID(),
			Path:     c.cookie.Path,
			Domain:   c.cookie.Domain,
			MaxAge:   c.cookie.MaxAge,
			HttpOnly: c.cookie.HttpOnly,
			Secure:   c.cookie.Secure,
		}

		if glog.V(2) {
			glog.Infof("Setting %v cookie with value: %v", cookie.Name, cookie.Value)
		}

		r.AddCookie(cookie) // 记录到请求，便于后续读取
		http.SetCookie(w, cookie)
	}

	return cookie
}

// newSessionID 生成一个 64 字符的不透明会话标识，用于 Cookie 的 Value。
// 使用 crypto/rand 生成随机字节并映射到字母数字字符集。
func newSessionID() string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 生成 64 字符字符串
	bytes := make([]byte, 64)
	rand.Read(bytes)

	for k, v := range bytes {
		bytes[k] = alphabet[v%byte(len(alphabet))]
	}

	return string(bytes)
}

// clearCookie 将指定 Cookie 设为过期并下发给客户端，达到清除的效果。
func clearCookie(w http.ResponseWriter, c *http.Cookie) {
	c.MaxAge = -1
	http.SetCookie(w, c)
}

// setSession 在 SessionStore 中记录会话 ID 与票据的映射关系。
func (c *Client) setSession(id string, ticket string) {
	if glog.V(2) {
		glog.Infof("Recording session, %v -> %v", id, ticket)
	}

	c.sessions.Set(id, ticket)
}

// clearSession 移除本地会话记录并下发过期 Cookie；同时尝试删除对应票据。
// 删除票据失败时仅打印错误日志，不影响 Cookie 清理。
func (c *Client) clearSession(w http.ResponseWriter, r *http.Request) {
	cookie := c.getCookie(w, r)

	if serviceTicket, ok := c.sessions.Get(cookie.Value); ok {
		if err := c.tickets.Delete(serviceTicket); err != nil {
			fmt.Printf("Failed to remove %v from %T: %v\n", cookie.Value, c.tickets, err)
			if glog.V(2) {
				glog.Errorf("Failed to remove %v from %T: %v", cookie.Value, c.tickets, err)
			}
		}

		c.deleteSession(cookie.Value)
	}

	clearCookie(w, cookie)
}

// deleteSession 从 SessionStore 中删除指定会话 ID。
func (c *Client) deleteSession(id string) {
	c.sessions.Delete(id)
}
