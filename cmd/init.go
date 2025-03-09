package cmd

import (
	"os"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/akarachen/magic-alias/pkg/ui"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Magic Alias in your shell environment",
	Long: `Sets up Magic Alias in your shell configuration file to enable automatic alias loading.

This command will:
- Create the necessary directory structure for storing aliases
- Add the required configuration to your shell's rc file (bash, zsh, or fish)
- Configure your PATH to include magic-alias commands

After initialization, you'll need to restart your shell or source your rc file
to apply the changes immediately.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Using UI package for styles

		// Show initialization message
		ui.LogTitle("Initializing Magic Alias")
		ui.LogInfo("Setting up Magic Alias...")

		// Get the shell rc path
		shellName, err := shell.GetShellName()
		rcPath, err := shell.GetShellRcPath()
		if err != nil {
			ui.LogError("Error getting shell information: %v", err)
			os.Exit(1)
		}

		// Create the magic-alias folder if it doesn't exist
		if err := os.MkdirAll(shell.MagicAliasPath, os.ModePerm); err != nil {
			ui.LogError("Error creating magic-alias folder: %v", err)
			os.Exit(1)
		}

		// Write the magic alias line to the rc file
		err = shell.WriteMagicAliasToRc(shellName)
		if err != nil {
			ui.LogError("Error writing to rc file: %v", err)
			os.Exit(1)
		}

		// Show success message
		ui.LogSuccess("Magic Alias successfully initialized!")
		ui.LogInfo("Added to %s", rcPath)
		ui.LogWarning("Please restart your shell or run 'source %s' to apply changes.", rcPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
