package ollama

import (
	"encoding/json"
	"log"
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

	body, err := json.Marshal(map[string]interface{}{
		"model":    "llama2",
		"messages": messages,
	})

	if err != nil {
		log.Printf("Error marshalling request body: %v", err)
		return "", err
	}

	url := "https://localhost:11434/api/generate"
}
