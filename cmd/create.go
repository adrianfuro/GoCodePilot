package cmd

import (
	"fmt"
	"github.com/adrianfuro/GoCodePilot/config"
	"github.com/adrianfuro/GoCodePilot/internal/openai"

	"github.com/spf13/cobra"
)

type createCommandParams struct {
	name          string
	assistantType string
	model         string
}

func init() {
	var params createCommandParams

	var createAssistant = &cobra.Command{
		Use:   `create -n/--name <name_of_assistant>`,
		Short: "Create OpenAI Assistant configuration",
		Run: func(cmd *cobra.Command, args []string) {
			createAssistantConfig(params)
		},
	}

	createAssistant.Flags().StringVarP(&params.name, "name", "n", "", "Name of the assistant")
	createAssistant.Flags().StringVar(&params.model, "model", "gpt-4o", "Model to use")
	createAssistant.Flags().StringVar(&params.assistantType, "type", "code_interpreter", "Type of assistant [code_interpreter/file_search]")

	createAssistant.MarkFlagRequired("name")

	RootCommand.AddCommand(createAssistant)
}

func createAssistantConfig(params createCommandParams) {
	envVars := config.LoadEnvVars()
	client := openai.NewClient(envVars.OpenAIKey)

	message := openai.CreateRequest{
		Instructions: envVars.SystemPrompt,
		Name:         params.name,
		Tools: []struct {
			Type string `json:"type"`
		}{
			{
				Type: params.assistantType,
			},
		},
		Model: params.model,
	}
	response, _ := client.CreateAssistantOpenAI(&message)
	fmt.Println(response)
}
