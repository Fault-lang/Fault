package tui

import (
	"fault/runner"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ViewState int

const (
	ViewSetup ViewState = iota
	ViewProgress
	ViewResults
	ViewError
)

type Model struct {
	state  ViewState
	width  int
	height int

	// Sub-models
	setup    SetupModel
	progress ProgressModel
	results  ResultsModel

	// Shared data
	config   *runner.CompilationConfig
	output   *runner.CompilationOutput
	err      error
	darkMode bool // Theme toggle
}

func NewModel() Model {
	return Model{
		state:    ViewSetup,
		setup:    NewSetupModel(),
		darkMode: true, // Start in dark mode
	}
}

func (m Model) Init() tea.Cmd {
	return m.setup.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Forward to current view
		var cmd tea.Cmd
		switch m.state {
		case ViewSetup:
			m.setup.width = msg.Width
			m.setup.height = msg.Height
		case ViewProgress:
			m.progress.width = msg.Width
			m.progress.height = msg.Height
		case ViewResults:
			m.results, cmd = m.results.Update(msg)
		}

		return m, cmd

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+q":
			return m, tea.Quit
		case "ctrl+t":
			// Toggle theme
			m.darkMode = !m.darkMode
			if m.darkMode {
				ApplyTheme(DarkTheme())
			} else {
				ApplyTheme(LightTheme())
			}
			return m, nil
		}

	case SetupCompleteMsg:
		// Transition to progress view and start compilation
		m.config = &msg.Config
		m.state = ViewProgress
		m.progress = NewProgressModel(msg.Config.Filepath)

		return m, tea.Batch(
			m.progress.Init(),
			startCompilation(msg.Config),
		)

	case ProgressUpdateMsg:
		// Handle progress update
		if m.state == ViewProgress {
			var cmd tea.Cmd
			m.progress, cmd = m.progress.Update(msg)
			return m, cmd
		}

	case CompilationCompleteMsg:
		// Transition to results view
		m.output = msg.Output
		m.state = ViewResults
		m.results = NewResultsModel(
			msg.Output.ResultLog,
			msg.Output.AST,
			msg.Output.SMT,
			msg.Output.IR,
			m.config.Mode,
		)

		// Use current dimensions or defaults if not set yet
		width := m.width
		height := m.height
		if width == 0 {
			width = 80
		}
		if height == 0 {
			height = 24
		}

		m.results.width = width
		m.results.height = height

		// Trigger window size update to initialize viewport
		var cmd tea.Cmd
		m.results, cmd = m.results.Update(tea.WindowSizeMsg{Width: width, Height: height})
		return m, cmd

	case CompilationErrorMsg:
		// Transition to error view
		m.err = msg.Error
		m.state = ViewError
		return m, nil
	}

	// Forward updates to current view
	var cmd tea.Cmd
	switch m.state {
	case ViewSetup:
		m.setup, cmd = m.setup.Update(msg)
	case ViewProgress:
		m.progress, cmd = m.progress.Update(msg)
	case ViewResults:
		m.results, cmd = m.results.Update(msg)
	}

	return m, cmd
}

func (m Model) View() string {
	switch m.state {
	case ViewSetup:
		return m.setup.View()
	case ViewProgress:
		return m.progress.View()
	case ViewResults:
		return m.results.View()
	case ViewError:
		return m.errorView()
	}
	return ""
}

func (m Model) errorView() string {
	title := TitleStyle.Render(" ✗ Compilation Failed ")
	errorMsg := ErrorStyle.Render(fmt.Sprintf("\nError: %v\n", m.err))
	help := InfoStyle.Render("\nPress q or Ctrl+C to exit")

	content := title + "\n\n" + errorMsg + help

	return lipgloss.NewStyle().
		Padding(2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF0000")).
		Render(content)
}

// startCompilation runs the compilation
// For now, we run synchronously and return the result
// TODO: Implement async progress updates using tea.Program.Send()
func startCompilation(config runner.CompilationConfig) tea.Cmd {
	return func() tea.Msg {
		// Run compilation without progress updates for now
		// The spinner in progress view will show activity
		r := runner.NewRunner(config, nil)
		output := r.Run()

		// Return final result
		if output.Error != nil {
			return CompilationErrorMsg{Error: output.Error}
		}

		return CompilationCompleteMsg{Output: output}
	}
}
