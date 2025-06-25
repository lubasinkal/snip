package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/lubasinkal/snip/internal/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snippets",
	Long:  `Display all saved snippets with their ID, title, and tags.`,
	Run: func(cmd *cobra.Command, args []string) {
		snippets, err := storage.ListAllSnippets()
		if err != nil {
			fmt.Println("❌ Error listing snippets:", err)
			return
		}

		if len(snippets) == 0 {
			fmt.Println("📝 No snippets found. Use 'snip save' to create your first snippet!")
			return
		}

		fmt.Printf("📚 Found %d snippet(s):\n\n", len(snippets))
		
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
			fmt.Printf("     Created %s\n\n", timeAgo)
		}
	},
}

func formatTimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)
	
	if diff < time.Minute {
		return "just now"
	} else if diff < time.Hour {
		minutes := int(diff.Minutes())
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	} else if diff < 24*time.Hour {
		hours := int(diff.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	} else if diff < 7*24*time.Hour {
		days := int(diff.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	} else {
		return t.Format("Jan 2, 2006")
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
