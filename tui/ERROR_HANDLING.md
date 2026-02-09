# TUI Error Handling

## Overview
The TUI now features comprehensive error handling with contextual information, recovery actions, and user-friendly feedback.

## Features

### 1. Structured Error Types
Errors are categorized into specific types:
- **File Errors**: File not found, permission issues, invalid file types
- **Parsing Errors**: Syntax errors in the specification
- **Type Check Errors**: Type compatibility issues
- **LLVM Errors**: IR generation problems
- **SMT Errors**: Constraint generation issues
- **Solver Errors**: Model checker configuration or runtime errors
- **Internal Errors**: Unexpected system errors

### 2. Enhanced Error Display
When an error occurs, the TUI shows:
- **Phase Context**: Which compilation phase failed
- **Error Message**: Clear description of what went wrong
- **Additional Details**: Context-specific information
- **Suggestions**: Actionable steps to resolve the issue
- **Pipeline Progress**: Visual representation showing completed/failed phases

Example:
```
✗ Compilation Failed at Type Checking

Error: undefined variable 'foo' at line 42

Details: Type checking found an issue with types or declarations

💡 Suggestion:
Check variable declarations and type compatibility

Pipeline Progress:

✓ Parsing (complete)
✓ Preprocessing (complete)
✗ Type Checking (failed)
⋯ LLVM IR Generation (not started)
⋯ SMT Generation (not started)
⋯ Model Checking (not started)
⋯ Results (not started)
```

### 3. Error Recovery Actions
Users can choose from three actions when an error occurs:
- **Retry**: Attempt compilation again with the same configuration
- **Back to Setup**: Return to the setup screen to modify settings
- **Quit**: Exit the application

Navigation:
- Use `↑/↓` or `j/k` to navigate between options
- Press `Enter` to select
- Quick keys: `r` for Retry, `b` for Back, `q` for Quit

### 4. Proactive Validation
The setup screen now validates files before starting compilation:
- Checks if file exists
- Verifies read permissions
- Validates solver configuration (for "check" mode)
- Shows inline error messages in the file input step

### 5. Context-Aware Suggestions
The error system provides specific suggestions based on error type:

| Error Type | Suggestion |
|------------|------------|
| File not found | "Check that the file path is correct and the file exists" |
| Permission denied | "Check file permissions - you may need read access" |
| Solver not configured | "Configure solver: set SOLVERCMD and SOLVERARG environment variables" |
| Parsing error | "Check for syntax errors in your specification" |
| Type error | "Check variable declarations and type compatibility" |
| Missing run block | "Ensure your specification has a valid run or start block" |

## Implementation Details

### New Files
- `tui/errors.go`: Error categorization, validation, and rendering logic

### Modified Files
- `tui/model.go`: Enhanced error state management and recovery actions
- `tui/messages.go`: Added phase information to error messages
- `tui/setup.go`: Proactive file validation in setup screen
- `runner/runner.go`: Phase tracking in error output

### Key Functions
- `CategorizeError()`: Analyzes errors and provides context
- `ValidateSetupConfig()`: Pre-compilation validation
- `RenderPhaseProgress()`: Visual pipeline progress display
- `GetPhaseName()`: Human-readable phase names

## Usage Examples

### File Not Found
```
✗ Compilation Failed at Parsing

Error: file not found: /path/to/missing.fspec

💡 Suggestion:
Check that the file path is correct and the file exists
```

### Solver Not Configured
```
✗ Compilation Failed at Model Checking

Error: solver not configured: set SOLVERCMD and SOLVERARG environment variables

Details: The SMT solver encountered an error

💡 Suggestion:
Configure solver: set SOLVERCMD and SOLVERARG environment variables
```

### Validation in Setup
When entering a non-existent file in the setup screen:
```
Enter the path to the file to compile:

/path/to/missing.fspec

⚠ file not found: /path/to/missing.fspec

Press Enter to continue
```

## Benefits

1. **Better User Experience**: Clear, actionable error messages
2. **Faster Debugging**: Phase context helps identify where issues occur
3. **Reduced Frustration**: Suggestions guide users to solutions
4. **Flexible Recovery**: Multiple recovery options instead of just quitting
5. **Early Detection**: Proactive validation catches common issues before compilation
