package cmd

import (
	"fmt"
	"strconv"

	"github.com/lubasinkal/snip/internal/storage"
	"github.com/lubasinkal/snip/internal/ui"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat [id]",
	Short: "Print snippet content to stdout",
	Long:  `Display the content of a snippet by its ID. Perfect for piping to other commands.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(ui.RenderError("Invalid snippet ID. Please provide a valid number."))
			return
		}

		snippet, err := storage.GetSnippetByID(id)
		if err != nil {
			fmt.Println(ui.RenderError(err.Error()))
			return
		}

		// Just print the content - no extra formatting for piping
		fmt.Print(snippet.Content)
	},
}

func init() {
	rootCmd.AddCommand(catCmd)
}
