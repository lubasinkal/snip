package cmd

import (
	"fmt"
	"github.com/lubasinkal/snip/internal/models"
	"github.com/lubasinkal/snip/internal/storage"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
	"time"
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
			fmt.Println("❌ Error saving snippet:", err)
			return
		}
		fmt.Println("✅ Snippet saved with ID:", id)
	},
}

func init() {
	saveCmd.Flags().StringVarP(&tags, "tags", "t", "", "Comma-separated tags")
	rootCmd.AddCommand(saveCmd)
}
