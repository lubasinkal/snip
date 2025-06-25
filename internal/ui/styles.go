package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// Color palette
var (
	// Primary colors
	Primary   = lipgloss.Color("#7C3AED") // Purple
	Secondary = lipgloss.Color("#06B6D4") // Cyan
	Success   = lipgloss.Color("#10B981") // Green
	Warning   = lipgloss.Color("#F59E0B") // Yellow
	Error     = lipgloss.Color("#EF4444") // Red
	
	// Neutral colors
	Text      = lipgloss.AdaptiveColor{Light: "#1F2937", Dark: "#F9FAFB"}
	TextMuted = lipgloss.AdaptiveColor{Light: "#6B7280", Dark: "#9CA3AF"}
	Border    = lipgloss.AdaptiveColor{Light: "#E5E7EB", Dark: "#374151"}
	
	// Background colors
	Background = lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#111827"}
	Surface    = lipgloss.AdaptiveColor{Light: "#F9FAFB", Dark: "#1F2937"}
)

// Base styles
var (
	// Text styles
	TitleStyle = lipgloss.NewStyle().
		Foreground(Primary).
		Bold(true).
		MarginBottom(1)
	
	SubtitleStyle = lipgloss.NewStyle().
		Foreground(TextMuted).
		Italic(true)
	
	BodyStyle = lipgloss.NewStyle().
		Foreground(Text)
	
	// Status styles
	SuccessStyle = lipgloss.NewStyle().
		Foreground(Success).
		Bold(true)
	
	ErrorStyle = lipgloss.NewStyle().
		Foreground(Error).
		Bold(true)
	
	WarningStyle = lipgloss.NewStyle().
		Foreground(Warning).
		Bold(true)
	
	InfoStyle = lipgloss.NewStyle().
		Foreground(Secondary).
		Bold(true)
	
	// Container styles
	BoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Border).
		Padding(1, 2).
		MarginBottom(1)
	
	HighlightBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(Primary).
		Padding(1, 2).
		MarginBottom(1)
	
	// List styles
	ListItemStyle = lipgloss.NewStyle().
		PaddingLeft(2).
		MarginBottom(0)
	
	ListNumberStyle = lipgloss.NewStyle().
		Foreground(Primary).
		Bold(true).
		Width(4).
		Align(lipgloss.Right)
	
	// Tag styles
	TagStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(Secondary).
		Padding(0, 1).
		MarginRight(1).
		Bold(true)
	
	// Code styles
	CodeStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E11D48")).
		Background(Surface).
		Padding(0, 1)
	
	CodeBlockStyle = lipgloss.NewStyle().
		Background(Surface).
		Border(lipgloss.NormalBorder()).
		BorderForeground(Border).
		Padding(1).
		MarginTop(1).
		MarginBottom(1)
)

// Icon constants
const (
	IconSuccess   = "✅"
	IconError     = "❌"
	IconWarning   = "⚠️"
	IconInfo      = "ℹ️"
	IconSnippet   = "📝"
	IconSearch    = "🔍"
	IconList      = "📚"
	IconCopy      = "📋"
	IconEdit      = "✏️"
	IconDelete    = "🗑️"
	IconTag       = "🏷️"
	IconTime      = "⏰"
	IconFolder    = "📁"
	IconDatabase  = "💾"
	IconRocket    = "🚀"
	IconSparkles  = "✨"
)

// Helper functions for common UI patterns
func RenderTitle(text string) string {
	return TitleStyle.Render(text)
}

func RenderSubtitle(text string) string {
	return SubtitleStyle.Render(text)
}

func RenderSuccess(text string) string {
	return SuccessStyle.Render(IconSuccess + " " + text)
}

func RenderError(text string) string {
	return ErrorStyle.Render(IconError + " " + text)
}

func RenderWarning(text string) string {
	return WarningStyle.Render(IconWarning + " " + text)
}

func RenderInfo(text string) string {
	return InfoStyle.Render(IconInfo + " " + text)
}

func RenderTag(text string) string {
	return TagStyle.Render(text)
}

func RenderCode(text string) string {
	return CodeStyle.Render(text)
}

func RenderBox(content string) string {
	return BoxStyle.Render(content)
}

func RenderHighlightBox(content string) string {
	return HighlightBoxStyle.Render(content)
}
