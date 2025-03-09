/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all aliases",
	Long:  `Display a list of all aliases created with magic-alias.`,
	Run: func(cmd *cobra.Command, args []string) {
		aliases, err := shell.ListAliases()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error listing aliases: %v\n", err)
			os.Exit(1)
		}
		
		if len(aliases) == 0 {
			fmt.Println("No aliases found.")
			return
		}
		
		fmt.Println("Available aliases:")
		for _, alias := range aliases {
			fmt.Printf("  %s\n", alias)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
