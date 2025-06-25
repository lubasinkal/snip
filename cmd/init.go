package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize snip configuration",
	Long:  `Set up the snip configuration directory and database. This is optional - snip will work without explicit initialization.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("❌ Error getting home directory:", err)
			return
		}

		// Create .snipdb directory
		snipDir := filepath.Join(homeDir, ".snipdb")
		err = os.MkdirAll(snipDir, 0755)
		if err != nil {
			fmt.Printf("❌ Error creating directory %s: %s\n", snipDir, err)
			return
		}

		// Check if directory already existed
		if _, err := os.Stat(snipDir); err == nil {
			fmt.Printf("📁 Configuration directory already exists: %s\n", snipDir)
		} else {
			fmt.Printf("📁 Created configuration directory: %s\n", snipDir)
		}

		// Database will be created automatically when first used
		dbPath := filepath.Join(snipDir, "snippets.db")
		fmt.Printf("💾 Database will be stored at: %s\n", dbPath)

		fmt.Println("✅ snip is ready to use!")
		fmt.Println()
		fmt.Println("Next steps:")
		fmt.Println("  • Save your first snippet: echo 'Hello World' | snip save 'My first snippet'")
		fmt.Println("  • List all snippets: snip list")
		fmt.Println("  • Search snippets: snip search 'hello'")
		fmt.Println("  • View help: snip --help")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
