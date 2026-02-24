package tui

// EXAMPLE: How to add manual theme toggling to the TUI
// This file shows the changes needed to add a 't' key to toggle themes

/*
// 1. Add to Model struct in model.go:
type Model struct {
    state  ViewState
    width  int
    height int

    // Theme support
    darkMode bool    // <-- ADD THIS

    setup    SetupModel
    progress ProgressModel
    results  ResultsModel
    config   *runner.CompilationConfig
    output   *runner.CompilationOutput
    err      error
}

// 2. Initialize in NewModel():
func NewModel() Model {
    return Model{
        state:    ViewSetup,
        setup:    NewSetupModel(),
        darkMode: true, // <-- ADD THIS (start in dark mode)
    }
}

// 3. Add theme toggle handling in Update():
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        case "t": // <-- ADD THIS CASE
            // Toggle theme
            m.darkMode = !m.darkMode
            if m.darkMode {
                ApplyTheme(DarkTheme())
            } else {
                ApplyTheme(LightTheme())
            }
            // Return to trigger re-render with new colors
            return m, nil
        }
    // ... rest of Update()
}

// 4. Add theme definitions to theme.go:
func DarkTheme() Theme {
    return Theme{
        Primary: lipgloss.AdaptiveColor{
            Light: "#FF00FF",
            Dark:  "#FF00FF",
        },
        Secondary: lipgloss.AdaptiveColor{
            Light: "#00FFFF",
            Dark:  "#00FFFF",
        },
        Success: lipgloss.AdaptiveColor{
            Light: "#00FF00",
            Dark:  "#00FF00",
        },
        Error: lipgloss.AdaptiveColor{
            Light: "#FF0000",
            Dark:  "#FF0000",
        },
        Warning: lipgloss.AdaptiveColor{
            Light: "#FFAA00",
            Dark:  "#FFAA00",
        },
        Muted: lipgloss.AdaptiveColor{
            Light: "#888888",
            Dark:  "#888888",
        },
        Border: lipgloss.AdaptiveColor{
            Light: "#FF00FF",
            Dark:  "#FF00FF",
        },
        Background: lipgloss.AdaptiveColor{
            Light: "#000000",
            Dark:  "#000000",
        },
        Foreground: lipgloss.AdaptiveColor{
            Light: "#FFFFFF",
            Dark:  "#FFFFFF",
        },
    }
}

func LightTheme() Theme {
    return Theme{
        Primary: lipgloss.AdaptiveColor{
            Light: "#AF00AF", // Darker magenta for light bg
            Dark:  "#AF00AF",
        },
        Secondary: lipgloss.AdaptiveColor{
            Light: "#0087AF", // Darker cyan for light bg
            Dark:  "#0087AF",
        },
        Success: lipgloss.AdaptiveColor{
            Light: "#00AA00", // Darker green for light bg
            Dark:  "#00AA00",
        },
        Error: lipgloss.AdaptiveColor{
            Light: "#D70000", // Darker red for light bg
            Dark:  "#D70000",
        },
        Warning: lipgloss.AdaptiveColor{
            Light: "#D78700", // Darker orange for light bg
            Dark:  "#D78700",
        },
        Muted: lipgloss.AdaptiveColor{
            Light: "#666666", // Dark gray for light bg
            Dark:  "#666666",
        },
        Border: lipgloss.AdaptiveColor{
            Light: "#AF00AF",
            Dark:  "#AF00AF",
        },
        Background: lipgloss.AdaptiveColor{
            Light: "#FFFFFF",
            Dark:  "#FFFFFF",
        },
        Foreground: lipgloss.AdaptiveColor{
            Light: "#000000",
            Dark:  "#000000",
        },
    }
}

// 5. Update help text in each view to show 't' key:
// In setup.go, progress.go, results.go, add to help text:
b.WriteString(InfoStyle.Render("Press t to toggle theme • Press q to quit"))
*/

// Example usage in a simple program:
/*
func main() {
    p := tea.NewProgram(
        tui.NewModel(),
        tea.WithAltScreen(),
    )

    if _, err := p.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}

// Now users can press 't' to toggle between light and dark themes!
*/
