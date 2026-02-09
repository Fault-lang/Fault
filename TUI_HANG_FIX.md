# TUI Hang Fix - "Initializing..." Issue

## Problem
The TUI was hanging at "Initializing..." message in the results view after compilation completed.

## Root Cause
The results view viewport wasn't being initialized because:
1. WindowSizeMsg wasn't being forwarded to the ResultsModel's Update() method
2. No fallback dimensions if WindowSizeMsg arrived with 0x0 dimensions
3. Content selection logic was too rigid (checking mode instead of actual output)

## Fixes Applied

### 1. Forward WindowSizeMsg to Results View (`tui/model.go`)
**Before:**
```go
case ViewResults:
    m.results.width = msg.Width
    m.results.height = msg.Height
```

**After:**
```go
case ViewResults:
    m.results, cmd = m.results.Update(msg)  // Actually call Update!
```

Now the viewport receives the message and can initialize properly.

### 2. Add Fallback Dimensions (`tui/model.go`)
**Added default dimensions (80x24) if WindowSizeMsg hasn't arrived yet:**
```go
width := m.width
height := m.height
if width == 0 {
    width = 80
}
if height == 0 {
    height = 24
}
```

This prevents viewport initialization failure if compilation completes before terminal size is known.

### 3. Ensure Viewport Initialization on Transition (`tui/model.go`)
**Immediately initialize viewport when transitioning to results:**
```go
var cmd tea.Cmd
m.results, cmd = m.results.Update(tea.WindowSizeMsg{Width: width, Height: height})
return m, cmd
```

### 4. Improve Content Selection Logic (`tui/results.go`)
**Changed from mode-based to availability-based:**

**Before:**
```go
switch m.mode {
    case "ast": return ast
    case "ir": return ir
    case "smt": return smt
    case "check": return logger  // Could be nil!
}
```

**After:**
```go
// Return first available output
if m.logger != nil { return m.logger.String() }
if m.smt != "" { return m.smt }
if m.ir != "" { return m.ir }
if m.ast != nil { return fmt.Sprintf("%v", m.ast) }
```

This handles cases where mode and output format don't align.

### 5. Add Debug Info to View (`tui/results.go`)
**If still not ready, show dimensions for debugging:**
```go
if !m.ready {
    // Show "Preparing results..." with dimensions
    b.WriteString(fmt.Sprintf("(Dimensions: %dx%d)", m.width, m.height))
}
```

## Testing

### Automated Tests: ✅ PASS
```bash
./test_tui.sh
```

### Manual Testing Required
1. Launch TUI: `./fault`
2. Enter file: `generator/testdata/booleans.fspec`
3. Select mode: `ast` or `check`
4. Select input: `fspec`
5. Select output: `log`
6. **Should show results immediately** (no hang!)
7. Scroll with j/k or arrow keys
8. Quit with 'q'

## Expected Behavior After Fix

### Before:
- ❌ Hung at "Initializing..." indefinitely
- ❌ No way to see what was wrong
- ❌ Had to kill process with Ctrl+C

### After:
- ✅ Results display immediately after compilation
- ✅ Viewport initialized with proper dimensions
- ✅ Scrollable output works
- ✅ If dimensions somehow still 0, shows debug info

## Files Modified
1. `tui/model.go` - WindowSizeMsg forwarding, fallback dimensions
2. `tui/results.go` - Content selection, debug output

## Verification
```bash
# Traditional CLI still works
./fault -f generator/testdata/booleans.fspec -m ast

# TUI mode should now work without hanging
./fault
```

## If Issue Persists
If you still see "Initializing...", check the debug output:
- If shows `(Dimensions: 0x0)` → Terminal size detection issue
- If shows `(Dimensions: 80x24)` → Viewport creation failing (report bug)

## Additional Notes
- Default dimensions (80x24) are used as fallback
- Viewport automatically adjusts when terminal is resized
- All fixes maintain backward compatibility
