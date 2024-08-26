package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/adrianfuro/GoCodePilot/internal/openai"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a message to send to OpenAI.")
	}

	userInput := strings.Join(os.Args[1:], " ")

	messages := []openai.Message{
		{Role: "system", Content: "You are a Software Engineer and you respond only with code without markdown format."},
		{Role: "user", Content: userInput},
	}

	response, err := openai.CallOpenAI(messages)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
