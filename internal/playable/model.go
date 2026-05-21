package playable

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/young-st511/advimture/internal/content"
	"github.com/young-st511/advimture/internal/e2estate"
	"github.com/young-st511/advimture/internal/progress"
	"github.com/young-st511/advimture/internal/progressadapter"
	"github.com/young-st511/advimture/internal/review"
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
	reviewQueue  []review.Candidate
	hintMessage  string
}

type gameEntry struct {
	PlaylistID      string
	PlaylistTitle   string
	ExerciseID      string
	ScenarioID      string
	IndexInPlaylist int
	TotalInPlaylist int
}

var actionPanelStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("63")).
	Padding(0, 1).
	Width(58)

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
	model.refreshReviewQueue()
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
			if m.run.State().Status == exerciseruntime.StatusFailed && (action.Key == vimengine.KeyEnter || action.Key == vimengine.KeyR) {
				m.retryCurrent()
				break
			}
			if m.run.State().Status == exerciseruntime.StatusSucceeded && action.Key == vimengine.KeyEnter {
				m.advanceToNext()
				break
			}
			m.hintMessage = ""
			m.run.ApplyKey(action.Key)
			m.applyProgressIfSucceeded()
		case tuiadapter.ActionHint:
			if hint, ok := m.run.RequestHint(); ok {
				m.hintMessage = hint
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
	if summary := m.reviewQueueSummary(); summary != "" {
		b.WriteString(summary + "\n")
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
	b.WriteString("\n")
	if view.Mode == string(vimengine.ModeCommand) || view.Mode == string(vimengine.ModeSearch) {
		prompt := ":"
		if view.Mode == string(vimengine.ModeSearch) {
			prompt = "/"
		}
		b.WriteString(prompt + view.CommandLine + "\n\n")
		b.WriteString(m.renderActionPanel(state, view))
		return b.String()
	}
	if view.LastCommand != "" {
		b.WriteString(fmt.Sprintf("Command: %s\n", view.LastCommand))
	}
	b.WriteString(m.renderActionPanel(state, view))
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
	m.refreshReviewQueue()
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
	m.hintMessage = ""
	m.startedAt = m.now()
}

func (m *Model) retryCurrent() {
	m.run.Retry()
	m.saved = false
	m.hintMessage = ""
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

func (m *Model) refreshReviewQueue() {
	if m.progress == nil || m.contentRoot == "" || len(m.entries) == 0 {
		m.reviewQueue = nil
		return
	}
	library, err := content.LoadLibrary(m.contentRoot)
	if err != nil {
		m.reviewQueue = nil
		return
	}
	m.reviewQueue = review.Candidates(library, *m.progress, review.Options{
		OrderedExerciseIDs: m.exerciseOrder(),
		Limit:              3,
	})
}

func (m Model) exerciseOrder() []string {
	ids := make([]string, 0, len(m.entries))
	seen := make(map[string]bool)
	for _, entry := range m.entries {
		if !seen[entry.ExerciseID] {
			seen[entry.ExerciseID] = true
			ids = append(ids, entry.ExerciseID)
		}
	}
	return ids
}

func (m Model) reviewQueueSummary() string {
	if len(m.reviewQueue) == 0 {
		return ""
	}
	return "재진단 큐: " + m.reviewQueue[0].Summary()
}

func (m Model) residualRiskSummary() string {
	for _, candidate := range m.reviewQueue {
		if candidate.ExerciseID != m.currentExerciseID() {
			return "잔류 리스크: " + candidate.Summary()
		}
	}
	return ""
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

func (m Model) renderActionPanel(state scenario.State, view tuiadapter.ViewModel) string {
	lines := []string{"ACTION"}
	switch {
	case view.Mode == string(vimengine.ModeCommand):
		lines = append(lines, "Keys: type command  enter: run  esc: normal")
	case view.Mode == string(vimengine.ModeSearch):
		lines = append(lines, "Keys: type search  enter: find  esc: normal")
	case state.Status == exerciseruntime.StatusSucceeded:
		lines = append(lines, m.successDebriefLines(state)...)
		if m.current+1 < len(m.entries) {
			if m.nextEntryStartsNewPlaylist() {
				lines = append(lines, "Next tutorial: enter")
			} else {
				lines = append(lines, "Next: enter")
			}
		} else {
			lines = append(lines, "Playlist complete")
			lines = append(lines, "q: quit")
		}
	case state.Status == exerciseruntime.StatusFailed:
		if state.Runtime.MaxInputs > 0 {
			lines = append(lines, fmt.Sprintf("Inputs left: %d/%d", state.Runtime.InputsLeft, state.Runtime.MaxInputs))
		}
		lines = append(lines, fmt.Sprintf("Attempts: %d/%s", state.Runtime.Attempts, attemptLimitLabel(state.Runtime.AttemptLimit)))
		if coach := coachingLine(state); coach != "" {
			lines = append(lines, coach)
		}
		if m.hintMessage != "" {
			lines = append(lines, "Hint: "+m.hintMessage)
		}
		lines = append(lines, "Retry: r or enter")
		lines = append(lines, "?: hint  q: quit")
	default:
		if state.Runtime.MaxInputs > 0 {
			lines = append(lines, fmt.Sprintf("Inputs left: %d/%d", state.Runtime.InputsLeft, state.Runtime.MaxInputs))
		}
		if coach := coachingLine(state); coach != "" {
			lines = append(lines, coach)
		}
		if m.hintMessage != "" {
			lines = append(lines, "Hint: "+m.hintMessage)
		}
		lines = append(lines, "?: hint  q: quit")
	}
	return actionPanelStyle.Render(strings.Join(lines, "\n"))
}

func coachingLine(state scenario.State) string {
	missing := missingRequiredKeys(state.Runtime.RequiredKeys, state.Runtime.KeyTrace)
	if len(missing) == 0 {
		return ""
	}
	return "Coach: 훈련 키 " + strings.Join(missing, " ")
}

func (m Model) successDebriefLines(state scenario.State) []string {
	if state.Score == nil {
		return nil
	}

	lines := []string{
		fmt.Sprintf("Debrief: grade %s, %d keys", state.Score.Grade, len(state.Runtime.KeyTrace)),
	}
	if best, ok := m.progress.Missions[m.currentExerciseID()]; ok && best.Completed {
		bestGrade := best.BestGrade
		if bestGrade == "" {
			bestGrade = "-"
		}
		bestKeys := "-"
		if best.BestKeystrokes > 0 {
			bestKeys = fmt.Sprintf("%d keys", best.BestKeystrokes)
		}
		lines = append(lines, fmt.Sprintf("Best: grade %s, %s", bestGrade, bestKeys))
	}
	if entry, ok := m.currentEntry(); ok {
		completed, total := m.playlistCompletion(entry.PlaylistID)
		lines = append(lines, fmt.Sprintf("Playlist: %d/%d complete", completed, total))
	}
	if residual := m.residualRiskSummary(); residual != "" {
		lines = append(lines, residual)
	}
	return lines
}

func (m Model) playlistCompletion(playlistID string) (int, int) {
	completed := 0
	total := 0
	for _, entry := range m.entries {
		if entry.PlaylistID != playlistID {
			continue
		}
		total++
		if m.progress.Missions[entry.ExerciseID].Completed {
			completed++
		}
	}
	return completed, total
}

func attemptLimitLabel(limit int) string {
	if limit <= 0 {
		return "unlimited"
	}
	return fmt.Sprintf("%d", limit)
}

func missingRequiredKeys(required []string, trace []string) []string {
	var missing []string
	for _, key := range required {
		if !containsString(trace, key) {
			missing = append(missing, key)
		}
	}
	return missing
}

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

var _ tea.Model = Model{}
