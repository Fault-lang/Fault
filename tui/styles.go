package tui

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// Initialize with adaptive theme
func init() {
	theme := DefaultTheme()
	ApplyTheme(theme)
}

var (
	// Title bar style
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF00FF")).
			Background(lipgloss.Color("#1a1a1a")).
			Padding(0, 1)

	// Subtitle/header style
	SubtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFFF")).
			Bold(true)

	// Phase status styles
	PhaseActiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#00FFFF")).
				Bold(true)

	PhaseDoneStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00"))

	PhasePendingStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#666666"))

	// Error style
	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Bold(true)

	// Info/help text style
	InfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Italic(true)

	WarningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f49e1b")).
			Italic(true)

	// Selected item style
	SelectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF00FF")).
			Bold(true).
			PaddingLeft(2)

	// Unselected item style
	UnselectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			PaddingLeft(2)

	// Input prompt style
	PromptStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFFF")).
			Bold(true)

	// Border style
	BorderStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF00FF")).
			Padding(1, 2)

	// Progress bar colors
	GradientStart color.Color = lipgloss.Color("#FF00FF") // Magenta
	GradientMid   color.Color = lipgloss.Color("#9F00FF") // Purple
	GradientEnd   color.Color = lipgloss.Color("#00FFFF") // Cyan

	// Success message
	SuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00")).
			Bold(true)

	// Border colors (will be set by ApplyTheme)
	SetupBorderColor    color.Color
	ProgressBorderColor color.Color
	ResultsBorderColor  color.Color
	ErrorBorderColor    color.Color
	ThemeBorder         color.Color

	// Divider style
	DividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#444444"))

	// File browser styles (set by ApplyTheme)
	BrowserDirStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#87bfc1")).PaddingLeft(2)
	BrowserFaultFileStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#b5d569")).PaddingLeft(2)
)
