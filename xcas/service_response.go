package xcas

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

// CAS AuthenticationFailure 响应中使用的错误码常量。
const (
	// INVALID_REQUEST 请求格式不合法。
	INVALID_REQUEST = "INVALID_REQUEST"
	// INVALID_TICKET_SPEC 票据规格不合法。
	INVALID_TICKET_SPEC = "INVALID_TICKET_SPEC"
	// UNAUTHORIZED_SERVICE 未被授权使用该 service。
	UNAUTHORIZED_SERVICE = "UNAUTHORIZED_SERVICE"
	// UNAUTHORIZED_SERVICE_PROXY 未被授权代理该 service。
	UNAUTHORIZED_SERVICE_PROXY = "UNAUTHORIZED_SERVICE_PROXY"
	// INVALID_PROXY_CALLBACK 代理回调地址不合法。
	INVALID_PROXY_CALLBACK = "INVALID_PROXY_CALLBACK"
	// INVALID_TICKET 票据无效。
	INVALID_TICKET = "INVALID_TICKET"
	// INVALID_SERVICE service 无效。
	INVALID_SERVICE = "INVALID_SERVICE"
	// INTERNAL_ERROR CAS 内部错误。
	INTERNAL_ERROR = "INTERNAL_ERROR"
)

// AuthenticationError 表示一次 CAS AuthenticationFailure 响应。
type AuthenticationError struct {
	Code    string
	Message string
}

// AuthenticationError 作为实现 error 的标识方法，便于类型断言。
func (e AuthenticationError) AuthenticationError() bool {
	return true
}

// Error 返回 "Code: Message" 形式的字符串。
func (e AuthenticationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// AuthenticationResponse 描述一次认证成功后的用户信息。
type AuthenticationResponse struct {
	User                string         // 用户登录名
	ProxyGrantingTicket string         // 代理授权票据（PGT）
	Proxies             []string       // 代理链列表
	AuthenticationDate  time.Time      // 认证完成的时间
	IsNewLogin          bool           // 是否通过新的认证流程获得 ST
	IsRememberedLogin   bool           // 是否使用了长期认证票据获得 ST
	MemberOf            []string       // 用户所属的组列表
	Attributes          UserAttributes // 用户的额外属性
}

// UserAttributes 以多值方式存储用户附加属性。
type UserAttributes map[string][]string

// Get 按属性名获取第一个值；属性不存在时返回空字符串。
//
// 属性在内部以数组存储，Get 仅返回首个元素。
func (a UserAttributes) Get(name string) string {
	if v, ok := a[name]; ok {
		return v[0]
	}

	return ""
}

// Add 向指定属性追加一个新值。
func (a UserAttributes) Add(name, value string) {
	a[name] = append(a[name], value)
}

// ParseServiceResponse 解析 CAS serviceValidate 响应字节流，
// 成功返回 AuthenticationResponse，失败返回 AuthenticationError。
func ParseServiceResponse(data []byte) (*AuthenticationResponse, error) {
	var x XmlServiceResponse

	if err := xml.Unmarshal(data, &x); err != nil {
		return nil, err
	}

	if x.Failure != nil {
		msg := strings.TrimSpace(x.Failure.Message)
		err := &AuthenticationError{Code: x.Failure.Code, Message: msg}
		return nil, err
	}

	r := &AuthenticationResponse{
		User:                x.Success.User,
		ProxyGrantingTicket: x.Success.ProxyGrantingTicket,
		Attributes:          make(UserAttributes),
	}

	if p := x.Success.Proxies; p != nil {
		r.Proxies = p.Proxies
	}

	if a := x.Success.Attributes; a != nil {
		r.AuthenticationDate = a.AuthenticationDate
		r.IsRememberedLogin = a.LongTermAuthenticationRequestTokenUsed
		r.IsNewLogin = a.IsFromNewLogin
		r.MemberOf = a.MemberOf

		if a.UserAttributes != nil {
			for _, ua := range a.UserAttributes.Attributes {
				if ua.Name == "" {
					continue
				}

				r.Attributes.Add(ua.Name, strings.TrimSpace(ua.Value))
			}

			for _, ea := range a.UserAttributes.AnyAttributes {
				r.Attributes.Add(ea.XMLName.Local, strings.TrimSpace(ea.Value))
			}
		}

		if a.ExtraAttributes != nil {
			for _, ea := range a.ExtraAttributes {
				r.Attributes.Add(ea.XMLName.Local, strings.TrimSpace(ea.Value))
			}
		}
	}

	for _, ea := range x.Success.ExtraAttributes {
		addRubycasAttribute(r.Attributes, ea.XMLName.Local, strings.TrimSpace(ea.Value))
	}

	return r, nil
}

// addRubycasAttribute 解析 RubyCAS 风格的附加属性值，
// 会将 YAML 字面量（如 "--- true"、列表等）转换为对应的字符串属性。
func addRubycasAttribute(attributes UserAttributes, key, value string) {
	if !strings.HasPrefix(value, "---") {
		attributes.Add(key, value)
		return
	}

	if value == "--- true" {
		attributes.Add(key, "true")
		return
	}

	if value == "--- false" {
		attributes.Add(key, "false")
		return
	}

	var decoded interface{}
	if err := yaml.Unmarshal([]byte(value), &decoded); err != nil {
		attributes.Add(key, err.Error())
		return
	}

	switch reflect.TypeOf(decoded).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(decoded)

		for i := 0; i < s.Len(); i++ {
			e := s.Index(i).Interface()

			switch reflect.TypeOf(e).Kind() {
			case reflect.String:
				attributes.Add(key, e.(string))
			}
		}
	case reflect.String:
		s := reflect.ValueOf(decoded).Interface()
		attributes.Add(key, s.(string))
	default:
		if glog.V(2) {
			kind := reflect.TypeOf(decoded).Kind()
			glog.Warningf("cas: service response: unable to parse %v value: %#v (kind: %v)", key, decoded, kind)
		}
	}
}
