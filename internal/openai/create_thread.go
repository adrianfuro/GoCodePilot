package openai

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ThreadsCreateResponse struct {
	ID            string            `json:"id"`
	Object        string            `json:"object"`
	CreatedAt     int64             `json:"created_at"`
	Metadata      map[string]string `json:"metadata"`
	ToolResources map[string]string `json:"tool_resources"`
}

func (c *Client) CreateThread() (*ThreadsCreateResponse, error) {
	client := resty.New()
	url := "https://api.openai.com/v1/threads"
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.APIKey).
		SetHeader("OpenAI-Beta", "assistants=v2").
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 && resp.StatusCode() != 201 {
		return nil, fmt.Errorf("received non-200/201 status code")
	}
	var response ThreadsCreateResponse
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return nil, err
	}
	return &response, nil
}
