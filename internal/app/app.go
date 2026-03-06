package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/young-st511/advimture/internal/editor"
)

// Screen represents the current screen state.
type Screen int

const (
	ScreenEditor Screen = iota
	ScreenMenu
)

// Model is the top-level Bubbletea model that manages screen transitions.
type Model struct {
	screen Screen
	editor editor.Model
	width  int
	height int
}

// New creates a new app model.
func New() Model {
	return Model{
		screen: ScreenEditor,
		editor: editor.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return m.editor.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	switch m.screen {
	case ScreenEditor:
		var cmd tea.Cmd
		m.editor, cmd = m.editor.Update(msg)
		if m.editor.Quitting() {
			return m, tea.Quit
		}
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	switch m.screen {
	case ScreenEditor:
		return m.editor.View()
	}
	return ""
}
