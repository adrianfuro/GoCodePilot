package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adrianfuro/GoCodePilot/internal/openai"
	"github.com/joho/godotenv"
)

type EnvVars struct {
	SystemPrompt string
	OpenAIKey    string
}

func LoadEnvVars() EnvVars {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return EnvVars{
		SystemPrompt: os.Getenv("SYSTEM_PROMPT"),
		OpenAIKey:    os.Getenv("OPENAI_KEY"),
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a message to send to OpenAI.")
	}

	envVars := LoadEnvVars()

	userInput := strings.Join(os.Args[1:], " ")

	messages := []openai.Message{
		{Role: "system", Content: envVars.SystemPrompt},
		{Role: "user", Content: userInput},
	}

	client := openai.NewClient(envVars.OpenAIKey)
	response, err := client.CallOpenAI(messages)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(response)
}
