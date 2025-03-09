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
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Using UI package for styles

		// Show initialization message
		ui.LogTitle("Initializing Magic Alias")
		ui.LogInfo("Setting up Magic Alias...")

		// Get the shell rc path
		shellName, err := shell.GetShellName()
		rcPath, err := shell.GetShellRcPath()
		if err != nil {
			ui.LogErrorAndExit("Get shell information failed.", "err", err)
		}

		// Create the magic-alias folder if it doesn't exist
		if err := os.MkdirAll(shell.MagicAliasPath, os.ModePerm); err != nil {
			ui.LogErrorAndExit("Create magic-alias folder failed.", "err", err)
		}

		// Write the magic alias line to the rc file
		err = shell.WriteMagicAliasToRc(shellName)
		if err != nil {
			ui.LogErrorAndExit("Write to rc file failed.", "err", err)
		}

		// Show success message
		ui.LogSuccess("Initialize Magic Alias successfully.")
		ui.LogInfo("Added to shell configuration.", "path", rcPath)
		ui.LogWarning("Please restart your shell or run source command to apply changes.", "path", rcPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
