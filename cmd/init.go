package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lubasinkal/snip/internal/ui"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize snip configuration",
	Long:  `Set up the snip configuration directory and database. This is optional - snip will work without explicit initialization.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Show welcome header
		fmt.Println(ui.RenderTitle(ui.IconRocket + " Initializing snip"))
		fmt.Println()

		// Get home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(ui.RenderError("Error getting home directory: " + err.Error()))
			return
		}

		// Create .snipdb directory
		snipDir := filepath.Join(homeDir, ".snipdb")
		err = os.MkdirAll(snipDir, 0755)
		if err != nil {
			fmt.Println(ui.RenderError(fmt.Sprintf("Error creating directory %s: %s", snipDir, err.Error())))
			return
		}

		// Check if directory already existed
		if _, err := os.Stat(snipDir); err == nil {
			fmt.Println(ui.RenderInfo(fmt.Sprintf("%s Configuration directory already exists: %s", ui.IconFolder, snipDir)))
		} else {
			fmt.Println(ui.RenderSuccess(fmt.Sprintf("%s Created configuration directory: %s", ui.IconFolder, snipDir)))
		}

		// Database will be created automatically when first used
		dbPath := filepath.Join(snipDir, "snippets.db")
		fmt.Println(ui.RenderInfo(fmt.Sprintf("%s Database will be stored at: %s", ui.IconDatabase, dbPath)))
		fmt.Println()

		fmt.Println(ui.RenderSuccess(ui.IconSparkles + " snip is ready to use!"))
		fmt.Println()

		// Show next steps in a beautiful box
		nextSteps := `Next steps:
  • Save your first snippet: echo 'Hello World' | snip save 'My first snippet'
  • List all snippets: snip list
  • Search snippets: snip search 'hello'
  • View help: snip --help`

		fmt.Println(ui.RenderBox(nextSteps))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
