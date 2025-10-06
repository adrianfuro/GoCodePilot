package openai

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type DeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

func (c *Client) DeleteAssistantOpenAI(assistant string) (string, error) {
	client := resty.New()

	url := "https://api.openai.com/v1/assistants/" + assistant

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+c.APIKey).
		SetHeader("OpenAI-Beta", "assistants=v2").
		Delete(url)

	if err != nil {
		log.Printf("Error sending request to %s: %v", url, err)
		return "", err
	}

	if resp.StatusCode() != 200 {
		log.Printf("Non-200 HTTP Response: %d %s", resp.StatusCode(), resp.String())
		return "", fmt.Errorf("received non-200 HTTP response")
	}

	var deleteResp DeleteResponse
	if err := json.Unmarshal(resp.Body(), &deleteResp); err != nil {
		log.Printf("Error unmarshalling JSON response: %v", err)
		return "", err
	}

	return fmt.Sprintf("Assistant with the id %s has been deleted", assistant), nil

}
