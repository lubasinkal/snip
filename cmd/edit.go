package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/lubasinkal/snip/internal/storage"
	"github.com/lubasinkal/snip/internal/ui"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [id]",
	Short: "Edit a snippet in your default editor",
	Long:  `Open a snippet in your default editor ($EDITOR) and save changes back to the database.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(ui.RenderError("Invalid snippet ID. Please provide a valid number."))
			return
		}

		// Get the snippet
		snippet, err := storage.GetSnippetByID(id)
		if err != nil {
			fmt.Println(ui.RenderError(err.Error()))
			return
		}

		// Get editor from environment
		editor := os.Getenv("EDITOR")
		if editor == "" {
			// Default editors by platform
			if _, err := exec.LookPath("code"); err == nil {
				editor = "code"
			} else if _, err := exec.LookPath("nano"); err == nil {
				editor = "nano"
			} else if _, err := exec.LookPath("vim"); err == nil {
				editor = "vim"
			} else if _, err := exec.LookPath("notepad"); err == nil {
				editor = "notepad"
			} else {
				fmt.Println(ui.RenderError("No editor found. Please set the EDITOR environment variable."))
				return
			}
		}

		// Create temporary file
		tmpFile, err := ioutil.TempFile("", fmt.Sprintf("snip_%d_*.txt", id))
		if err != nil {
			fmt.Println(ui.RenderError("Error creating temporary file: " + err.Error()))
			return
		}
		defer os.Remove(tmpFile.Name())

		// Write current content to temp file
		_, err = tmpFile.WriteString(snippet.Content)
		if err != nil {
			fmt.Println(ui.RenderError("Error writing to temporary file: " + err.Error()))
			return
		}
		tmpFile.Close()

		// Show what we're editing
		fmt.Println(ui.RenderInfo(fmt.Sprintf("%s Opening snippet '%s' in %s...", ui.IconEdit, snippet.Title, editor)))
		fmt.Println()
		fmt.Println(ui.RenderSnippetCard(*snippet, false))
		
		var editorCmd *exec.Cmd
		if editor == "code" {
			editorCmd = exec.Command(editor, "--wait", tmpFile.Name())
		} else {
			editorCmd = exec.Command(editor, tmpFile.Name())
		}
		
		editorCmd.Stdin = os.Stdin
		editorCmd.Stdout = os.Stdout
		editorCmd.Stderr = os.Stderr

		err = editorCmd.Run()
		if err != nil {
			fmt.Println(ui.RenderError("Error running editor: " + err.Error()))
			return
		}

		// Read the modified content
		modifiedContent, err := ioutil.ReadFile(tmpFile.Name())
		if err != nil {
			fmt.Println(ui.RenderError("Error reading modified file: " + err.Error()))
			return
		}

		// Check if content changed
		newContent := string(modifiedContent)
		if newContent == snippet.Content {
			fmt.Println(ui.RenderInfo("No changes made."))
			return
		}

		// Update the snippet
		snippet.Content = strings.TrimRight(newContent, "\n\r")
		err = storage.UpdateSnippet(*snippet)
		if err != nil {
			fmt.Println(ui.RenderError("Error saving changes: " + err.Error()))
			return
		}

		fmt.Println(ui.RenderSuccess(fmt.Sprintf("Updated snippet '%s'", snippet.Title)))
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
