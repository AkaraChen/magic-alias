package cmd

import (
	"fmt"
	"os"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/akarachen/magic-alias/pkg/ui"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <alias> <command>",
	Short: "Create a new shell alias for any command",
	Long: `Create a new shell alias that will execute the specified command when invoked.

You can provide the alias and command as arguments or use the interactive prompt:
- With arguments: ma add g git
- Without arguments: An interactive form will guide you through the process

The alias will be available in your shell immediately after creation.
Aliases are stored as executable files in the magic-alias directory.`,
	Args:    cobra.MinimumNArgs(0),
	Example: "ma add m git",
	Run: func(cmd *cobra.Command, args []string) {
		var alias, command string

		// If arguments are provided, use them directly
		if len(args) >= 2 {
			alias = args[0]
			command = args[1]
		} else {
			// Otherwise, use interactive form
			ui.LogTitle("Create New Alias")

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
				ui.LogError("Error: %v", err)
				os.Exit(1)
			}
		}

		// Add the alias
		err := shell.AddAlias(alias, command)
		if err != nil {
			ui.LogError("Error adding alias: %v", err)
			os.Exit(1)
		}

		ui.LogSuccess("Successfully added alias: %s", alias)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
