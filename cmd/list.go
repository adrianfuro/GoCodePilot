package cmd

import (
	"fmt"
	"log"

	"github.com/adrianfuro/GoCodePilot/config"
	"github.com/adrianfuro/GoCodePilot/internal/openai"
	"github.com/spf13/cobra"
)

func init() {
	var listAssistant = &cobra.Command{
		Use:   `list`,
		Short: "List OpenAI Assistants configuration",
		Run: func(cmd *cobra.Command, args []string) {
			envVars := config.LoadEnvVars()
			client := openai.NewClient(envVars.OpenAIKey)

			response, err := client.ListAssistantOpenAI()
			if err != nil {
				log.Fatalf("Error listing assistants: %v", err)
			}

			fmt.Println(response)
		},
	}

	RootCommand.AddCommand(listAssistant)
}
