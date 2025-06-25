package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/lubasinkal/snip/internal/models"
	"github.com/lubasinkal/snip/internal/storage"
	"github.com/lubasinkal/snip/internal/ui"
	"github.com/spf13/cobra"
)

var (
	importFile   string
	importFormat string
	skipConfirm  bool
)

var importCmd = &cobra.Command{
	Use:   "import [file]",
	Short: "Import snippets from a file",
	Long:  `Import snippets from a JSON export file. Use this to restore backups or migrate snippets.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Determine input file
		inputFile := importFile
		if len(args) > 0 {
			inputFile = args[0]
		}

		if inputFile == "" {
			fmt.Println(ui.RenderError("Please specify a file to import using --file or as an argument"))
			return
		}

		// Check if file exists
		if _, err := os.Stat(inputFile); os.IsNotExist(err) {
			fmt.Println(ui.RenderError(fmt.Sprintf("File not found: %s", inputFile)))
			return
		}

		// Read file
		data, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Println(ui.RenderError("Error reading file: " + err.Error()))
			return
		}

		// Parse JSON
		var importData struct {
			ExportedAt time.Time        `json:"exported_at"`
			Version    string           `json:"version"`
			Count      int              `json:"count"`
			Snippets   []models.Snippet `json:"snippets"`
		}

		err = json.Unmarshal(data, &importData)
		if err != nil {
			fmt.Println(ui.RenderError("Error parsing JSON: " + err.Error()))
			return
		}

		if len(importData.Snippets) == 0 {
			fmt.Println(ui.RenderInfo("No snippets found in the import file."))
			return
		}

		// Show import preview
		fmt.Println(ui.RenderTitle(ui.IconDatabase + " Import Preview"))
		fmt.Println()

		absPath, _ := filepath.Abs(inputFile)
		previewInfo := fmt.Sprintf(`📄 Import Details:
  • File: %s
  • Exported: %s
  • Version: %s
  • Snippets to import: %d`,
			filepath.Base(absPath),
			importData.ExportedAt.Format("January 2, 2006"),
			importData.Version,
			len(importData.Snippets))

		fmt.Println(ui.RenderBox(previewInfo))
		fmt.Println()

		// Show first few snippets as preview
		fmt.Println(ui.RenderSubtitle("📝 Preview of snippets:"))
		fmt.Println()

		previewCount := 3
		if len(importData.Snippets) < previewCount {
			previewCount = len(importData.Snippets)
		}

		for i := 0; i < previewCount; i++ {
			snippet := importData.Snippets[i]
			
			// Format tags
			tagsStr := ""
			if len(snippet.Tags) > 0 {
				var formattedTags []string
				for _, tag := range snippet.Tags {
					formattedTags = append(formattedTags, ui.RenderTag(tag))
				}
				tagsStr = " " + strings.Join(formattedTags, " ")
			}

			fmt.Printf("  %s %s%s\n", 
				ui.IconSnippet, 
				snippet.Title,
				tagsStr)
		}

		if len(importData.Snippets) > previewCount {
			fmt.Printf("  ... and %d more snippets\n", len(importData.Snippets)-previewCount)
		}
		fmt.Println()

		// Confirmation
		if !skipConfirm {
			var confirm bool
			
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewConfirm().
						Title(fmt.Sprintf("Import %d snippets?", len(importData.Snippets))).
						Description("This will add the snippets to your collection. Existing snippets will not be affected.").
						Affirmative("Yes, import them!").
						Negative("Cancel").
						Value(&confirm),
				),
			)

			err := form.Run()
			if err != nil {
				fmt.Println(ui.RenderError("Error running confirmation: " + err.Error()))
				return
			}

			if !confirm {
				fmt.Println(ui.RenderInfo("Import cancelled."))
				return
			}
		}

		// Import snippets
		fmt.Println(ui.RenderInfo("Importing snippets..."))
		fmt.Println()

		imported := 0
		failed := 0

		for _, snippet := range importData.Snippets {
			// Create new snippet (without ID to get auto-generated ID)
			newSnippet := models.Snippet{
				Title:     snippet.Title,
				Tags:      snippet.Tags,
				CreatedAt: time.Now(), // Use current time for imported snippets
				Content:   snippet.Content,
			}

			_, err := storage.SaveSnippet(newSnippet)
			if err != nil {
				fmt.Printf("  %s Failed to import: %s\n", ui.IconError, snippet.Title)
				failed++
			} else {
				fmt.Printf("  %s Imported: %s\n", ui.IconSuccess, snippet.Title)
				imported++
			}
		}

		fmt.Println()

		// Show results
		if failed == 0 {
			fmt.Println(ui.RenderSuccess(fmt.Sprintf("Successfully imported all %d snippets!", imported)))
		} else {
			fmt.Println(ui.RenderWarning(fmt.Sprintf("Imported %d snippets, %d failed", imported, failed)))
		}

		// Show next steps
		if imported > 0 {
			nextSteps := `🎉 Import completed!
  • Use 'snip list' to see all your snippets
  • Use 'snip stats' to see updated statistics
  • Use 'snip search' to find specific snippets`
			
			fmt.Println()
			fmt.Println(ui.RenderBox(nextSteps))
		}
	},
}

func init() {
	importCmd.Flags().StringVarP(&importFile, "file", "f", "", "File to import from")
	importCmd.Flags().BoolVar(&skipConfirm, "yes", false, "Skip confirmation prompt")
	rootCmd.AddCommand(importCmd)
}
