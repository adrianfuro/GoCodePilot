package main

import (
	"github.com/adrianfuro/GoCodePilot/cmd"
	"os"
)

func main() {
	if err := cmd.RootCommand.Execute(); err != nil {
		os.Exit(1)
	}
	// envVars := LoadEnvVars()
	//
	//	if len(os.Args) < 2 {
	//		log.Fatalf("Please provide a message to send to %s.", envVars.LLMProvider)
	//	}
	//
	// userInput := strings.Join(os.Args[1:], " ")
	//
	// if envVars.LLMProvider == "openai" {
	//
	//		client := openai.NewClient(envVars.OpenAIKey)
	//
	//		if envVars.Assistant == true {
	//			message := openai.CreateRequest{
	//				Instructions: envVars.SystemPrompt,
	//				Name:         userInput,
	//				Tools: []struct {
	//					Type string `json:"type"`
	//				}{
	//					{
	//						Type: "code_interpreter",
	//					},
	//				},
	//				Model: envVars.Model,
	//			}
	//			response, _ := client.CreateAssistantOpenAI(&message)
	//			fmt.Println(response)
	//
	//		} else {
	//			messages := []openai.Message{
	//				{Role: "system", Content: envVars.SystemPrompt},
	//				{Role: "user", Content: userInput},
	//			}
	//
	//			temperature, err := strconv.ParseFloat(envVars.Temperature, 64)
	//			if err != nil {
	//				log.Fatalf("Error parsing TEMPERATURE: %v", err)
	//			}
	//
	//			tokens, err := strconv.Atoi(envVars.Tokens)
	//			if err != nil {
	//				log.Fatalf("Error parsing TOKENS: %v", err)
	//			}
	//
	//			modelConfig := openai.ModelConfig{
	//				MaxTokens:   tokens,
	//				Model:       envVars.Model,
	//				Temperature: temperature,
	//			}
	//
	//			response, _ := client.CallOpenAI(messages, &modelConfig)
	//
	//			fmt.Println(response)
	//		}
	//	} else if envVars.LLMProvider == "ollama" {
	//
	//	messages := []ollama.Message{
	//		{Model: "llama2", Prompt: userInput}, // Use envVars.LLMModel
	//	}
	//
	//	response, err := ollama.CallOllama(messages)
	//	if err != nil {
	//		log.Fatalf("Error calling Ollama: %v", err)
	//	}
	//
	//	fmt.Println(response)
	//
	//	} else {
	//		log.Fatalf("Invalid LLM provider: %s", envVars.LLMProvider)
	//	}
}
