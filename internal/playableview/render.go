package playableview

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/young-st511/advimture/internal/tuiadapter"
)

const SelectionLinewise = "linewise"

type Screen struct {
	PlaylistTitle   string
	ReviewSummary   string
	DailyRoute      string
	Title           string
	Message         string
	BufferLines     []string
	Mode            string
	Status          string
	CursorRow       int
	CursorCol       int
	Selection       *tuiadapter.SelectionView
	ExerciseIndex   int
	ExerciseTotal   int
	Grade           string
	CommandLine     string
	LastCommand     string
	ActionLines     []string
	CommandPrompt   string
	ShowCommandLine bool
	ShowLastCommand bool
}

var actionPanelStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("63")).
	Padding(0, 1).
	Width(58)

func Render(screen Screen) string {
	var b strings.Builder
	b.WriteString(renderHeader(screen))
	b.WriteString("\n\n")
	b.WriteString(screen.Title + "\n")
	b.WriteString(screen.Message + "\n\n")
	if screen.ReviewSummary != "" || screen.DailyRoute != "" {
		b.WriteString("OPS\n")
		if screen.ReviewSummary != "" {
			b.WriteString(screen.ReviewSummary + "\n")
		}
		if screen.DailyRoute != "" {
			b.WriteString(screen.DailyRoute + "\n")
		}
		b.WriteString("\n")
	}
	b.WriteString("RUNBOOK CONSOLE\n")
	for row, line := range screen.BufferLines {
		b.WriteString(RenderLine(line, row, screen.CursorRow, screen.CursorCol, screen.Selection))
		b.WriteString("\n")
	}
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Mode: %s  Status: %s  Cursor: %d,%d\n", screen.Mode, screen.Status, screen.CursorRow, screen.CursorCol))
	if screen.Selection != nil && screen.Selection.Active {
		b.WriteString(fmt.Sprintf("Selection: %s %d,%d -> %d,%d\n", screen.Selection.Kind, screen.Selection.Start.Row, screen.Selection.Start.Col, screen.Selection.End.Row, screen.Selection.End.Col))
	}
	if screen.ExerciseTotal > 0 {
		b.WriteString(fmt.Sprintf("Exercise: %d/%d\n", screen.ExerciseIndex+1, screen.ExerciseTotal))
	}
	if screen.Grade != "" {
		b.WriteString(fmt.Sprintf("Grade: %s\n", screen.Grade))
	}
	b.WriteString("\n")
	if screen.ShowCommandLine {
		b.WriteString(screen.CommandPrompt + screen.CommandLine + "\n\n")
	}
	if screen.ShowLastCommand && screen.LastCommand != "" {
		b.WriteString(fmt.Sprintf("Command: %s\n", screen.LastCommand))
	}
	b.WriteString(RenderActionPanel(screen.ActionLines))
	return b.String()
}

func renderHeader(screen Screen) string {
	parts := []string{"ADVIMTURE"}
	if screen.PlaylistTitle != "" {
		parts = append(parts, screen.PlaylistTitle)
	}
	if screen.ExerciseTotal > 0 {
		parts = append(parts, fmt.Sprintf("Exercise: %d/%d", screen.ExerciseIndex+1, screen.ExerciseTotal))
	}
	if screen.Status != "" {
		parts = append(parts, fmt.Sprintf("Status: %s", screen.Status))
	}
	return strings.Join(parts, " | ")
}

func RenderActionPanel(lines []string) string {
	if len(lines) == 0 {
		return ""
	}
	return actionPanelStyle.Render(strings.Join(lines, "\n"))
}

func RenderLine(line string, row int, cursorRow int, cursorCol int, selection *tuiadapter.SelectionView) string {
	prefix := "  "
	if row != cursorRow {
		return prefix + renderLineCells(line, row, -1, selection)
	}
	runes := []rune(line)
	if len(runes) == 0 {
		return "> []"
	}
	if cursorCol < 0 {
		cursorCol = 0
	}
	if cursorCol >= len(runes) {
		cursorCol = len(runes) - 1
	}
	return "> " + renderLineCells(line, row, cursorCol, selection)
}

func renderLineCells(line string, row int, cursorCol int, selection *tuiadapter.SelectionView) string {
	runes := []rune(line)
	if len(runes) == 0 {
		return line
	}
	var b strings.Builder
	for col, r := range runes {
		if col == cursorCol {
			b.WriteString("[")
			b.WriteRune(r)
			b.WriteString("]")
			continue
		}
		if cellSelected(row, col, selection) {
			b.WriteString("{")
			b.WriteRune(r)
			b.WriteString("}")
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

func cellSelected(row int, col int, selection *tuiadapter.SelectionView) bool {
	if selection == nil || !selection.Active {
		return false
	}
	if selection.Kind == SelectionLinewise {
		return row >= selection.Start.Row && row <= selection.End.Row
	}
	if selection.Kind != "charwise" {
		return false
	}
	if row < selection.Start.Row || row > selection.End.Row {
		return false
	}
	startCol := 0
	if row == selection.Start.Row {
		startCol = selection.Start.Col
	}
	endCol := int(^uint(0) >> 1)
	if row == selection.End.Row {
		endCol = selection.End.Col
	}
	return col >= startCol && col <= endCol
}
