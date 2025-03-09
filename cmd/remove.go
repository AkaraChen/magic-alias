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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
