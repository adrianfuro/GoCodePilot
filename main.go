package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adrianfuro/GoCodePilot/internal/ollama"
	"github.com/adrianfuro/GoCodePilot/internal/openai"
	"github.com/joho/godotenv"
)

type EnvVars struct {
	LLMProvider  string
	SystemPrompt string
	OpenAIKey    string
}

func LoadEnvVars() EnvVars {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return EnvVars{
		LLMProvider:  os.Getenv("LLM_PROVIDER"),
		SystemPrompt: os.Getenv("SYSTEM_PROMPT"),
		OpenAIKey:    os.Getenv("OPENAI_KEY"),
	}
}

func main() {
	envVars := LoadEnvVars()

	if len(os.Args) < 2 {
		log.Fatalf("Please provide a message to send to %s.", envVars.LLMProvider)
	}

	userInput := strings.Join(os.Args[1:], " ")

	if envVars.LLMProvider == "openai" {
		messages := []openai.Message{
			{Role: "system", Content: envVars.SystemPrompt},
			{Role: "user", Content: userInput},
		}

		client := openai.NewClient(envVars.OpenAIKey)
		response, _ := client.CallOpenAI(messages)

		fmt.Println(response)

	} else if envVars.LLMProvider == "ollama" {
		messages := []ollama.Message{
			{Model: "llama2", Prompt: userInput}, // Use envVars.LLMModel
		}

		response, err := ollama.CallOllama(messages)
		if err != nil {
			log.Fatalf("Error calling Ollama: %v", err)
		}

		fmt.Println(response)

	} else {
		log.Fatalf("Invalid LLM provider: %s", envVars.LLMProvider)
	}

}
