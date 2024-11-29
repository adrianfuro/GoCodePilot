package cmd

import (
	"os"
	"path"

	"github.com/spf13/cobra"
)

// RootCommand is the base CLI command that all subcommands are added to.
var RootCommand = &cobra.Command{
	Use:   path.Base(os.Args[0]),
	Short: "GoCodePilot",
	Long:  "An open-source AI Assistant for your terminal",
}
