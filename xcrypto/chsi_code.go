package xcrypto

// 以下常量定义了接口操作结果码（Chsi_OptResult_*），用于标识各类操作的结果状态。
const (
	Chsi_OptResult_Success   = "00000000" // 操作成功
	Chsi_OptResult_Parameter = "10000001" // 参数错误
	Chsi_OptResult_Accept    = "10000002" // 请求已受理
	Chsi_OptResult_Retry     = "10000003" // 请求受理失败，请重试
	Chsi_OptResult_Error     = "10000004" // 数据上报失败，详看errorDesc
	Chsi_OptResult_Failure   = "99999999" // 服务异常，请联系管理员
)
