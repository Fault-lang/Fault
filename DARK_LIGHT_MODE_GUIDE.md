# Dark/Light Mode Implementation Guide

## Quick Answer

**The TUI already uses adaptive colors!** Lipgloss automatically adjusts colors based on your terminal's background. No action needed.

## What's Already Implemented

✅ **Adaptive Color Infrastructure**
- File: `tui/theme.go`
- Colors automatically adjust to terminal background
- Works seamlessly in any terminal

✅ **Theme System**
- `Theme` type with adaptive colors
- `ApplyTheme()` function
- `DefaultTheme()` with light/dark variants

## How It Works

Lipgloss uses `AdaptiveColor`:
```go
lipgloss.AdaptiveColor{
    Light: "#0087AF", // For light terminals
    Dark:  "#00FFFF", // For dark terminals
}
```

When you use this color in a style, lipgloss automatically picks the right variant based on your terminal's background.

## Testing

### In Dark Terminal
```bash
./fault
# Colors will be bright (cyan, magenta, green)
```

### In Light Terminal
```bash
# Change your terminal's theme to light mode, then:
./fault
# Colors will be darker (for better contrast)
```

## Options for Enhancement

### Option 1: Manual Toggle (Recommended)
Let users press 't' to switch themes:

**Implementation:** See `tui/theme_toggle_example.go` for complete code

**Steps:**
1. Add `darkMode bool` to Model
2. Handle 't' key in Update()
3. Call `ApplyTheme(DarkTheme())` or `ApplyTheme(LightTheme())`
4. Update help text to show 't' key

**Effort:** ~30 lines of code
**Benefit:** User control without terminal changes

### Option 2: Environment Variable
Read theme from environment:

```bash
FAULT_THEME=light ./fault
FAULT_THEME=dark ./fault
```

**Implementation:**
```go
func (m Model) Init() tea.Cmd {
    theme := os.Getenv("FAULT_THEME")
    switch theme {
    case "light":
        ApplyTheme(LightTheme())
    case "dark":
        ApplyTheme(DarkTheme())
    default:
        ApplyTheme(DefaultTheme()) // Adaptive
    }
    return m.setup.Init()
}
```

**Effort:** ~15 lines of code
**Benefit:** Easy to set per-session

### Option 3: Config File
Save persistent preference:

```bash
# ~/.config/fault/config.json
{
  "theme": "dark"
}
```

**Implementation:** See `tui/THEMING.md` Option 4

**Effort:** ~100 lines of code (config loading/saving)
**Benefit:** Persistent user preference

### Option 4: Multiple Named Themes
Offer various themes (Solarized, Dracula, Nord):

```go
var Themes = map[string]Theme{
    "default": DefaultTheme(),
    "dracula": DraculaTheme(),
    "solarized": SolarizedTheme(),
    "nord": NordTheme(),
}
```

**Implementation:** See `tui/THEMING.md` Advanced section

**Effort:** ~50 lines per theme
**Benefit:** Visual variety, popular themes

## Popular Theme Examples

All included in `tui/THEMING.md`:

- **Solarized Dark/Light** - Ethan Schoonover's famous theme
- **Dracula** - Purple/pink aesthetic
- **Nord** - Arctic, bluish palette
- **Monokai** - Sublime Text classic

## Files Created

1. ✅ `tui/theme.go` - Theme system implementation
2. ✅ `tui/THEMING.md` - Complete theming guide (19 KB)
3. ✅ `tui/theme_toggle_example.go` - Manual toggle example
4. ✅ `DARK_LIGHT_MODE_GUIDE.md` - This file

## Recommendation

**For now:** The adaptive colors work great, no changes needed!

**To add:** Manual theme toggle (Option 1) - gives users immediate control with minimal code.

**Later:** Config file (Option 3) - for users who want persistent preferences.

## Next Steps

### If You Want Manual Toggle
1. Copy code from `tui/theme_toggle_example.go`
2. Add to `model.go` as shown
3. Add `DarkTheme()` and `LightTheme()` to `theme.go`
4. Update help text in views
5. Rebuild and test

### If You Want Config File
1. Follow Option 4 in `tui/THEMING.md`
2. Create config directory structure
3. Add load/save functions
4. Hook into Model.Init()
5. Add settings view (optional)

### If Current Adaptive Colors Are Enough
**Do nothing!** It already works. 🎉

## Color Contrast Tips

For accessibility:
- **Dark mode:** Use bright colors (#00FFFF, #FF00FF)
- **Light mode:** Use darker variants (#0087AF, #AF00AF)
- **Test both:** Ensure text is readable in both modes
- **Minimum contrast:** 4.5:1 for normal text, 3:1 for large text

## Testing Checklist

- [ ] Dark terminal - bright colors visible?
- [ ] Light terminal - dark colors visible?
- [ ] Toggle works (if implemented)?
- [ ] Config persists (if implemented)?
- [ ] Help text mentions theme key?
- [ ] All views use theme colors?

## Resources

- **Lipgloss Docs:** https://github.com/charmbracelet/lipgloss
- **Adaptive Colors:** See lipgloss `AdaptiveColor` type
- **Color Contrast Checker:** https://webaim.org/resources/contrastchecker/
- **Terminal Color Guide:** https://github.com/mbadolato/iTerm2-Color-Schemes

## Summary

✅ **Current state:** Adaptive colors working automatically
🎯 **Recommended next:** Add manual toggle (press 't')
🚀 **Future:** Config file with named themes

The foundation is already in place - you can enhance with just a few lines of code or leave it as-is (it already works great!).
