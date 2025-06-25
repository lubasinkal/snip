package cmd

import (
	"fmt"

	"github.com/lubasinkal/snip/internal/storage"
	"github.com/lubasinkal/snip/internal/ui"
	"github.com/spf13/cobra"
)

var tagFilter string

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search snippets by title, tags, or content",
	Long:  `Search through your snippets by title, tags, or content. Use --tag to filter by specific tags.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]

		snippets, err := storage.SearchSnippets(query, tagFilter)
		if err != nil {
			fmt.Println(ui.RenderError("Error searching snippets: " + err.Error()))
			return
		}

		// Render beautiful search results
		fmt.Println(ui.RenderSearchResults(snippets, query, tagFilter))
	},
}

func init() {
	searchCmd.Flags().StringVarP(&tagFilter, "tag", "t", "", "Filter by tag")
	rootCmd.AddCommand(searchCmd)
}
