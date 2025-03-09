package cmd

import (
	"fmt"
	"os"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/akarachen/magic-alias/pkg/ui"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// Variables for command flags
var (
	skipConfirm bool
	removeAllAliases bool
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall Magic Alias from your system",
	Long: `Removes Magic Alias configuration from your shell environment.

This command will:
- Remove the Magic Alias configuration from your shell's rc file (bash, zsh, or fish)
- Optionally remove all your saved aliases

Note: This will not remove the Magic Alias binary. To completely remove Magic Alias,
you'll need to uninstall it using your package manager after running this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Using UI package for styles
		fmt.Println(ui.Title("Uninstalling Magic Alias"))

		var confirmed bool = skipConfirm // If skipConfirm is true, we don't need to ask
		var removeAliases bool = removeAllAliases // Use flag value if provided

		// Ask for confirmation if not skipping
		if !skipConfirm {
			confirmForm := huh.NewForm(
				huh.NewGroup(
					huh.NewConfirm().
						Title("Are you sure you want to uninstall Magic Alias?").
						Value(&confirmed),
				),
			)

			err := confirmForm.Run()
			if err != nil {
				fmt.Println(ui.Error("Error: " + err.Error()))
				os.Exit(1)
			}

			if !confirmed {
				fmt.Println(ui.Info("Uninstall cancelled."))
				return
			}

			// Ask if user wants to remove aliases (only if not specified in flags)
			if !removeAllAliases {
				removeAliasesForm := huh.NewForm(
					huh.NewGroup(
						huh.NewConfirm().
							Title("Do you want to remove all your saved aliases?").
							Value(&removeAliases),
					),
				)

				err = removeAliasesForm.Run()
				if err != nil {
					fmt.Println(ui.Error("Error: " + err.Error()))
					os.Exit(1)
				}
			}
		}

		// Create a loading indicator
		var complete bool
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title("Uninstalling Magic Alias..."),
			),
		)

		// Process in a goroutine
		go func() {
			// Get the shell rc path
			rcPath, err := shell.GetShellRcPath()
			if err != nil {
				fmt.Println(ui.Error("Error getting shell rc path: " + err.Error()))
				os.Exit(1)
			}

			// Remove Magic Alias from rc file
			err = shell.RemoveMagicAliasFromRc(rcPath)
			if err != nil {
				fmt.Println(ui.Error("Error removing from rc file: " + err.Error()))
				os.Exit(1)
			}

			// Remove aliases if requested
			if removeAliases {
				if err := os.RemoveAll(shell.MagicAliasPath); err != nil {
					fmt.Println(ui.Error("Error removing aliases directory: " + err.Error()))
					os.Exit(1)
				}
			}

			// Mark as complete
			complete = true

			// Show success message
			fmt.Println(ui.Success("Magic Alias successfully uninstalled!"))
			fmt.Println(ui.Info("Removed from " + rcPath))
			if removeAliases {
				fmt.Println(ui.Info("All aliases have been removed."))
			}
			fmt.Println(ui.Warning("Please restart your shell or run 'source " + rcPath + "' to apply changes."))
			fmt.Println(ui.Info("To completely remove Magic Alias, uninstall it using your package manager."))
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
	uninstallCmd.Flags().BoolVarP(&skipConfirm, "yes", "y", false, "Skip confirmation prompts")
	uninstallCmd.Flags().BoolVarP(&removeAllAliases, "remove-aliases", "r", false, "Remove all saved aliases")
	rootCmd.AddCommand(uninstallCmd)
}
