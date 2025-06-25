package cmd

import (
    "fmt"
    "os"

    "github.com/lubasinkal/snip/internal/ui"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "snip",
    Short: "snip is a fast CLI code snippet manager",
    Long: ui.RenderTitle("🚀 snip — A Terminal Code Snippet Manager") + "\n\n" +
        ui.RenderSubtitle("A superfast command-line tool to save, search, view, and reuse your code snippets — all from your terminal.") + "\n\n" +
        ui.RenderBox(`✨ Features:
  • 🚀 Lightning fast - Built in Go for speed
  • 💾 Local storage - Your snippets stay on your machine
  • 🔍 Powerful search - Search by title, content, or tags
  • 🏷️ Tag support - Organize snippets with multiple tags
  • 📋 Clipboard integration - Copy snippets directly to clipboard
  • ✏️ Edit in place - Open snippets in your favorite editor
  • 🎯 Simple CLI - Intuitive commands that just work`),
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
