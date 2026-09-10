package xgpt

// OpenAIOption 是一个用于配置 OpenAI 的可选参数函数类型。
type OpenAIOption func(*OpenAI)

// NewOpenAI 创建一个新的 OpenAI 实例，并可传入若干选项函数进行初始化配置。
func NewOpenAI(opts ...OpenAIOption) (openai *OpenAI) {
	openai = &OpenAI{}
	for _, opt := range opts {
		opt(openai)
	}
	return
}

// WithOpenAIPrompt 设置 OpenAI 请求的提示词 Prompt。
func WithOpenAIPrompt(prompt string) func(*OpenAI) {
	return func(openai *OpenAI) {
		openai.Prompt = prompt
	}
}

// WithOpenAIMaxTokens 设置 OpenAI 请求的最大生成 token 数 MaxTokens。
func WithOpenAIMaxTokens(maxtokens int) func(*OpenAI) {
	return func(openai *OpenAI) {
		openai.MaxTokens = maxtokens
	}
}

// WithOpenAITemperature 设置 OpenAI 请求的采样温度 Temperature。
func WithOpenAITemperature(temperature float32) func(*OpenAI) {
	return func(openai *OpenAI) {
		openai.Temperature = temperature
	}
}
