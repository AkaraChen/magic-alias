package cmd

import (
	"os"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/akarachen/magic-alias/pkg/ui"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all your configured aliases",
	Long: `Display a table of all aliases created with Magic Alias.

Features:
- Shows all available aliases in a table format
- Displays alias names and their corresponding paths
- Empty state handling with helpful guidance when no aliases exist

Use this command to review your existing aliases.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Using UI package for styles

		aliases, err := shell.ListAliases()
		if err != nil {
			ui.LogErrorAndExit("List aliases failed.", "err", err)
		}

		if len(aliases) == 0 {
			ui.LogEmpty("No aliases found. Use 'ma add' to create one.")
			return
		}

		ui.LogTitle("Available Aliases")

		// Create a new table
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		
		// Configure table style
		t.SetStyle(table.StyleRounded)
		
		// Set table headers
		t.AppendHeader(table.Row{"ALIAS", "PATH"})
		
		// Configure header colors and style
		t.Style().Format.Header = text.FormatDefault
		
		// Add rows to the table
		for _, alias := range aliases {
			aliasPath := shell.GetAliasPath(alias)
			t.AppendRow(table.Row{alias, aliasPath})
		}
		
		// Render the table
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
