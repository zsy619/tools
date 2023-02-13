package xgpt

type OpenAIResponseOption func(*OpenAIResponse)

func NewOpenAIResponse(opts ...OpenAIResponseOption) (openairesponse *OpenAIResponse) {
	openairesponse = &OpenAIResponse{}
	for _, opt := range opts {
		opt(openairesponse)
	}
	return
}

func WithOpenAIResponseID(id string) func(*OpenAIResponse) {
	return func(openairesponse *OpenAIResponse) {
		openairesponse.ID = id
	}
}

func WithOpenAIResponseResponse(response string) func(*OpenAIResponse) {
	return func(openairesponse *OpenAIResponse) {
		openairesponse.Response = response
	}
}
