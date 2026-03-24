package tui

import (
	"fault/runner"
	"fmt"
	"strings"

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
	config       *runner.CompilationConfig
	output       *runner.CompilationOutput
	err          error
	enhancedErr  *EnhancedError
	errorPhase   runner.ProgressPhase
	errorCursor  int // For error view actions
	darkMode     bool // Theme toggle
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
		// Validate configuration before starting compilation
		if err := ValidateSetupConfig(msg.Config); err != nil {
			return m, func() tea.Msg {
				return ValidationErrorMsg{Error: err}
			}
		}

		// Transition to progress view and start compilation
		m.config = &msg.Config
		m.state = ViewProgress
		m.progress = NewProgressModel(msg.Config.Filepath)

		return m, tea.Batch(
			m.progress.Init(),
			startCompilation(msg.Config),
		)

	case ProgressUpdateMsg:
		// Handle progress update and schedule the next read from the channel.
		if m.state == ViewProgress {
			var cmd tea.Cmd
			m.progress, cmd = m.progress.Update(msg)
			return m, tea.Batch(cmd, waitForProgress(msg.progressCh, msg.resultCh))
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
		// Transition to error view with enhanced error info
		m.err = msg.Error
		m.errorPhase = msg.Phase
		m.enhancedErr = CategorizeError(msg.Error, msg.Phase)
		m.state = ViewError
		m.errorCursor = 0 // Default to "Retry"
		return m, nil

	case ValidationErrorMsg:
		// Show validation error but stay in setup view
		m.err = msg.Error
		m.errorPhase = runner.PhaseParsing
		m.enhancedErr = CategorizeError(msg.Error, runner.PhaseParsing)
		m.state = ViewError
		m.errorCursor = 1 // Default to "Back to Setup"
		return m, nil

	case RetryCompilationMsg:
		// Retry compilation with same config
		if m.config != nil {
			m.state = ViewProgress
			m.progress = NewProgressModel(m.config.Filepath)
			m.err = nil
			m.enhancedErr = nil
			return m, tea.Batch(
				m.progress.Init(),
				startCompilation(*m.config),
			)
		}
		return m, nil

	case BackToSetupMsg:
		// Go back to setup view
		m.state = ViewSetup
		m.setup = NewSetupModel()
		m.err = nil
		m.enhancedErr = nil
		m.config = nil
		return m, m.setup.Init()
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
	case ViewError:
		// Handle error view navigation
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			switch keyMsg.String() {
			case "up", "k":
				m.errorCursor--
				if m.errorCursor < 0 {
					m.errorCursor = 2 // Wrap to "Quit"
				}
			case "down", "j":
				m.errorCursor++
				if m.errorCursor > 2 {
					m.errorCursor = 0 // Wrap to "Retry"
				}
			case "enter":
				switch m.errorCursor {
				case 0: // Retry
					return m, func() tea.Msg { return RetryCompilationMsg{} }
				case 1: // Back to Setup
					return m, func() tea.Msg { return BackToSetupMsg{} }
				case 2: // Quit
					return m, tea.Quit
				}
			case "r":
				return m, func() tea.Msg { return RetryCompilationMsg{} }
			case "b":
				return m, func() tea.Msg { return BackToSetupMsg{} }
			case "q":
				return m, tea.Quit
			}
		}
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
	var b strings.Builder

	// Title with phase information
	phaseName := GetPhaseName(m.errorPhase)
	title := TitleStyle.Render(fmt.Sprintf(" ✗ Compilation Failed at %s ", phaseName))
	b.WriteString(title)
	b.WriteString("\n\n")

	// Error message
	if m.enhancedErr != nil {
		b.WriteString(ErrorStyle.Render("Error:"))
		b.WriteString(" ")
		b.WriteString(m.enhancedErr.Message)
		b.WriteString("\n\n")

		// Additional detail if available
		if m.enhancedErr.Detail != "" {
			b.WriteString(InfoStyle.Render("Details: "))
			b.WriteString(m.enhancedErr.Detail)
			b.WriteString("\n\n")
		}

		// Suggestion
		if m.enhancedErr.Suggestion != "" {
			b.WriteString(SubtitleStyle.Render("💡 Suggestion:"))
			b.WriteString("\n")
			b.WriteString(m.enhancedErr.Suggestion)
			b.WriteString("\n\n")
		}
	} else {
		b.WriteString(ErrorStyle.Render(fmt.Sprintf("Error: %v", m.err)))
		b.WriteString("\n\n")
	}

	// Phase progress visualization
	b.WriteString(RenderPhaseProgress(m.errorPhase))
	b.WriteString("\n\n")

	// Action buttons
	b.WriteString(SubtitleStyle.Render("What would you like to do?"))
	b.WriteString("\n\n")

	actions := []string{"Retry", "Back to Setup", "Quit"}
	for i, action := range actions {
		cursor := "  "
		if i == m.errorCursor {
			cursor = "❯ "
		}

		if i == m.errorCursor {
			b.WriteString(SelectedStyle.Render(cursor + action))
		} else {
			b.WriteString(UnselectedStyle.Render(cursor + action))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(InfoStyle.Render("↑/↓ or j/k to navigate • Enter to select • [R]etry • [B]ack • [Q]uit"))

	return lipgloss.NewStyle().
		Padding(2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF0000")).
		Render(b.String())
}

// startCompilation runs the runner in a goroutine and returns a Cmd that
// begins reading progress updates. Each ProgressUpdateMsg chains the next
// read, keeping the UI responsive throughout compilation.
func startCompilation(config runner.CompilationConfig) tea.Cmd {
	progressCh := make(chan runner.ProgressUpdate, 20)
	resultCh := make(chan *runner.CompilationOutput, 1)

	go func() {
		r := runner.NewRunner(config, progressCh)
		output := r.Run()
		resultCh <- output
		close(progressCh)
	}()

	return waitForProgress(progressCh, resultCh)
}

// waitForProgress returns a Cmd that blocks on the next value from progressCh.
// When progressCh is closed the runner is done, so it reads the final output
// from resultCh and returns the appropriate completion or error message.
func waitForProgress(progressCh <-chan runner.ProgressUpdate, resultCh <-chan *runner.CompilationOutput) tea.Cmd {
	return func() tea.Msg {
		update, ok := <-progressCh
		if !ok {
			output := <-resultCh
			if output.Error != nil {
				return CompilationErrorMsg{Error: output.Error, Phase: output.ErrorPhase}
			}
			return CompilationCompleteMsg{Output: output}
		}
		return ProgressUpdateMsg{Update: update, progressCh: progressCh, resultCh: resultCh}
	}
}
