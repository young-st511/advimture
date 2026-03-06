package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// MenuAction represents what the user selected from the menu.
type MenuAction int

const (
	MenuNone MenuAction = iota
	MenuFreeMode
	MenuTutorial
	MenuQuit
)

// MenuItem represents a single menu entry.
type MenuItem struct {
	Label    string
	Action   MenuAction
	Locked   bool
	LockText string
}

// MenuModel is the Bubbletea model for the main menu.
type MenuModel struct {
	Items    []MenuItem
	Selected int
	Rank     string
	Width    int
	Height   int

	chosen MenuAction
}

// NewMenu creates a new menu model with default items.
func NewMenu() MenuModel {
	return MenuModel{
		Items: []MenuItem{
			{Label: "Tutorial", Action: MenuTutorial},
			{Label: "Mission", Locked: true, LockText: "Tutorial 완료 후 해금"},
			{Label: "Time Attack", Locked: true, LockText: "Mission 3 완료 후 해금"},
			{Label: "Free Mode", Action: MenuFreeMode},
			{Label: "Progress", Locked: true, LockText: "준비 중"},
			{Label: "Cheatsheet", Locked: true, LockText: "준비 중"},
			{Label: "Quit", Action: MenuQuit},
		},
		Selected: 0,
		Rank:     "Intern",
	}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

// Update handles menu key input.
func (m MenuModel) Update(msg tea.Msg) (MenuModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Selected++
			if m.Selected >= len(m.Items) {
				m.Selected = len(m.Items) - 1
			}
		case "k", "up":
			m.Selected--
			if m.Selected < 0 {
				m.Selected = 0
			}
		case "enter":
			item := m.Items[m.Selected]
			if !item.Locked {
				m.chosen = item.Action
			}
		case "q":
			m.chosen = MenuQuit
		}
	}
	return m, nil
}

// Chosen returns the menu action chosen by the user (resets after read).
func (m *MenuModel) Chosen() MenuAction {
	action := m.chosen
	m.chosen = MenuNone
	return action
}

// View renders the menu.
func (m MenuModel) View() string {
	var b strings.Builder

	// Title
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#5B7FFF")).
		MarginBottom(1).
		Render("⌨  A D V I M T U R E")

	subtitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Render(fmt.Sprintf("Rank: %s", m.Rank))

	b.WriteString(title + "\n")
	b.WriteString(subtitle + "\n\n")

	// Menu items
	for i, item := range m.Items {
		cursor := "  "
		if i == m.Selected {
			cursor = "> "
		}

		style := lipgloss.NewStyle()
		if item.Locked {
			label := fmt.Sprintf("%s%s  [%s]", cursor, item.Label, item.LockText)
			b.WriteString(style.Foreground(lipgloss.Color("#555555")).Render(label))
		} else if i == m.Selected {
			label := cursor + item.Label
			b.WriteString(style.Foreground(lipgloss.Color("#5B7FFF")).Bold(true).Render(label))
		} else {
			label := cursor + item.Label
			b.WriteString(style.Foreground(lipgloss.Color("#CCCCCC")).Render(label))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(DimStyle.Render("j/k: 이동  Enter: 선택  q: 종료"))

	return b.String()
}
