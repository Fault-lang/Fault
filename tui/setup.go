package tui

import (
	"fault/runner"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SetupModel struct {
	step          int // 0=file, 1=mode, 2=input, 3=output
	fileInput     textinput.Model
	cursor        int
	config        runner.CompilationConfig
	width         int
	height        int
	validationErr string // Error message for file validation
}

var (
	modes   = []string{"model (recommended)", "ast", "ir", "smt"}
	inputs  = []string{"fault (recommended)", "ll", "smt2"}
	outputs = []string{"text (recommended)", "smt"}
)

func NewSetupModel() SetupModel {
	ti := textinput.New()
	ti.Placeholder = "path/to/file.fspec"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 50

	return SetupModel{
		step:      0,
		fileInput: ti,
		cursor:    0,
		config: runner.CompilationConfig{
			Mode:   "model",
			Input:  "fault",
			Output: "text",
			Reach:  false,
		},
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

		case "enter":
			return m.handleEnter()

		case "up":
			if m.step > 0 {
				m.cursor--
				if m.cursor < 0 {
					m.cursor = m.getOptionsCount() - 1
				}
			}

		case "down":
			if m.step > 0 {
				m.cursor++
				if m.cursor >= m.getOptionsCount() {
					m.cursor = 0
				}
			}
		}
	}

	// Update file input if on step 0
	if m.step == 0 {
		m.fileInput, cmd = m.fileInput.Update(msg)
	}

	return m, cmd
}

func (m SetupModel) handleEnter() (SetupModel, tea.Cmd) {
	switch m.step {
	case 0: // File path
		if m.fileInput.Value() != "" {
			filepath := m.fileInput.Value()

			// Basic file validation
			testConfig := runner.CompilationConfig{
				Filepath: filepath,
				Mode:     "model",
			}
			if err := ValidateSetupConfig(testConfig); err != nil {
				m.validationErr = err.Error()
				return m, nil
			}

			// Clear validation error and proceed
			m.validationErr = ""
			m.config.Filepath = filepath
			m.step++
			m.cursor = 0 // Default to "model (recommended)"
		}
	case 1: // Mode
		switch m.cursor {
		case 0:
			m.config.Mode = "model"
		case 1:
			m.config.Mode = "ast"
		case 2:
			m.config.Mode = "ir"
		case 3:
			m.config.Mode = "smt"
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
		return len(modes)
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

		// Show validation error if any
		if m.validationErr != "" {
			b.WriteString(ErrorStyle.Render("⚠ " + m.validationErr))
			b.WriteString("\n\n")
		}

	case 1:
		b.WriteString(SuccessStyle.Render("✓ File: " + m.config.Filepath))
		b.WriteString("\n\n")
		b.WriteString(PromptStyle.Render("Select compilation mode:"))
		b.WriteString("\n\n")
		b.WriteString(m.renderOptions(modes))
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

	// Help text based on step
	var helpText string
	if m.step == 0 {
		helpText = " Enter: continue • ctrl+c/ctrl+q: quit • ctrl+t: theme "
	} else {
		helpText = " ↑↓: navigate • Enter: select • ctrl+c/ctrl+q: quit • ctrl+t: theme "
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
