package cmd

import (
	"fmt"
	"github.com/adrianfuro/GoCodePilot/config"
	"github.com/adrianfuro/GoCodePilot/internal/openai"
	"github.com/spf13/cobra"
	"log"
)

type deleteCommandParams struct {
	Name string
}

func init() {
	var params deleteCommandParams

	var deleteAssistant = &cobra.Command{
		Use:   `delete -n/--name <name_of_assistant>`,
		Short: "Delete OpenAI Assistant configuration",
		Run: func(cmd *cobra.Command, args []string) {
			deleteAssistantConfig(params)
		},
	}

	deleteAssistant.Flags().StringVarP(&params.Name, "name", "n", "", "Name of the assistant")

	err := deleteAssistant.MarkFlagRequired("name")
	if err != nil {
		log.Fatal(err)
	}

	RootCommand.AddCommand(deleteAssistant)
}

func deleteAssistantConfig(params deleteCommandParams) {
	envVars := config.LoadEnvVars()
	client := openai.NewClient(envVars.OpenAIKey)

	response, _ := client.DeleteAssistantOpenAI(params.Name)
	fmt.Println(response)
}
