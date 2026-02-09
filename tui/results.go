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

		if !m.ready {
			m.viewport = viewport.New(msg.Width-4, msg.Height-10)
			m.viewport.YPosition = 0
			m.ready = true

			// Set content based on mode
			m.content = m.getContent()
			m.viewport.SetContent(m.content)
		} else {
			m.viewport.Width = msg.Width - 4
			m.viewport.Height = msg.Height - 10
		}

		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+q":
			return m, tea.Quit
		}
	}

	// Update viewport
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m ResultsModel) getContent() string {
	// Check what output is actually available
	if m.logger != nil {
		return m.logger.String()
	}
	if m.smt != "" {
		return m.smt
	}
	if m.ir != "" {
		return m.ir
	}
	if m.ast != nil {
		return fmt.Sprintf("%v", m.ast)
	}

	return "No output available"
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
		return lipgloss.NewStyle().Padding(2).Render(b.String())
	}

	var b strings.Builder

	// Title
	title := TitleStyle.Render(" ✓ Compilation Complete ")
	b.WriteString(title)
	b.WriteString("\n\n")

	// Viewport
	b.WriteString(m.viewport.View())
	b.WriteString("\n\n")

	// Status bar
	scrollPercent := m.viewport.ScrollPercent()
	statusBar := InfoStyle.Render(fmt.Sprintf("%.0f%% • ↑/↓ or j/k to scroll • Ctrl+T for theme • Ctrl+C to quit", scrollPercent*100))
	b.WriteString(statusBar)

	return lipgloss.NewStyle().Padding(2).Render(b.String())
}
