package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// faultFileExts are extensions highlighted as fault-related files
var faultFileExts = map[string]bool{
	".fspec": true,
	".ll":    true,
	".smt2":  true,
}

type FileBrowserModel struct {
	currentDir string
	entries    []os.DirEntry
	cursor     int
	offset     int
	height     int  // max visible rows
	showHidden bool
	selected   string // non-empty when user confirms a file
	err        string
}

func NewFileBrowserModel(startDir string) FileBrowserModel {
	m := FileBrowserModel{
		height: 10,
	}
	if startDir == "" {
		var err error
		startDir, err = os.Getwd()
		if err != nil {
			startDir = "."
		}
	}
	m.loadDir(startDir)
	return m
}

func (m *FileBrowserModel) loadDir(dir string) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		m.err = err.Error()
		return
	}
	entries, err := os.ReadDir(abs)
	if err != nil {
		m.err = err.Error()
		return
	}
	m.err = ""
	m.currentDir = abs
	m.cursor = 0
	m.offset = 0

	if m.showHidden {
		m.entries = entries
	} else {
		filtered := entries[:0]
		for _, e := range entries {
			if !strings.HasPrefix(e.Name(), ".") {
				filtered = append(filtered, e)
			}
		}
		m.entries = filtered
	}
}

func (m FileBrowserModel) Update(msg tea.Msg) (FileBrowserModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.cursor++
			if m.cursor >= len(m.entries) {
				m.cursor = len(m.entries) - 1
			}
			m.clampOffset()

		case "k", "up":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = 0
			}
			m.clampOffset()

		case "g":
			m.cursor = 0
			m.offset = 0

		case "G":
			m.cursor = len(m.entries) - 1
			m.clampOffset()

		case "h", "left":
			parent := filepath.Dir(m.currentDir)
			if parent != m.currentDir {
				m.loadDir(parent)
			}

		case "l", "right", "enter":
			if len(m.entries) == 0 {
				break
			}
			entry := m.entries[m.cursor]
			target := filepath.Join(m.currentDir, entry.Name())
			if entry.IsDir() {
				m.loadDir(target)
			} else {
				m.selected = target
			}

		case "~":
			home, err := os.UserHomeDir()
			if err == nil {
				m.loadDir(home)
			}

		case ".":
			m.showHidden = !m.showHidden
			m.loadDir(m.currentDir)
		}
	}
	return m, nil
}

// clampOffset adjusts m.offset so cursor stays within the visible window.
func (m *FileBrowserModel) clampOffset() {
	if m.cursor < m.offset {
		m.offset = m.cursor
	}
	if m.cursor >= m.offset+m.height {
		m.offset = m.cursor - m.height + 1
	}
}

func (m FileBrowserModel) View() string {
	var b strings.Builder

	// Header: current directory
	b.WriteString(PromptStyle.Render(m.currentDir))
	b.WriteString("\n")

	if m.err != "" {
		b.WriteString(ErrorStyle.Render(m.err))
		return b.String()
	}

	if len(m.entries) == 0 {
		b.WriteString(InfoStyle.Render("  (empty directory)"))
		return b.String()
	}

	// Visible slice
	end := m.offset + m.height
	if end > len(m.entries) {
		end = len(m.entries)
	}

	for i := m.offset; i < end; i++ {
		entry := m.entries[i]
		name := entry.Name()
		if entry.IsDir() {
			name = name + "/"
		}

		if i == m.cursor {
			b.WriteString(SelectedStyle.Render("❯ " + name))
		} else if entry.IsDir() {
			b.WriteString(BrowserDirStyle.Render("  " + name))
		} else if faultFileExts[filepath.Ext(entry.Name())] {
			b.WriteString(BrowserFaultFileStyle.Render("  " + name))
		} else {
			b.WriteString(UnselectedStyle.Render("  " + name))
		}
		b.WriteString("\n")
	}

	// Scroll indicator
	if len(m.entries) > m.height {
		lo := m.offset + 1
		hi := end
		total := len(m.entries)
		b.WriteString(InfoStyle.Render(fmt.Sprintf("  %d-%d / %d", lo, hi, total)))
		b.WriteString("\n")
	}

	return b.String()
}
