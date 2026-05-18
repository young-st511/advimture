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
	ContentRoot  string
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

	run, err := newRunFromContent(contentRoot(options.ContentRoot))
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
	switch msg := msg.(type) {
	case tea.KeyMsg:
		action := tuiadapter.MapInputForMode(msg.String(), m.inputMode())
		if m.err != nil {
			if action.Type == tuiadapter.ActionQuit {
				_ = m.writeE2EState()
				return m, tea.Quit
			}
			return m, nil
		}
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

	if m.err != nil {
		return m, nil
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
	if view.Mode == string(vimengine.ModeCommand) {
		b.WriteString(":" + view.CommandLine + "\n")
		b.WriteString("Keys: type command  enter: run  esc: normal  ctrl+c: quit")
		return b.String()
	}
	if view.LastCommand != "" {
		b.WriteString(fmt.Sprintf("Command: %s\n", view.LastCommand))
	}
	b.WriteString("Keys: h/j/k/l/w/b/e/g/G/0/$ move  ?: hint  r: retry  q: quit  :: command")
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
		Command:  state.Runtime.Vim.LastCommand,
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

func (m Model) inputMode() vimengine.Mode {
	if m.run == nil {
		return ""
	}
	return m.run.State().Runtime.Vim.Mode
}

func newRunFromContent(root string) (*scenario.Run, error) {
	library, err := content.LoadLibrary(root)
	if err != nil {
		return nil, err
	}
	playable := library.PlayableExercises()
	if len(playable) == 0 {
		return nil, fmt.Errorf("no playable exercises in %s", root)
	}
	exercise := playable[0]
	scenarioDoc, err := scenarioForExercise(library, exercise.ID)
	if err != nil {
		return nil, err
	}
	compiled, err := library.CompileExercise(exercise.ID)
	if err != nil {
		return nil, err
	}
	return scenario.NewRun(scenario.Spec{
		ID:          scenarioDoc.ID,
		Title:       scenarioDoc.MissionTitle,
		Briefing:    scenarioDoc.Briefing,
		SuccessText: scenarioDoc.MentorSuccess,
		Exercise:    compiled,
	})
}

func scenarioForExercise(library content.Library, exerciseID string) (content.ScenarioDocument, error) {
	var selected content.ScenarioDocument
	for _, scenarioDoc := range library.Scenarios {
		if scenarioDoc.ExerciseID != exerciseID {
			continue
		}
		if scenarioDoc.EngineSupport != content.EngineSupportImplemented {
			continue
		}
		if scenarioDoc.Status != content.StatusApproved && scenarioDoc.Status != content.StatusImplemented {
			continue
		}
		if selected.ID == "" || scenarioDoc.ID < selected.ID {
			selected = scenarioDoc
		}
	}
	if selected.ID == "" {
		return content.ScenarioDocument{}, fmt.Errorf("no playable scenario for exercise %q", exerciseID)
	}
	return selected, nil
}

func contentRoot(root string) string {
	if root != "" {
		return root
	}
	return "content"
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
