package cmd

import (
	"fmt"

	"github.com/akarachen/magic-alias/pkg/shell"
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
		if len(args) == 0 {
			fmt.Println("Error: alias name is required")
			return
		}

		aliasName := args[0]

		err := shell.RemoveAlias(aliasName)
		if err != nil {
			fmt.Printf("Error removing alias: %v\n", err)
			return
		}

		fmt.Printf("Alias '%s' removed successfully\n", aliasName)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
