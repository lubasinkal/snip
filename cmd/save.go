package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/lubasinkal/snip/internal/models"
	"github.com/lubasinkal/snip/internal/storage"
	"github.com/lubasinkal/snip/internal/ui"
	"github.com/spf13/cobra"
)

var tags string

var saveCmd = &cobra.Command{
	Use:   "save [title]",
	Short: "Save a snippet from stdin",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		content, _ := io.ReadAll(os.Stdin)

		// Parse tags
		var tagList []string
		if tags != "" {
			tagList = strings.Split(tags, ",")
			// Trim whitespace from each tag
			for i, tag := range tagList {
				tagList[i] = strings.TrimSpace(tag)
			}
		}

		snippet := models.Snippet{
			Title:     title,
			Tags:      tagList,
			CreatedAt: time.Now(),
			Content:   string(content),
		}
		id, err := storage.SaveSnippet(snippet)
		if err != nil {
			fmt.Println(ui.RenderError("Error saving snippet: " + err.Error()))
			return
		}

		// Show success message with snippet details
		successMsg := fmt.Sprintf("Snippet saved with ID: %d", id)
		fmt.Println(ui.RenderSuccess(successMsg))

		// Show a preview of what was saved
		fmt.Println()
		fmt.Println(ui.RenderSnippetCard(models.Snippet{
			ID:        int(id),
			Title:     snippet.Title,
			Tags:      snippet.Tags,
			CreatedAt: snippet.CreatedAt,
			Content:   snippet.Content,
		}, false))
	},
}

func init() {
	saveCmd.Flags().StringVarP(&tags, "tags", "t", "", "Comma-separated tags")
	rootCmd.AddCommand(saveCmd)
}
