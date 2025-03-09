package cmd

import (
	"fmt"
	"os"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/akarachen/magic-alias/pkg/ui"
	"github.com/charmbracelet/huh"
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
		fmt.Println(ui.Title("Initializing Magic Alias"))

		// Create a loading indicator
		var complete bool
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title("Setting up Magic Alias..."),
			),
		)

		// Process in a goroutine
		go func() {
			// Get the shell rc path
			shellName, err := shell.GetShellName()
			rcPath, err := shell.GetShellRcPath()
			if err != nil {
				panic(err)
			}

			// Create the magic-alias folder if it doesn't exist
			if err := os.MkdirAll(shell.MagicAliasPath, os.ModePerm); err != nil {
				fmt.Println(ui.Error("Error creating magic-alias folder: " + err.Error()))
				os.Exit(1)
			}

			// Write the magic alias line to the rc file
			err = shell.WriteMagicAliasToRc(shellName)
			if err != nil {
				fmt.Println(ui.Error("Error writing to rc file: " + err.Error()))
				os.Exit(1)
			}

			// Mark as complete
			complete = true

			// Show success message
			fmt.Println(ui.Success("Magic Alias successfully initialized!"))
			fmt.Println(ui.Info("Added to " + rcPath))
			fmt.Println(ui.Warning("Please restart your shell or run 'source " + rcPath + "' to apply changes."))
		}()

		// Run the form
		form.Run()

		// Wait for the goroutine to complete if it hasn't already
		for !complete {
			// Small pause to avoid CPU spinning
			fmt.Print("")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
