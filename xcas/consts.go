package xcas

const (
	// CASPrefix CAS 服务统一路径前缀
	CASPrefix = "/cas"

	// CASLoginURI 凭证请求/接收端点 URI（登录入口）
	CASLoginURI = CASPrefix + "/login"

	// CASLogoutURI 销毁 CAS 会话（登出）URI
	CASLogoutURI = CASPrefix + "/logout"

	// CASValidateURI CAS 1.0 服务票据（ST）校验 URI
	CASValidateURI = CASPrefix + "/validate"

	// CASVersion2ServiceValidateURI [CAS 2.0] 服务票据校验 URI
	CASVersion2ServiceValidateURI = CASPrefix + "/serviceValidate"

	// CASVersion2ProxyValidateURI [CAS 2.0] 服务票据/代理票据校验 URI
	CASVersion2ProxyValidateURI = CASPrefix + "/proxyValidate"

	// CASVersion2ProxyURI [CAS 2.0] 代理票据服务 URI
	CASVersion2ProxyURI = CASPrefix + "/proxy"

	// CASVersion3ServiceValidateURI [CAS 3.0] 服务票据校验 URI
	CASVersion3ServiceValidateURI = CASPrefix + "/p3/serviceValidate"

	// CASVersion3ProxyValidateURI [CAS 3.0] 服务票据/代理票据校验 URI
	CASVersion3ProxyValidateURI = CASPrefix + "/p3/proxyValidate"

	// CASV1TicketsURI REST 版票据管理 URI
	CASV1TicketsURI = CASPrefix + "/v1/tickets"

	// UserTypeField 用户类型字段名（lx）
	UserTypeField = "lx"
	// ZhanghaoField 账号字段名（zh）
	ZhanghaoField = "zh"
)
