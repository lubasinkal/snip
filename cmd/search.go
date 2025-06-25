package cmd

import (
	"fmt"
	"strings"

	"github.com/lubasinkal/snip/internal/storage"
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
			fmt.Println("❌ Error searching snippets:", err)
			return
		}

		if len(snippets) == 0 {
			if tagFilter != "" {
				fmt.Printf("🔍 No snippets found matching '%s' with tag '%s'\n", query, tagFilter)
			} else {
				fmt.Printf("🔍 No snippets found matching '%s'\n", query)
			}
			return
		}

		if tagFilter != "" {
			fmt.Printf("🔍 Found %d snippet(s) matching '%s' with tag '%s':\n\n", len(snippets), query, tagFilter)
		} else {
			fmt.Printf("🔍 Found %d snippet(s) matching '%s':\n\n", len(snippets), query)
		}
		
		for _, snippet := range snippets {
			// Format tags
			tagsStr := ""
			if len(snippet.Tags) > 0 {
				tagsStr = fmt.Sprintf(" [%s]", strings.Join(snippet.Tags, ", "))
			}
			
			// Format date
			timeAgo := formatTimeAgo(snippet.CreatedAt)
			
			// Print snippet info
			fmt.Printf("  %d. %s%s\n", snippet.ID, snippet.Title, tagsStr)
			fmt.Printf("     Created %s\n", timeAgo)
			
			// Show a preview of the content (first 100 chars)
			preview := strings.ReplaceAll(snippet.Content, "\n", " ")
			if len(preview) > 100 {
				preview = preview[:100] + "..."
			}
			if preview != "" {
				fmt.Printf("     Preview: %s\n", preview)
			}
			fmt.Println()
		}
	},
}

func init() {
	searchCmd.Flags().StringVarP(&tagFilter, "tag", "t", "", "Filter by tag")
	rootCmd.AddCommand(searchCmd)
}
