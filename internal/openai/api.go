package openai

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
)

type ModelConfig struct {
	Model       string  `json:"model"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type Client struct {
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func (c *Client) CallOpenAI(messages []Message, modelconfig *ModelConfig) (string, error) {
	client := resty.New()

	body, err := json.Marshal(map[string]interface{}{
		"model":       modelconfig.Model,
		"messages":    messages,
		"max_tokens":  modelconfig.MaxTokens,
		"temperature": modelconfig.Temperature,
	})

	if err != nil {
		log.Printf("Error marshalling request body: %v", err)
		return "", err
	}

	url := "https://api.openai.com/v1/chat/completions"

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.APIKey).
		SetBody(body).
		Post(url)

	if err != nil {
		log.Printf("Error sending request to %s: %v", url, err)
		return "", err
	}

	if resp.StatusCode() != 200 {
		log.Printf("Non-200 HTTP Response: %d %s", resp.StatusCode(), resp.String())
		return "", fmt.Errorf("received non-200 HTTP response")
	}

	var result OpenAIResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		log.Printf("Error unmarshalling JSON response: %v", err)
		return "", err
	}

	var contents []string
	for _, choice := range result.Choices {
		contents = append(contents, choice.Message.Content)
	}

	return strings.Join(contents, "\n"), nil
}
