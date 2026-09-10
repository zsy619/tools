package xcas

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/golang/glog"
)

// 参考协议：https://apereo.github.io/cas/4.2.x/protocol/REST-Protocol.html

// TicketGrantingTicket 表示用户的 SSO 会话，又称 TGT。
type TicketGrantingTicket string

// ServiceTicket 表示 CAS 服务器为特定用户向特定应用签发的访问凭证，又称 ST。
type ServiceTicket string

// RestOptions 为 RestClient 提供配置项。
type RestOptions struct {
	CasURL     *url.URL     // CAS 服务根 URL
	ServiceURL *url.URL     // 当前应用自身的服务 URL
	Client     *http.Client // 自定义 HTTP 客户端
	URLScheme  URLScheme    // 自定义 URL Scheme
}

// RestClient 基于 CAS 提供的 REST 协议与 CAS 服务器交互。
type RestClient struct {
	urlScheme   URLScheme
	serviceURL  *url.URL
	client      *http.Client
	stValidator *ServiceTicketValidator
}

// NewRestClient 根据 RestOptions 创建一个 CAS REST 协议客户端。
// options.CasURL 与 options.ServiceURL 不能为 nil。
func NewRestClient(options *RestOptions) *RestClient {
	if glog.V(2) {
		glog.Infof("cas: new rest client with options %v", options)
	}

	var client *http.Client
	if options.Client != nil {
		client = options.Client
	} else {
		client = &http.Client{}
	}

	var urlScheme URLScheme
	if options.URLScheme != nil {
		urlScheme = options.URLScheme
	} else {
		urlScheme = NewDefaultURLScheme(options.CasURL)
	}

	return &RestClient{
		urlScheme:   urlScheme,
		serviceURL:  options.ServiceURL,
		client:      client,
		stValidator: NewServiceTicketValidator(client, options.CasURL),
	}
}

// Handle 将 http.Handler 包装为带 CAS REST 认证能力的处理器。
func (c *RestClient) Handle(h http.Handler) http.Handler {
	return &restClientHandler{
		c: c,
		h: h,
	}
}

// HandleFunc 将一个函数包装为带 CAS REST 认证能力的处理器。
func (c *RestClient) HandleFunc(h func(http.ResponseWriter, *http.Request)) http.Handler {
	return c.Handle(http.HandlerFunc(h))
}

// RequestGrantingTicket 使用用户名密码向 CAS 请求一个新的 TGT；
// 期望服务端返回 201 并在 Location 头中给出 TGT 路径。
// 状态码非 201 或缺少 Location 头时返回错误。
func (c *RestClient) RequestGrantingTicket(username string, password string) (TicketGrantingTicket, error) {
	// 请求：
	// POST /cas/v1/tickets HTTP/1.0
	// username=battags&password=password&additionalParam1=paramvalue

	endpoint, err := c.urlScheme.RestGrantingTicket()
	if err != nil {
		return "", err
	}

	values := url.Values{}
	values.Set("username", username)
	values.Set("password", password)

	resp, err := c.client.PostForm(endpoint.String(), values)
	if err != nil {
		return "", err
	}

	// 响应：
	// 201 Created
	// Location: http://www.whatever.com/cas/v1/tickets/{TGT id}

	if resp.StatusCode != 201 {
		return "", fmt.Errorf("ticket endoint returned status code %v", resp.StatusCode)
	}

	tgt := path.Base(resp.Header.Get("Location"))
	if tgt == "" {
		return "", fmt.Errorf("does not return a valid location header")
	}

	return TicketGrantingTicket(tgt), nil
}

// RequestServiceTicket 使用 TGT 为当前配置的 service URL 申请一个 Service Ticket；
// 期望返回 200 且响应体即为 ST。状态码非 200 或读取响应体失败时返回错误。
func (c *RestClient) RequestServiceTicket(tgt TicketGrantingTicket) (ServiceTicket, error) {
	// 请求：
	// POST /cas/v1/tickets/{TGT id} HTTP/1.0
	// service={form encoded parameter for the service url}
	endpoint, err := c.urlScheme.RestServiceTicket(string(tgt))
	if err != nil {
		return "", err
	}

	values := url.Values{}
	values.Set("service", c.serviceURL.String())

	resp, err := c.client.PostForm(endpoint.String(), values)
	if err != nil {
		return "", err
	}

	// 响应：
	// 200 OK
	// ST-1-FFDFHDSJKHSDFJKSDHFJKRUEYREWUIFSD2132

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("service ticket endoint returned status code %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return ServiceTicket(data), nil
}

// ValidateServiceTicket 校验给定的 ServiceTicket 并返回 AuthenticationResponse。
// 校验失败或网络错误时会透传底层错误。
func (c *RestClient) ValidateServiceTicket(st ServiceTicket) (*AuthenticationResponse, error) {
	return c.stValidator.ValidateTicket(c.serviceURL, string(st))
}

// Logout 通过 DELETE 请求销毁指定的 TGT，使 CAS 服务端的会话失效；
// 期望状态码为 200 或 204，否则返回错误。
func (c *RestClient) Logout(tgt TicketGrantingTicket) error {
	// DELETE /cas/v1/tickets/TGT-fdsjfsdfjkalfewrihfdhfaie HTTP/1.0
	endpoint, err := c.urlScheme.RestLogout(string(tgt))
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", endpoint.String(), nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		return fmt.Errorf("could not destroy granting ticket %v, server returned %v", tgt, resp.StatusCode)
	}

	return nil
}
