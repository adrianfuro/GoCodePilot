package main

import (
	"github.com/adrianfuro/GoCodePilot/cmd"
	"os"
)

func main() {
	if err := cmd.RootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
