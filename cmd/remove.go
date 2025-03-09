package cmd

import (
	"fmt"
	"os"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove <alias>",
	Aliases: []string{"rm"},
	Short:   "Remove an existing alias",
	Long: `Remove an existing alias by deleting its corresponding file.

Example:
  magic-alias remove myalias`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create styles for the UI
		warningStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFAF00")).
			Bold(true).
			Margin(1, 0)

		var aliasName string

		// If an alias name is provided as an argument, use it directly
		if len(args) > 0 {
			aliasName = args[0]
		} else {
			// Otherwise, show an interactive selection of available aliases
			aliases, err := shell.ListAliases()
			if err != nil {
				fmt.Println(errorStyle.Render("Error listing aliases: " + err.Error()))
				os.Exit(1)
			}

			if len(aliases) == 0 {
				fmt.Println(warningStyle.Render("No aliases found to remove."))
				return
			}

			fmt.Println(titleStyle.Render("✨ Remove Alias"))

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
				fmt.Println(errorStyle.Render("Error: " + err.Error()))
				os.Exit(1)
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
				fmt.Println(errorStyle.Render("Error: " + err.Error()))
				os.Exit(1)
			}

			if !confirmed {
				fmt.Println(warningStyle.Render("Removal cancelled."))
				return
			}
		}

		// Remove the alias
		err := shell.RemoveAlias(aliasName)
		if err != nil {
			fmt.Println(errorStyle.Render("Error removing alias: " + err.Error()))
			return
		}

		fmt.Println(successStyle.Render("✓ Alias '" + aliasName + "' removed successfully"))
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
