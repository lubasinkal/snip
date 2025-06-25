package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lubasinkal/snip/internal/storage"
	"github.com/spf13/cobra"
)

var forceDelete bool

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a snippet",
	Long:  `Remove a snippet from your collection by its ID. Use --force to skip confirmation.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("❌ Invalid snippet ID. Please provide a valid number.")
			return
		}

		// First, get the snippet to show what we're deleting
		snippet, err := storage.GetSnippetByID(id)
		if err != nil {
			fmt.Printf("❌ %s\n", err.Error())
			return
		}

		// Show confirmation unless --force is used
		if !forceDelete {
			fmt.Printf("🗑️  Are you sure you want to delete snippet '%s' (ID: %d)? [y/N]: ", snippet.Title, id)
			
			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("❌ Error reading input:", err)
				return
			}
			
			response = strings.ToLower(strings.TrimSpace(response))
			if response != "y" && response != "yes" {
				fmt.Println("❌ Deletion cancelled.")
				return
			}
		}

		// Delete the snippet
		err = storage.DeleteSnippet(id)
		if err != nil {
			fmt.Printf("❌ Error deleting snippet: %s\n", err.Error())
			return
		}

		fmt.Printf("✅ Deleted snippet '%s' (ID: %d)\n", snippet.Title, id)
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&forceDelete, "force", "f", false, "Skip confirmation prompt")
	rootCmd.AddCommand(deleteCmd)
}
