package xcas

import (
	"encoding/xml"
	"time"
)

type XmlServiceResponse struct {
	XMLName xml.Name `xml:"http://www.yale.edu/tp/cas serviceResponse"`

	Failure *XmlAuthenticationFailure
	Success *XmlAuthenticationSuccess
}

type XmlAuthenticationFailure struct {
	XMLName xml.Name `xml:"authenticationFailure"`
	Code    string   `xml:"code,attr"`
	Message string   `xml:",innerxml"`
}

type XmlAuthenticationSuccess struct {
	XMLName             xml.Name           `xml:"authenticationSuccess"`
	User                string             `xml:"user"`
	ProxyGrantingTicket string             `xml:"proxyGrantingTicket,omitempty"`
	Proxies             *XmlProxies        `xml:"proxies"`
	Attributes          *XmlAttributes     `xml:"attributes"`
	ExtraAttributes     []*XmlAnyAttribute `xml:",any"`
}

type XmlProxies struct {
	XMLName xml.Name `xml:"proxies"`
	Proxies []string `xml:"proxy"`
}

func (p *XmlProxies) AddProxy(proxy string) {
	p.Proxies = append(p.Proxies, proxy)
}

type XmlAttributes struct {
	XMLName                                xml.Name  `xml:"attributes"`
	AuthenticationDate                     time.Time `xml:"authenticationDate"`
	LongTermAuthenticationRequestTokenUsed bool      `xml:"longTermAuthenticationRequestTokenUsed"`
	IsFromNewLogin                         bool      `xml:"isFromNewLogin"`
	MemberOf                               []string  `xml:"memberOf"`
	UserAttributes                         *XmlUserAttributes
	ExtraAttributes                        []*XmlAnyAttribute `xml:",any"`
}

type XmlUserAttributes struct {
	XMLName       xml.Name             `xml:"userAttributes"`
	Attributes    []*XmlNamedAttribute `xml:"attribute"`
	AnyAttributes []*XmlAnyAttribute   `xml:",any"`
}

type XmlNamedAttribute struct {
	XMLName xml.Name `xml:"attribute"`
	Name    string   `xml:"name,attr,omitempty"`
	Value   string   `xml:",innerxml"`
}

type XmlAnyAttribute struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

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

func FailureServiceResponse(code, message string) *XmlServiceResponse {
	return &XmlServiceResponse{
		Failure: &XmlAuthenticationFailure{
			Code:    code,
			Message: message,
		},
	}
}

func SuccessServiceResponse(username, pgt string) *XmlServiceResponse {
	return &XmlServiceResponse{
		Success: &XmlAuthenticationSuccess{
			User:                username,
			ProxyGrantingTicket: pgt,
		},
	}
}
