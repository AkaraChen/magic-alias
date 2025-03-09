package cmd

import (
	"fmt"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/akarachen/magic-alias/pkg/ui"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove <alias>",
	Aliases: []string{"rm"},
	Short:   "Delete an existing alias from your configuration",
	Long: `Remove an existing alias from your Magic Alias configuration.

This command offers two ways to remove an alias:
1. Directly via command line: ma remove myalias
2. Interactively: Run without arguments to select from a list of existing aliases

The interactive mode includes:
- A selection menu of all available aliases
- Confirmation prompt to prevent accidental deletion
- Clear success/error feedback

Once removed, the alias will no longer be available in your shell.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Using UI package for styles

		var aliasName string

		// If an alias name is provided as an argument, use it directly
		if len(args) > 0 {
			aliasName = args[0]
		} else {
			// Otherwise, show an interactive selection of available aliases
			aliases, err := shell.ListAliases()
			if err != nil {
				ui.LogErrorAndExit("Error listing aliases: %v", "err", err)
			}

			if len(aliases) == 0 {
				ui.LogWarning("No aliases found to remove.")
				return
			}

			ui.LogTitle("Remove Alias")

			// Create a form to select which alias to remove
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[string]().
						Title("Select alias to remove").
						Options(
							huh.NewOptions(aliases...)...,
						).
						Value(&aliasName),
				),
			)

			err = form.Run()
			if err != nil {
				ui.LogErrorAndExit("Error: %v", "err", err)
			}

			if aliasName == "" {
				return // User cancelled
			}

			// Confirm removal
			var confirmed bool
			confirmForm := huh.NewForm(
				huh.NewGroup(
					huh.NewConfirm().
						Title(fmt.Sprintf("Are you sure you want to remove alias '%s'?", aliasName)).
						Value(&confirmed),
				),
			)

			err = confirmForm.Run()
			if err != nil {
				ui.LogErrorAndExit("Error: %v", "err", err)
			}

			if !confirmed {
				ui.LogWarning("Removal cancelled.")
				return
			}
		}

		// Remove the alias
		err := shell.RemoveAlias(aliasName)
		if err != nil {
			ui.LogError("Error removing alias: %v", "err", err)
			return
		}

		ui.LogSuccess("Alias '%s' removed successfully", "alias", aliasName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
