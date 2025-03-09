package cmd

import (
	"github.com/akarachen/magic-alias/pkg/ui"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ma",
	Short: "A simple and powerful shell alias manager",
	Long: `Magic Alias is a modern, user-friendly tool for managing shell aliases.
It allows you to create, list, and remove aliases with an interactive interface
or command-line arguments.

Features:
- Create aliases with a simple command
- List all your aliases in an interactive menu
- Remove aliases easily
- Automatic shell integration

Get started with 'ma init' to set up the tool in your shell.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		ui.LogErrorAndExit("Execute command failed.", "err", err)
	}
}

func init() {}
