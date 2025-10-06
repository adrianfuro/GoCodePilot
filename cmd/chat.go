package cmd

import (
	"github.com/adrianfuro/GoCodePilot/pkg/ui"
	"github.com/spf13/cobra"
)

func init() {
	var chatAssistant = &cobra.Command{
		Use:   "chat",
		Short: "Chat with OpenAI Assistants",
		Run: func(cmd *cobra.Command, args []string) {
			ui.StartTUI()
		},
	}
	RootCommand.AddCommand(chatAssistant)
}
