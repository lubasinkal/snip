package cmd

import (
	"fmt"

	"github.com/lubasinkal/snip/internal/storage"
	"github.com/lubasinkal/snip/internal/ui"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snippets",
	Long:  `Display all saved snippets with their ID, title, and tags.`,
	Run: func(cmd *cobra.Command, args []string) {
		snippets, err := storage.ListAllSnippets()
		if err != nil {
			fmt.Println(ui.RenderError("Error listing snippets: " + err.Error()))
			return
		}

		// Show header
		fmt.Println(ui.RenderTitle(ui.IconList + " Your Code Snippets"))
		fmt.Println()

		// Render the beautiful table
		fmt.Println(ui.RenderSnippetsTable(snippets))
	},
}



func init() {
	rootCmd.AddCommand(listCmd)
}
