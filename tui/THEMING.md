# Lipgloss Dark/Light Mode Guide

## Overview
There are several ways to implement dark/light mode themes with lipgloss:

1. **Adaptive Colors** (Automatic) - Recommended
2. **Manual Theme Switching** (User toggle)
3. **Environment Detection** (Based on terminal settings)
4. **User Configuration** (Persistent preference)

## Option 1: Adaptive Colors (Implemented)

Lipgloss's `AdaptiveColor` type automatically adjusts based on terminal background:

```go
lipgloss.AdaptiveColor{
    Light: "#0087AF", // Used when terminal has light background
    Dark:  "#00FFFF", // Used when terminal has dark background
}
```

### How It Works
- Lipgloss detects terminal's background color
- Automatically selects appropriate color variant
- No user action required
- Works seamlessly across different terminals

### Implementation
See `tui/theme.go` for the full adaptive theme implementation.

### Usage
```go
// Colors adapt automatically
theme := DefaultTheme()
ApplyTheme(theme)

// Styles will use appropriate colors for terminal background
style := lipgloss.NewStyle().Foreground(theme.Primary)
```

## Option 2: Manual Theme Switching

Allow users to toggle between themes with a keybinding:

```go
// In model.go
type Model struct {
    // ... existing fields
    darkMode bool
}

// Add to Update()
case tea.KeyMsg:
    switch msg.String() {
    case "t": // Toggle theme
        m.darkMode = !m.darkMode
        if m.darkMode {
            ApplyTheme(DarkTheme())
        } else {
            ApplyTheme(LightTheme())
        }
        return m, nil
    }
```

### Define Separate Themes
```go
func DarkTheme() Theme {
    return Theme{
        Primary:    lipgloss.Color("#FF00FF"),
        Secondary:  lipgloss.Color("#00FFFF"),
        Success:    lipgloss.Color("#00FF00"),
        Error:      lipgloss.Color("#FF0000"),
        Background: lipgloss.Color("#000000"),
        Foreground: lipgloss.Color("#FFFFFF"),
        // ...
    }
}

func LightTheme() Theme {
    return Theme{
        Primary:    lipgloss.Color("#AF00AF"),
        Secondary:  lipgloss.Color("#0087AF"),
        Success:    lipgloss.Color("#00AA00"),
        Error:      lipgloss.Color("#D70000"),
        Background: lipgloss.Color("#FFFFFF"),
        Foreground: lipgloss.Color("#000000"),
        // ...
    }
}
```

## Option 3: Environment Detection

Detect terminal settings to choose theme:

```go
func DetectTheme() Theme {
    // Method 1: Check COLORFGBG environment variable
    if colorfgbg := os.Getenv("COLORFGBG"); colorfgbg != "" {
        // Format is usually "foreground;background"
        // Example: "15;0" = light text on dark bg (dark mode)
        //          "0;15" = dark text on light bg (light mode)
        parts := strings.Split(colorfgbg, ";")
        if len(parts) == 2 {
            bg, _ := strconv.Atoi(parts[1])
            if bg >= 8 {
                return LightTheme() // High number = light background
            }
        }
    }

    // Method 2: Check terminal type
    if term := os.Getenv("TERM_PROGRAM"); term != "" {
        // Some terminals set specific env vars
        // Could check iTerm2 settings, etc.
    }

    // Method 3: Check if TERM contains "light"
    if strings.Contains(os.Getenv("TERM"), "light") {
        return LightTheme()
    }

    // Default to dark
    return DarkTheme()
}
```

## Option 4: User Configuration

Save theme preference to config file:

```go
type Config struct {
    Theme string `json:"theme"` // "dark", "light", or "auto"
}

func LoadConfig() (*Config, error) {
    home, _ := os.UserHomeDir()
    configPath := filepath.Join(home, ".config", "fault", "config.json")

    data, err := os.ReadFile(configPath)
    if err != nil {
        return &Config{Theme: "auto"}, nil // Default
    }

    var cfg Config
    json.Unmarshal(data, &cfg)
    return &cfg, nil
}

func (m Model) Init() tea.Cmd {
    cfg, _ := LoadConfig()

    switch cfg.Theme {
    case "dark":
        ApplyTheme(DarkTheme())
    case "light":
        ApplyTheme(LightTheme())
    case "auto":
        ApplyTheme(DefaultTheme()) // Adaptive
    }

    return m.setup.Init()
}
```

## Recommended Color Palettes

### Dark Mode (Current)
```go
Primary:   "#FF00FF" // Magenta
Secondary: "#00FFFF" // Cyan
Success:   "#00FF00" // Green
Error:     "#FF0000" // Red
Warning:   "#FFAA00" // Orange
Muted:     "#888888" // Light gray
```

### Light Mode
```go
Primary:   "#AF00AF" // Darker magenta
Secondary: "#0087AF" // Darker cyan
Success:   "#00AA00" // Darker green
Error:     "#D70000" // Darker red
Warning:   "#D78700" // Darker orange
Muted:     "#666666" // Dark gray
```

### High Contrast Dark
```go
Primary:   "#FFFFFF" // White
Secondary: "#00FFFF" // Cyan
Success:   "#00FF00" // Green
Error:     "#FF0000" // Red
```

### High Contrast Light
```go
Primary:   "#000000" // Black
Secondary: "#0000FF" // Blue
Success:   "#008000" // Green
Error:     "#FF0000" // Red
```

## Testing Different Modes

### Test in Dark Terminal
```bash
# Most terminals default to dark mode
./fault
```

### Test in Light Terminal
```bash
# Set terminal background to light color
# Or use a terminal that supports theme switching (iTerm2, etc.)
./fault
```

### Force Specific Theme (if implementing manual switching)
```bash
FAULT_THEME=light ./fault
# or
FAULT_THEME=dark ./fault
```

## Best Practices

1. **Use Adaptive Colors** - Best user experience, works automatically
2. **Provide Toggle** - Let power users override if needed
3. **Test Both Modes** - Ensure readability in light and dark terminals
4. **Consider Accessibility** - Use sufficient contrast ratios
5. **Respect Terminal Settings** - Don't override user's terminal theme unnecessarily

## Color Contrast Guidelines

For accessibility (WCAG AA compliance):
- **Normal text**: 4.5:1 contrast ratio minimum
- **Large text**: 3:1 contrast ratio minimum
- **UI elements**: 3:1 contrast ratio minimum

### Tools
- WebAIM Contrast Checker: https://webaim.org/resources/contrastchecker/
- Colorable: https://colorable.jxnblk.com/

## Advanced: Multiple Named Themes

```go
var Themes = map[string]Theme{
    "default": DefaultTheme(),
    "solarized-dark": SolarizedDarkTheme(),
    "solarized-light": SolarizedLightTheme(),
    "dracula": DraculaTheme(),
    "nord": NordTheme(),
    "monokai": MonokaiTheme(),
}

func ApplyNamedTheme(name string) {
    if theme, ok := Themes[name]; ok {
        ApplyTheme(theme)
    }
}
```

## Example Themes

### Solarized Dark
```go
func SolarizedDarkTheme() Theme {
    return Theme{
        Primary:    lipgloss.Color("#268BD2"), // Blue
        Secondary:  lipgloss.Color("#2AA198"), // Cyan
        Success:    lipgloss.Color("#859900"), // Green
        Error:      lipgloss.Color("#DC322F"), // Red
        Warning:    lipgloss.Color("#B58900"), // Yellow
        Muted:      lipgloss.Color("#586E75"), // Gray
        Background: lipgloss.Color("#002B36"), // Base03
        Foreground: lipgloss.Color("#839496"), // Base0
    }
}
```

### Dracula
```go
func DraculaTheme() Theme {
    return Theme{
        Primary:    lipgloss.Color("#FF79C6"), // Pink
        Secondary:  lipgloss.Color("#8BE9FD"), // Cyan
        Success:    lipgloss.Color("#50FA7B"), // Green
        Error:      lipgloss.Color("#FF5555"), // Red
        Warning:    lipgloss.Color("#FFB86C"), // Orange
        Muted:      lipgloss.Color("#6272A4"), // Comment
        Background: lipgloss.Color("#282A36"), // Background
        Foreground: lipgloss.Color("#F8F8F2"), // Foreground
    }
}
```

## Current Implementation Status

✅ **Implemented:**
- Adaptive color infrastructure (`theme.go`)
- Theme type definition
- `ApplyTheme()` function
- Default adaptive theme

⏳ **To Implement:**
- Manual theme toggle (press 't' to switch)
- Config file support
- Additional named themes
- Theme selection in setup view

## Quick Start

### Use Current Adaptive Theme (No changes needed)
The TUI already uses adaptive colors! Just run:
```bash
./fault
```

Colors will automatically adjust to your terminal's background.

### Add Manual Theme Toggle
1. Add `darkMode bool` to Model struct
2. Handle 't' key in Update()
3. Call `ApplyTheme()` with appropriate theme
4. Rebuild and test

See examples above for implementation details.
