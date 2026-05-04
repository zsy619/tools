package xgpt

type OpenAIOption func(*OpenAI)

func NewOpenAI(opts ...OpenAIOption) (openai *OpenAI) {
	openai = &OpenAI{}
	for _, opt := range opts {
		opt(openai)
	}
	return
}

func WithOpenAIPrompt(prompt string) func(*OpenAI) {
	return func(openai *OpenAI) {
		openai.Prompt = prompt
	}
}

func WithOpenAIMaxTokens(maxtokens int) func(*OpenAI) {
	return func(openai *OpenAI) {
		openai.MaxTokens = maxtokens
	}
}

func WithOpenAITemperature(temperature float32) func(*OpenAI) {
	return func(openai *OpenAI) {
		openai.Temperature = temperature
	}
}
