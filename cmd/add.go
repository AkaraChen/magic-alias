package cmd

import (
	"fmt"
	"os"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// Styles for UI elements
var (
	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF5F87")).
		Bold(true).
		Margin(1, 0)

	successStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#5AF78E")).
		Bold(true)

	errorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF5F87")).
		Bold(true)
)

var addCmd = &cobra.Command{
	Use:   "add <alias> <command>",
	Short: "Add a new alias",
	Long: `Add a new alias that will execute the specified command.
Example: magic-alias add m git`,
	Args: cobra.MinimumNArgs(0),
	Example: "magic-alias add m git",
	Run: func(cmd *cobra.Command, args []string) {
		var alias, command string

		// If arguments are provided, use them directly
		if len(args) >= 2 {
			alias = args[0]
			command = args[1]
		} else {
			// Otherwise, use interactive form
			fmt.Println(titleStyle.Render("✨ Create New Alias"))

			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Alias Name").
						Description("What would you like to call your alias?").
						Placeholder("e.g., g").
						Validate(func(s string) error {
							if s == "" {
								return fmt.Errorf("alias cannot be empty")
							}
							return nil
						}).
						Value(&alias),
					huh.NewInput().
						Title("Command").
						Description("What command should this alias run?").
						Placeholder("e.g., git").
						Validate(func(s string) error {
							if s == "" {
								return fmt.Errorf("command cannot be empty")
							}
							return nil
						}).
						Value(&command),
				),
			)

			err := form.Run()
			if err != nil {
				fmt.Println(errorStyle.Render("Error: " + err.Error()))
				os.Exit(1)
			}
		}

		// Add the alias
		err := shell.AddAlias(alias, command)
		if err != nil {
			fmt.Println(errorStyle.Render("Error adding alias: " + err.Error()))
			os.Exit(1)
		}

		fmt.Println(successStyle.Render("✓ Successfully added alias: " + alias))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
