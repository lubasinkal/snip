package cmd

import (
	"fmt"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/lubasinkal/snip/internal/storage"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy [id]",
	Short: "Copy snippet content to clipboard",
	Long:  `Copy the content of a snippet to your system clipboard by its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("❌ Invalid snippet ID. Please provide a valid number.")
			return
		}

		snippet, err := storage.GetSnippetByID(id)
		if err != nil {
			fmt.Printf("❌ %s\n", err.Error())
			return
		}

		err = clipboard.WriteAll(snippet.Content)
		if err != nil {
			fmt.Println("❌ Error copying to clipboard:", err)
			return
		}

		fmt.Printf("📋 Copied snippet '%s' to clipboard!\n", snippet.Title)
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
