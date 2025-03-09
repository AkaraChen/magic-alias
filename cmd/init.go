package cmd

import (
	"fmt"
	"os"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize magic-alias in the shell",
	Long: `Adds magic-alias to the shell rc file and PATH
to enable automatic loading of aliases.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the shell rc path
		rcPath, err := shell.GetShellRcPath()
		if err != nil {
			fmt.Printf("Error getting shell rc path: %v\n", err)
			os.Exit(1)
		}

		// Create the magic-alias folder if it doesn't exist
		if err := os.MkdirAll(shell.MagicAliasPath, os.ModePerm); err != nil {
			fmt.Printf("Error creating magic-alias folder: %v\n", err)
			os.Exit(1)
		}

		// Write the magic alias line to the rc file
		err = shell.WriteMagicAliasToRc(rcPath)
		if err != nil {
			fmt.Printf("Error writing to rc file: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("magic-alias successfully initialized!")
		fmt.Printf("Added to %s\n", rcPath)
		fmt.Println("\nPlease restart your shell or run 'source " + rcPath + "' to apply changes.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
