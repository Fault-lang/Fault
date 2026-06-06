package tui

import (
	"fault/runner"
	"fmt"
	"strings"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type ViewState int

const (
	ViewSetup ViewState = iota
	ViewProgress
	ViewResults
	ViewError
	ViewConfirmLargeSMT
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

	// Error view scrollable body
	errorViewport viewport.Model
	errorReady    bool

	// Large SMT confirmation state
	largeSMT LargeSMTWarningMsg
	largeSMTCursor int // 0 = Proceed, 1 = Abort
}

func NewModel() Model {
	return Model{
		state:    ViewSetup,
		setup:    NewSetupModel(),
		darkMode: true, // Overridden by BackgroundColorMsg on first frame
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.setup.Init(), func() tea.Msg { return tea.RequestBackgroundColor() })
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
		case ViewError:
			m = m.initErrorViewport(msg.Width, msg.Height)
		}

		return m, cmd

	case tea.BackgroundColorMsg:
		m.darkMode = msg.IsDark()
		if m.darkMode {
			ApplyTheme(DarkTheme())
		} else {
			ApplyTheme(LightTheme())
		}
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+q":
			return m, tea.Quit
		case "ctrl+t":
			// Manual override
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
		// Large SMT warning: pause and ask the user before sending to solver.
		if msg.Update.Phase == runner.PhaseConfirmLargeSMT {
			m.largeSMT = LargeSMTWarningMsg{
				SMTLines:   msg.Update.SMTLines,
				ConfirmCh:  msg.Update.ConfirmCh,
				ProgressCh: msg.progressCh,
				ResultCh:   msg.resultCh,
			}
			m.largeSMTCursor = 0
			m.state = ViewConfirmLargeSMT
			return m, nil
		}
		// Handle progress update and schedule the next read from the channel.
		if m.state == ViewProgress {
			var cmd tea.Cmd
			m.progress, cmd = m.progress.Update(msg)
			return m, tea.Batch(cmd, waitForProgress(msg.progressCh, msg.resultCh))
		}

	case CompilationCompleteMsg:
		// Ignore the drain message when the user cancelled a large-SMT run.
		if msg.Output.LargeSMTLines > 0 {
			return m, nil
		}
		// Transition to results view
		m.output = msg.Output
		m.state = ViewResults
		m.results = NewResultsModel(
			msg.Output.ResultLog,
			msg.Output.Asserts,
			msg.Output.Warnings,
			msg.Output.AST,
			msg.Output.SMT,
			msg.Output.IR,
			msg.Output.Message,
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
		w, h := m.width, m.height
		if w == 0 {
			w = 80
		}
		if h == 0 {
			h = 24
		}
		m = m.initErrorViewport(w, h)
		return m, nil

	case ValidationErrorMsg:
		// Show validation error but stay in setup view
		m.err = msg.Error
		m.errorPhase = runner.PhaseParsing
		m.enhancedErr = CategorizeError(msg.Error, runner.PhaseParsing)
		m.state = ViewError
		m.errorCursor = 1 // Default to "Back to Setup"
		w, h := m.width, m.height
		if w == 0 {
			w = 80
		}
		if h == 0 {
			h = 24
		}
		m = m.initErrorViewport(w, h)
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
	case ViewConfirmLargeSMT:
		if keyMsg, ok := msg.(tea.KeyPressMsg); ok {
			switch keyMsg.String() {
			case "up", "k":
				if m.largeSMTCursor > 0 {
					m.largeSMTCursor--
				}
			case "down", "j":
				if m.largeSMTCursor < 1 {
					m.largeSMTCursor++
				}
			case "enter":
				if m.largeSMTCursor == 0 {
					// Proceed: unblock runner, resume progress view.
					m.largeSMT.ConfirmCh <- true
					m.state = ViewProgress
					return m, waitForProgress(m.largeSMT.ProgressCh, m.largeSMT.ResultCh)
				}
				// Abort: unblock runner (it returns early), go back to setup.
				m.largeSMT.ConfirmCh <- false
				m.state = ViewSetup
				m.setup = NewSetupModel()
				m.config = nil
				return m, m.setup.Init()
			case "y":
				m.largeSMT.ConfirmCh <- true
				m.state = ViewProgress
				return m, waitForProgress(m.largeSMT.ProgressCh, m.largeSMT.ResultCh)
			case "n", "q":
				m.largeSMT.ConfirmCh <- false
				m.state = ViewSetup
				m.setup = NewSetupModel()
				m.config = nil
				return m, m.setup.Init()
			}
		}
	case ViewError:
		// Handle error view navigation
		if keyMsg, ok := msg.(tea.KeyPressMsg); ok {
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
			default:
				// Forward page/scroll keys to the error body viewport
				if m.errorReady {
					m.errorViewport, _ = m.errorViewport.Update(keyMsg)
				}
			}
		}
	}

	return m, cmd
}

func (m Model) View() tea.View {
	var content string
	switch m.state {
	case ViewSetup:
		content = m.setup.View()
	case ViewProgress:
		content = m.progress.View()
	case ViewResults:
		content = m.results.View()
	case ViewError:
		content = m.errorView()
	case ViewConfirmLargeSMT:
		content = m.confirmLargeSMTView()
	}
	v := tea.NewView(content)
	v.AltScreen = true
	return v
}

func (m Model) confirmLargeSMTView() string {
	var b strings.Builder

	title := TitleStyle.Render(" ⚠ Large SMT Formula ")
	b.WriteString(title)
	b.WriteString("\n\n")

	b.WriteString(fmt.Sprintf("The generated SMT formula is %s lines long.\n", WarningStyle.Render(fmt.Sprintf("%d", m.largeSMT.SMTLines))))
	b.WriteString("Sending a formula this large to the solver may take a very long time.\n\n")

	actions := []string{"Proceed with model checking", "Abort and go back"}
	for i, action := range actions {
		cursor := "  "
		if i == m.largeSMTCursor {
			cursor = "❯ "
		}
		if i == m.largeSMTCursor {
			b.WriteString(SelectedStyle.Render(cursor + action))
		} else {
			b.WriteString(UnselectedStyle.Render(cursor + action))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(InfoStyle.Render("↑/↓ or j/k to navigate • Enter to select • [Y]es • [N]o"))

	return lipgloss.NewStyle().
		Padding(2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FFA500")).
		Render(b.String())
}

// initErrorViewport (re)initialises the error body viewport for the given
// terminal dimensions. It is called on WindowSizeMsg and whenever a new error
// is set (so the content is always fresh).
func (m Model) initErrorViewport(w, h int) Model {
	// Build the scrollable body: error details + phase progress.
	var body strings.Builder
	if m.enhancedErr != nil {
		body.WriteString(ErrorStyle.Render("Error:"))
		body.WriteString(" ")
		body.WriteString(m.enhancedErr.Message)
		body.WriteString("\n\n")
		if m.enhancedErr.Detail != "" {
			body.WriteString(InfoStyle.Render("Details: "))
			body.WriteString(m.enhancedErr.Detail)
			body.WriteString("\n\n")
		}
		if m.enhancedErr.Suggestion != "" {
			body.WriteString(SubtitleStyle.Render("💡 Suggestion:"))
			body.WriteString("\n")
			body.WriteString(m.enhancedErr.Suggestion)
			body.WriteString("\n\n")
		}
	} else if m.err != nil {
		body.WriteString(ErrorStyle.Render(fmt.Sprintf("Error: %v", m.err)))
		body.WriteString("\n\n")
	}
	body.WriteString(RenderPhaseProgress(m.errorPhase))

	// Reserve lines for: outer border(2) + outer padding(4) + title(1) + blank(1)
	// + actions header(1) + blank(1) + 3 actions(3) + blank(1) + help(1) = 15
	const fixedLines = 15
	vpH := h - fixedLines
	if vpH < 3 {
		vpH = 3
	}
	vpW := w - 8 // account for border + padding
	if vpW < 20 {
		vpW = 20
	}

	vp := viewport.New(viewport.WithWidth(vpW), viewport.WithHeight(vpH))
	vp.SetContent(body.String())
	m.errorViewport = vp
	m.errorReady = true
	return m
}

func (m Model) errorView() string {
	var b strings.Builder

	// Title with phase information
	phaseName := GetPhaseName(m.errorPhase)
	title := TitleStyle.Render(fmt.Sprintf(" ✗ Compilation Failed at %s ", phaseName))
	b.WriteString(title)
	b.WriteString("\n\n")

	// Scrollable error body
	if m.errorReady {
		b.WriteString(m.errorViewport.View())
	} else {
		// Fallback before first WindowSizeMsg
		if m.enhancedErr != nil {
			b.WriteString(ErrorStyle.Render("Error: "))
			b.WriteString(m.enhancedErr.Message)
		} else if m.err != nil {
			b.WriteString(ErrorStyle.Render(fmt.Sprintf("Error: %v", m.err)))
		}
	}
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
	b.WriteString(InfoStyle.Render("↑/↓/j/k: menu • pgup/pgdn/u/d: scroll error • Enter to select • [R]etry • [B]ack • [Q]uit"))

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
