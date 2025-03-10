package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/akarachen/magic-alias/pkg/shell"
	"github.com/akarachen/magic-alias/pkg/ui"
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

		// Create a tabwriter for aligned output
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		
		// Print table header
		fmt.Fprintln(w, "ALIAS	PATH	")
		fmt.Fprintln(w, "-----	----	")
		
		// Print each alias and its path
		for _, alias := range aliases {
			aliasPath := shell.GetAliasPath(alias)
			fmt.Fprintf(w, "%s	%s	\n", alias, aliasPath)
		}
		
		// Flush the tabwriter buffer
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
