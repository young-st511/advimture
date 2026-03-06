package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/young-st511/advimture/internal/editor"
	"github.com/young-st511/advimture/internal/ui"
)

// Screen represents the current screen state.
type Screen int

const (
	ScreenMenu Screen = iota
	ScreenEditor
)

// Model is the top-level Bubbletea model that manages screen transitions.
type Model struct {
	screen Screen
	menu   ui.MenuModel
	editor editor.Model
	width  int
	height int
}

// New creates a new app model starting at the menu.
func New() Model {
	return Model{
		screen: ScreenMenu,
		menu:   ui.NewMenu(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	switch m.screen {
	case ScreenMenu:
		return m.updateMenu(msg)
	case ScreenEditor:
		return m.updateEditor(msg)
	}

	return m, nil
}

func (m Model) updateMenu(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.menu, cmd = m.menu.Update(msg)

	switch m.menu.Chosen() {
	case ui.MenuFreeMode:
		m.editor = editor.New()
		m.screen = ScreenEditor
	case ui.MenuTutorial:
		// TODO: tutorial screen
		m.editor = editor.New()
		m.screen = ScreenEditor
	case ui.MenuQuit:
		return m, tea.Quit
	}

	return m, cmd
}

func (m Model) updateEditor(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.editor, cmd = m.editor.Update(msg)

	if m.editor.Quitting() {
		// Return to menu instead of quitting the program
		m.screen = ScreenMenu
		m.menu = ui.NewMenu()
		return m, nil
	}

	return m, cmd
}

func (m Model) View() string {
	switch m.screen {
	case ScreenMenu:
		return m.menu.View()
	case ScreenEditor:
		return m.editor.View()
	}
	return ""
}
