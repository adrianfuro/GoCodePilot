package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
)

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

func callOpenAI(messages []Message) (string, error) {
	client := resty.New()

	body, err := json.Marshal(map[string]interface{}{
		"model":      "gpt-4o",
		"messages":   messages,
		"max_tokens": 4000,
	})

	url := "https://api.openai.com//v1/chat/completions"

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY")).
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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a message to send to OpenAI.")
	}

	userInput := strings.Join(os.Args[1:], " ")

	messages := []Message{
		{Role: "system", Content: "You are a helpful assistant."},
		{Role: "user", Content: userInput},
	}

	response, err := callOpenAI(messages)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response from OpenAI: ", response)
}
