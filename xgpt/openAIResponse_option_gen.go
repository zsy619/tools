package xgpt

// OpenAIResponseOption 是一个用于配置 OpenAIResponse 的可选参数函数类型。
type OpenAIResponseOption func(*OpenAIResponse)

// NewOpenAIResponse 创建一个新的 OpenAIResponse 实例，并可传入若干选项函数进行初始化配置。
func NewOpenAIResponse(opts ...OpenAIResponseOption) (openairesponse *OpenAIResponse) {
	openairesponse = &OpenAIResponse{}
	for _, opt := range opts {
		opt(openairesponse)
	}
	return
}

// WithOpenAIResponseID 设置 OpenAIResponse 的 ID 字段。
func WithOpenAIResponseID(id string) func(*OpenAIResponse) {
	return func(openairesponse *OpenAIResponse) {
		openairesponse.ID = id
	}
}

// WithOpenAIResponseResponse 设置 OpenAIResponse 的 Response 字段。
func WithOpenAIResponseResponse(response string) func(*OpenAIResponse) {
	return func(openairesponse *OpenAIResponse) {
		openairesponse.Response = response
	}
}
