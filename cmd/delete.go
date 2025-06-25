package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lubasinkal/snip/internal/storage"
	"github.com/lubasinkal/snip/internal/ui"
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
			fmt.Println(ui.RenderError("Invalid snippet ID. Please provide a valid number."))
			return
		}

		// First, get the snippet to show what we're deleting
		snippet, err := storage.GetSnippetByID(id)
		if err != nil {
			fmt.Println(ui.RenderError(err.Error()))
			return
		}

		// Show confirmation unless --force is used
		if !forceDelete {
			// Show the snippet that will be deleted
			fmt.Println(ui.RenderWarning("You are about to delete this snippet:"))
			fmt.Println()
			fmt.Println(ui.RenderSnippetCard(*snippet, false))
			fmt.Println()

			fmt.Printf("%s Are you sure you want to delete this snippet? [y/N]: ", ui.IconDelete)

			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(ui.RenderError("Error reading input: " + err.Error()))
				return
			}

			response = strings.ToLower(strings.TrimSpace(response))
			if response != "y" && response != "yes" {
				fmt.Println(ui.RenderInfo("Deletion cancelled."))
				return
			}
		}

		// Delete the snippet
		err = storage.DeleteSnippet(id)
		if err != nil {
			fmt.Println(ui.RenderError("Error deleting snippet: " + err.Error()))
			return
		}

		successMsg := fmt.Sprintf("Deleted snippet '%s' (ID: %d)", snippet.Title, id)
		fmt.Println(ui.RenderSuccess(successMsg))
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&forceDelete, "force", "f", false, "Skip confirmation prompt")
	rootCmd.AddCommand(deleteCmd)
}
