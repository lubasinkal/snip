package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/lubasinkal/snip/internal/models"
)

// Table styles
var (
	headerStyle = lipgloss.NewStyle().
		Foreground(Primary).
		Bold(true).
		Align(lipgloss.Center).
		Padding(0, 1)
	
	cellStyle = lipgloss.NewStyle().
		Padding(0, 1).
		Align(lipgloss.Left)
	
	idCellStyle = lipgloss.NewStyle().
		Foreground(Primary).
		Bold(true).
		Padding(0, 1).
		Align(lipgloss.Center).
		Width(4)
	
	titleCellStyle = lipgloss.NewStyle().
		Foreground(Text).
		Bold(true).
		Padding(0, 1).
		Width(30)
	
	tagsCellStyle = lipgloss.NewStyle().
		Foreground(TextMuted).
		Padding(0, 1).
		Width(25)
	
	timeCellStyle = lipgloss.NewStyle().
		Foreground(TextMuted).
		Padding(0, 1).
		Width(15)
)

// RenderSnippetsTable creates a beautiful table for displaying snippets
func RenderSnippetsTable(snippets []models.Snippet) string {
	if len(snippets) == 0 {
		return RenderInfo("No snippets found. Use 'snip save' to create your first snippet!")
	}

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(Border)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case col == 0: // ID column
				return idCellStyle
			case col == 1: // Title column
				return titleCellStyle
			case col == 2: // Tags column
				return tagsCellStyle
			case col == 3: // Time column
				return timeCellStyle
			default:
				return cellStyle
			}
		}).
		Headers("ID", "Title", "Tags", "Created")

	// Add rows
	for _, snippet := range snippets {
		// Format tags
		tagsStr := ""
		if len(snippet.Tags) > 0 {
			var formattedTags []string
			for _, tag := range snippet.Tags {
				formattedTags = append(formattedTags, RenderTag(tag))
			}
			tagsStr = strings.Join(formattedTags, " ")
		} else {
			tagsStr = lipgloss.NewStyle().Foreground(TextMuted).Render("no tags")
		}

		// Format time
		timeStr := formatTimeAgo(snippet.CreatedAt)

		// Truncate title if too long
		title := snippet.Title
		if len(title) > 28 {
			title = title[:25] + "..."
		}

		t.Row(
			fmt.Sprintf("%d", snippet.ID),
			title,
			tagsStr,
			timeStr,
		)
	}

	return t.Render()
}

// RenderSnippetCard creates a detailed card view for a single snippet
func RenderSnippetCard(snippet models.Snippet, showContent bool) string {
	var content strings.Builder
	
	// Header with ID and title
	header := fmt.Sprintf("%s %d: %s", IconSnippet, snippet.ID, snippet.Title)
	content.WriteString(TitleStyle.Render(header))
	content.WriteString("\n")
	
	// Tags
	if len(snippet.Tags) > 0 {
		content.WriteString(lipgloss.NewStyle().Foreground(TextMuted).Render(IconTag + " Tags: "))
		for i, tag := range snippet.Tags {
			if i > 0 {
				content.WriteString(" ")
			}
			content.WriteString(RenderTag(tag))
		}
		content.WriteString("\n")
	}

	// Created time
	timeStr := formatTimeAgo(snippet.CreatedAt)
	content.WriteString(lipgloss.NewStyle().Foreground(TextMuted).Render(IconTime + " Created: " + timeStr))
	content.WriteString("\n")

	// Content if requested
	if showContent {
		content.WriteString("\n")
		content.WriteString(lipgloss.NewStyle().Foreground(TextMuted).Render("Content:"))
		content.WriteString("\n")
		content.WriteString(CodeBlockStyle.Render(snippet.Content))
	}
	
	return HighlightBoxStyle.Render(content.String())
}

// RenderSearchResults creates a formatted display for search results
func RenderSearchResults(snippets []models.Snippet, query string, tagFilter string) string {
	var content strings.Builder
	
	// Header
	if tagFilter != "" {
		header := fmt.Sprintf("%s Found %d snippet(s) matching '%s' with tag '%s':", 
			IconSearch, len(snippets), query, tagFilter)
		content.WriteString(InfoStyle.Render(header))
	} else {
		header := fmt.Sprintf("%s Found %d snippet(s) matching '%s':", 
			IconSearch, len(snippets), query)
		content.WriteString(InfoStyle.Render(header))
	}
	content.WriteString("\n\n")
	
	if len(snippets) == 0 {
		if tagFilter != "" {
			content.WriteString(lipgloss.NewStyle().Foreground(TextMuted).Render(fmt.Sprintf("No snippets found matching '%s' with tag '%s'", query, tagFilter)))
		} else {
			content.WriteString(lipgloss.NewStyle().Foreground(TextMuted).Render(fmt.Sprintf("No snippets found matching '%s'", query)))
		}
		return content.String()
	}
	
	// Results
	for i, snippet := range snippets {
		if i > 0 {
			content.WriteString("\n")
		}
		
		// Snippet header
		header := fmt.Sprintf("%d. %s", snippet.ID, snippet.Title)
		content.WriteString(BodyStyle.Bold(true).Render(header))
		
		// Tags
		if len(snippet.Tags) > 0 {
			content.WriteString(" ")
			for _, tag := range snippet.Tags {
				content.WriteString(RenderTag(tag))
				content.WriteString(" ")
			}
		}
		content.WriteString("\n")
		
		// Time
		timeStr := formatTimeAgo(snippet.CreatedAt)
		content.WriteString(lipgloss.NewStyle().Foreground(TextMuted).Render("     Created " + timeStr))
		content.WriteString("\n")

		// Preview
		preview := strings.ReplaceAll(snippet.Content, "\n", " ")
		if len(preview) > 80 {
			preview = preview[:77] + "..."
		}
		if preview != "" {
			content.WriteString(lipgloss.NewStyle().Foreground(TextMuted).Render("     Preview: "))
			content.WriteString(CodeStyle.Render(preview))
		}
		content.WriteString("\n")
	}
	
	return content.String()
}

// formatTimeAgo formats a time as a human-readable "time ago" string
func formatTimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)
	
	if diff < time.Minute {
		return "just now"
	} else if diff < time.Hour {
		minutes := int(diff.Minutes())
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	} else if diff < 24*time.Hour {
		hours := int(diff.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	} else if diff < 7*24*time.Hour {
		days := int(diff.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	} else {
		return t.Format("Jan 2, 2006")
	}
}
