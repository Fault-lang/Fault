package tui

import (
	"os"

	"github.com/charmbracelet/lipgloss"
)

// Theme represents a color theme
type Theme struct {
	Primary        lipgloss.AdaptiveColor
	Secondary      lipgloss.AdaptiveColor
	Success        lipgloss.AdaptiveColor
	Error          lipgloss.AdaptiveColor
	Warning        lipgloss.AdaptiveColor
	Muted          lipgloss.AdaptiveColor
	Border         lipgloss.AdaptiveColor
	Background     lipgloss.AdaptiveColor
	Foreground     lipgloss.AdaptiveColor
	SetupBorder    lipgloss.AdaptiveColor
	ProgressBorder lipgloss.AdaptiveColor
	ResultsBorder  lipgloss.AdaptiveColor
	ErrorBorder    lipgloss.AdaptiveColor
}

// DefaultTheme returns an adaptive theme that works in both dark and light terminals
func DefaultTheme() Theme {
	return Theme{
		Primary: lipgloss.AdaptiveColor{
			Light: "#567b02", // Matcha
			Dark:  "#e2ecba", // Familiarity
		},
		Secondary: lipgloss.AdaptiveColor{
			Light: "#87bfc1", // Unrequited
			Dark:  "#87bfc1", // Unrequited
		},
		Success: lipgloss.AdaptiveColor{
			Light: "#72ae2c", // Darker green for light backgrounds
			Dark:  "#b5d569", // Bright green for dark backgrounds
		},
		Error: lipgloss.AdaptiveColor{
			Light: "#bb0009", // Darker red for light backgrounds
			Dark:  "#E73825", // Bright red for dark backgrounds
		},
		Warning: lipgloss.AdaptiveColor{
			Light: "#f49e1b", // Darker orange for light backgrounds
			Dark:  "#F5af7d", // Bright orange for dark backgrounds
		},
		Muted: lipgloss.AdaptiveColor{
			Light: "#6e7156", // Dark gray for light backgrounds
			Dark:  "#bbb4a5", // Light gray for dark backgrounds
		},
		Border: lipgloss.AdaptiveColor{
			Light: "#5b73b0ff", // Darker blue for light backgrounds
			Dark:  "#73b0edff", // Bright blue for dark backgrounds
		},
		Foreground: lipgloss.AdaptiveColor{
			Light: "#685936ff", // Russet
			Dark:  "#b6ae9bff", // White text
		},
		SetupBorder: lipgloss.AdaptiveColor{
			Light: "#87bfc1", // Unrequited
			Dark:  "#87bfc1", // Unrequited
		},
		ProgressBorder: lipgloss.AdaptiveColor{
			Light: "#567b02", // Matcha
			Dark:  "#e2ecba", // Familiarity
		},
		ResultsBorder: lipgloss.AdaptiveColor{
			Light: "#87bfc1", // Unrequited
			Dark:  "#87bfc1", // Unrequited
		},
		ErrorBorder: lipgloss.AdaptiveColor{
			Light: "#bb0009", // Darker red for light backgrounds
			Dark:  "#E73825", // Bright red for dark backgrounds
		},
	}
}

// DarkTheme returns a theme optimized for dark terminals
func DarkTheme() Theme {
	return Theme{
		Primary: lipgloss.AdaptiveColor{
			Light: "#e2ecba", // Familiarity
			Dark:  "#e2ecba", // Familiarity
		},
		Secondary: lipgloss.AdaptiveColor{
			Light: "#87bfc1", // Unrequited
			Dark:  "#87bfc1", // Unrequited
		},
		Success: lipgloss.AdaptiveColor{
			Light: "#b5d569",
			Dark:  "#b5d569",
		},
		Error: lipgloss.AdaptiveColor{
			Light: "#E73825",
			Dark:  "#E73825",
		},
		Warning: lipgloss.AdaptiveColor{
			Light: "#F5af7d",
			Dark:  "#F5af7d",
		},
		Muted: lipgloss.AdaptiveColor{
			Light: "#bbb4a5",
			Dark:  "#bbb4a5",
		},
		Border: lipgloss.AdaptiveColor{
			Light: "#73b0edff",
			Dark:  "#73b0edff",
		},
		Background: lipgloss.AdaptiveColor{
			Light: "#000000",
		},
		Foreground: lipgloss.AdaptiveColor{
			Light: "#b6ae9bff",
			Dark:  "#b6ae9bff",
		},
		SetupBorder: lipgloss.AdaptiveColor{
			Light: "#87bfc1", // Unrequited
			Dark:  "#87bfc1", // Unrequited
		},
		ProgressBorder: lipgloss.AdaptiveColor{
			Light: "#567b02", // Matcha
			Dark:  "#e2ecba", // Familiarity
		},
		ResultsBorder: lipgloss.AdaptiveColor{
			Light: "#87bfc1", // Unrequited
			Dark:  "#87bfc1", // Unrequited
		},
		ErrorBorder: lipgloss.AdaptiveColor{
			Light: "#bb0009", // Darker red for light backgrounds
			Dark:  "#E73825", // Bright red for dark backgrounds
		},
	}
}

// LightTheme returns a theme optimized for light terminals
func LightTheme() Theme {
	return Theme{
		Primary: lipgloss.AdaptiveColor{
			Light: "#567b02", // Matcha
			Dark:  "#567b02", // Matcha
		},
		Secondary: lipgloss.AdaptiveColor{
			Light: "#87bfc1", // Unrequited
			Dark:  "#87bfc1", // Unrequited
		},
		Success: lipgloss.AdaptiveColor{
			Light: "#72ae2c", // Darker green
			Dark:  "#72ae2c",
		},
		Error: lipgloss.AdaptiveColor{
			Light: "#bb0009", // Darker red
			Dark:  "#bb0009",
		},
		Warning: lipgloss.AdaptiveColor{
			Light: "#f49e1b", // Darker orange
			Dark:  "#f49e1b",
		},
		Muted: lipgloss.AdaptiveColor{
			Light: "#6e7156", // Dark gray
			Dark:  "#6e7156",
		},
		Border: lipgloss.AdaptiveColor{
			Light: "#5b73b0ff",
			Dark:  "#5b73b0ff",
		},
		Background: lipgloss.AdaptiveColor{
			Dark: "#FFFFFF",
		},
		Foreground: lipgloss.AdaptiveColor{
			Light: "#685936ff",
			Dark:  "#685936ff",
		},
		SetupBorder: lipgloss.AdaptiveColor{
			Light: "#87bfc1", // Unrequited
			Dark:  "#87bfc1", // Unrequited
		},
		ProgressBorder: lipgloss.AdaptiveColor{
			Light: "#567b02", // Matcha
			Dark:  "#e2ecba", // Familiarity
		},
		ResultsBorder: lipgloss.AdaptiveColor{
			Light: "#87bfc1", // Unrequited
			Dark:  "#87bfc1", // Unrequited
		},
		ErrorBorder: lipgloss.AdaptiveColor{
			Light: "#bb0009", // Darker red for light backgrounds
			Dark:  "#E73825", // Bright red for dark backgrounds
		},
	}
}

// DetectColorScheme attempts to detect if the terminal is using dark or light mode
func DetectColorScheme() string {
	// Check COLORFGBG environment variable (format: "foreground;background")
	// Light background usually has high number (15), dark has low (0)
	if colorfgbg := os.Getenv("COLORFGBG"); colorfgbg != "" {
		// Parse and determine, but this is terminal-specific
		// For simplicity, we'll rely on lipgloss's adaptive colors
	}

	// Check if TERM_PROGRAM is set (some terminals set this)
	// Could also check terminal capabilities

	// Default to dark (most developer terminals are dark)
	return "dark"
}

// ApplyTheme applies the theme to all styles
func ApplyTheme(theme Theme) {
	// Update global styles to use adaptive colors
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

	// Update border colors
	SetupBorderColor = theme.SetupBorder
	ProgressBorderColor = theme.ProgressBorder
	ResultsBorderColor = theme.ResultsBorder
	ErrorBorderColor = theme.ErrorBorder
	ThemeBorder = theme.Border

	// Update gradient colors for progress bars
	GradientStart = theme.Primary.Dark // Use dark variant
	GradientMid = theme.Secondary.Dark
	GradientEnd = theme.Secondary.Dark
}
