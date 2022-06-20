package xcas

const (
	// CASPrefix = "/authserver"
	CASPrefix = "/cas"

	// CASLoginURI represents credential requestor / acceptor
	CASLoginURI = CASPrefix + "/login"

	// CASLogoutURI represents destroy CAS session (logout)
	CASLogoutURI = CASPrefix + "/logout"

	// CASValidateURI represents service ticket validation
	CASValidateURI = CASPrefix + "/validate"

	// CASVersion2ServiceValidateURI represents service ticket validation [CAS 2.0]
	CASVersion2ServiceValidateURI = CASPrefix + "/serviceValidate"

	// CASVersion2ProxyValidateURI represents service/proxy ticket validation [CAS 2.0]
	CASVersion2ProxyValidateURI = CASPrefix + "/proxyValidate"

	// CASVersion2ProxyURI represents proxy ticket service [CAS 2.0]
	CASVersion2ProxyURI = CASPrefix + "/proxy"

	// CASVersion3ServiceValidateURI represents service ticket validation [CAS 3.0]
	CASVersion3ServiceValidateURI = CASPrefix + "/p3/serviceValidate"

	// CASVersion3ProxyValidateURI represents service/proxy ticket validation [CAS 3.0]
	CASVersion3ProxyValidateURI = CASPrefix + "/p3/proxyValidate"

	CASV1TicketsURI = CASPrefix + "/v1/tickets"

	UserTypeField = "lx"
	ZhanghaoField = "zh"
)
