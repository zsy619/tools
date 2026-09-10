package xcas

import (
	"encoding/xml"
	"time"
)

// XmlServiceResponse CAS XML 服务响应结构体，包含认证失败或认证成功两种结果之一。
type XmlServiceResponse struct {
	XMLName xml.Name `xml:"http://www.yale.edu/tp/cas serviceResponse"`

	Failure *XmlAuthenticationFailure
	Success *XmlAuthenticationSuccess
}

// XmlAuthenticationFailure CAS 认证失败响应结构体。
type XmlAuthenticationFailure struct {
	XMLName xml.Name `xml:"authenticationFailure"`
	Code    string   `xml:"code,attr"`
	Message string   `xml:",innerxml"`
}

// XmlAuthenticationSuccess CAS 认证成功响应结构体。
type XmlAuthenticationSuccess struct {
	XMLName             xml.Name           `xml:"authenticationSuccess"`
	User                string             `xml:"user"`
	ProxyGrantingTicket string             `xml:"proxyGrantingTicket,omitempty"`
	Proxies             *XmlProxies        `xml:"proxies"`
	Attributes          *XmlAttributes     `xml:"attributes"`
	ExtraAttributes     []*XmlAnyAttribute `xml:",any"`
}

// XmlProxies CAS 响应中的代理服务器列表结构体。
type XmlProxies struct {
	XMLName xml.Name `xml:"proxies"`
	Proxies []string `xml:"proxy"`
}

// AddProxy 向代理服务器列表中添加一个代理地址。
func (p *XmlProxies) AddProxy(proxy string) {
	p.Proxies = append(p.Proxies, proxy)
}

// XmlAttributes CAS 认证成功响应中的用户属性结构体。
type XmlAttributes struct {
	XMLName                                xml.Name  `xml:"attributes"`
	AuthenticationDate                     time.Time `xml:"authenticationDate"`
	LongTermAuthenticationRequestTokenUsed bool      `xml:"longTermAuthenticationRequestTokenUsed"`
	IsFromNewLogin                         bool      `xml:"isFromNewLogin"`
	MemberOf                               []string  `xml:"memberOf"`
	UserAttributes                         *XmlUserAttributes
	ExtraAttributes                        []*XmlAnyAttribute `xml:",any"`
}

// XmlUserAttributes CAS 用户自定义属性结构体。
type XmlUserAttributes struct {
	XMLName       xml.Name             `xml:"userAttributes"`
	Attributes    []*XmlNamedAttribute `xml:"attribute"`
	AnyAttributes []*XmlAnyAttribute   `xml:",any"`
}

// XmlNamedAttribute 带名称的单个属性结构体。
type XmlNamedAttribute struct {
	XMLName xml.Name `xml:"attribute"`
	Name    string   `xml:"name,attr,omitempty"`
	Value   string   `xml:",innerxml"`
}

// XmlAnyAttribute 任意属性结构体，用于保存未显式映射的扩展属性。
type XmlAnyAttribute struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

// MarshalXML 将 XmlServiceResponse 序列化为 XML；当 indent 大于 0 时，以 indent 个空格作为缩进进行格式化输出。
func (xsr *XmlServiceResponse) MarshalXML(indent int) ([]byte, error) {
	if indent == 0 {
		return xml.Marshal(xsr)
	}

	indentStr := ""
	for i := 0; i < indent; i++ {
		indentStr += " "
	}

	return xml.MarshalIndent(xsr, "", indentStr)
}

// FailureServiceResponse 构建并返回一个认证失败的 CAS 服务响应。
func FailureServiceResponse(code, message string) *XmlServiceResponse {
	return &XmlServiceResponse{
		Failure: &XmlAuthenticationFailure{
			Code:    code,
			Message: message,
		},
	}
}

// SuccessServiceResponse 构建并返回一个认证成功的 CAS 服务响应，pgt 为代理授权票据。
func SuccessServiceResponse(username, pgt string) *XmlServiceResponse {
	return &XmlServiceResponse{
		Success: &XmlAuthenticationSuccess{
			User:                username,
			ProxyGrantingTicket: pgt,
		},
	}
}
