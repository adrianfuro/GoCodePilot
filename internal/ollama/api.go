package ollama

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OllamaResponse struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func CallOllama(messages []Message) (string, error) {
	client := resty.New()

	body, err := json.Marshal(map[string]interface{}{
		"model":  "llama2",
		"prompt": messages,
	})

	if err != nil {
		log.Printf("Error marshalling request body: %v", err)
		return "", err
	}

	url := "https://localhost:11434/api/generate"

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
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

	var result OllamaResponse

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		log.Printf("Error unmarshalling response body: %v", err)
		return "", err
	}

	var contents []string
	for _, choice := range result.Choices {
		contents = append(contents, choice.Message.Content)
	}

	return strings.Join(contents, "\n"), nil

}
