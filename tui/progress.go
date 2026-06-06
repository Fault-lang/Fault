package tui

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/progress"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type PhaseStatus struct {
	name    string
	status  string
	percent float64
	done    bool
}

type ProgressModel struct {
	phases       [7]PhaseStatus
	overallBar   progress.Model
	spinner      spinner.Model
	currentPhase int
	filepath     string
	width        int
	height       int
}

func NewProgressModel(filepath string) ProgressModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = PhaseActiveStyle

	p := progress.New(progress.WithColors(GradientStart, GradientEnd), progress.WithWidth(40))

	return ProgressModel{
		phases: [7]PhaseStatus{
			{name: "Parsing", status: "Pending", percent: 0.0, done: false},
			{name: "Preprocessing", status: "Pending", percent: 0.0, done: false},
			{name: "Type Checking", status: "Pending", percent: 0.0, done: false},
			{name: "LLVM IR Generation", status: "Pending", percent: 0.0, done: false},
			{name: "SMT Generation", status: "Pending", percent: 0.0, done: false},
			{name: "Model Checking", status: "Pending", percent: 0.0, done: false},
			{name: "Results Processing", status: "Pending", percent: 0.0, done: false},
		},
		overallBar:   p,
		spinner:      s,
		currentPhase: -1,
		filepath:     filepath,
	}
}

func (m ProgressModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m ProgressModel) Update(msg tea.Msg) (ProgressModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.overallBar.SetWidth(min(60, msg.Width-20))
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case ProgressUpdateMsg:
		update := msg.Update
		phaseIdx := int(update.Phase)

		if phaseIdx >= 0 && phaseIdx < len(m.phases) {
			m.phases[phaseIdx].status = update.Status
			m.phases[phaseIdx].percent = update.Percent
			m.phases[phaseIdx].done = update.Done

			if !update.Done {
				m.currentPhase = phaseIdx
			}
		}

		return m, nil
	}

	return m, nil
}

func (m ProgressModel) View() string {
	var b strings.Builder

	// Title
	title := TitleStyle.Render(" Fault Compiler ")
	b.WriteString(title)
	b.WriteString("\n\n")

	// Filepath
	b.WriteString(InfoStyle.Render(fmt.Sprintf("Compiling: %s", m.filepath)))
	b.WriteString("\n\n")

	// Phase list
	for i, phase := range m.phases {
		var symbol, statusText string
		var style lipgloss.Style

		if phase.done {
			symbol = "✓"
			statusText = phase.name
			style = PhaseDoneStyle
		} else if i == m.currentPhase {
			symbol = m.spinner.View()
			statusText = fmt.Sprintf("%s... %s", phase.name, phase.status)
			style = PhaseActiveStyle
		} else {
			symbol = "⋯"
			statusText = phase.name
			style = PhasePendingStyle
		}

		line := fmt.Sprintf("%s %s", symbol, statusText)
		b.WriteString(style.Render(line))
		b.WriteString("\n")
	}

	b.WriteString("\n")

	// Overall progress bar
	overallPercent := 0.0
	for _, phase := range m.phases {
		overallPercent += phase.percent
	}
	overallPercent = overallPercent / float64(len(m.phases))

	b.WriteString(SubtitleStyle.Render("Overall Progress:"))
	b.WriteString("\n")
	b.WriteString(m.overallBar.ViewAs(overallPercent))
	b.WriteString(fmt.Sprintf(" %.0f%%\n", overallPercent*100))

	return lipgloss.NewStyle().
		Padding(2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ProgressBorderColor).
		Render(b.String())
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
