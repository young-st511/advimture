package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/young-st511/advimture/internal/data"
)

// ResultData holds the data needed to render a mission result screen.
// It is separate from game.GradeResult to avoid circular imports (ui ← game ← ui).
type ResultData struct {
	Grade       string
	EffKeys     int
	OptimalKeys int
	TimeMs      int64
	Accuracy    float64
	Success     bool
	MentorMsg   string
}

// RenderResult renders the mission result screen and returns a string.
func RenderResult(result *ResultData, diffs []data.DiffLine, totalDiff int, width int) string {
	var b strings.Builder

	b.WriteString(HeaderStyle.Render("RESULT"))
	b.WriteString("\n\n")

	if !result.Success {
		b.WriteString(renderFailure(diffs, totalDiff))
	} else {
		b.WriteString(renderSuccess(result))
	}

	b.WriteString("\n\n")
	b.WriteString(DimStyle.Render("[Enter] 다시하기  [b] 메뉴"))

	return b.String()
}

func renderSuccess(result *ResultData) string {
	var b strings.Builder

	// Grade badge
	gc := resultGradeColor(result.Grade)
	gradeBadge := lipgloss.NewStyle().
		Bold(true).
		Foreground(gc).
		Padding(0, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(gc).
		Render(result.Grade)

	b.WriteString(gradeBadge)
	b.WriteString("\n\n")

	// Stats
	b.WriteString(fmt.Sprintf("  유효 키스트로크: %d  (최적: %d)\n", result.EffKeys, result.OptimalKeys))
	b.WriteString(fmt.Sprintf("  소요 시간:       %s\n", formatMs(result.TimeMs)))
	b.WriteString(fmt.Sprintf("  정확도:          %.0f%%\n", result.Accuracy))

	// Mentor message
	if result.MentorMsg != "" {
		b.WriteString("\n")
		mentorBox := lipgloss.NewStyle().
			Foreground(MentorColor).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(MentorColor).
			Padding(0, 1).
			Render("Vi 선배: " + result.MentorMsg)
		b.WriteString(mentorBox)
	}

	return b.String()
}

func renderFailure(diffs []data.DiffLine, totalDiff int) string {
	var b strings.Builder

	failStyle := lipgloss.NewStyle().Bold(true).Foreground(ErrorColor)
	b.WriteString(failStyle.Render("✗ 텍스트가 일치하지 않습니다"))
	b.WriteString("\n\n")

	for _, d := range diffs {
		lineLabel := fmt.Sprintf("Line %d:", d.LineNum)
		b.WriteString(lipgloss.NewStyle().Foreground(ErrorColor).Render(
			fmt.Sprintf("  ✗ %s  yours:    %s", lineLabel, d.Yours),
		))
		b.WriteString("\n")
		b.WriteString(lipgloss.NewStyle().Foreground(SuccessColor).Render(
			fmt.Sprintf("  ✓ %s  expected: %s", lineLabel, d.Expected),
		))
		b.WriteString("\n\n")
	}

	if totalDiff > len(diffs) {
		extra := totalDiff - len(diffs)
		b.WriteString(DimStyle.Render(fmt.Sprintf("  ... 외 %d줄 더 불일치", extra)))
		b.WriteString("\n")
	}

	return b.String()
}

func resultGradeColor(grade string) lipgloss.Color {
	switch grade {
	case "S":
		return lipgloss.Color("#FFD700") // gold
	case "A":
		return SuccessColor // green
	case "B":
		return NormalColor // blue
	default: // C or F
		return DimColor // gray
	}
}

func formatMs(ms int64) string {
	if ms < 1000 {
		return fmt.Sprintf("%dms", ms)
	}
	secs := float64(ms) / 1000.0
	if secs < 60 {
		return fmt.Sprintf("%.1f초", secs)
	}
	mins := int(secs) / 60
	s := int(secs) % 60
	return fmt.Sprintf("%d분 %02d초", mins, s)
}
