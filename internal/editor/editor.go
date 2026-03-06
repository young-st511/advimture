package editor

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/young-st511/advimture/internal/ui"
)

// Mode represents the editor mode.
type Mode int

const (
	ModeNormal Mode = iota
	ModeInsert
	ModeCommand
	ModeVisual
	ModeOperatorPending
)

func (m Mode) String() string {
	switch m {
	case ModeNormal:
		return "NORMAL"
	case ModeInsert:
		return "INSERT"
	case ModeCommand:
		return "COMMAND"
	case ModeVisual:
		return "VISUAL"
	case ModeOperatorPending:
		return "OP-PENDING"
	default:
		return "NORMAL"
	}
}

// Model is the Bubbletea model for the Vim editor.
type Model struct {
	lines    []string
	row      int
	col      int
	mode     Mode
	width    int
	height   int
	quitting bool
}

// New creates a new editor model with sample text.
func New() Model {
	sample := []string{
		"Hello, Advimture!",
		"",
		"This is a Vim practice game.",
		"Use hjkl to move around.",
		"Press i to enter Insert mode.",
		"Press :wq to save and quit.",
	}
	return Model{
		lines:  sample,
		row:    0,
		col:    0,
		mode:   ModeNormal,
		width:  80,
		height: 24,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles key input and returns the updated model.
// Returns (Model, tea.Cmd) with concrete type for the parent model.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		if m.mode == ModeNormal {
			switch msg.String() {
			case "q":
				m.quitting = true
			case "h":
				if m.col > 0 {
					m.col--
				}
			case "l":
				lineRunes := []rune(m.lines[m.row])
				if m.col < len(lineRunes)-1 {
					m.col++
				}
			case "j":
				if m.row < len(m.lines)-1 {
					m.row++
					m.clampCol()
				}
			case "k":
				if m.row > 0 {
					m.row--
					m.clampCol()
				}
			}
		}
	}

	return m, nil
}

// clampCol ensures the cursor column doesn't exceed the current line length (rune-aware).
func (m *Model) clampCol() {
	lineRunes := []rune(m.lines[m.row])
	maxCol := len(lineRunes) - 1
	if maxCol < 0 {
		maxCol = 0
	}
	if m.col > maxCol {
		m.col = maxCol
	}
}

// Quitting returns whether the editor is quitting.
func (m Model) Quitting() bool {
	return m.quitting
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder

	// Editor area
	editorHeight := m.height - 3 // reserve for status bar + help line
	for i := 0; i < editorHeight; i++ {
		if i < len(m.lines) {
			lineNum := fmt.Sprintf("%3d ", i+1)
			b.WriteString(ui.LineNumberStyle.Render(lineNum))

			line := m.lines[i]
			if i == m.row {
				b.WriteString(renderLineWithCursor(line, m.col))
			} else {
				b.WriteString(line)
			}
		} else {
			b.WriteString(ui.EmptyLineStyle.Render("  ~ "))
		}
		b.WriteString("\n")
	}

	// Status bar
	modeStr := m.mode.String()
	modeLabel := ui.StatusBarStyle(modeStr).Render(fmt.Sprintf(" -- %s -- ", modeStr))
	position := fmt.Sprintf("Ln %d, Col %d", m.row+1, m.col+1)
	gap := m.width - lipgloss.Width(modeLabel) - len(position) - 2
	if gap < 1 {
		gap = 1
	}
	statusLine := modeLabel + strings.Repeat(" ", gap) + position
	b.WriteString(statusLine + "\n")

	// Help line
	helpLine := ui.DimStyle.Render("Ctrl+c ×2: 메뉴 | q: 종료")
	b.WriteString(helpLine)

	return b.String()
}

// renderLineWithCursor renders a line with a block cursor at the given column.
func renderLineWithCursor(line string, col int) string {
	runes := []rune(line)
	if len(runes) == 0 {
		// Empty line: show block cursor on space
		return lipgloss.NewStyle().Reverse(true).Render(" ")
	}
	if col >= len(runes) {
		col = len(runes) - 1
	}

	var b strings.Builder
	if col > 0 {
		b.WriteString(string(runes[:col]))
	}
	b.WriteString(lipgloss.NewStyle().Reverse(true).Render(string(runes[col])))
	if col < len(runes)-1 {
		b.WriteString(string(runes[col+1:]))
	}
	return b.String()
}
