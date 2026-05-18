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
	renderEditor(&b, m.buf, m.cursor, m.mode, editorHeight, m.searchState)
	renderStatusBar(&b, m.mode, m.cursor, m.keystrokes, m.width)
	renderBottomLine(&b, m.mode, m.cmdInput, m.statusMsg, m.statusError, m.searchMode, m.searchInput)

	return b.String()
}

func renderEditor(b *strings.Builder, buf *Buffer, cur *Cursor, mode Mode, height int, search *SearchState) {
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
				rowMatches := searchMatchesForRow(search, i)
				b.WriteString(renderLineWithHighlights(line, cur.Col, insertMode, rowMatches))
			} else {
				rowMatches := searchMatchesForRow(search, i)
				if len(rowMatches) > 0 {
					b.WriteString(renderLineSearchOnly(line, rowMatches))
				} else {
					b.WriteString(line)
				}
			}
		} else {
			padding := strings.Repeat(" ", lineNumWidth)
			b.WriteString(ui.EmptyLineStyle.Render(padding + "~"))
		}
		b.WriteString("\n")
	}
}

// searchMatchesForRow returns the matches for the given row from the SearchState.
func searchMatchesForRow(search *SearchState, row int) []MatchPos {
	if search == nil {
		return nil
	}
	var result []MatchPos
	for _, m := range search.Matches {
		if m.Row == row {
			result = append(result, m)
		}
	}
	return result
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

func renderBottomLine(b *strings.Builder, mode Mode, cmdInput string, statusMsg string, statusError bool, searchMode bool, searchInput string) {
	if searchMode {
		b.WriteString("/" + searchInput)
		b.WriteString(lipgloss.NewStyle().Reverse(true).Render(" "))
		return
	}

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

// renderLineWithHighlights renders a line with both search highlights and cursor.
func renderLineWithHighlights(line string, col int, insertMode bool, searchMatches []MatchPos) string {
	runes := []rune(line)
	n := len(runes)

	if n == 0 {
		// Empty line: just cursor
		return lipgloss.NewStyle().Reverse(true).Render(" ")
	}

	if len(searchMatches) == 0 {
		return renderLineWithCursor(line, col, insertMode)
	}

	// Build per-character style: 0=normal, 1=search match, 2=cursor
	charStyle := make([]int, n)
	for _, m := range searchMatches {
		for i := 0; i < m.Len && m.Col+i < n; i++ {
			charStyle[m.Col+i] = 1
		}
	}

	// Determine cursor position
	cursorPos := col
	if insertMode {
		if cursorPos > n {
			cursorPos = n
		}
	} else {
		if cursorPos >= n {
			cursorPos = n - 1
		}
		charStyle[cursorPos] = 2
	}

	var b strings.Builder

	if insertMode {
		// Render each character; cursor is a bar before cursorPos
		for i, r := range runes {
			if i == cursorPos {
				// Cursor bar: show reverse on this character
				b.WriteString(lipgloss.NewStyle().Reverse(true).Render(string(r)))
			} else if charStyle[i] == 1 {
				b.WriteString(ui.SearchHighlightStyle.Render(string(r)))
			} else {
				b.WriteRune(r)
			}
		}
		if cursorPos == n {
			b.WriteString(lipgloss.NewStyle().Reverse(true).Render(" "))
		}
	} else {
		// Block cursor: render each character with appropriate style
		for i, r := range runes {
			switch charStyle[i] {
			case 2: // cursor
				b.WriteString(lipgloss.NewStyle().Reverse(true).Render(string(r)))
			case 1: // search match
				b.WriteString(ui.SearchHighlightStyle.Render(string(r)))
			default:
				b.WriteRune(r)
			}
		}
	}

	return b.String()
}

// renderLineSearchOnly renders a non-cursor line with search highlights.
func renderLineSearchOnly(line string, searchMatches []MatchPos) string {
	runes := []rune(line)
	n := len(runes)
	if n == 0 || len(searchMatches) == 0 {
		return line
	}

	// Build per-character highlight flag
	inMatch := make([]bool, n)
	for _, m := range searchMatches {
		for i := 0; i < m.Len && m.Col+i < n; i++ {
			inMatch[m.Col+i] = true
		}
	}

	// Render in segments
	var b strings.Builder
	i := 0
	for i < n {
		highlighted := inMatch[i]
		j := i + 1
		for j < n && inMatch[j] == highlighted {
			j++
		}
		segment := string(runes[i:j])
		if highlighted {
			b.WriteString(ui.SearchHighlightStyle.Render(segment))
		} else {
			b.WriteString(segment)
		}
		i = j
	}
	return b.String()
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
