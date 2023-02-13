package xgpt

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	openAIAPI = "https://api.openai.com/v1/engines/text-davinci/jobs"
)

// OpenAI OpenAI API结构体
type OpenAI struct {
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float32 `json:"temperature"`
}

// OpenAIResponse OpenAI API返回结构体
type OpenAIResponse struct {
	ID       string `json:"id"`
	Response string `json:"response"`
}

// GenerateText
/**
 * @description: 调取OpenAI API返回文本
 * @return {string, error}
 */
func (ai *OpenAI) GenerateText() (string, error) {
	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(ai)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", openAIAPI, &buf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	var response OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.Response, nil
}

func main() {
	ai := &OpenAI{
		Prompt:      "Hello, how can I help you today?",
		MaxTokens:   100,
		Temperature: 0.5,
	}

	response, err := ai.GenerateText()
	if err != nil {
		fmt.Printf("Error generating text: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(response)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("You: ")
		if !scanner.Scan() {
			break
		}

		ai.Prompt = scanner.Text()
		response, err := ai.GenerateText()
		if err != nil {
			fmt.Printf("Error generating text: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("ChatGPT:", response)
	}
}
