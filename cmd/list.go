package cmd

import (
	"fmt"
	"os"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all aliases",
	Long:  `Display a list of all aliases created with magic-alias.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create styles for the list display
		headerStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5F87")).
			Bold(true).
			Margin(1, 0)

		aliasStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5AF78E")).
			Margin(0, 2)

		emptyStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5F87")).
			Italic(true).
			Margin(1, 0)

		aliases, err := shell.ListAliases()
		if err != nil {
			fmt.Println(errorStyle.Render("Error listing aliases: " + err.Error()))
			os.Exit(1)
		}

		if len(aliases) == 0 {
			fmt.Println(emptyStyle.Render("No aliases found. Use 'magic-alias add' to create one."))
			return
		}

		fmt.Println(headerStyle.Render("✨ Available Aliases"))

		// Create a list of aliases with selection capability
		var selectedAlias string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Your aliases").
					Description("Select an alias to see more details").
					Options(
						huh.NewOptions(aliases...)...,
					).
					Value(&selectedAlias),
			),
		)

		// Run the form
		err = form.Run()
		if err != nil {
			fmt.Println(errorStyle.Render("Error: " + err.Error()))
			return
		}

		// If an alias was selected, show details
		if selectedAlias != "" {
			aliasPath := shell.GetAliasPath(selectedAlias)
			fmt.Println(headerStyle.Render("Alias Details"))
			fmt.Println(aliasStyle.Render("Name: " + selectedAlias))
			fmt.Println(aliasStyle.Render("Path: " + aliasPath))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
