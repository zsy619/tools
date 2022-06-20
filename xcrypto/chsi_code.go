package xcrypto

const (
	Chsi_OptResult_Success   = "00000000" // 操作成功
	Chsi_OptResult_Parameter = "10000001" // 参数错误
	Chsi_OptResult_Accept    = "10000002" // 请求已受理
	Chsi_OptResult_Retry     = "10000003" // 请求受理失败，请重试
	Chsi_OptResult_Error     = "10000004" // 数据上报失败，详看errorDesc
	Chsi_OptResult_Failure   = "99999999" // 服务异常，请联系管理员
)
