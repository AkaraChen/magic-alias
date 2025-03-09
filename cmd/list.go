package cmd

import (
	"fmt"
	"os"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/akarachen/magic-alias/pkg/ui"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all aliases",
	Long:  `Display a list of all aliases created with magic-alias.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Using UI package for styles

		aliases, err := shell.ListAliases()
		if err != nil {
			fmt.Println(ui.Error("Error listing aliases: " + err.Error()))
			os.Exit(1)
		}

		if len(aliases) == 0 {
			fmt.Println(ui.Empty("No aliases found. Use 'magic-alias add' to create one."))
			return
		}

		fmt.Println(ui.Title("Available Aliases"))

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
			fmt.Println(ui.Error("Error: " + err.Error()))
			return
		}

		// If an alias was selected, show details
		if selectedAlias != "" {
			aliasPath := shell.GetAliasPath(selectedAlias)
			fmt.Println(ui.Title("Alias Details"))
			fmt.Println(ui.Info("Name: " + selectedAlias))
			fmt.Println(ui.Info("Path: " + aliasPath))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
