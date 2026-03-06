package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/young-st511/advimture/internal/data"
	"github.com/young-st511/advimture/internal/editor"
	"github.com/young-st511/advimture/internal/game"
	"github.com/young-st511/advimture/internal/progress"
	"github.com/young-st511/advimture/internal/ui"
)

// Screen represents the current screen state.
type Screen int

const (
	ScreenMenu Screen = iota
	ScreenEditor
	ScreenFTUE
	ScreenTutorial
)

// Model is the top-level Bubbletea model that manages screen transitions.
type Model struct {
	screen   Screen
	menu     ui.MenuModel
	editor   editor.Model
	ftue     game.FTUEModel
	tutorial game.TutorialModel
	progress *progress.Progress
	width    int
	height   int
}

// New creates a new app model.
func New() Model {
	p, _ := progress.Load()

	m := Model{
		progress: p,
	}

	// First run: show FTUE
	if p.CompletedTutorialCount() == 0 {
		m.screen = ScreenFTUE
		m.ftue = game.NewFTUE()
	} else {
		m.screen = ScreenMenu
		m.menu = ui.NewMenu()
		m.menu.Rank = p.CurrentRank().String()
	}

	return m
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
	case ScreenFTUE:
		return m.updateFTUE(msg)
	case ScreenTutorial:
		return m.updateTutorial(msg)
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
		m.startTutorial("t01_survival.yaml")
	case ui.MenuQuit:
		return m, tea.Quit
	}

	return m, cmd
}

func (m Model) updateEditor(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.editor, cmd = m.editor.Update(msg)

	if m.editor.Quitting() {
		m.screen = ScreenMenu
		m.menu = ui.NewMenu()
		m.menu.Rank = m.progress.CurrentRank().String()
		return m, nil
	}

	return m, cmd
}

func (m Model) updateFTUE(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.ftue, cmd = m.ftue.Update(msg)

	if m.ftue.Quitting() {
		return m, tea.Quit
	}

	if m.ftue.Done() {
		m.startTutorial("t01_survival.yaml")
	}

	return m, cmd
}

func (m Model) updateTutorial(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.tutorial, cmd = m.tutorial.Update(msg)

	if m.tutorial.Quitting() {
		if m.tutorial.Completed() {
			// Save progress
			m.progress.CompleteTutorial(
				m.tutorial.TutorialID(),
				m.tutorial.ElapsedSeconds(),
				m.tutorial.Keystrokes(),
			)
			_ = progress.Save(m.progress)
		}

		m.screen = ScreenMenu
		m.menu = ui.NewMenu()
		m.menu.Rank = m.progress.CurrentRank().String()
		return m, nil
	}

	return m, cmd
}

func (m *Model) startTutorial(filename string) {
	td, err := data.LoadTutorial(filename)
	if err != nil {
		// Fallback to menu on error
		m.screen = ScreenMenu
		m.menu = ui.NewMenu()
		return
	}
	m.tutorial = game.NewTutorial(td)
	m.screen = ScreenTutorial
}

func (m Model) View() string {
	switch m.screen {
	case ScreenMenu:
		return m.menu.View()
	case ScreenEditor:
		return m.editor.View()
	case ScreenFTUE:
		return m.ftue.View()
	case ScreenTutorial:
		return m.tutorial.View()
	}
	return ""
}
