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

type Options struct {
	Progress     *progress.Progress
	SaveProgress func(*progress.Progress) error
	E2EStatePath string
	ContentRoot  string
	Now          func() time.Time
}

type Model struct {
	run          *scenario.Run
	entries      []gameEntry
	current      int
	contentRoot  string
	progress     *progress.Progress
	saveProgress func(*progress.Progress) error
	e2eStatePath string
	now          func() time.Time
	startedAt    time.Time
	saved        bool
	err          error
}

type gameEntry struct {
	PlaylistID      string
	PlaylistTitle   string
	ExerciseID      string
	ScenarioID      string
	IndexInPlaylist int
	TotalInPlaylist int
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

	root := contentRoot(options.ContentRoot)
	entries, current, run, err := newGameFromContent(root, p)
	model := Model{
		run:          run,
		entries:      entries,
		current:      current,
		contentRoot:  root,
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
			if m.run.State().Status == exerciseruntime.StatusFailed && action.Key == vimengine.KeyEnter {
				m.retryCurrent()
				break
			}
			if m.run.State().Status == exerciseruntime.StatusSucceeded && action.Key == vimengine.KeyEnter {
				m.advanceToNext()
				break
			}
			m.run.ApplyKey(action.Key)
			m.applyProgressIfSucceeded()
		case tuiadapter.ActionHint:
			if hint, ok := m.run.RequestHint(); ok {
				_ = hint
			}
		case tuiadapter.ActionRetry:
			m.retryCurrent()
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
	if entry, ok := m.currentEntry(); ok {
		b.WriteString(entry.PlaylistTitle + "\n")
	}
	b.WriteString(view.Title + "\n")
	b.WriteString(view.Message + "\n\n")
	for row, line := range view.BufferLines {
		b.WriteString(renderLine(line, row, view.CursorRow, view.CursorCol))
		b.WriteString("\n")
	}
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Mode: %s  Status: %s  Cursor: %d,%d\n", view.Mode, view.Status, view.CursorRow, view.CursorCol))
	if entry, ok := m.currentEntry(); ok {
		b.WriteString(fmt.Sprintf("Exercise: %d/%d\n", entry.IndexInPlaylist+1, entry.TotalInPlaylist))
	}
	if view.Grade != "" {
		b.WriteString(fmt.Sprintf("Grade: %s\n", view.Grade))
	}
	if state.Runtime.MaxInputs > 0 {
		b.WriteString(fmt.Sprintf("Inputs left: %d/%d\n", state.Runtime.InputsLeft, state.Runtime.MaxInputs))
	}
	if state.Status == exerciseruntime.StatusSucceeded {
		if m.current+1 < len(m.entries) {
			if m.nextEntryStartsNewPlaylist() {
				b.WriteString("Next tutorial: enter\n")
			} else {
				b.WriteString("Next: enter\n")
			}
		} else {
			b.WriteString("Playlist complete\n")
		}
	} else if state.Status == exerciseruntime.StatusFailed {
		b.WriteString("Retry: r or enter\n")
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
		MissionID: m.currentExerciseID(),
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

func (m *Model) applyProgressIfSucceeded() {
	if m.saved || m.run.State().Status != exerciseruntime.StatusSucceeded {
		return
	}
	completion, err := progressadapter.MissionCompletionFromScenario(m.currentExerciseID(), m.run.State(), m.now().Sub(m.startedAt))
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

func (m *Model) advanceToNext() {
	if m.current+1 >= len(m.entries) {
		return
	}
	nextIndex := m.current + 1
	run, err := runForEntry(m.contentRoot, m.entries[nextIndex])
	if err != nil {
		m.err = err
		return
	}
	m.current = nextIndex
	m.run = run
	m.saved = false
	m.startedAt = m.now()
}

func (m *Model) retryCurrent() {
	m.run.Retry()
	m.saved = false
	m.startedAt = m.now()
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

func (m Model) currentExerciseID() string {
	if len(m.entries) == 0 || m.current < 0 || m.current >= len(m.entries) {
		return ""
	}
	return m.entries[m.current].ExerciseID
}

func (m Model) currentEntry() (gameEntry, bool) {
	if len(m.entries) == 0 || m.current < 0 || m.current >= len(m.entries) {
		return gameEntry{}, false
	}
	return m.entries[m.current], true
}

func (m Model) nextEntryStartsNewPlaylist() bool {
	if m.current+1 >= len(m.entries) {
		return false
	}
	return m.entries[m.current].PlaylistID != m.entries[m.current+1].PlaylistID
}

func newGameFromContent(root string, progressState *progress.Progress) ([]gameEntry, int, *scenario.Run, error) {
	library, err := content.LoadLibrary(root)
	if err != nil {
		return nil, 0, nil, err
	}
	entries, err := playlistEntries(library)
	if err != nil {
		return nil, 0, nil, err
	}
	if len(entries) == 0 {
		return nil, 0, nil, fmt.Errorf("no playable exercises in %s", root)
	}
	current := firstIncompleteIndex(entries, progressState)
	run, err := runForEntry(root, entries[current])
	if err != nil {
		return nil, 0, nil, err
	}
	return entries, current, run, nil
}

func runForEntry(root string, entry gameEntry) (*scenario.Run, error) {
	library, err := content.LoadLibrary(root)
	if err != nil {
		return nil, err
	}
	scenarioDoc := library.Scenarios[entry.ScenarioID]
	compiled, err := library.CompileExercise(entry.ExerciseID)
	if err != nil {
		return nil, err
	}
	return scenario.NewRun(scenario.Spec{
		ID:          scenarioDoc.ID,
		Title:       scenarioDoc.MissionTitle,
		Briefing:    scenarioDoc.Briefing,
		SuccessText: scenarioDoc.MentorSuccess,
		FailureText: scenarioDoc.MentorFailure,
		Exercise:    compiled,
	})
}

func playlistEntries(library content.Library) ([]gameEntry, error) {
	playlists := library.PlayablePlaylists()
	if len(playlists) == 0 {
		return nil, fmt.Errorf("no playlists found")
	}

	var entries []gameEntry
	for _, playlist := range playlists {
		var playlistEntries []gameEntry
		for _, beat := range playlist.Beats {
			exercise, ok := library.Exercises[beat.ExerciseID]
			if !ok || !isPlayableExercise(exercise) {
				continue
			}
			scenarioDoc, ok := library.Scenarios[beat.ScenarioID]
			if !ok || !isPlayableScenario(scenarioDoc) {
				continue
			}
			playlistEntries = append(playlistEntries, gameEntry{
				PlaylistID:    playlist.ID,
				PlaylistTitle: playlist.Title,
				ExerciseID:    beat.ExerciseID,
				ScenarioID:    beat.ScenarioID,
			})
		}
		total := len(playlistEntries)
		for index := range playlistEntries {
			playlistEntries[index].IndexInPlaylist = index
			playlistEntries[index].TotalInPlaylist = total
		}
		entries = append(entries, playlistEntries...)
	}
	return entries, nil
}

func isPlayableExercise(exercise content.ExerciseDocument) bool {
	return (exercise.Status == content.StatusApproved || exercise.Status == content.StatusImplemented) &&
		exercise.EngineSupport == content.EngineSupportImplemented &&
		exercise.ReplayStatus == content.ReplayStatusPass
}

func isPlayableScenario(scenarioDoc content.ScenarioDocument) bool {
	return (scenarioDoc.Status == content.StatusApproved || scenarioDoc.Status == content.StatusImplemented) &&
		scenarioDoc.EngineSupport == content.EngineSupportImplemented
}

func firstIncompleteIndex(entries []gameEntry, progressState *progress.Progress) int {
	if progressState == nil {
		return 0
	}
	for index, entry := range entries {
		if !progressState.Missions[entry.ExerciseID].Completed {
			return index
		}
	}
	return len(entries) - 1
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
