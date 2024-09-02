package ollama

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Message struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type OllamaResponse struct {
	Model              string `json:"model"`
	CreatedAt          string `json:"created_at"`
	Response           string `json:"response"`
	Done               bool   `json:"done"`
	DoneReason         string `json:"done_reason"`
	Context            []int  `json:"context"`
	TotalDuration      int64  `json:"total_duration"`
	LoadDuration       int64  `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int64  `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int64  `json:"eval_duration"`
}

func CallOllama(messages []Message) (string, error) {
	client := resty.New()

	// Concatenate all prompts into a single string
	var prompts []string
	for _, message := range messages {
		prompts = append(prompts, message.Prompt)
	}
	concatenatedPrompt := strings.Join(prompts, "\n")

	body, err := json.Marshal(map[string]interface{}{
		"model":  "llama2",
		"prompt": concatenatedPrompt,
		"stream": true, // Enable streaming
	})

	if err != nil {
		log.Printf("Error marshalling request body: %v", err)
		return "", err
	}

	url := "http://localhost:11434/api/generate"

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetDoNotParseResponse(true). // Disable automatic response parsing
		Post(url)

	if err != nil {
		log.Printf("Error sending request to %s: %v", url, err)
		return "", err
	}
	defer resp.RawBody().Close()

	if resp.StatusCode() != 200 {
		log.Printf("Non-200 HTTP Response: %d %s", resp.StatusCode(), resp.String())
		return "", fmt.Errorf("received non-200 HTTP response")
	}

	var result OllamaResponse

	reader := bufio.NewReader(resp.RawBody())
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		// Remove any leading or trailing whitespace
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if err := json.Unmarshal(line, &result); err != nil {
			log.Printf("Error unmarshalling response body: %v", err)
			continue
		}

		// Append the response part to the list

		// Print the response part
		fmt.Print(result.Response)
	}

	return "", nil
}
