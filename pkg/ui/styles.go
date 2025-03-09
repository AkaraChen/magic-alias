package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// Colors
const (
	ColorPrimary   = "#FF5F87" // Pink
	ColorSuccess   = "#5AF78E" // Green
	ColorWarning   = "#FFAF00" // Yellow/Orange
	ColorError     = "#FF5F87" // Same as primary for errors
	ColorInfo      = "#5AF78E" // Same as success for info
)

// Emoji icons
const (
	IconSuccess = "✓"
	IconSparkle = "✨"
	IconWarning = "⚠"
)

// Title style for section headers
var TitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorPrimary)).
	Bold(true).
	Margin(1, 0)

// Success style for success messages
var SuccessStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorSuccess)).
	Bold(true)

// Error style for error messages
var ErrorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorError)).
	Bold(true)

// Warning style for warning messages
var WarningStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorWarning)).
	Bold(true).
	Margin(1, 0)

// Info style for general information
var InfoStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorInfo)).
	Margin(0, 2)

// Empty style for when no items are found
var EmptyStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorError)).
	Italic(true).
	Margin(1, 0)

// Success formats a message with success styling and icon
func Success(message string) string {
	return SuccessStyle.Render(IconSuccess + " " + message)
}

// Error formats a message with error styling
func Error(message string) string {
	return ErrorStyle.Render(message)
}

// Warning formats a message with warning styling and icon
func Warning(message string) string {
	return WarningStyle.Render(IconWarning + " " + message)
}

// Title formats a message with title styling and sparkle icon
func Title(message string) string {
	return TitleStyle.Render(IconSparkle + " " + message)
}

// Info formats a message with info styling
func Info(message string) string {
	return InfoStyle.Render(message)
}

// Empty formats a message for empty results
func Empty(message string) string {
	return EmptyStyle.Render(message)
}
