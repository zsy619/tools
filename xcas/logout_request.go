package xcas

import (
	"crypto/rand"
	"encoding/xml"
	"strings"
	"time"
)

// logoutRequest 表示 CAS 单点登出（SLO）请求的 XML 数据结构。
type logoutRequest struct {
	XMLName         xml.Name  `xml:"urn:oasis:names:tc:SAML:2.0:protocol LogoutRequest"`
	Version         string    `xml:"Version,attr"`
	IssueInstant    time.Time `xml:"-"`
	RawIssueInstant string    `xml:"IssueInstant,attr"`
	ID              string    `xml:"ID,attr"`
	NameID          string    `xml:"urn:oasis:names:tc:SAML:2.0:assertion NameID"`
	SessionIndex    string    `xml:"SessionIndex"`
}

// parseLogoutRequest 解析 SLO 请求的 XML 数据，转换 IssueInstant 时间格式并清理空白字符。
// 解析失败或时间格式无法识别时返回错误。
func parseLogoutRequest(data []byte) (*logoutRequest, error) {
	l := &logoutRequest{}
	if err := xml.Unmarshal(data, &l); err != nil {
		return nil, err
	}

	t, err := parseDate(l.RawIssueInstant)
	if err != nil {
		return nil, err
	}

	l.IssueInstant = t
	l.NameID = strings.TrimSpace(l.NameID)
	l.SessionIndex = strings.TrimSpace(l.SessionIndex)

	return l, nil
}

// parseDate 解析 CAS SLO 中的 IssueInstant 时间字符串，先尝试 RFC1123Z，再尝试 ISO8601 格式。
// 两种格式都不匹配时返回错误。
func parseDate(raw string) (time.Time, error) {
	t, err := time.Parse(time.RFC1123Z, raw)
	if err != nil {
		// 若 RFC1123Z 不匹配，则尝试 ISO8601
		t, err = time.Parse("2006-01-02T15:04:05Z0700", raw)
		if err != nil {
			return t, err
		}
	}
	return t, nil
}

// newLogoutRequestID 生成 64 字符的十六进制字符串，用作 LogoutRequest 的 ID。
func newLogoutRequestID() string {
	const alphabet = "abcdef0123456789"

	// 生成 64 字符字符串
	bytes := make([]byte, 64)
	rand.Read(bytes)

	for k, v := range bytes {
		bytes[k] = alphabet[v%byte(len(alphabet))]
	}

	return string(bytes)
}

// xmlLogoutRequest 为指定票据生成 LogoutRequest 的 XML 字节流。
// 返回的字节可直接作为 SLO 请求的 logoutRequest 表单字段值。
func xmlLogoutRequest(ticket string) ([]byte, error) {
	l := &logoutRequest{
		Version:      "2.0",
		IssueInstant: time.Now().UTC(),
		ID:           newLogoutRequestID(),
		NameID:       "@NOT_USED@",
		SessionIndex: ticket,
	}

	l.RawIssueInstant = l.IssueInstant.Format(time.RFC1123Z)

	return xml.MarshalIndent(l, "", "  ")
}
