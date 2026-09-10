package xemail

import (
	log "github.com/golang/glog"

	sendgrid "github.com/sendgrid/sendgrid-go"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// NewEmailer 是一个便捷函数，使用指定的 Sendgrid apiKey 返回一个新的 Emailer 实例。
// 生产环境请使用此函数；若需使用 Sendgrid 沙箱，请改用 NewEmailerWithSandbox。
func NewEmailer(apiKey string) *Emailer {
	return NewEmailerWithSandbox(apiKey, false)
}

// NewEmailerWithSandbox 是一个便捷函数，使用指定的 apiKey 返回一个新的 Emailer 实例，
// 并可通过 useSandbox 控制是否走 Sendgrid 沙箱模式。
//
// 警告：仅用于测试，生产环境请使用 NewEmailer。
func NewEmailerWithSandbox(apiKey string, useSandbox bool) *Emailer {
	return &Emailer{
		sendGridClient: sendgrid.NewSendClient(apiKey),
		useSandbox:     useSandbox,
	}
}

// Emailer 是一个对邮件服务（Sendgrid）的封装，让发送邮件更简单。
type Emailer struct {
	sendGridClient *sendgrid.Client
	useSandbox     bool
}

// SendEmailRequest 描述 SendEmail 发送一封普通邮件所需的全部参数。
type SendEmailRequest struct {
	ToName    string // 收件人姓名
	ToEmail   string // 收件人邮箱
	FromName  string // 发件人姓名
	FromEmail string // 发件人邮箱
	Subject   string // 邮件主题
	Text      string // 纯文本正文
	HTML      string // HTML 正文
}

// SendEmail 通过底层邮件服务商发送一封普通邮件。
// 发送失败时记录错误日志并返回 err；成功时记录响应详情并返回 nil。
func (e *Emailer) SendEmail(req *SendEmailRequest) error {
	from := mail.NewEmail(req.FromName, req.FromEmail)
	to := mail.NewEmail(req.ToName, req.ToEmail)
	msg := mail.NewSingleEmail(from, req.Subject, to, req.Text, req.HTML)
	msg.MailSettings = &mail.MailSettings{
		SandboxMode: &mail.Setting{
			Enable: &e.useSandbox,
		},
	}

	resp, err := e.sendGridClient.Send(msg)
	if err != nil {
		log.Errorf("Error sending email: err: %v", err)
		return err
	}
	log.Infof("sendemail: useSandbox: %v", e.useSandbox)
	log.Infof("sendemail: response status: %v", resp.StatusCode)
	log.Infof("sendemail: response body: %v", resp.Body)
	log.Infof("sendemail: response headers: %v", resp.Headers)
	return nil
}

// TemplateData 表示邮件模板使用的键值对数据。
type TemplateData map[string]interface{}

// SendTemplateEmailRequest 描述 SendTemplateEmail 发送一封模板邮件所需的全部参数。
type SendTemplateEmailRequest struct {
	ToName       string       // 收件人姓名
	ToEmail      string       // 收件人邮箱
	FromName     string       // 发件人姓名
	FromEmail    string       // 发件人邮箱
	TemplateID   string       // Sendgrid 中的模板 ID
	TemplateData TemplateData // 模板动态数据
	AsmGroupID   int          // 抑制组 ID（0 表示不设置）
}

// SendTemplateEmail 使用邮件服务商中的模板发送一封邮件。
// 模板动态字段通过 TemplateData 提供；AsmGroupID 非 0 时会设置邮件抑制组。
// 发送失败时记录错误日志并返回 err；成功时记录响应详情并返回 nil。
func (e *Emailer) SendTemplateEmail(req *SendTemplateEmailRequest) error {
	msg := mail.NewV3Mail()
	msg.MailSettings = &mail.MailSettings{
		SandboxMode: &mail.Setting{
			Enable: &e.useSandbox,
		},
	}

	from := mail.NewEmail(req.FromName, req.FromEmail)
	to := mail.NewEmail(req.ToName, req.ToEmail)

	msg.SetFrom(from)
	msg.SetReplyTo(from)
	msg.SetTemplateID(req.TemplateID)

	p := mail.NewPersonalization()
	p.AddTos(to)
	for key, val := range req.TemplateData {
		p.SetDynamicTemplateData(key, val)
	}

	msg.AddPersonalizations(p)

	if req.AsmGroupID != 0 {
		a := mail.NewASM()
		a.SetGroupID(req.AsmGroupID)
		msg.SetASM(a)
	}

	resp, err := e.sendGridClient.Send(msg)
	if err != nil {
		log.Errorf("Error sending email: err: %v", err)
		return err
	}
	log.Infof("sendemail: useSandbox: %v", e.useSandbox)
	log.Infof("sendemail: response status: %v", resp.StatusCode)
	log.Infof("sendemail: response body: %v", resp.Body)
	log.Infof("sendemail: response headers: %v", resp.Headers)
	return nil
}
