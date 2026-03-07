package game

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/young-st511/advimture/internal/data"
	"github.com/young-st511/advimture/internal/editor"
	"github.com/young-st511/advimture/internal/ui"
)

// TutorialState represents the current phase of the tutorial.
type TutorialState int

const (
	StateInstruction TutorialState = iota
	StatePractice
	StateComplete
	StateAllDone
)

// TutorialModel is the Bubbletea model for a tutorial session.
type TutorialModel struct {
	tutorial    *data.TutorialData
	substepIdx  int
	state       TutorialState
	editor      editor.Model
	checker     *data.GoalChecker
	startTime   time.Time
	hintLevel   int // 0=none, 1-3=progressive hints
	width       int
	height      int
	quitting    bool
	completed   bool
	keystrokes  int
	elapsedSecs float64
	errorMsg    string // shown when :wq is attempted but goal not met
}

// NewTutorial creates a new tutorial model from tutorial data.
func NewTutorial(t *data.TutorialData) TutorialModel {
	m := TutorialModel{
		tutorial: t,
		checker:  data.NewGoalChecker(),
	}
	m.initSubstep()
	return m
}

func (m *TutorialModel) initSubstep() {
	if m.substepIdx >= len(m.tutorial.Substeps) {
		m.state = StateAllDone
		return
	}

	sub := m.tutorial.Substeps[m.substepIdx]
	lines := strings.Split(sub.InitialText, "\n")
	m.editor = editor.NewWithLines(lines)
	m.checker.Reset()
	m.state = StateInstruction
	m.startTime = time.Now()
	m.hintLevel = 0
	m.keystrokes = 0
}

func (m TutorialModel) Init() tea.Cmd {
	return nil
}

// Update handles input for the tutorial.
func (m TutorialModel) Update(msg tea.Msg) (TutorialModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		key := msg.String()

		// Global controls
		if key == "ctrl+c" {
			m.quitting = true
			return m, nil
		}

		switch m.state {
		case StateInstruction:
			// Any key to start practice
			if key == "enter" || key == " " {
				m.state = StatePractice
				m.startTime = time.Now()
			}

		case StatePractice:
			// ? for hints
			if key == "?" && m.editor.GetMode() == editor.ModeNormal {
				m.hintLevel++
				if m.hintLevel > 3 {
					m.hintLevel = 3
				}
				return m, nil
			}

			// Check allowed keys
			sub := m.tutorial.Substeps[m.substepIdx]
			if len(sub.AllowedKeys) > 0 && !isKeyAllowed(key, sub.AllowedKeys) {
				// Key not allowed — ignore
				return m, nil
			}

			// Forward to editor
			m.keystrokes++
			prevMode := m.editor.GetMode()
			var cmd tea.Cmd
			m.editor, cmd = m.editor.Update(msg)

			// Track esc from insert mode
			if prevMode == editor.ModeInsert && m.editor.GetMode() == editor.ModeNormal {
				m.checker.RecordCommand("esc_from_insert")
			}

			// Handle editor quit attempts
			if m.editor.Quitting() {
				if m.editor.SavedAndQuit() {
					// :wq — check goal
					m.checker.RecordSaveQuit()
					if m.checkGoal() {
						m.elapsedSecs = time.Since(m.startTime).Seconds()
						m.state = StateComplete
						m.errorMsg = ""
					} else {
						// Goal not met — block exit, reset editor quit state
						m.editor.ResetQuit()
						m.errorMsg = "아직 완료되지 않았어. 내용을 다시 확인해봐!"
					}
				} else {
					// :q! — abort tutorial, return to menu
					m.quitting = true
				}
				return m, cmd
			}

			// Check goal (for non-wq goals like cursor_position, command_used)
			if m.checkGoal() {
				m.elapsedSecs = time.Since(m.startTime).Seconds()
				m.state = StateComplete
				m.errorMsg = ""
			}

			return m, cmd

		case StateComplete:
			if key == "enter" || key == " " {
				m.substepIdx++
				m.initSubstep()
			}

		case StateAllDone:
			if key == "enter" || key == " " || key == "q" {
				m.completed = true
				m.quitting = true
			}
		}
	}

	return m, nil
}

func (m TutorialModel) checkGoal() bool {
	sub := m.tutorial.Substeps[m.substepIdx]
	bufText := m.editor.GetBuffer().GetText()
	curRow := m.editor.GetCursor().Row
	curCol := m.editor.GetCursor().Col
	mode := m.editor.GetMode().String()
	return m.checker.CheckGoal(sub.Goal, bufText, curRow, curCol, mode)
}

func isKeyAllowed(key string, allowed []string) bool {
	for _, k := range allowed {
		if k == key {
			return true
		}
	}
	// Always allow esc and ctrl+c
	return key == "esc" || key == "ctrl+c"
}

// Quitting returns whether the tutorial should exit.
func (m TutorialModel) Quitting() bool { return m.quitting }

// Completed returns whether the tutorial was fully completed.
func (m TutorialModel) Completed() bool { return m.completed }

// TutorialID returns the tutorial ID.
func (m TutorialModel) TutorialID() string { return m.tutorial.ID }

// ElapsedSeconds returns the time spent.
func (m TutorialModel) ElapsedSeconds() float64 { return m.elapsedSecs }

// Keystrokes returns total keystrokes.
func (m TutorialModel) Keystrokes() int { return m.keystrokes }

// View renders the tutorial screen.
func (m TutorialModel) View() string {
	if m.quitting && m.state != StateAllDone {
		return ""
	}

	var b strings.Builder

	sub := m.currentSubstep()

	// Header
	header := fmt.Sprintf("Tutorial: %s", m.tutorial.Title)
	if sub != nil {
		header += fmt.Sprintf(" — %s (%d/%d)", sub.Title, m.substepIdx+1, len(m.tutorial.Substeps))
	}
	b.WriteString(ui.HeaderStyle.Render(header))
	b.WriteString("\n")

	// Mentor message area
	if sub != nil && sub.MentorMsg != "" && m.state == StateInstruction {
		mentorBox := lipgloss.NewStyle().
			Foreground(ui.MentorColor).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ui.MentorColor).
			Padding(0, 1).
			Width(m.width - 4).
			Render(fmt.Sprintf("Vi 선배: %s", sub.MentorMsg))
		b.WriteString(mentorBox)
		b.WriteString("\n")
	}

	switch m.state {
	case StateInstruction:
		if sub != nil {
			b.WriteString("\n")
			b.WriteString(sub.Instruction)
			b.WriteString("\n\n")
			b.WriteString(ui.DimStyle.Render("Enter 키를 눌러 시작하세요"))
		}

	case StatePractice:
		// Editor view
		b.WriteString(m.editor.View())

		// Error message (goal not met on :wq)
		if m.errorMsg != "" {
			errStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4444")).Bold(true)
			b.WriteString("\n")
			b.WriteString(errStyle.Render("✗ " + m.errorMsg))
		}

		// Hint area
		if sub != nil && m.hintLevel > 0 && m.hintLevel <= len(sub.Hints) {
			b.WriteString("\n")
			hint := sub.Hints[m.hintLevel-1]
			hintStyle := lipgloss.NewStyle().
				Foreground(ui.MentorColor).
				Italic(true)
			b.WriteString(hintStyle.Render(fmt.Sprintf("💡 힌트 %d: %s", m.hintLevel, hint)))
		}
		if sub != nil {
			b.WriteString("\n")
			b.WriteString(ui.DimStyle.Render("? 키로 힌트 보기 | Ctrl+c: 종료"))
		}

	case StateComplete:
		b.WriteString("\n")
		successStyle := lipgloss.NewStyle().Foreground(ui.SuccessColor).Bold(true)
		b.WriteString(successStyle.Render("✓ 완료!"))
		b.WriteString(fmt.Sprintf("  시간: %.1f초  키 입력: %d", m.elapsedSecs, m.keystrokes))
		b.WriteString("\n\n")
		b.WriteString(ui.DimStyle.Render("Enter 키를 눌러 다음으로 진행하세요"))

	case StateAllDone:
		b.WriteString("\n")
		doneStyle := lipgloss.NewStyle().Foreground(ui.SuccessColor).Bold(true)
		b.WriteString(doneStyle.Render("🎉 모든 단계를 완료했습니다!"))
		b.WriteString("\n\n")
		b.WriteString(ui.DimStyle.Render("Enter 키를 눌러 돌아가세요"))
	}

	return b.String()
}

func (m TutorialModel) currentSubstep() *data.SubstepData {
	if m.substepIdx < len(m.tutorial.Substeps) {
		return &m.tutorial.Substeps[m.substepIdx]
	}
	return nil
}
