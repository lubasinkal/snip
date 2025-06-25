package cmd

import (
	"fmt"
	"runtime"

	"github.com/lubasinkal/snip/internal/ui"
	"github.com/spf13/cobra"
)

const (
	Version   = "1.0.0"
	BuildDate = "2025-06-25"
	GitCommit = "main"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  `Display version information about snip including build details and system info.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Show beautiful version header
		fmt.Println(ui.RenderTitle(ui.IconRocket + " snip — Terminal Code Snippet Manager"))
		fmt.Println()

		// Version info
		versionInfo := fmt.Sprintf(`🏷️ Version Information:
  • Version: %s
  • Build Date: %s
  • Git Commit: %s
  • Go Version: %s
  • Platform: %s/%s`,
			Version,
			BuildDate,
			GitCommit,
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH)

		fmt.Println(ui.RenderBox(versionInfo))
		fmt.Println()

		// Features
		features := `✨ Features:
  • 🚀 Lightning fast - Built in Go for speed
  • 💾 Local storage - Your snippets stay on your machine
  • 🔍 Powerful search - Search by title, content, or tags
  • 🏷️ Tag support - Organize snippets with multiple tags
  • 📋 Clipboard integration - Copy snippets directly to clipboard
  • ✏️ Edit in place - Open snippets in your favorite editor
  • 🎯 Simple CLI - Intuitive commands that just work
  • 📊 Statistics - Beautiful analytics about your snippets
  • 📤 Export/Import - Backup and restore your collection
  • 🎨 Beautiful UI - Powered by Charm's Lipgloss & Huh`

		fmt.Println(ui.RenderBox(features))
		fmt.Println()

		// Links and info
		links := `🔗 Links:
  • GitHub: https://github.com/lubasinkal/snip
  • Documentation: https://github.com/lubasinkal/snip#readme
  • Issues: https://github.com/lubasinkal/snip/issues
  • License: MIT`

		fmt.Println(ui.RenderBox(links))
		fmt.Println()

		// Credits
		fmt.Println(ui.RenderSubtitle("🙏 Built with amazing open source libraries:"))
		fmt.Println()

		credits := `  • Cobra - CLI framework by spf13
  • Lipgloss - Terminal styling by Charm
  • Huh - Terminal forms by Charm
  • SQLite - Database by modernc.org
  • Clipboard - Cross-platform clipboard by atotto`

		fmt.Println(ui.BodyStyle.Render(credits))
		fmt.Println()

		fmt.Println(ui.RenderSuccess("Thank you for using snip! " + ui.IconSparkles))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
