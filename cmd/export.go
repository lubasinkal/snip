package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lubasinkal/snip/internal/models"
	"github.com/lubasinkal/snip/internal/storage"
	"github.com/lubasinkal/snip/internal/ui"
	"github.com/spf13/cobra"
)

var (
	exportFormat string
	exportOutput string
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export snippets to various formats",
	Long:  `Export your snippets to JSON, Markdown, or plain text format for backup or sharing.`,
	Run: func(cmd *cobra.Command, args []string) {
		snippets, err := storage.ListAllSnippets()
		if err != nil {
			fmt.Println(ui.RenderError("Error loading snippets: " + err.Error()))
			return
		}

		if len(snippets) == 0 {
			fmt.Println(ui.RenderInfo("No snippets found to export."))
			return
		}

		var content string
		var fileExt string

		switch exportFormat {
		case "json":
			content, err = exportToJSON(snippets)
			fileExt = ".json"
		case "markdown", "md":
			content, err = exportToMarkdown(snippets)
			fileExt = ".md"
		case "text", "txt":
			content, err = exportToText(snippets)
			fileExt = ".txt"
		default:
			fmt.Println(ui.RenderError("Unsupported format. Use: json, markdown, or text"))
			return
		}

		if err != nil {
			fmt.Println(ui.RenderError("Error generating export: " + err.Error()))
			return
		}

		// Determine output file
		outputFile := exportOutput
		if outputFile == "" {
			timestamp := time.Now().Format("20060102_150405")
			outputFile = fmt.Sprintf("snip_export_%s%s", timestamp, fileExt)
		}

		// Write to file
		err = os.WriteFile(outputFile, []byte(content), 0644)
		if err != nil {
			fmt.Println(ui.RenderError("Error writing file: " + err.Error()))
			return
		}

		// Show success
		absPath, _ := filepath.Abs(outputFile)
		fmt.Println(ui.RenderSuccess(fmt.Sprintf("Exported %d snippets to: %s", len(snippets), absPath)))
		
		// Show file info
		fileInfo, _ := os.Stat(outputFile)
		infoText := fmt.Sprintf(`📄 Export Details:
  • Format: %s
  • File: %s
  • Size: %d bytes
  • Snippets: %d`,
			strings.ToUpper(exportFormat),
			filepath.Base(outputFile),
			fileInfo.Size(),
			len(snippets))
		
		fmt.Println()
		fmt.Println(ui.RenderBox(infoText))
	},
}

func exportToJSON(snippets []models.Snippet) (string, error) {
	type ExportData struct {
		ExportedAt time.Time `json:"exported_at"`
		Version    string    `json:"version"`
		Count      int       `json:"count"`
		Snippets   []models.Snippet `json:"snippets"`
	}

	data := ExportData{
		ExportedAt: time.Now(),
		Version:    "1.0",
		Count:      len(snippets),
		Snippets:   snippets,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func exportToMarkdown(snippets []models.Snippet) (string, error) {
	var content strings.Builder
	
	content.WriteString("# Code Snippets Export\n\n")
	content.WriteString(fmt.Sprintf("Exported on: %s\n", time.Now().Format("January 2, 2006 at 3:04 PM")))
	content.WriteString(fmt.Sprintf("Total snippets: %d\n\n", len(snippets)))
	content.WriteString("---\n\n")

	for i, snippet := range snippets {
		content.WriteString(fmt.Sprintf("## %d. %s\n\n", snippet.ID, snippet.Title))
		
		if len(snippet.Tags) > 0 {
			content.WriteString("**Tags:** ")
			for j, tag := range snippet.Tags {
				if j > 0 {
					content.WriteString(", ")
				}
				content.WriteString(fmt.Sprintf("`%s`", tag))
			}
			content.WriteString("\n\n")
		}
		
		content.WriteString(fmt.Sprintf("**Created:** %s\n\n", snippet.CreatedAt.Format("January 2, 2006")))
		
		content.WriteString("```\n")
		content.WriteString(snippet.Content)
		content.WriteString("\n```\n\n")
		
		if i < len(snippets)-1 {
			content.WriteString("---\n\n")
		}
	}

	return content.String(), nil
}

func exportToText(snippets []models.Snippet) (string, error) {
	var content strings.Builder
	
	content.WriteString("CODE SNIPPETS EXPORT\n")
	content.WriteString("===================\n\n")
	content.WriteString(fmt.Sprintf("Exported on: %s\n", time.Now().Format("January 2, 2006 at 3:04 PM")))
	content.WriteString(fmt.Sprintf("Total snippets: %d\n\n", len(snippets)))

	for i, snippet := range snippets {
		content.WriteString(fmt.Sprintf("[%d] %s\n", snippet.ID, snippet.Title))
		content.WriteString(strings.Repeat("-", len(snippet.Title)+10) + "\n")
		
		if len(snippet.Tags) > 0 {
			content.WriteString(fmt.Sprintf("Tags: %s\n", strings.Join(snippet.Tags, ", ")))
		}
		
		content.WriteString(fmt.Sprintf("Created: %s\n\n", snippet.CreatedAt.Format("January 2, 2006")))
		content.WriteString(snippet.Content)
		content.WriteString("\n\n")
		
		if i < len(snippets)-1 {
			content.WriteString(strings.Repeat("=", 50) + "\n\n")
		}
	}

	return content.String(), nil
}

func init() {
	exportCmd.Flags().StringVarP(&exportFormat, "format", "f", "json", "Export format (json, markdown, text)")
	exportCmd.Flags().StringVarP(&exportOutput, "output", "o", "", "Output file path (default: auto-generated)")
	rootCmd.AddCommand(exportCmd)
}
