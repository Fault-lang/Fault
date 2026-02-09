# Fault TUI - Interactive Terminal User Interface

This package implements a colorful, interactive TUI for the Fault compiler using [Bubble Tea](https://github.com/charmbracelet/bubbletea).

## Features

- **Hybrid Mode**: Automatically switches between traditional CLI and interactive TUI
  - Traditional CLI: When `-f` flag is provided
  - Interactive TUI: When no flags are provided
- **Setup View**: Interactive menu for selecting compilation options
- **Progress View**: Shows compilation progress with animated spinners (basic version - spinner only)
- **Results View**: Scrollable output with vim-style keybindings
- **Error View**: Colorful error display with helpful messages

## Architecture

### File Structure

```
tui/
├── model.go      # Main Bubble Tea model and state machine
├── setup.go      # Setup view: file selection and compilation options
├── progress.go   # Progress view: compilation status with spinner
├── results.go    # Results view: scrollable output
├── styles.go     # Lipgloss color styles
├── messages.go   # Bubble Tea message types
└── README.md     # This file
```

### State Machine

```
ViewSetup → ViewProgress → ViewResults
    ↓
ViewError (on any error)
```

### Views

#### Setup View
- File path input (textinput)
- Mode selection: check, ast, ir, smt
- Input format selection: fspec, ll, smt2
- Output format selection: log, smt, static, legacy, visualize
- Navigation: ↑/↓ or j/k, Enter to select

#### Progress View
- Shows compilation phase with animated spinner
- Simple version: just a spinner while compilation runs
- Future enhancement: Real-time progress updates for 7 phases

#### Results View
- Scrollable viewport with compilation output
- Vim-style navigation: j/k or ↑/↓
- Scroll percentage indicator
- Press 'q' to quit

#### Error View
- Red-bordered error message
- Clear error description
- Press 'q' to quit

## Usage

### Interactive Mode
```bash
./fault
# Launches TUI - follow on-screen prompts
```

### Traditional CLI Mode
```bash
./fault -f examples/battery.fspec -m check
./fault -f examples/battery.fspec -m ast
./fault -f examples/battery.fspec -m smt -format smt
```

## Keybindings

### Global
- `Ctrl+C` or `q`: Quit

### Setup View
- `↑`/`↓` or `j`/`k`: Navigate options
- `Enter`: Select/Continue

### Progress View
- Compilation runs automatically
- No user input during compilation

### Results View
- `↑`/`↓` or `j`/`k`: Scroll
- `q` or `Ctrl+C`: Exit

## Color Scheme

The TUI uses a vibrant color palette:
- **Magenta (#FF00FF)**: Titles, selected items, borders
- **Cyan (#00FFFF)**: Subtitles, active phases, prompts
- **Green (#00FF00)**: Completed phases, success messages
- **Red (#FF0000)**: Errors
- **Gray (#666666, #888888)**: Pending items, info text

## Future Enhancements

1. **Real-time Progress Updates**
   - Currently: Single spinner during compilation
   - Future: 7-phase progress bars with percentage tracking
   - Implementation: Use tea.Program.Send() for async updates

2. **Search in Results**
   - Press `/` to search
   - Highlight matches
   - Navigate between matches

3. **History**
   - Remember last used file path
   - Quick access to recent compilations

4. **Themes**
   - Allow users to customize colors
   - Light/dark theme toggle

5. **Export Results**
   - Save output to file from results view
   - Export as different formats

## Implementation Notes

### Progress Updates

The current implementation runs compilation synchronously in a tea.Cmd. Progress updates are not yet implemented in the TUI (though the infrastructure exists in `runner.go`).

To enable real-time progress:
1. Use `tea.Program.Send()` in a goroutine
2. Forward `runner.ProgressUpdate` messages to the tea program
3. Update progress view on each `ProgressUpdateMsg`

See commented code in `model.go` for the pattern.

### Async Compilation

The compilation runs in a Bubble Tea command goroutine, which blocks until complete. This is intentional for simplicity. Real-time progress would require a more complex message passing setup.

## Dependencies

- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/lipgloss` - Styling
- `github.com/charmbracelet/bubbles` - Pre-built components (textinput, viewport, spinner, progress)
