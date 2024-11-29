package openai

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type ListResponse struct {
	Data []struct {
		Id           string `json:"id"`
		Object       string `json:"object"`
		Name         string `json:"name"`
		Model        string `json:"model"`
		Instructions string `json:"instructions"`
		Tools        []struct {
			Type string `json:"type"`
		} `json:"tools"`
	} `json:"data"`
}

func (c *Client) ListAssistantOpenAI() (string, error) {
	client := resty.New()

	url := "https://api.openai.com/v1/assistants?order=desc&limit=20"

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.APIKey).
		SetHeader("OpenAI-Beta", "assistants=v2").
		Get(url)

	if err != nil {
		log.Printf("Error sending request to %s: %v", url, err)
		return "", err
	}

	if resp.StatusCode() != 200 {
		log.Printf("Non-200 HTTP Response: %d %s", resp.StatusCode(), resp.String())
		return "", fmt.Errorf("received non-200 HTTP response")
	}

	var listResp ListResponse
	if err := json.Unmarshal(resp.Body(), &listResp); err != nil {
		log.Printf("Error unmarshalling JSON response: %v", err)
		return "", err
	}
	// Marshal the data to JSON
	jsonData, err := json.Marshal(listResp.Data)
	if err != nil {
		log.Printf("Error marshalling data to JSON: %v", err)
		return "", err
	}

	return string(jsonData), nil
}

