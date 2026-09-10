package xcas

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/golang/glog"
)

// NewServiceTicketValidator 创建一个新的 ServiceTicketValidator。
func NewServiceTicketValidator(client *http.Client, casURL *url.URL) *ServiceTicketValidator {
	return &ServiceTicketValidator{
		client: client,
		casURL: casURL,
	}
}

// ServiceTicketValidator 负责校验 service ticket。
type ServiceTicketValidator struct {
	client *http.Client
	casURL *url.URL
}

// ValidateTicket 校验给定 service 的 service ticket。
// 默认使用 CAS 2.0+ 的 serviceValidate 接口；若服务端返回 404，则回退到 CAS 1.0 的 validate 接口。
// 任意网络或解析失败会返回错误；CAS 1.0 接口下若响应为 "no\n\n" 则返回 (nil, nil) 表示未登录。
func (validator *ServiceTicketValidator) ValidateTicket(serviceURL *url.URL, ticket string) (*AuthenticationResponse, error) {
	if glog.V(2) {
		glog.Infof("Validating ticket %v for service %v", ticket, serviceURL)
	}

	u, err := validator.ServiceValidateUrl(serviceURL, ticket)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("User-Agent", "Golang CAS client gopkg.in/cas")

	if glog.V(2) {
		glog.Infof("Attempting ticket validation with %v", r.URL)
	}

	resp, err := validator.client.Do(r)
	if err != nil {
		return nil, err
	}

	if glog.V(2) {
		glog.Infof("Request %v %v returned %v",
			r.Method, r.URL,
			resp.Status)
	}

	if resp.StatusCode == http.StatusNotFound {
		return validator.validateTicketCas1(serviceURL, ticket)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cas: validate ticket: %v", string(body))
	}

	if glog.V(2) {
		glog.Infof("Received authentication response\n%v", string(body))
	}

	success, err := ParseServiceResponse(body)
	if err != nil {
		return nil, err
	}

	if glog.V(2) {
		glog.Infof("Parsed ServiceResponse: %#v", success)
	}

	return success, nil
}

// ServiceValidateUrl 构造 CAS 2.0+ serviceValidate 接口的完整 URL。
// TODO: 仅因 Client.ServiceValidateUrl 调用而对外暴露。
func (validator *ServiceTicketValidator) ServiceValidateUrl(serviceURL *url.URL, ticket string) (string, error) {
	u, err := validator.casURL.Parse(path.Join(validator.casURL.Path, "serviceValidate"))
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Add("service", sanitisedURLString(serviceURL))
	q.Add("ticket", ticket)
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// validateTicketCas1 使用 CAS 1.0 协议校验 service ticket。
// 响应体格式为 "yes\n<username>\n"，未登录时为 "no\n\n"。
// 当响应为未登录时返回 (nil, nil)；其它状态码或读取错误返回相应错误。
func (validator *ServiceTicketValidator) validateTicketCas1(serviceURL *url.URL, ticket string) (*AuthenticationResponse, error) {
	u, err := validator.ValidateUrl(serviceURL, ticket)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Add("User-Agent", "Golang CAS client gopkg.in/cas")

	if glog.V(2) {
		glog.Infof("Attempting ticket validation with %v", r.URL)
	}

	resp, err := validator.client.Do(r)
	if err != nil {
		return nil, err
	}

	if glog.V(2) {
		glog.Infof("Request %v %v returned %v",
			r.Method, r.URL,
			resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil, err
	}

	body := string(data)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cas: validate ticket: %v", body)
	}

	if glog.V(2) {
		glog.Infof("Received authentication response\n%v", body)
	}

	if body == "no\n\n" {
		return nil, nil // 未登录
	}

	success := &AuthenticationResponse{
		User: body[4 : len(body)-1],
	}

	if glog.V(2) {
		glog.Infof("Parsed ServiceResponse: %#v", success)
	}

	return success, nil
}

// ValidateUrl 构造 CAS 1.0 validate 接口的完整 URL。
// TODO: 仅因 Client.ValidateUrl 调用而对外暴露。
func (validator *ServiceTicketValidator) ValidateUrl(serviceURL *url.URL, ticket string) (string, error) {
	u, err := validator.casURL.Parse(path.Join(validator.casURL.Path, "validate"))
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Add("service", sanitisedURLString(serviceURL))
	q.Add("ticket", ticket)
	u.RawQuery = q.Encode()

	return u.String(), nil
}
