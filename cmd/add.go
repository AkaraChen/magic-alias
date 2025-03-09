package cmd

import (
	"fmt"
	"os"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <alias> <command>",
	Short: "Add a new alias",
	Long: `Add a new alias that will execute the specified command.
Example: magic-alias add m git`,
	Args: cobra.ExactArgs(2),
	Example: "magic-alias add m git",
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]
		command := args[1]

		err := shell.AddAlias(alias, command)
		if err != nil {
			fmt.Printf("Error adding alias: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully added alias: %s\n", alias)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
