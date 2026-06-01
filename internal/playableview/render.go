package playableview

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/young-st511/advimture/internal/tuiadapter"
)

const SelectionLinewise = "linewise"

type Screen struct {
	Width            int
	Height           int
	PlaylistTitle    string
	PlaylistCategory string
	ReviewSummary    string
	DailyRoute       string
	ReviewCount      int
	ReviewPrimary    string
	Title            string
	Message          string
	BufferLines      []string
	Mode             string
	Status           string
	CursorRow        int
	CursorCol        int
	Selection        *tuiadapter.SelectionView
	ExerciseIndex    int
	ExerciseTotal    int
	Grade            string
	CommandLine      string
	LastCommand      string
	FocusPanel       *FocusPanel
	ActionLines      []string
	CommandPrompt    string
	ShowCommandLine  bool
	ShowLastCommand  bool
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

var floatingModalStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("63")).
	Padding(0, 1).
	Width(58)

func Render(screen Screen) string {
	if screen.Width > 0 && screen.Height > 0 {
		return renderHUD(screen)
	}
	return renderLegacy(screen)
}

func renderLegacy(screen Screen) string {
	var b strings.Builder
	b.WriteString(renderHeader(screen))
	b.WriteString("\n\n")
	b.WriteString(screen.Title + "\n")
	b.WriteString(screen.Message + "\n\n")
	if screen.ReviewSummary != "" || screen.DailyRoute != "" {
		b.WriteString("복구 현황\n")
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

func renderHUD(screen Screen) string {
	var b strings.Builder
	b.WriteString(renderHeader(screen))
	b.WriteString("\n\n")
	b.WriteString("MISSION\n")
	b.WriteString(screen.Title + "\n")
	for _, line := range wrapHUDMessage(screen.Message, screen.Width) {
		b.WriteString(line + "\n")
	}
	if cue := renderMissionCue(screen.FocusPanel, screen.Width); cue != "" {
		b.WriteString(cue)
	}
	if status := recoveryStatusLine(screen); status != "" {
		b.WriteString(status + "\n")
	}
	b.WriteString("\n")
	b.WriteString("RUNBOOK CONSOLE\n")
	for row, line := range screen.BufferLines {
		b.WriteString(RenderLine(line, row, screen.CursorRow, screen.CursorCol, screen.Selection))
		b.WriteString("\n")
	}
	if screen.FocusPanel != nil && isFloatingPanel(*screen.FocusPanel) {
		b.WriteString("\n")
		b.WriteString(RenderFloatingModal(*screen.FocusPanel, screen.Width))
		b.WriteString("\n")
	}
	b.WriteString("\n")
	b.WriteString(renderHUDStatusLine(screen))
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
	return b.String()
}

func renderHUDStatusLine(screen Screen) string {
	mode := strings.ToUpper(screen.Mode)
	if mode == "" {
		mode = "-"
	}
	status := screen.Status
	if status == "" {
		status = "-"
	}
	return fmt.Sprintf("%s · %s · cursor %d:%d\n", mode, status, screen.CursorRow, screen.CursorCol)
}

func recoveryStatusLine(screen Screen) string {
	if screen.FocusPanel != nil && isFloatingPanel(*screen.FocusPanel) {
		return ""
	}
	if screen.FocusPanel != nil && !isFloatingPanel(*screen.FocusPanel) && screen.ReviewCount > 0 {
		primary := screen.ReviewPrimary
		if primary == "" {
			primary = "대기 중"
		}
		switch {
		case screen.FocusPanel.Kind == "training" || screen.PlaylistCategory == "tutorial":
			return fmt.Sprintf("복구 메모: 재점검 %d건 · 다음: %s", screen.ReviewCount, primary)
		case screen.FocusPanel.Kind == "incident" || screen.PlaylistCategory == "incident":
			return fmt.Sprintf("복구 현황: 재점검 %d건 · 잔류: %s", screen.ReviewCount, primary)
		}
	}
	parts := []string{}
	if screen.ReviewSummary != "" {
		parts = append(parts, screen.ReviewSummary)
	}
	if screen.DailyRoute != "" {
		parts = append(parts, screen.DailyRoute)
	}
	return strings.Join(parts, " · ")
}

func wrapHUDMessage(message string, screenWidth int) []string {
	if message == "" {
		return []string{""}
	}
	width := hudTextWidth(screenWidth)
	words := strings.Fields(message)
	if len(words) == 0 {
		return []string{message}
	}
	lines := make([]string, 0, 2)
	current := ""
	for _, word := range words {
		next := word
		if current != "" {
			next = current + " " + word
		}
		if displayWidth(next) <= width {
			current = next
			continue
		}
		if current != "" {
			lines = append(lines, current)
			current = word
		} else {
			lines = append(lines, trimRunes(word, width))
		}
		if len(lines) == 2 {
			lines[1] = forceEllipsis(lines[1], width)
			return lines
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	if len(lines) > 2 {
		lines = lines[:2]
		lines[1] = ellipsize(lines[1], width)
	}
	return lines
}

func forceEllipsis(text string, width int) string {
	if strings.HasSuffix(text, "...") {
		return text
	}
	if width <= 3 {
		return trimDisplayWidth(text, width)
	}
	return trimDisplayWidth(text, width-3) + "..."
}

func hudTextWidth(screenWidth int) int {
	const fallback = 88
	const minWidth = 36
	const maxWidth = 88
	if screenWidth <= 0 {
		return fallback
	}
	width := screenWidth - 8
	if width < minWidth {
		width = minWidth
	}
	if width > maxWidth {
		width = maxWidth
	}
	return width
}

func ellipsize(text string, width int) string {
	if displayWidth(text) <= width {
		return text
	}
	if width <= 3 {
		return trimDisplayWidth(text, width)
	}
	return trimDisplayWidth(text, width-3) + "..."
}

func trimRunes(text string, limit int) string {
	if limit <= 0 {
		return ""
	}
	runes := []rune(text)
	if len(runes) <= limit {
		return text
	}
	return string(runes[:limit])
}

func runeLen(text string) int {
	return len([]rune(text))
}

func displayWidth(text string) int {
	width := 0
	for _, r := range text {
		if r <= 127 {
			width++
		} else {
			width += 2
		}
	}
	return width
}

func trimDisplayWidth(text string, limit int) string {
	if limit <= 0 {
		return ""
	}
	width := 0
	var b strings.Builder
	for _, r := range text {
		cellWidth := 1
		if r > 127 {
			cellWidth = 2
		}
		if width+cellWidth > limit {
			break
		}
		b.WriteRune(r)
		width += cellWidth
	}
	return b.String()
}

func renderMissionCue(panel *FocusPanel, screenWidth int) string {
	if panel == nil || isFloatingPanel(*panel) {
		return ""
	}
	lines := append([]string{panel.Title}, panel.Lines...)
	return strings.Join(wrapMissionCueLines(lines, hudTextWidth(screenWidth)), "\n") + "\n"
}

func isFloatingPanel(panel FocusPanel) bool {
	return panel.Kind == "failure" || panel.Kind == "success"
}

func wrapMissionCueLines(parts []string, width int) []string {
	if len(parts) == 0 {
		return nil
	}
	if width <= 0 {
		return []string{strings.Join(parts, " · ")}
	}
	out := []string{}
	current := ""
	for _, part := range parts {
		wrapped := wrapTextByDisplayWidth(part, width)
		for i, line := range wrapped {
			if i > 0 && current != "" {
				out = append(out, current)
				current = ""
			}
			if current == "" {
				current = line
				continue
			}
			next := current + " · " + line
			if displayWidth(next) <= width {
				current = next
				continue
			}
			out = append(out, current)
			current = line
		}
	}
	if current != "" {
		out = append(out, current)
	}
	return out
}

func wrapTextByDisplayWidth(text string, width int) []string {
	if text == "" {
		return []string{""}
	}
	if displayWidth(text) <= width {
		return []string{text}
	}
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{trimDisplayWidth(text, width)}
	}
	lines := []string{}
	current := ""
	for _, word := range words {
		if displayWidth(word) > width {
			if current != "" {
				lines = append(lines, current)
				current = ""
			}
			lines = append(lines, trimDisplayWidth(word, width))
			continue
		}
		next := word
		if current != "" {
			next = current + " " + word
		}
		if displayWidth(next) <= width {
			current = next
			continue
		}
		if current != "" {
			lines = append(lines, current)
		}
		current = word
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
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

func RenderFloatingModal(panel FocusPanel, screenWidth int) string {
	lines := floatingModalLines(panel)
	panelWidth := focusPanelWidth(screenWidth)
	rendered := floatingModalStyle.Width(panelWidth).Render(strings.Join(lines, "\n"))
	renderedLines := strings.Split(rendered, "\n")
	renderedLines = fitFocusPanelLines(renderedLines, 12)
	rendered = strings.Join(renderedLines, "\n")
	if screenWidth <= 0 || screenWidth <= panelWidth {
		return rendered
	}
	return lipgloss.PlaceHorizontal(screenWidth, lipgloss.Center, rendered)
}

func floatingModalLines(panel FocusPanel) []string {
	switch panel.Kind {
	case "failure":
		lines := []string{floatingModalTitle(panel), panel.Title}
		return append(lines, failureModalLines(panel.Lines)...)
	case "success":
		lines := []string{floatingModalTitle(panel)}
		return append(lines, successModalLines(panel.Lines)...)
	default:
		lines := []string{floatingModalTitle(panel), panel.Title}
		return append(lines, panel.Lines...)
	}
}

func floatingModalTitle(panel FocusPanel) string {
	switch panel.Kind {
	case "failure":
		return "RECOVERY CHECK"
	case "success":
		return "RUNBOOK SEALED"
	default:
		return panel.Title
	}
}

func failureModalLines(lines []string) []string {
	out := []string{}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		switch {
		case i == 0:
			out = append(out, "Mistake  "+line)
		case strings.HasPrefix(line, "Inputs left:") && i+1 < len(lines) && strings.HasPrefix(lines[i+1], "Attempts:"):
			out = append(out, line+" · "+lines[i+1])
			i++
		case strings.HasPrefix(line, "Coach:") || strings.HasPrefix(line, "복구 힌트:"):
			out = append(out, "Next     "+line)
		default:
			out = append(out, line)
		}
	}
	return out
}

func successModalLines(lines []string) []string {
	out := []string{}
	for i, line := range lines {
		switch {
		case i == 0:
			out = append(out, "Learned  "+line)
		case strings.HasPrefix(line, "복구 기록:") || strings.HasPrefix(line, "이번 복구:"):
			out = append(out, "Result   "+line)
		default:
			out = append(out, line)
		}
	}
	return out
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
	for _, marker := range []string{"Retry:", "Next:", "Next tutorial:", "Next runbook:", "Next dispatch:", "Playlist complete", "Dispatch complete", "q: quit"} {
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
