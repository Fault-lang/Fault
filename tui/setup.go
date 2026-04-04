package tui

import (
	"fault/runner"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SetupModel struct {
	step            int // 0=file, 1=mode, 2=input, 3=output
	fileInput       textinput.Model
	cursor          int
	config          runner.CompilationConfig
	width           int
	height          int
	validationErr   string // Error message for file validation
	browseMode      bool
	fileBrowser     FileBrowserModel
	solverAvailable bool
}

var (
	modes   = []string{"model (recommended)", "ast", "ir", "smt"}
	inputs  = []string{"fault (recommended)", "ll", "smt2"}
	outputs = []string{"text (recommended)", "smt"}
)

// activeModes returns the list of selectable modes. When no solver is
// configured, "model" is excluded because it requires SOLVERCMD/SOLVERARG.
func (m SetupModel) activeModes() []string {
	if m.solverAvailable {
		return modes
	}
	return []string{"ast", "ir", "smt"}
}

func NewSetupModel() SetupModel {
	ti := textinput.New()
	ti.Placeholder = "path/to/file.fspec"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 50

	solverAvailable := os.Getenv("SOLVERCMD") != "" && os.Getenv("SOLVERARG") != ""
	defaultMode := "smt"
	if solverAvailable {
		defaultMode = "model"
	}
	return SetupModel{
		step:      0,
		fileInput: ti,
		cursor:    0,
		config: runner.CompilationConfig{
			Mode:   defaultMode,
			Input:  "fault",
			Output: "text",
			Reach:  false,
		},
		solverAvailable: solverAvailable,
	}
}

func (m SetupModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m SetupModel) Update(msg tea.Msg) (SetupModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+q":
			return m, tea.Quit
		}

		if m.step == 0 {
			return m.updateFileStep(msg)
		}

		switch msg.String() {
		case "enter":
			return m.handleEnter()

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = m.getOptionsCount() - 1
			}

		case "down", "j":
			m.cursor++
			if m.cursor >= m.getOptionsCount() {
				m.cursor = 0
			}
		}
	}

	// Update file input for blink tick when on step 0 and not in browse mode
	if m.step == 0 && !m.browseMode {
		m.fileInput, cmd = m.fileInput.Update(msg)
	}

	return m, cmd
}

// updateFileStep handles all key input for step 0.
func (m SetupModel) updateFileStep(msg tea.KeyMsg) (SetupModel, tea.Cmd) {
	switch msg.String() {
	case "tab":
		if m.browseMode {
			// Exit browse mode, re-focus text input
			m.browseMode = false
			m.fileInput.Focus()
		} else {
			// Enter browse mode: derive start dir from current typed value
			startDir := startDirFromPath(m.fileInput.Value())
			m.fileBrowser = NewFileBrowserModel(startDir)
			m.browseMode = true
			m.fileInput.Blur()
		}
		return m, nil

	case "esc":
		if m.browseMode {
			m.browseMode = false
			m.fileInput.Focus()
		}
		return m, nil

	case "enter":
		if m.browseMode {
			var browserCmd tea.Cmd
			m.fileBrowser, browserCmd = m.fileBrowser.Update(msg)
			if m.fileBrowser.selected != "" {
				m.fileInput.SetValue(m.fileBrowser.selected)
				m.fileBrowser.selected = ""
				m.browseMode = false
				m.fileInput.Focus()
			}
			return m, browserCmd
		}
		return m.handleEnter()

	default:
		if m.browseMode {
			var browserCmd tea.Cmd
			m.fileBrowser, browserCmd = m.fileBrowser.Update(msg)
			if m.fileBrowser.selected != "" {
				m.fileInput.SetValue(m.fileBrowser.selected)
				m.fileBrowser.selected = ""
				m.browseMode = false
				m.fileInput.Focus()
			}
			return m, browserCmd
		}
		var cmd tea.Cmd
		m.fileInput, cmd = m.fileInput.Update(msg)
		return m, cmd
	}
}

// startDirFromPath returns a sensible start directory for the file browser
// based on the currently typed path.
func startDirFromPath(path string) string {
	if path == "" {
		return ""
	}
	// If the path points to an existing file, use its directory
	info, err := os.Stat(path)
	if err == nil && !info.IsDir() {
		return filepath.Dir(path)
	}
	// If it points to an existing directory, use it directly
	if err == nil && info.IsDir() {
		return path
	}
	// Otherwise try the parent directory
	parent := filepath.Dir(path)
	if _, err := os.Stat(parent); err == nil {
		return parent
	}
	return ""
}

func (m SetupModel) handleEnter() (SetupModel, tea.Cmd) {
	switch m.step {
	case 0: // File path
		if m.fileInput.Value() != "" {
			fp := m.fileInput.Value()

			// Basic file validation — use "smt" mode so ValidateSetupConfig
			// checks only file existence/permissions, not solver configuration.
			testConfig := runner.CompilationConfig{
				Filepath: fp,
				Mode:     "smt",
			}
			if err := ValidateSetupConfig(testConfig); err != nil {
				m.validationErr = err.Error()
				return m, nil
			}

			// Clear validation error and proceed
			m.validationErr = ""
			m.config.Filepath = fp
			m.step++
			m.cursor = 0 // Default to "model (recommended)"
		}
	case 1: // Mode
		active := m.activeModes()
		if m.cursor < len(active) {
			label := active[m.cursor]
			switch {
			case strings.HasPrefix(label, "model"):
				m.config.Mode = "model"
			case strings.HasPrefix(label, "ast"):
				m.config.Mode = "ast"
			case strings.HasPrefix(label, "ir"):
				m.config.Mode = "ir"
			case strings.HasPrefix(label, "smt"):
				m.config.Mode = "smt"
			}
		}
		m.step++
		m.cursor = 0 // Default to "fault (recommended)"
	case 2: // Input
		switch m.cursor {
		case 0:
			m.config.Input = "fault"
		case 1:
			m.config.Input = "ll"
		case 2:
			m.config.Input = "smt2"
		}
		m.step++
		m.cursor = 0 // Default to "text (recommended)"
	case 3: // Output
		switch m.cursor {
		case 0:
			m.config.Output = "text"
		case 1:
			m.config.Output = "smt"
		}
		// Setup complete, send message to parent
		return m, func() tea.Msg {
			return SetupCompleteMsg{Config: m.config}
		}
	}
	return m, nil
}

func (m SetupModel) getOptionsCount() int {
	switch m.step {
	case 1:
		return len(m.activeModes())
	case 2:
		return len(inputs)
	case 3:
		return len(outputs)
	}
	return 0
}

func (m SetupModel) View() string {
	var b strings.Builder

	// Header with step indicator
	stepNames := []string{"File Selection", "Compilation Mode", "Input Format", "Output Format"}
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(SetupBorderColor).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderBottom(true).
		BorderForeground(SetupBorderColor)

	header := headerStyle.Render(fmt.Sprintf(" Fault Setup ⚙ Step %d/4: %s ", m.step+1, stepNames[m.step]))
	b.WriteString(header)
	b.WriteString("\n\n")

	// Content area
	switch m.step {
	case 0:
		b.WriteString(PromptStyle.Render("Enter the path to the file to compile:"))
		b.WriteString("\n\n")
		b.WriteString(m.fileInput.View())
		b.WriteString("\n\n")

		if m.browseMode {
			browserBox := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(SetupBorderColor).
				Padding(0, 1)
			b.WriteString(browserBox.Render(m.fileBrowser.View()))
			b.WriteString("\n")
		} else if m.validationErr != "" {
			// Show validation error only when not browsing
			b.WriteString(ErrorStyle.Render("⚠ " + m.validationErr))
			b.WriteString("\n\n")
		}

	case 1:
		b.WriteString(SuccessStyle.Render("✓ File: " + m.config.Filepath))
		b.WriteString("\n\n")
		if !m.solverAvailable {
			b.WriteString(ErrorStyle.Render("⚠ No solver configured — model mode unavailable."))
			b.WriteString("\n")
			b.WriteString(UnselectedStyle.Render("  Set SOLVERCMD=z3 and SOLVERARG=-in to enable it."))
			b.WriteString("\n\n")
		}
		b.WriteString(PromptStyle.Render("Select compilation mode:"))
		b.WriteString("\n\n")
		b.WriteString(m.renderOptions(m.activeModes()))
		b.WriteString("\n")

	case 2:
		b.WriteString(SuccessStyle.Render("✓ File: " + m.config.Filepath))
		b.WriteString("\n")
		b.WriteString(SuccessStyle.Render("✓ Mode: " + m.config.Mode))
		b.WriteString("\n\n")
		b.WriteString(PromptStyle.Render("Select input format:"))
		b.WriteString("\n\n")
		b.WriteString(m.renderOptions(inputs))
		b.WriteString("\n")

	case 3:
		b.WriteString(SuccessStyle.Render("✓ File: " + m.config.Filepath))
		b.WriteString("\n")
		b.WriteString(SuccessStyle.Render("✓ Mode: " + m.config.Mode))
		b.WriteString("\n")
		b.WriteString(SuccessStyle.Render("✓ Input: " + m.config.Input))
		b.WriteString("\n\n")
		b.WriteString(PromptStyle.Render("Select output format:"))
		b.WriteString("\n\n")
		b.WriteString(m.renderOptions(outputs))
		b.WriteString("\n")
	}

	// Footer
	width := m.width
	if width == 0 {
		width = 80
	}
	line := strings.Repeat("─", max(width-8, 40))
	footerTop := lipgloss.NewStyle().
		Foreground(SetupBorderColor).
		Render(line)
	b.WriteString("\n")
	b.WriteString(footerTop)
	b.WriteString("\n")

	// Help text based on step and mode
	var helpText string
	switch {
	case m.step == 0 && m.browseMode:
		helpText = " j/k: navigate • l/Enter: open • h: parent • Tab/Esc: back to input • .: hidden • ~: home "
	case m.step == 0:
		helpText = " Enter: continue • Tab: browse files • ctrl+c/ctrl+q: quit • ctrl+t: theme "
	default:
		helpText = " ↑↓/j/k: navigate • Enter: select • ctrl+c/ctrl+q: quit • ctrl+t: theme "
	}

	footerStyle := lipgloss.NewStyle().
		Foreground(InfoStyle.GetForeground())

	footer := footerStyle.Render(helpText)
	b.WriteString(footer)

	return lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(SetupBorderColor).
		Render(b.String())
}

func (m SetupModel) renderOptions(options []string) string {
	var b strings.Builder

	for i, option := range options {
		cursor := "  "
		if i == m.cursor {
			cursor = "❯ "
		}

		if i == m.cursor {
			b.WriteString(SelectedStyle.Render(cursor + option))
		} else {
			b.WriteString(UnselectedStyle.Render(cursor + option))
		}
		b.WriteString("\n")
	}

	return b.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
