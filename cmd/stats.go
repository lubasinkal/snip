package cmd

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/lubasinkal/snip/internal/storage"
	"github.com/lubasinkal/snip/internal/ui"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show statistics about your snippets",
	Long:  `Display beautiful statistics about your code snippet collection including counts by language, most used tags, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		snippets, err := storage.ListAllSnippets()
		if err != nil {
			fmt.Println(ui.RenderError("Error loading snippets: " + err.Error()))
			return
		}

		if len(snippets) == 0 {
			fmt.Println(ui.RenderInfo("No snippets found. Use 'snip save' to create your first snippet!"))
			return
		}

		// Show header
		fmt.Println(ui.RenderTitle(ui.IconSparkles + " Snippet Statistics"))
		fmt.Println()

		// Basic stats
		totalSnippets := len(snippets)
		
		// Count tags
		tagCounts := make(map[string]int)
		totalTags := 0
		
		for _, snippet := range snippets {
			for _, tag := range snippet.Tags {
				tagCounts[tag]++
				totalTags++
			}
		}

		// Create basic stats box
		basicStats := fmt.Sprintf(`📊 Overview:
  • Total Snippets: %d
  • Total Tags: %d
  • Unique Tags: %d
  • Average Tags per Snippet: %.1f`,
			totalSnippets,
			totalTags,
			len(tagCounts),
			float64(totalTags)/float64(totalSnippets))

		fmt.Println(ui.RenderBox(basicStats))
		fmt.Println()

		// Show top tags if we have any
		if len(tagCounts) > 0 {
			fmt.Println(ui.RenderSubtitle("🏷️ Most Popular Tags"))
			fmt.Println()

			// Sort tags by count
			type tagCount struct {
				tag   string
				count int
			}
			
			var sortedTags []tagCount
			for tag, count := range tagCounts {
				sortedTags = append(sortedTags, tagCount{tag, count})
			}
			
			sort.Slice(sortedTags, func(i, j int) bool {
				return sortedTags[i].count > sortedTags[j].count
			})

			// Create table for top tags
			t := table.New().
				Border(lipgloss.RoundedBorder()).
				BorderStyle(lipgloss.NewStyle().Foreground(ui.Border)).
				StyleFunc(func(row, col int) lipgloss.Style {
					if row == table.HeaderRow {
						return lipgloss.NewStyle().
							Foreground(ui.Primary).
							Bold(true).
							Align(lipgloss.Center).
							Padding(0, 1)
					}
					if col == 0 {
						return lipgloss.NewStyle().
							Foreground(ui.Text).
							Bold(true).
							Padding(0, 1).
							Width(20)
					}
					return lipgloss.NewStyle().
						Foreground(ui.TextMuted).
						Padding(0, 1).
						Align(lipgloss.Center).
						Width(10)
				}).
				Headers("Tag", "Count", "Usage %")

			// Show top 10 tags
			maxTags := 10
			if len(sortedTags) < maxTags {
				maxTags = len(sortedTags)
			}

			for i := 0; i < maxTags; i++ {
				tag := sortedTags[i]
				percentage := float64(tag.count) / float64(totalSnippets) * 100
				
				t.Row(
					ui.RenderTag(tag.tag),
					fmt.Sprintf("%d", tag.count),
					fmt.Sprintf("%.1f%%", percentage),
				)
			}

			fmt.Println(t.Render())
			fmt.Println()
		}

		// Show recent activity
		fmt.Println(ui.RenderSubtitle("⏰ Recent Activity"))
		fmt.Println()

		// Show last 5 snippets
		recentCount := 5
		if len(snippets) < recentCount {
			recentCount = len(snippets)
		}

		for i := 0; i < recentCount; i++ {
			snippet := snippets[i]
			
			// Format tags
			tagsStr := ""
			if len(snippet.Tags) > 0 {
				var formattedTags []string
				for _, tag := range snippet.Tags {
					formattedTags = append(formattedTags, ui.RenderTag(tag))
				}
				tagsStr = " " + strings.Join(formattedTags, " ")
			}

			// Format time
			timeStr := formatTimeAgo(snippet.CreatedAt)
			
			fmt.Printf("  %s %d. %s%s\n", 
				ui.IconSnippet, 
				snippet.ID, 
				lipgloss.NewStyle().Bold(true).Render(snippet.Title),
				tagsStr)
			fmt.Printf("     %s Created %s\n", 
				ui.IconTime, 
				lipgloss.NewStyle().Foreground(ui.TextMuted).Render(timeStr))
			
			if i < recentCount-1 {
				fmt.Println()
			}
		}

		fmt.Println()
		
		// Show helpful tips
		tips := `💡 Pro Tips:
  • Use 'snip search --tag=<tag>' to find snippets by tag
  • Use 'snip save-interactive' for a guided snippet creation
  • Use 'snip copy <id>' to quickly copy snippets to clipboard`
		
		fmt.Println(ui.RenderBox(tips))
	},
}

// Helper function for time formatting (reused from table.go)
func formatTimeAgo(t time.Time) string {
	// This function is already implemented in table.go
	// For now, let's use a simple implementation
	return t.Format("Jan 2, 2006")
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
