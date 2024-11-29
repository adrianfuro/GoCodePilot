package openai

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type CreateRequest struct {
	Instructions string `json:"instructions"`
	Name         string `json:"name"`
	Tools        []struct {
		Type string `json:"type"`
	} `json:"tools"`
	Model string `json:"model"`
}

type CreateResponse struct {
	Id           string `json:"id"`
	Object       string `json:"object"`
	Name         string `json:"name"`
	Model        string `json:"model"`
	Instructions string `json:"instructions"`
	Tools        []struct {
		Type string `json:"type"`
	} `json:"tools"`
}

func (c *Client) CreateAssistantOpenAI(createReq *CreateRequest) (string, error) {
	client := resty.New()

	body, err := json.Marshal(map[string]interface{}{
		"instructions": createReq.Instructions,
		"name":         createReq.Name,
		"tools":        createReq.Tools,
		"model":        createReq.Model,
	})

	if err != nil {
		log.Printf("Error marshalling request body: %v", err)
		return "", err
	}

	url := "https://api.openai.com/v1/assistants"

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.APIKey).
		SetHeader("OpenAI-Beta", "assistants=v2").
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

	var createResp CreateResponse
	if err := json.Unmarshal(resp.Body(), &createResp); err != nil {
		log.Printf("Error unmarshalling JSON response: %v", err)
		return "", err
	}

	return createResp.Id, nil

}
