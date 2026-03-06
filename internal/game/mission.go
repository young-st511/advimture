package game

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/young-st511/advimture/internal/data"
	"github.com/young-st511/advimture/internal/editor"
	"github.com/young-st511/advimture/internal/ui"
)

// MissionState represents the phase of a mission session.
type MissionState int

const (
	MissionPreview   MissionState = iota
	MissionChallenge              // editor active
	MissionResult                 // result/grade screen
	MissionAborted                // user quit without completing
)

// MissionModel is the Bubbletea model for a mission session.
type MissionModel struct {
	mission   *data.MissionData
	state     MissionState
	editor    editor.Model
	startTime time.Time
	width     int
	height    int
	quitting  bool
	result    *GradeResult
}

// NewMission creates a new MissionModel from mission data.
func NewMission(m *data.MissionData) MissionModel {
	return MissionModel{
		mission: m,
		state:   MissionPreview,
	}
}

func (m MissionModel) Init() tea.Cmd { return nil }

// Update handles key input for the mission.
func (m MissionModel) Update(msg tea.Msg) (MissionModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if m.state == MissionChallenge {
			m.editor.SetSize(msg.Width, msg.Height-3)
		}

	case tea.KeyMsg:
		key := msg.String()

		if key == "ctrl+c" {
			m.quitting = true
			return m, nil
		}

		switch m.state {
		case MissionPreview:
			switch key {
			case "enter":
				m = m.startChallenge()
			case "b", "q":
				m.state = MissionAborted
				m.quitting = true
			}

		case MissionChallenge:
			// cursor_on_line: Enter in Normal mode on matching line completes mission
			if m.mission.GoalType == "cursor_on_line" && key == "enter" &&
				m.editor.GetMode() == editor.ModeNormal {
				curRow := m.editor.GetCursor().Row
				line := m.editor.GetBuffer().GetLine(curRow)
				if strings.Contains(line, m.mission.GoalPattern) {
					m.completeCursorGoal()
					return m, nil
				}
			}

			var cmd tea.Cmd
			m.editor, cmd = m.editor.Update(msg)

			if m.editor.Quitting() {
				if m.editor.SavedAndQuit() {
					m.validateAndComplete()
				} else {
					// :q! → abort
					m.state = MissionAborted
					m.quitting = true
				}
			}
			return m, cmd

		case MissionResult:
			switch key {
			case "enter":
				m = m.startChallenge()
			case "b", "q":
				m.quitting = true
			}
		}
	}

	return m, nil
}

func (m MissionModel) startChallenge() MissionModel {
	lines := strings.Split(m.mission.InitialText, "\n")
	if len(lines) > 1 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	m.editor = editor.NewWithLines(lines)
	m.editor.SetSize(m.width, m.height-3)
	m.editor.GetCursor().Row = m.mission.CursorStart.Row
	m.editor.GetCursor().Col = m.mission.CursorStart.Col
	m.startTime = time.Now()
	m.state = MissionChallenge
	m.result = nil
	return m
}

func (m *MissionModel) validateAndComplete() {
	elapsed := time.Since(m.startTime)
	effKeys := m.editor.GetEffectiveKeystrokes()
	totalKeys := m.editor.GetTotalKeystrokes()
	actual := m.editor.GetBuffer().GetText()

	ok, diffs, totalDiff := data.CompareText(m.mission.ExpectedText, actual)
	if ok {
		grade := CalcGrade(effKeys, m.mission.OptimalKeystrokes)
		m.result = &GradeResult{
			Grade:       grade,
			EffKeys:     effKeys,
			OptimalKeys: m.mission.OptimalKeystrokes,
			TimeMs:      elapsed.Milliseconds(),
			Accuracy:    CalcAccuracy(effKeys, totalKeys),
			Success:     true,
		}
	} else {
		m.result = &GradeResult{
			Success:   false,
			Diffs:     diffs,
			TotalDiff: totalDiff,
		}
	}
	m.state = MissionResult
}

func (m *MissionModel) completeCursorGoal() {
	elapsed := time.Since(m.startTime)
	effKeys := m.editor.GetEffectiveKeystrokes()
	totalKeys := m.editor.GetTotalKeystrokes()
	grade := CalcGrade(effKeys, m.mission.OptimalKeystrokes)
	m.result = &GradeResult{
		Grade:       grade,
		EffKeys:     effKeys,
		OptimalKeys: m.mission.OptimalKeystrokes,
		TimeMs:      elapsed.Milliseconds(),
		Accuracy:    CalcAccuracy(effKeys, totalKeys),
		Success:     true,
	}
	m.state = MissionResult
}

// activeTip returns the first matching tip message, or "".
func (m MissionModel) activeTip() string {
	effKeys := m.editor.GetEffectiveKeystrokes()
	for _, tip := range m.mission.Tips {
		if strings.HasPrefix(tip.Trigger, "keystroke_over_") {
			nStr := strings.TrimPrefix(tip.Trigger, "keystroke_over_")
			n, err := strconv.Atoi(nStr)
			if err == nil && effKeys > n {
				return tip.Message
			}
		}
	}
	return ""
}

// ─── Exported accessors ───────────────────────────────────────────────────────

func (m MissionModel) Quitting() bool       { return m.quitting }
func (m MissionModel) Completed() bool      { return m.result != nil && m.result.Success }
func (m MissionModel) Aborted() bool        { return m.state == MissionAborted }
func (m MissionModel) MissionID() string    { return m.mission.ID }
func (m MissionModel) Result() *GradeResult { return m.result }

// ─── View ─────────────────────────────────────────────────────────────────────

func (m MissionModel) View() string {
	if m.quitting {
		return ""
	}
	switch m.state {
	case MissionPreview:
		return m.viewPreview()
	case MissionChallenge:
		return m.viewChallenge()
	case MissionResult:
		return m.viewResult()
	}
	return ""
}

func (m MissionModel) viewPreview() string {
	var b strings.Builder
	b.WriteString(ui.HeaderStyle.Render("[ MISSION ]"))
	b.WriteString("\n\n")

	stars := strings.Repeat("★", m.mission.Difficulty) + strings.Repeat("☆", 3-m.mission.Difficulty)
	width := 60
	if m.width > 0 && m.width-4 < width {
		width = m.width - 4
	}
	ticketStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ui.NormalColor).
		Padding(1, 2).
		Width(width)

	ticket := fmt.Sprintf("%s  [%s]\n\n%s\n\n최적 키스트로크: %d",
		lipgloss.NewStyle().Bold(true).Foreground(ui.NormalColor).Render(m.mission.Title),
		stars,
		lipgloss.NewStyle().Italic(true).Foreground(ui.MentorColor).Render(m.mission.Story),
		m.mission.OptimalKeystrokes,
	)
	b.WriteString(ticketStyle.Render(ticket))
	b.WriteString("\n\n")
	b.WriteString(ui.DimStyle.Render("[Enter] 시작   [b] 돌아가기"))
	return b.String()
}

func (m MissionModel) viewChallenge() string {
	var b strings.Builder
	b.WriteString(ui.HeaderStyle.Render(fmt.Sprintf("Mission: %s", m.mission.Title)))
	b.WriteString("\n")
	b.WriteString(m.editor.View())
	if tip := m.activeTip(); tip != "" {
		b.WriteString("\n")
		b.WriteString(ui.MentorStyle.Render("💡 TIP: " + tip))
	}
	return b.String()
}

func (m MissionModel) viewResult() string {
	if m.result == nil {
		return ""
	}
	var b strings.Builder
	b.WriteString(ui.HeaderStyle.Render("[ RESULT ]"))
	b.WriteString("\n\n")

	if !m.result.Success {
		b.WriteString(lipgloss.NewStyle().Foreground(ui.ErrorColor).Bold(true).Render("✗ 텍스트가 일치하지 않습니다"))
		b.WriteString("\n\n")
		for _, d := range m.result.Diffs {
			b.WriteString(fmt.Sprintf("Line %d:\n", d.LineNum))
			b.WriteString(lipgloss.NewStyle().Foreground(ui.ErrorColor).Render("  ✗ 입력: "+d.Yours))
			b.WriteString("\n")
			b.WriteString(lipgloss.NewStyle().Foreground(ui.SuccessColor).Render("  ✓ 정답: "+d.Expected))
			b.WriteString("\n")
		}
		if m.result.TotalDiff > 3 {
			b.WriteString(ui.DimStyle.Render(fmt.Sprintf("  ... 외 %d줄 불일치", m.result.TotalDiff-3)))
			b.WriteString("\n")
		}
		b.WriteString("\n")
		b.WriteString(ui.DimStyle.Render("[Enter] 다시하기   [b] 메뉴"))
		return b.String()
	}

	gradeStyle := lipgloss.NewStyle().Bold(true).Foreground(gradeColor(m.result.Grade))
	b.WriteString(fmt.Sprintf("등급: %s\n", gradeStyle.Render(m.result.Grade)))
	b.WriteString(fmt.Sprintf("키스트로크: %d / 최적: %d\n", m.result.EffKeys, m.result.OptimalKeys))
	b.WriteString(fmt.Sprintf("시간: %.1f초\n", float64(m.result.TimeMs)/1000.0))
	b.WriteString(fmt.Sprintf("정확도: %.0f%%\n", m.result.Accuracy))
	b.WriteString("\n")

	width := 60
	if m.width > 0 && m.width-4 < width {
		width = m.width - 4
	}
	mentorBox := lipgloss.NewStyle().
		Foreground(ui.MentorColor).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ui.MentorColor).
		Padding(0, 1).
		Width(width).
		Render("Vi 선배: " + MentorMessage(m.result.Grade))
	b.WriteString(mentorBox)
	b.WriteString("\n\n")

	if len(m.mission.OptimalSolutions) > 0 {
		b.WriteString(ui.DimStyle.Render("최적 풀이: " + m.mission.OptimalSolutions[0].Keys))
		b.WriteString("\n\n")
	}
	b.WriteString(ui.DimStyle.Render("[Enter] 다시하기   [b] 메뉴"))
	return b.String()
}

func gradeColor(grade string) lipgloss.Color {
	switch grade {
	case "S":
		return lipgloss.Color("#FFD700")
	case "A":
		return lipgloss.Color("#50C878")
	case "B":
		return lipgloss.Color("#5B7FFF")
	default:
		return lipgloss.Color("#888888")
	}
}
