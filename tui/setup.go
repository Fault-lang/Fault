package tui

import (
	"fault/runner"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SetupModel struct {
	step      int // 0=file, 1=mode, 2=input, 3=output
	fileInput textinput.Model
	cursor    int
	config    runner.CompilationConfig
	width     int
	height    int
}

var (
	modes   = []string{"check (recommended)", "ast", "ir", "smt"}
	inputs  = []string{"fspec (recommended)", "ll", "smt2"}
	outputs = []string{"log (recommended)", "smt", "static", "legacy", "visualize"}
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
			Mode:   "check",
			Input:  "fspec",
			Output: "log",
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

		case "up", "k":
			if m.step > 0 {
				m.cursor--
				if m.cursor < 0 {
					m.cursor = m.getOptionsCount() - 1
				}
			}

		case "down", "j":
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
			m.config.Filepath = m.fileInput.Value()
			m.step++
			m.cursor = 0 // Default to "check (recommended)"
		}
	case 1: // Mode
		switch m.cursor {
		case 0:
			m.config.Mode = "check"
		case 1:
			m.config.Mode = "ast"
		case 2:
			m.config.Mode = "ir"
		case 3:
			m.config.Mode = "smt"
		}
		m.step++
		m.cursor = 0 // Default to "fspec (recommended)"
	case 2: // Input
		switch m.cursor {
		case 0:
			m.config.Input = "fspec"
		case 1:
			m.config.Input = "ll"
		case 2:
			m.config.Input = "smt2"
		}
		m.step++
		m.cursor = 0 // Default to "log (recommended)"
	case 3: // Output
		switch m.cursor {
		case 0:
			m.config.Output = "log"
		case 1:
			m.config.Output = "smt"
		case 2:
			m.config.Output = "static"
		case 3:
			m.config.Output = "legacy"
		case 4:
			m.config.Output = "visualize"
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

	// Title
	title := TitleStyle.Render(" Fault Interactive Compiler ")
	b.WriteString(title)
	b.WriteString("\n\n")

	switch m.step {
	case 0:
		b.WriteString(PromptStyle.Render("Enter the path to the file to compile:"))
		b.WriteString("\n\n")
		b.WriteString(m.fileInput.View())
		b.WriteString("\n\n")
		b.WriteString(InfoStyle.Render("Press Enter to continue"))

	case 1:
		b.WriteString(SubtitleStyle.Render("✓ File: " + m.config.Filepath))
		b.WriteString("\n\n")
		b.WriteString(PromptStyle.Render("Select compilation mode:"))
		b.WriteString("\n\n")
		b.WriteString(m.renderOptions(modes))
		b.WriteString("\n\n")
		b.WriteString(InfoStyle.Render("Use ↑/↓ or j/k to navigate, Enter to select"))

	case 2:
		b.WriteString(SubtitleStyle.Render("✓ File: " + m.config.Filepath))
		b.WriteString("\n")
		b.WriteString(SubtitleStyle.Render("✓ Mode: " + m.config.Mode))
		b.WriteString("\n\n")
		b.WriteString(PromptStyle.Render("Select input format:"))
		b.WriteString("\n\n")
		b.WriteString(m.renderOptions(inputs))
		b.WriteString("\n\n")
		b.WriteString(InfoStyle.Render("Use ↑/↓ or j/k to navigate, Enter to select"))

	case 3:
		b.WriteString(SubtitleStyle.Render("✓ File: " + m.config.Filepath))
		b.WriteString("\n")
		b.WriteString(SubtitleStyle.Render("✓ Mode: " + m.config.Mode))
		b.WriteString("\n")
		b.WriteString(SubtitleStyle.Render("✓ Input: " + m.config.Input))
		b.WriteString("\n\n")
		b.WriteString(PromptStyle.Render("Select output format:"))
		b.WriteString("\n\n")
		b.WriteString(m.renderOptions(outputs))
		b.WriteString("\n\n")
		b.WriteString(InfoStyle.Render("Use ↑/↓ or j/k to navigate, Enter to select"))
	}

	b.WriteString("\n\n")
	b.WriteString(InfoStyle.Render("Ctrl+T to toggle theme • Ctrl+C or Ctrl+Q to quit"))

	return lipgloss.NewStyle().Padding(2).Render(b.String())
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
