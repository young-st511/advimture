package editor

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/young-st511/advimture/internal/ui"
)

// View renders the editor to a string.
func (m Model) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder

	editorHeight := m.height - 3 // reserve for status bar + command/message line
	renderEditor(&b, m.buf, m.cursor, m.mode, editorHeight)
	renderStatusBar(&b, m.mode, m.cursor, m.keystrokes, m.width)
	renderBottomLine(&b, m.mode, m.cmdInput, m.statusMsg, m.statusError)

	return b.String()
}

func renderEditor(b *strings.Builder, buf *Buffer, cur *Cursor, mode Mode, height int) {
	lineCount := buf.LineCount()
	lineNumWidth := len(fmt.Sprintf("%d", lineCount))
	if lineNumWidth < 3 {
		lineNumWidth = 3
	}

	for i := 0; i < height; i++ {
		if i < lineCount {
			lineNum := fmt.Sprintf("%*d ", lineNumWidth, i+1)
			b.WriteString(ui.LineNumberStyle.Render(lineNum))

			line := buf.GetLine(i)
			if i == cur.Row {
				insertMode := mode == ModeInsert
				b.WriteString(renderLineWithCursor(line, cur.Col, insertMode))
			} else {
				b.WriteString(line)
			}
		} else {
			padding := strings.Repeat(" ", lineNumWidth)
			b.WriteString(ui.EmptyLineStyle.Render(padding + "~"))
		}
		b.WriteString("\n")
	}
}

func renderStatusBar(b *strings.Builder, mode Mode, cur *Cursor, keystrokes int, width int) {
	modeStr := mode.String()
	modeLabel := ui.StatusBarStyle(modeStr).Render(fmt.Sprintf(" %s ", modeStr))
	position := fmt.Sprintf("Ln %d, Col %d", cur.Row+1, cur.Col+1)
	keystrokeInfo := fmt.Sprintf("Keys: %d", keystrokes)
	rightInfo := position + "  " + keystrokeInfo

	gap := width - lipgloss.Width(modeLabel) - len(rightInfo) - 2
	if gap < 1 {
		gap = 1
	}
	b.WriteString(modeLabel + strings.Repeat(" ", gap) + rightInfo + "\n")
}

func renderBottomLine(b *strings.Builder, mode Mode, cmdInput string, statusMsg string, statusError bool) {
	if mode == ModeCommand {
		b.WriteString(":" + cmdInput)
		// Show blinking cursor indicator
		b.WriteString(lipgloss.NewStyle().Reverse(true).Render(" "))
		return
	}

	if statusMsg != "" {
		if statusError {
			b.WriteString(lipgloss.NewStyle().Foreground(ui.ErrorColor).Render(statusMsg))
		} else {
			b.WriteString(lipgloss.NewStyle().Foreground(ui.SuccessColor).Render(statusMsg))
		}
		return
	}

	b.WriteString(ui.DimStyle.Render("Ctrl+c ×2: 종료 | :help"))
}

// renderLineWithCursor renders a line with a block or bar cursor at the given column.
func renderLineWithCursor(line string, col int, insertMode bool) string {
	runes := []rune(line)

	if insertMode {
		// Bar cursor: render between characters
		if col > len(runes) {
			col = len(runes)
		}
		var b strings.Builder
		if col > 0 {
			b.WriteString(string(runes[:col]))
		}
		// Bar cursor as a reverse space or thin marker
		if col < len(runes) {
			b.WriteString(lipgloss.NewStyle().Reverse(true).Render(string(runes[col])))
			if col < len(runes)-1 {
				b.WriteString(string(runes[col+1:]))
			}
		} else {
			b.WriteString(lipgloss.NewStyle().Reverse(true).Render(" "))
		}
		return b.String()
	}

	// Block cursor (Normal mode)
	if len(runes) == 0 {
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
