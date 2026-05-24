package playableview

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/young-st511/advimture/internal/tuiadapter"
)

const SelectionLinewise = "linewise"

type Screen struct {
	Width           int
	Height          int
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
	FocusPanel      *FocusPanel
	ActionLines     []string
	CommandPrompt   string
	ShowCommandLine bool
	ShowLastCommand bool
}

type FocusPanel struct {
	Kind  string
	Title string
	Lines []string
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
	if screen.Height > 0 {
		b.WriteString(RenderFocusLayer(screen.FocusPanel, screen.Width))
	} else if screen.FocusPanel != nil {
		b.WriteString(RenderFocusPanel(*screen.FocusPanel, screen.Width))
		b.WriteString("\n\n")
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
	if screen.FocusPanel == nil {
		b.WriteString(RenderActionPanel(screen.ActionLines))
	}
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

func RenderFocusPanel(panel FocusPanel, screenWidth int) string {
	lines := append([]string{panel.Title}, panel.Lines...)
	panelWidth := focusPanelWidth(screenWidth)
	rendered := actionPanelStyle.Width(panelWidth).Render(strings.Join(lines, "\n"))
	if screenWidth <= 0 || screenWidth <= panelWidth {
		return rendered
	}
	return lipgloss.PlaceHorizontal(screenWidth, lipgloss.Center, rendered)
}

func RenderFocusLayer(panel *FocusPanel, screenWidth int) string {
	const layerHeight = 9
	lines := make([]string, layerHeight)
	if panel == nil {
		return strings.Join(lines, "\n") + "\n"
	}
	rendered := strings.Split(RenderFocusPanel(*panel, screenWidth), "\n")
	rendered = fitFocusPanelLines(rendered, layerHeight)
	start := (layerHeight - len(rendered)) / 2
	if start < 0 {
		start = 0
	}
	for i, line := range rendered {
		target := start + i
		if target >= len(lines) {
			break
		}
		lines[target] = line
	}
	return strings.Join(lines, "\n") + "\n"
}

func fitFocusPanelLines(rendered []string, maxHeight int) []string {
	if len(rendered) <= maxHeight || maxHeight < 3 {
		return rendered
	}
	interiorLimit := maxHeight - 2
	fitted := make([]string, 0, maxHeight)
	fitted = append(fitted, rendered[0])
	interior := append([]string(nil), rendered[1:len(rendered)-1]...)
	if len(interior) > interiorLimit {
		interior = interior[:interiorLimit]
	}
	if priority := focusPanelPriorityLine(rendered[1 : len(rendered)-1]); priority != "" && !containsLine(interior, priority) {
		interior[len(interior)-1] = priority
	}
	fitted = append(fitted, interior...)
	fitted = append(fitted, rendered[len(rendered)-1])
	return fitted
}

func focusPanelPriorityLine(lines []string) string {
	for _, marker := range []string{"Retry:", "Next:", "Next tutorial:", "Playlist complete", "q: quit"} {
		for i := len(lines) - 1; i >= 0; i-- {
			line := lines[i]
			if strings.Contains(line, marker) {
				return line
			}
		}
	}
	return ""
}

func containsLine(lines []string, target string) bool {
	for _, line := range lines {
		if line == target {
			return true
		}
	}
	return false
}

func focusPanelWidth(screenWidth int) int {
	const fallback = 58
	const minWidth = 32
	const maxWidth = 72
	if screenWidth <= 0 {
		return fallback
	}
	width := screenWidth - 4
	if width < minWidth {
		width = screenWidth
	}
	if width > maxWidth {
		width = maxWidth
	}
	if width <= 0 {
		return fallback
	}
	return width
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
