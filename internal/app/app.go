package app

import (
	"io/fs"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/young-st511/advimture/internal/e2estate"
	"github.com/young-st511/advimture/internal/playable"
	"github.com/young-st511/advimture/internal/progress"
)

type Model struct {
	playable playable.Model
}

type Options struct {
	ContentFS fs.FS
}

func New() Model {
	return NewWithOptions(Options{})
}

func NewWithOptions(options Options) Model {
	p, _ := progress.Load()
	return Model{
		playable: playable.New(playable.Options{
			Progress:     p,
			SaveProgress: progress.Save,
			E2EStatePath: e2eStatePath(),
			ContentRoot:  "content",
			ContentFS:    options.ContentFS,
		}),
	}
}

func (m Model) Init() tea.Cmd {
	return m.playable.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	updated, cmd := m.playable.Update(msg)
	if next, ok := updated.(playable.Model); ok {
		m.playable = next
	}
	return m, cmd
}

func (m Model) View() string {
	return m.playable.View()
}

func e2eStatePath() string {
	if os.Getenv("ADVIMTURE_E2E") != "1" {
		return ""
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return e2estate.DefaultPath(home)
}
