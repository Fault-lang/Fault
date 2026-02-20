package tui

import (
	"fault/ast"
	"fault/generator/scenario"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ResultsModel struct {
	viewport viewport.Model
	logger   *scenario.Logger
	ast      *ast.Spec
	smt      string
	ir       string
	content  string
	ready    bool
	width    int
	height   int
	mode     string
}

func NewResultsModel(logger *scenario.Logger, astSpec *ast.Spec, smt string, ir string, mode string) ResultsModel {
	return ResultsModel{
		logger: logger,
		ast:    astSpec,
		smt:    smt,
		ir:     ir,
		mode:   mode,
	}
}

func (m ResultsModel) Init() tea.Cmd {
	return nil
}

func (m ResultsModel) Update(msg tea.Msg) (ResultsModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		headerHeight := 4
		footerHeight := 3
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width-4, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.ready = true

			// Set content based on mode
			m.content = m.getContent()
			m.viewport.SetContent(m.content)
		} else {
			m.viewport.Width = msg.Width - 4
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+q", "q", "esc":
			return m, tea.Quit
		case "g", "home":
			m.viewport.GotoTop()
			return m, nil
		case "G", "end":
			m.viewport.GotoBottom()
			return m, nil
		case "h", "pgup", "b":
			m.viewport.PageUp()
			return m, nil
		case "l", "pgdown", "f":
			m.viewport.PageDown()
			return m, nil
		}
	}

	// Update viewport for other keys (arrow keys, j/k, etc.)
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m ResultsModel) getContent() string {
	var content strings.Builder

	// Add styled section header based on output type
	sectionStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(ResultsBorderColor).
		MarginBottom(1)

	divider := DividerStyle.Render(strings.Repeat("─", 80))

	// Check what output is actually available and format accordingly
	if m.logger != nil {
		content.WriteString(sectionStyle.Render("🌋 Fault found the following scenario"))
		content.WriteString("\n")
		content.WriteString(divider)
		content.WriteString("\n\n")
		content.WriteString(m.formatLoggerOutput(m.logger.String()))
	} else if m.smt != "" {
		content.WriteString(sectionStyle.Render("⛰️ SMT Output"))
		content.WriteString("\n")
		content.WriteString(divider)
		content.WriteString("\n\n")
		content.WriteString(m.smt)
	} else if m.ir != "" {
		content.WriteString(sectionStyle.Render("❄️ LLVM IR Output"))
		content.WriteString("\n")
		content.WriteString(divider)
		content.WriteString("\n\n")
		content.WriteString(m.ir)
	} else if m.ast != nil {
		content.WriteString(sectionStyle.Render("🌳 Abstract Syntax Tree"))
		content.WriteString("\n")
		content.WriteString(divider)
		content.WriteString("\n\n")
		content.WriteString(fmt.Sprintf("%v", m.ast))
	} else {
		content.WriteString(InfoStyle.Render("No output available"))
	}

	return content.String()
}

func (m ResultsModel) formatLoggerOutput(output string) string {
	// Add color coding to logger output using global theme-based styles
	lines := strings.Split(output, "\n")
	var formatted strings.Builder

	for _, line := range lines {
		// Color code based on keywords using global styles
		if strings.Contains(line, "✓") || strings.Contains(line, "PASS") || strings.Contains(line, "Success") {
			formatted.WriteString(SuccessStyle.Render(line))
		} else if strings.Contains(line, "✗") || strings.Contains(line, "FAIL") || strings.Contains(line, "Error") {
			formatted.WriteString(ErrorStyle.Render(line))
		} else if strings.Contains(line, "⚠") || strings.Contains(line, "Warning") {
			formatted.WriteString(WarningStyle.Render(line))
		} else if strings.Contains(line, "→") || strings.Contains(line, "->") {
			formatted.WriteString(SubtitleStyle.Render(line))
		} else if strings.Contains(line, "Run function") {
			formatted.WriteString(TitleStyle.Render(line))
		} else {
			formatted.WriteString(line)
		}
		formatted.WriteString("\n")
	}

	return formatted.String()
}

func (m ResultsModel) getOutputType() string {
	if m.logger != nil {
		return "Model Checking Results"
	}
	if m.smt != "" {
		return "SMT Output"
	}
	if m.ir != "" {
		return "LLVM IR"
	}
	if m.ast != nil {
		return "Abstract Syntax Tree"
	}
	return "Output"
}

func (m ResultsModel) getResultsSummary() string {
	if m.logger == nil {
		return ""
	}

	output := m.logger.String()
	passCount := strings.Count(output, "✓")
	failCount := strings.Count(output, "✗")

	if passCount == 0 && failCount == 0 {
		return ""
	}

	summary := fmt.Sprintf(" Tests: %d passed", passCount)
	if failCount > 0 {
		summary += fmt.Sprintf(", %d failed", failCount)
	}

	return summary
}

func (m ResultsModel) View() string {
	if !m.ready {
		var b strings.Builder
		title := TitleStyle.Render(" ✓ Compilation Complete ")
		b.WriteString(title)
		b.WriteString("\n\n")
		b.WriteString(InfoStyle.Render("Preparing results..."))
		b.WriteString("\n\n")
		b.WriteString(InfoStyle.Render(fmt.Sprintf("(Dimensions: %dx%d)", m.width, m.height)))
		return lipgloss.NewStyle().
			Padding(2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ResultsBorderColor).
			Render(b.String())
	}

	var b strings.Builder

	// Header with output type and summary
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(ResultsBorderColor).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderBottom(true).
		BorderForeground(ResultsBorderColor)

	outputType := m.getOutputType()
	header := headerStyle.Render(fmt.Sprintf(" ✓ Compilation Complete • %s ", outputType))
	b.WriteString(header)
	b.WriteString("\n")

	// Add summary if model results are available
	if m.logger != nil {
		summary := m.getResultsSummary()
		if summary != "" {
			summaryStyle := lipgloss.NewStyle().
				Foreground(InfoStyle.GetForeground()).
				Italic(true).
				MarginTop(1).
				MarginBottom(1)
			b.WriteString(summaryStyle.Render(summary))
			b.WriteString("\n")
		}
	}

	// Viewport content
	b.WriteString(m.viewport.View())
	b.WriteString("\n")

	// Footer with status bar
	line := strings.Repeat("─", m.width-4)
	footerTop := lipgloss.NewStyle().
		Foreground(ResultsBorderColor).
		Render(line)
	b.WriteString(footerTop)
	b.WriteString("\n")

	scrollPercent := m.viewport.ScrollPercent()
	helpText := " ↑↓/jk: scroll • g/G: top/bottom • f/b: page • q/esc: quit • ctrl+t: theme "
	percentText := fmt.Sprintf(" %.0f%% ", scrollPercent*100)

	footerStyle := lipgloss.NewStyle().
		Foreground(InfoStyle.GetForeground())

	percentStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(ResultsBorderColor)

	footer := footerStyle.Render(helpText) + percentStyle.Render(percentText)
	b.WriteString(footer)

	return lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ResultsBorderColor).
		Render(b.String())
}
