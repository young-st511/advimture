package game

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/young-st511/advimture/internal/ui"
)

// FTUEPhase represents the current phase of the first-time user experience.
type FTUEPhase int

const (
	FTUEIntro FTUEPhase = iota
	FTUESSHTyping
	FTUEVimPanic
	FTUEMentorAppear
	FTUEQuitPrompt
	FTUEDone
)

// FTUEModel is the Bubbletea model for the first-time user experience.
type FTUEModel struct {
	phase    FTUEPhase
	width    int
	height   int
	quitting bool
	done     bool
	input    string // for :q! input tracking
}

// NewFTUE creates a new FTUE model.
func NewFTUE() FTUEModel {
	return FTUEModel{
		phase: FTUEIntro,
	}
}

func (m FTUEModel) Init() tea.Cmd {
	return nil
}

// Update handles FTUE input.
func (m FTUEModel) Update(msg tea.Msg) (FTUEModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		key := msg.String()

		if key == "ctrl+c" {
			m.quitting = true
			return m, nil
		}

		switch m.phase {
		case FTUEIntro:
			if key == "enter" || key == " " {
				m.phase = FTUESSHTyping
			}
		case FTUESSHTyping:
			if key == "enter" || key == " " {
				m.phase = FTUEVimPanic
			}
		case FTUEVimPanic:
			if key == "enter" || key == " " {
				m.phase = FTUEMentorAppear
			}
		case FTUEMentorAppear:
			if key == "enter" || key == " " {
				m.phase = FTUEQuitPrompt
				m.input = ""
			}
		case FTUEQuitPrompt:
			// Track :q! input
			switch key {
			case ":":
				m.input = ":"
			case "enter":
				if m.input == ":q!" {
					m.phase = FTUEDone
				}
				m.input = ""
			default:
				if len(m.input) > 0 && m.input[0] == ':' {
					m.input += key
				}
			}
		case FTUEDone:
			if key == "enter" || key == " " {
				m.done = true
			}
		}
	}

	return m, nil
}

// Quitting returns whether the FTUE should exit.
func (m FTUEModel) Quitting() bool { return m.quitting }

// Done returns whether the FTUE is complete and should transition to tutorial.
func (m FTUEModel) Done() bool { return m.done }

// View renders the FTUE screen.
func (m FTUEModel) View() string {
	var b strings.Builder

	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#5B7FFF"))
	mentorStyle := lipgloss.NewStyle().Foreground(ui.MentorColor).Italic(true)
	dimStyle := ui.DimStyle
	panicStyle := lipgloss.NewStyle().Foreground(ui.ErrorColor)

	switch m.phase {
	case FTUEIntro:
		b.WriteString(titleStyle.Render("⌨  A D V I M T U R E"))
		b.WriteString("\n\n")
		b.WriteString("당신은 신입 DevOps 엔지니어.\n")
		b.WriteString("첫 출근 날, 프로덕션 서버에 긴급 수정이 필요합니다.\n\n")
		b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Render(
			"팀장: \"이 서버에는 nano가 없어. vim으로 수정해줘.\""))
		b.WriteString("\n\n")
		b.WriteString(dimStyle.Render("Enter 키를 눌러 계속..."))

	case FTUESSHTyping:
		b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#50C878")).Render(
			"$ ssh production-server-01"))
		b.WriteString("\n")
		b.WriteString("Connected.\n\n")
		b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#50C878")).Render(
			"$ vim /etc/nginx/nginx.conf"))
		b.WriteString("\n\n")
		b.WriteString(dimStyle.Render("Enter 키를 눌러 계속..."))

	case FTUEVimPanic:
		b.WriteString("~\n~\n~\n~\n")
		b.WriteString(panicStyle.Render("-- 이게 뭐지? 어떻게 나가지?? --"))
		b.WriteString("\n\n")
		b.WriteString("타이핑해도 아무것도 안 되고...\n")
		b.WriteString("마우스도 안 먹히고...\n\n")
		b.WriteString(panicStyle.Render("😱 Vim에 갇혔다!"))
		b.WriteString("\n\n")
		b.WriteString(dimStyle.Render("Enter 키를 눌러 계속..."))

	case FTUEMentorAppear:
		b.WriteString(mentorStyle.Render("???: \"이봐, 신입. 당황하지 마.\""))
		b.WriteString("\n\n")
		b.WriteString(mentorStyle.Render("Vi 선배: \"나는 Vi 선배. Vim의 세계를 안내해 줄게.\""))
		b.WriteString("\n\n")
		b.WriteString(mentorStyle.Render("Vi 선배: \"일단 여기서 나가는 법부터 알려줄게.\""))
		b.WriteString("\n")
		b.WriteString(mentorStyle.Render("          :q! 를 입력하면 저장 없이 나갈 수 있어.\""))
		b.WriteString("\n\n")
		b.WriteString(dimStyle.Render("Enter 키를 눌러 직접 해보세요..."))

	case FTUEQuitPrompt:
		b.WriteString("~\n~\n~\n")
		b.WriteString(fmt.Sprintf("  %d   # nginx configuration\n", 1))
		b.WriteString(fmt.Sprintf("  %d   server {\n", 2))
		b.WriteString(fmt.Sprintf("  %d       listen 80;\n", 3))
		b.WriteString(fmt.Sprintf("  %d   }\n", 4))
		b.WriteString("~\n~\n")

		if len(m.input) > 0 {
			b.WriteString(m.input)
		}
		b.WriteString("\n\n")
		b.WriteString(mentorStyle.Render("Vi 선배: \":q! 를 입력하고 Enter를 눌러봐.\""))

	case FTUEDone:
		b.WriteString("\n")
		b.WriteString(lipgloss.NewStyle().Foreground(ui.SuccessColor).Bold(true).Render(
			"✓ 첫 번째 관문 통과!"))
		b.WriteString("\n\n")
		b.WriteString(mentorStyle.Render("Vi 선배: \"잘했어! 이제 본격적으로 Vim을 배워보자.\""))
		b.WriteString("\n\n")
		b.WriteString(dimStyle.Render("Enter 키를 눌러 튜토리얼을 시작하세요"))
	}

	return b.String()
}
