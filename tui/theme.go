package tui

import (
	"image/color"
	"os"

	"charm.land/lipgloss/v2"
)

// Theme represents a color theme
type Theme struct {
	Primary        color.Color
	Secondary      color.Color
	Success        color.Color
	Error          color.Color
	Warning        color.Color
	Muted          color.Color
	Border         color.Color
	Background     color.Color
	Foreground     color.Color
	SetupBorder    color.Color
	ProgressBorder color.Color
	ResultsBorder  color.Color
	ErrorBorder    color.Color
}

// DefaultTheme returns a dark theme (app starts in dark mode)
func DefaultTheme() Theme {
	return Theme{
		Primary:        lipgloss.Color("#e2ecba"), // Familiarity
		Secondary:      lipgloss.Color("#87bfc1"), // Unrequited
		Success:        lipgloss.Color("#b5d569"),
		Error:          lipgloss.Color("#E73825"),
		Warning:        lipgloss.Color("#F5af7d"),
		Muted:          lipgloss.Color("#bbb4a5"),
		Border:         lipgloss.Color("#73b0edff"),
		Foreground:     lipgloss.Color("#b6ae9bff"),
		SetupBorder:    lipgloss.Color("#87bfc1"),
		ProgressBorder: lipgloss.Color("#e2ecba"),
		ResultsBorder:  lipgloss.Color("#87bfc1"),
		ErrorBorder:    lipgloss.Color("#E73825"),
	}
}

// DarkTheme returns a theme optimized for dark terminals
func DarkTheme() Theme {
	return Theme{
		Primary:        lipgloss.Color("#e2ecba"),
		Secondary:      lipgloss.Color("#87bfc1"),
		Success:        lipgloss.Color("#b5d569"),
		Error:          lipgloss.Color("#E73825"),
		Warning:        lipgloss.Color("#F5af7d"),
		Muted:          lipgloss.Color("#bbb4a5"),
		Border:         lipgloss.Color("#73b0edff"),
		Background:     lipgloss.Color("#000000"),
		Foreground:     lipgloss.Color("#b6ae9bff"),
		SetupBorder:    lipgloss.Color("#87bfc1"),
		ProgressBorder: lipgloss.Color("#e2ecba"),
		ResultsBorder:  lipgloss.Color("#87bfc1"),
		ErrorBorder:    lipgloss.Color("#E73825"),
	}
}

// LightTheme returns a theme optimized for light terminals
func LightTheme() Theme {
	return Theme{
		Primary:        lipgloss.Color("#567b02"), // Matcha
		Secondary:      lipgloss.Color("#87bfc1"),
		Success:        lipgloss.Color("#72ae2c"),
		Error:          lipgloss.Color("#bb0009"),
		Warning:        lipgloss.Color("#f49e1b"),
		Muted:          lipgloss.Color("#6e7156"),
		Border:         lipgloss.Color("#5b73b0ff"),
		Background:     lipgloss.Color("#FFFFFF"),
		Foreground:     lipgloss.Color("#685936ff"),
		SetupBorder:    lipgloss.Color("#87bfc1"),
		ProgressBorder: lipgloss.Color("#567b02"),
		ResultsBorder:  lipgloss.Color("#87bfc1"),
		ErrorBorder:    lipgloss.Color("#bb0009"),
	}
}

// DetectColorScheme attempts to detect if the terminal is using dark or light mode
func DetectColorScheme() string {
	// Check COLORFGBG environment variable (format: "foreground;background")
	if colorfgbg := os.Getenv("COLORFGBG"); colorfgbg != "" {
		// For simplicity, rely on manual toggle
		_ = colorfgbg
	}
	// Default to dark (most developer terminals are dark)
	return "dark"
}

// ApplyTheme applies the theme to all styles
func ApplyTheme(theme Theme) {
	TitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.Primary).
		Background(theme.Background).
		Padding(0, 1)

	SubtitleStyle = lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Bold(true)

	PhaseActiveStyle = lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Bold(true)

	PhaseDoneStyle = lipgloss.NewStyle().
		Foreground(theme.Success)

	PhasePendingStyle = lipgloss.NewStyle().
		Foreground(theme.Muted)

	ErrorStyle = lipgloss.NewStyle().
		Foreground(theme.Error).
		Bold(true)

	InfoStyle = lipgloss.NewStyle().
		Foreground(theme.Muted).
		Italic(true)

	SelectedStyle = lipgloss.NewStyle().
		Foreground(theme.Primary).
		Bold(true).
		PaddingLeft(2)

	UnselectedStyle = lipgloss.NewStyle().
		Foreground(theme.Foreground).
		PaddingLeft(2)

	PromptStyle = lipgloss.NewStyle().
		Foreground(theme.Secondary).
		Bold(true)

	BorderStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(theme.Border).
		Padding(1, 2)

	SuccessStyle = lipgloss.NewStyle().
		Foreground(theme.Success).
		Bold(true)

	WarningStyle = lipgloss.NewStyle().
		Foreground(theme.Warning).
		Bold(true)

	DividerStyle = lipgloss.NewStyle().
		Foreground(theme.Muted)

	BrowserDirStyle = lipgloss.NewStyle().
		Foreground(theme.Secondary).
		PaddingLeft(2)

	BrowserFaultFileStyle = lipgloss.NewStyle().
		Foreground(theme.Success).
		PaddingLeft(2)

	// Update border colors
	SetupBorderColor = theme.SetupBorder
	ProgressBorderColor = theme.ProgressBorder
	ResultsBorderColor = theme.ResultsBorder
	ErrorBorderColor = theme.ErrorBorder
	ThemeBorder = theme.Border

	// Update gradient colors for progress bars
	GradientStart = theme.Primary
	GradientMid = theme.Secondary
	GradientEnd = theme.Secondary
}
