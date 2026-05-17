package playable

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/young-st511/advimture/internal/content"
	"github.com/young-st511/advimture/internal/e2estate"
	"github.com/young-st511/advimture/internal/progress"
	"github.com/young-st511/advimture/internal/progressadapter"
	exerciseruntime "github.com/young-st511/advimture/internal/runtime"
	"github.com/young-st511/advimture/internal/scenario"
	"github.com/young-st511/advimture/internal/tuiadapter"
	"github.com/young-st511/advimture/internal/vimengine"
)

const missionID = "mission-1"

type Options struct {
	Progress     *progress.Progress
	SaveProgress func(*progress.Progress) error
	E2EStatePath string
	Now          func() time.Time
}

type Model struct {
	run          *scenario.Run
	progress     *progress.Progress
	saveProgress func(*progress.Progress) error
	e2eStatePath string
	now          func() time.Time
	startedAt    time.Time
	saved        bool
	err          error
}

func New(options Options) Model {
	now := options.Now
	if now == nil {
		now = time.Now
	}
	p := options.Progress
	if p == nil {
		p = progress.NewProgress()
	}

	run, err := newRun()
	model := Model{
		run:          run,
		progress:     p,
		saveProgress: options.SaveProgress,
		e2eStatePath: options.E2EStatePath,
		now:          now,
		startedAt:    now(),
		err:          err,
	}
	_ = model.writeE2EState()
	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.err != nil {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		action := tuiadapter.MapInput(msg.String())
		switch action.Type {
		case tuiadapter.ActionKey:
			m.run.ApplyKey(action.Key)
			m.applyProgressIfSucceeded()
		case tuiadapter.ActionHint:
			if hint, ok := m.run.RequestHint(); ok {
				_ = hint
			}
		case tuiadapter.ActionRetry:
			m.run.Retry()
			m.saved = false
		case tuiadapter.ActionQuit:
			_ = m.writeE2EState()
			return m, tea.Quit
		}
	}

	_ = m.writeE2EState()
	return m, nil
}

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Playable error: %v\nq: quit", m.err)
	}

	state := m.run.State()
	view := tuiadapter.RenderState(state)
	var b strings.Builder
	b.WriteString("A D V I M T U R E - Playable Slice\n\n")
	b.WriteString(view.Title + "\n")
	b.WriteString(view.Message + "\n\n")
	for row, line := range view.BufferLines {
		b.WriteString(renderLine(line, row, view.CursorRow, view.CursorCol))
		b.WriteString("\n")
	}
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Mode: %s  Status: %s  Cursor: %d,%d\n", view.Mode, view.Status, view.CursorRow, view.CursorCol))
	if view.Grade != "" {
		b.WriteString(fmt.Sprintf("Grade: %s\n", view.Grade))
	}
	b.WriteString("Keys: h/j/k/l move  ?: hint  r: retry  q: quit")
	return b.String()
}

func (m Model) State() e2estate.State {
	if m.run == nil {
		return e2estate.State{}
	}
	state := m.run.State()
	var score e2estate.Score
	if state.Score != nil {
		score = e2estate.Score{
			Grade:  string(state.Score.Grade),
			Passed: state.Score.Passed,
		}
	}
	progressState := e2estate.Progress{
		MissionID: missionID,
		Completed: state.Status == exerciseruntime.StatusSucceeded,
	}
	return e2estate.State{
		Buffer: append([]string(nil), state.Runtime.Vim.Lines...),
		Cursor: e2estate.Cursor{
			Row: state.Runtime.Vim.Cursor.Row,
			Col: state.Runtime.Vim.Cursor.Col,
		},
		Mode:     string(state.Runtime.Vim.Mode),
		Status:   string(state.Status),
		Score:    score,
		Progress: progressState,
	}
}

func (m Model) applyProgressIfSucceeded() {
	if m.saved || m.run.State().Status != exerciseruntime.StatusSucceeded {
		return
	}
	completion, err := progressadapter.MissionCompletionFromScenario(missionID, m.run.State(), m.now().Sub(m.startedAt))
	if err != nil {
		return
	}
	updated := progressadapter.ApplyMissionCompletion(*m.progress, completion)
	m.progress = &updated
	if m.saveProgress != nil {
		_ = m.saveProgress(m.progress)
	}
	m.saved = true
}

func (m Model) writeE2EState() error {
	return e2estate.Write(m.e2eStatePath, m.State())
}

func newRun() (*scenario.Run, error) {
	compiled, err := content.CompileExercise(content.ExerciseSpec{
		ID:               "move-right",
		CommandClusterID: "normal-motion-basic",
		Title:            "Move right twice",
		Initial: content.StateSpec{
			Lines: []string{"abc"},
			Mode:  string(vimengine.ModeNormal),
		},
		Goal: content.GoalSpec{
			Cursor: content.CursorSpecPtr(0, 2),
			Mode:   string(vimengine.ModeNormal),
		},
		Hints: []content.HintSpec{
			{AfterKeys: 1, Text: "Use l twice."},
		},
		ExpectedKeys: []string{vimengine.KeyL, vimengine.KeyL},
		AllowedKeys:  []string{vimengine.KeyH, vimengine.KeyJ, vimengine.KeyK, vimengine.KeyL},
	})
	if err != nil {
		return nil, err
	}
	return scenario.NewRun(scenario.Spec{
		ID:          "door",
		Title:       "Open the door",
		Briefing:    "Reach the marked column.",
		SuccessText: "Door opened.",
		Exercise:    compiled,
	})
}

func renderLine(line string, row int, cursorRow int, cursorCol int) string {
	if row != cursorRow {
		return "  " + line
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
	return fmt.Sprintf("> %s[%s]%s", string(runes[:cursorCol]), string(runes[cursorCol]), string(runes[cursorCol+1:]))
}

var _ tea.Model = Model{}
