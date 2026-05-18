package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/young-st511/advimture/internal/data"
	"github.com/young-st511/advimture/internal/progress"
)

// MissionListItem is one row in the mission list.
type MissionListItem struct {
	Mission  *data.MissionData
	Unlocked bool
	Grade    string // best grade achieved, "" if not played
	Filename string // yaml filename for loading
}

// MissionListModel is the Bubbletea model for the mission selection screen.
type MissionListModel struct {
	items    []MissionListItem
	selected int
	width    int
	height   int
	quitting bool
	chosen   string // filename of chosen mission, "" if none
}

// NewMissionList creates a MissionListModel from a list of missions and player progress.
func NewMissionList(missions []*data.MissionData, p *progress.Progress) MissionListModel {
	items := make([]MissionListItem, len(missions))
	for i, m := range missions {
		grade := ""
		if p != nil {
			if mp, ok := p.Missions[m.ID]; ok && mp.Completed {
				grade = mp.BestGrade
			}
		}
		items[i] = MissionListItem{
			Mission:  m,
			Unlocked: progress.IsMissionUnlocked(m, p),
			Grade:    grade,
			Filename: missionFilename(m.ID),
		}
	}
	return MissionListModel{items: items}
}

// missionFilename converts mission ID to filename (e.g. "m-01" → "m01.yaml").
func missionFilename(id string) string {
	// "m-01" → "m01.yaml"
	cleaned := strings.ReplaceAll(id, "-", "")
	return cleaned + ".yaml"
}

func (m MissionListModel) Init() tea.Cmd { return nil }

// Update handles key input for the mission list.
func (m MissionListModel) Update(msg tea.Msg) (MissionListModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.selected < len(m.items)-1 {
				m.selected++
			}
		case "k", "up":
			if m.selected > 0 {
				m.selected--
			}
		case "enter":
			if m.selected < len(m.items) {
				item := m.items[m.selected]
				if item.Unlocked {
					m.chosen = item.Filename
				}
			}
		case "b", "q":
			m.quitting = true
		case "ctrl+c":
			m.quitting = true
		}
	}
	return m, nil
}

// Chosen returns the filename of the chosen mission (resets after read).
func (m *MissionListModel) Chosen() string {
	f := m.chosen
	m.chosen = ""
	return f
}

// Quitting returns whether the user wants to leave this screen.
func (m MissionListModel) Quitting() bool { return m.quitting }

// View renders the mission list.
func (m MissionListModel) View() string {
	var b strings.Builder

	b.WriteString(HeaderStyle.Render("[ MISSION LIST ]"))
	b.WriteString("\n\n")

	if len(m.items) == 0 {
		b.WriteString(DimStyle.Render("미션이 없습니다."))
		b.WriteString("\n\n")
		b.WriteString(DimStyle.Render("[b] 돌아가기"))
		return b.String()
	}

	for i, item := range m.items {
		cursor := "  "
		if i == m.selected {
			cursor = "> "
		}

		stars := strings.Repeat("★", item.Mission.Difficulty) + strings.Repeat("☆", 3-item.Mission.Difficulty)
		gradeTag := ""
		if item.Grade != "" {
			gradeTag = fmt.Sprintf(" [%s]", item.Grade)
		}

		label := fmt.Sprintf("%s%s  %s%s", cursor, item.Mission.Title, stars, gradeTag)

		if !item.Unlocked {
			lockMsg := "  (잠김: Tutorial " + strings.Join(item.Mission.RequiredTutorials, ", ") + " 필요)"
			b.WriteString(lipgloss.NewStyle().Foreground(DimColor).Render(label + lockMsg))
		} else if i == m.selected {
			b.WriteString(lipgloss.NewStyle().Foreground(NormalColor).Bold(true).Render(label))
		} else {
			b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC")).Render(label))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(DimStyle.Render("j/k: 이동  Enter: 선택  b: 돌아가기"))

	return b.String()
}
