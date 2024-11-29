package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvVars struct {
	LLMProvider  string
	SystemPrompt string
	OpenAIKey    string
	Model        string
	Temperature  string
	Tokens       string
}

func LoadEnvVars() *EnvVars {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &EnvVars{
		LLMProvider:  os.Getenv("LLM_PROVIDER"),
		SystemPrompt: os.Getenv("SYSTEM_PROMPT"),
		OpenAIKey:    os.Getenv("OPENAI_KEY"),
		Model:        os.Getenv("MODEL"),
		Temperature:  os.Getenv("TEMPERATURE"),
		Tokens:       os.Getenv("TOKENS"),
	}
}
