package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/lubasinkal/snip/internal/models"
	"github.com/lubasinkal/snip/internal/storage"
	"github.com/lubasinkal/snip/internal/ui"
	"github.com/spf13/cobra"
)

var saveInteractiveCmd = &cobra.Command{
	Use:   "save-interactive",
	Short: "Save a snippet using an interactive form",
	Long:  `Launch an interactive form to save a code snippet with title, content, and tags.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			title       string
			content     string
			tagsInput   string
			language    string
			description string
		)

		// Create the interactive form
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("What's the title of your snippet?").
					Placeholder("e.g., 'HTTP Server in Go'").
					Value(&title).
					Validate(func(str string) error {
						if strings.TrimSpace(str) == "" {
							return fmt.Errorf("title cannot be empty")
						}
						return nil
					}),

				huh.NewSelect[string]().
					Title("What programming language is this?").
					Options(
						huh.NewOption("Go", "go"),
						huh.NewOption("JavaScript", "javascript"),
						huh.NewOption("TypeScript", "typescript"),
						huh.NewOption("Python", "python"),
						huh.NewOption("Rust", "rust"),
						huh.NewOption("Java", "java"),
						huh.NewOption("C++", "cpp"),
						huh.NewOption("C#", "csharp"),
						huh.NewOption("PHP", "php"),
						huh.NewOption("Ruby", "ruby"),
						huh.NewOption("Shell/Bash", "bash"),
						huh.NewOption("SQL", "sql"),
						huh.NewOption("HTML", "html"),
						huh.NewOption("CSS", "css"),
						huh.NewOption("Other", "other"),
					).
					Value(&language),

				huh.NewText().
					Title("Enter your code snippet:").
					Placeholder("Paste or type your code here...").
					CharLimit(5000).
					Value(&content).
					Validate(func(str string) error {
						if strings.TrimSpace(str) == "" {
							return fmt.Errorf("content cannot be empty")
						}
						return nil
					}),
			),

			huh.NewGroup(
				huh.NewText().
					Title("Description (optional):").
					Placeholder("Brief description of what this code does...").
					CharLimit(200).
					Value(&description),

				huh.NewInput().
					Title("Tags (optional):").
					Placeholder("e.g., 'web, api, server' (comma-separated)").
					Value(&tagsInput),
			),
		)

		// Run the form
		err := form.Run()
		if err != nil {
			fmt.Println(ui.RenderError("Error running form: " + err.Error()))
			return
		}

		// Process the input
		var tags []string
		if tagsInput != "" {
			tagList := strings.Split(tagsInput, ",")
			for _, tag := range tagList {
				trimmed := strings.TrimSpace(tag)
				if trimmed != "" {
					tags = append(tags, trimmed)
				}
			}
		}

		// Add language as a tag if specified
		if language != "" && language != "other" {
			tags = append([]string{language}, tags...)
		}

		// Add description to content if provided
		finalContent := content
		if description != "" {
			finalContent = fmt.Sprintf("// %s\n%s", description, content)
		}

		// Create and save the snippet
		snippet := models.Snippet{
			Title:     title,
			Tags:      tags,
			CreatedAt: time.Now(),
			Content:   finalContent,
		}

		id, err := storage.SaveSnippet(snippet)
		if err != nil {
			fmt.Println(ui.RenderError("Error saving snippet: " + err.Error()))
			return
		}

		// Show success message
		fmt.Println()
		fmt.Println(ui.RenderSuccess(fmt.Sprintf("Snippet saved with ID: %d", id)))
		fmt.Println()

		// Show the saved snippet
		fmt.Println(ui.RenderSnippetCard(models.Snippet{
			ID:        int(id),
			Title:     snippet.Title,
			Tags:      snippet.Tags,
			CreatedAt: snippet.CreatedAt,
			Content:   snippet.Content,
		}, true))
	},
}

func init() {
	rootCmd.AddCommand(saveInteractiveCmd)
}
