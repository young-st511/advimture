package playable

import (
	"fmt"
	"io/fs"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/young-st511/advimture/internal/content"
	"github.com/young-st511/advimture/internal/e2estate"
	"github.com/young-st511/advimture/internal/playableview"
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
	ContentFS    fs.FS
	StartupError error
	Now          func() time.Time
}

type Model struct {
	run          *scenario.Run
	entries      []gameEntry
	current      int
	contentRoot  string
	contentFS    fs.FS
	progress     *progress.Progress
	saveProgress func(*progress.Progress) error
	e2eStatePath string
	now          func() time.Time
	startedAt    time.Time
	saved        bool
	saveErr      error
	err          error
	reviewQueue  []review.Candidate
	hintMessage  string
	width        int
	height       int
}

type gameEntry struct {
	PlaylistID       string
	PlaylistTitle    string
	PlaylistCategory string
	ExerciseID       string
	ScenarioID       string
	IndexInPlaylist  int
	TotalInPlaylist  int
	TrainedCommands  []string
	ReviewedCommands []string
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
	entries, current, run, err := newGameFromContent(root, options.ContentFS, p)
	if options.StartupError != nil {
		if err != nil {
			err = fmt.Errorf("%v; %w", options.StartupError, err)
		} else {
			err = options.StartupError
		}
	}
	model := Model{
		run:          run,
		entries:      entries,
		current:      current,
		contentRoot:  root,
		contentFS:    options.ContentFS,
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
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
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
				m.advanceAfterSuccess()
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
	screen := playableview.Screen{
		Width:           m.width,
		Height:          m.height,
		Title:           view.Title,
		Message:         view.Message,
		BufferLines:     view.BufferLines,
		Mode:            view.Mode,
		Status:          view.Status,
		CursorRow:       view.CursorRow,
		CursorCol:       view.CursorCol,
		Selection:       view.Selection,
		Grade:           view.Grade,
		CommandLine:     view.CommandLine,
		LastCommand:     view.LastCommand,
		FocusPanel:      m.focusPanel(state, view),
		ShowLastCommand: state.Status == exerciseruntime.StatusRunning,
	}
	if entry, ok := m.currentEntry(); ok {
		screen.PlaylistTitle = entry.PlaylistTitle
		screen.PlaylistCategory = entry.PlaylistCategory
		screen.ExerciseIndex = entry.IndexInPlaylist
		screen.ExerciseTotal = entry.TotalInPlaylist
	}
	screen.ReviewSummary = m.reviewQueueSummary()
	screen.DailyRoute = m.dailyRouteSummary()
	screen.ReviewCount = len(m.reviewQueue)
	if len(m.reviewQueue) > 0 {
		screen.ReviewPrimary = m.reviewQueue[0].Title
	}
	if view.Mode == string(vimengine.ModeCommand) || view.Mode == string(vimengine.ModeSearch) {
		prompt := ":"
		if view.Mode == string(vimengine.ModeSearch) {
			prompt = "/"
		}
		screen.CommandPrompt = prompt
		screen.ShowCommandLine = true
		screen.ShowLastCommand = false
	}
	return playableview.Render(screen)
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
	focusPanel := m.focusPanel(state, tuiadapter.RenderState(state))
	return e2estate.State{
		Buffer: append([]string(nil), state.Runtime.Vim.Lines...),
		Cursor: e2estate.Cursor{
			Row: state.Runtime.Vim.Cursor.Row,
			Col: state.Runtime.Vim.Cursor.Col,
		},
		Mode:      string(state.Runtime.Vim.Mode),
		Command:   state.Runtime.Vim.LastCommand,
		Status:    string(state.Status),
		Score:     score,
		Progress:  progressState,
		Review:    m.e2eReview(),
		UI:        e2eUI(focusPanel),
		Selection: e2eSelection(state.Runtime.Vim.Selection),
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
	if m.saveProgress != nil {
		if err := m.saveProgress(&updated); err != nil {
			m.saveErr = err
			return
		}
	}
	m.progress = &updated
	m.refreshReviewQueue()
	m.saved = true
	m.saveErr = nil
}

func (m *Model) advanceToNext() {
	if m.current+1 >= len(m.entries) {
		return
	}
	m.jumpToEntry(m.current + 1)
}

func (m *Model) advanceAfterSuccess() {
	if !m.saved {
		m.applyProgressIfSucceeded()
	}
	if !m.saved {
		return
	}
	if m.current+1 < len(m.entries) {
		m.advanceToNext()
		return
	}
	if index, ok := m.reviewDispatchTargetIndex(); ok {
		m.jumpToEntry(index)
	}
}

func (m *Model) jumpToEntry(index int) {
	if index < 0 || index >= len(m.entries) {
		return
	}
	run, err := runForEntry(m.contentRoot, m.contentFS, m.entries[index])
	if err != nil {
		m.err = err
		return
	}
	m.current = index
	m.run = run
	m.saved = false
	m.saveErr = nil
	m.hintMessage = ""
	m.startedAt = m.now()
}

func (m *Model) retryCurrent() {
	m.run.Retry()
	m.saved = false
	m.saveErr = nil
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
	library, err := loadLibrary(m.contentRoot, m.contentFS)
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
	return "재점검 대상: " + m.reviewQueue[0].Summary()
}

func (m Model) dailyRouteSummary() string {
	if len(m.reviewQueue) == 0 {
		return ""
	}
	primary := m.reviewQueue[0].DailyRouteLabel()
	if len(m.reviewQueue) == 1 {
		return "오늘의 복구 루트: " + primary
	}
	return fmt.Sprintf("오늘의 복구 루트: %s 외 %d건 대기", primary, len(m.reviewQueue)-1)
}

func (m Model) nextDispatchSummary() string {
	if len(m.reviewQueue) == 0 {
		return ""
	}
	primary := m.reviewQueue[0].DailyRouteLabel()
	label := "다음 출격 후보"
	if m.currentEntryIsTutorial() {
		label = "나중에 다시 풀기"
	}
	if len(m.reviewQueue) == 1 {
		return label + ": " + primary
	}
	return fmt.Sprintf("%s: %s 외 %d건 대기", label, primary, len(m.reviewQueue)-1)
}

func (m Model) residualRiskSummary() string {
	label := "잔류 리스크"
	if m.currentEntryIsTutorial() {
		label = "재점검 메모"
	}
	for _, candidate := range m.reviewQueue {
		if candidate.ExerciseID != m.currentExerciseID() {
			return label + ": " + candidate.Summary()
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

func (m Model) reviewDispatchTargetIndex() (int, bool) {
	if len(m.reviewQueue) == 0 {
		return 0, false
	}
	target := m.reviewQueue[0].ExerciseID
	for index, entry := range m.entries {
		if entry.ExerciseID == target {
			return index, true
		}
	}
	return 0, false
}

func (m Model) successActionLines() []string {
	return actionLabels(m.successActions())
}

func (m Model) successActions() []playableview.ActionLine {
	if m.saveErr != nil {
		return []playableview.ActionLine{
			focusAction("save_retry", "저장 재시도: enter"),
			focusAction("quit", "종료: q"),
		}
	}
	if m.current+1 < len(m.entries) {
		if !m.nextEntryStartsNewPlaylist() {
			return []playableview.ActionLine{focusAction("next", "다음 단계: enter")}
		}
		switch m.entries[m.current+1].PlaylistCategory {
		case "tutorial":
			return []playableview.ActionLine{focusAction("next_tutorial", "다음 튜토리얼: enter")}
		case "incident":
			return []playableview.ActionLine{focusAction("next_runbook", "다음 런북: enter")}
		default:
			return []playableview.ActionLine{focusAction("next_playlist", "다음 플레이리스트: enter")}
		}
	}
	if _, ok := m.reviewDispatchTargetIndex(); ok {
		return []playableview.ActionLine{
			focusAction("next_dispatch", "다음 출격: enter"),
			focusAction("quit", "종료: q"),
		}
	}
	if m.currentEntryIsIncident() {
		return []playableview.ActionLine{
			focusAction("dispatch_complete", "출격 완료"),
			focusAction("quit", "종료: q"),
		}
	}
	return []playableview.ActionLine{
		focusAction("playlist_complete", "플레이리스트 완료"),
		focusAction("quit", "종료: q"),
	}
}

func failureActions() []playableview.ActionLine {
	return []playableview.ActionLine{
		focusAction("retry", "다시 시도: r 또는 enter"),
		focusAction("hint", "힌트: ?"),
		focusAction("quit", "종료: q"),
	}
}

func focusAction(id, label string) playableview.ActionLine {
	return playableview.ActionLine{ID: id, Label: label}
}

func actionLabels(actions []playableview.ActionLine) []string {
	labels := make([]string, 0, len(actions))
	for _, action := range actions {
		if action.Label != "" {
			labels = append(labels, action.Label)
		}
	}
	return labels
}

func newGameFromContent(root string, files fs.FS, progressState *progress.Progress) ([]gameEntry, int, *scenario.Run, error) {
	library, err := loadLibrary(root, files)
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
	run, err := runForEntry(root, files, entries[current])
	if err != nil {
		return nil, 0, nil, err
	}
	return entries, current, run, nil
}

func runForEntry(root string, files fs.FS, entry gameEntry) (*scenario.Run, error) {
	library, err := loadLibrary(root, files)
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

func loadLibrary(root string, files fs.FS) (content.Library, error) {
	if files != nil {
		return content.LoadLibraryFS(files, root)
	}
	return content.LoadLibrary(root)
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
			if !ok {
				return nil, fmt.Errorf("playlist %q beat %q references missing exercise %q", playlist.ID, beat.ID, beat.ExerciseID)
			}
			if !isPlayableExercise(exercise) {
				return nil, fmt.Errorf("playlist %q beat %q references non-playable exercise %q", playlist.ID, beat.ID, beat.ExerciseID)
			}
			scenarioDoc, ok := library.Scenarios[beat.ScenarioID]
			if !ok {
				return nil, fmt.Errorf("playlist %q beat %q references missing scenario %q", playlist.ID, beat.ID, beat.ScenarioID)
			}
			if !isPlayableScenario(scenarioDoc) {
				return nil, fmt.Errorf("playlist %q beat %q references non-playable scenario %q", playlist.ID, beat.ID, beat.ScenarioID)
			}
			playlistEntries = append(playlistEntries, gameEntry{
				PlaylistID:       playlist.ID,
				PlaylistTitle:    playlist.Title,
				PlaylistCategory: playlist.Category,
				ExerciseID:       beat.ExerciseID,
				ScenarioID:       beat.ScenarioID,
				TrainedCommands:  append([]string(nil), exercise.TrainedCommands...),
				ReviewedCommands: append([]string(nil), exercise.ReviewedCommands...),
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

func (m Model) focusPanel(state scenario.State, view tuiadapter.ViewModel) *playableview.FocusPanel {
	kind, title := focusPanelIdentity(state, view, m.currentEntryIsIncident())
	return &playableview.FocusPanel{
		Kind:    kind,
		Title:   title,
		Lines:   m.focusPanelLines(state, view),
		Actions: m.focusPanelActions(state, view),
	}
}

func focusPanelIdentity(state scenario.State, view tuiadapter.ViewModel, incident bool) (string, string) {
	switch {
	case view.Mode == string(vimengine.ModeVisual):
		return "mode", "선택 모드"
	case view.Mode == string(vimengine.ModeInsert):
		return "mode", "입력 모드"
	case view.Mode == string(vimengine.ModeCommand):
		return "mode", "명령 모드"
	case view.Mode == string(vimengine.ModeSearch):
		return "mode", "검색 모드"
	case state.Status == exerciseruntime.StatusSucceeded:
		return "success", "STEP SEALED"
	case state.Status == exerciseruntime.StatusFailed:
		return "failure", "RECOVERY REQUIRED"
	case incident:
		return "incident", "OPERATOR JUDGMENT"
	default:
		return "training", "TRAINING BRIEF"
	}
}

func (m Model) focusPanelLines(state scenario.State, view tuiadapter.ViewModel) []string {
	lines := []string{}
	switch {
	case view.Mode == string(vimengine.ModeVisual):
		lines = append(lines, "선택: 이동 키로 범위 조정  esc/v: normal")
	case view.Mode == string(vimengine.ModeInsert):
		lines = append(lines, "입력: 텍스트 작성  esc: normal")
	case view.Mode == string(vimengine.ModeCommand):
		lines = append(lines, "명령: 입력 후 enter 실행  esc: normal")
	case view.Mode == string(vimengine.ModeSearch):
		lines = append(lines, "검색: 입력 후 enter 찾기  esc: normal")
	case state.Status == exerciseruntime.StatusSucceeded:
		if feedback := scenarioFeedbackLine(state); feedback != "" {
			lines = append(lines, feedback)
		}
		if m.saveErr != nil {
			lines = append(lines, fmt.Sprintf("진행도 저장 실패: %v", m.saveErr))
			lines = append(lines, "저장 후 계속하려면 enter")
		}
		lines = append(lines, m.successDebriefLines(state)...)
	case state.Status == exerciseruntime.StatusFailed:
		if feedback := scenarioFeedbackLine(state); feedback != "" {
			lines = append(lines, feedback)
		}
		if state.Runtime.MaxInputs > 0 {
			lines = append(lines, fmt.Sprintf("Inputs left: %d/%d", state.Runtime.InputsLeft, state.Runtime.MaxInputs))
		}
		lines = append(lines, fmt.Sprintf("Attempts: %d/%s", state.Runtime.Attempts, attemptLimitLabel(state.Runtime.AttemptLimit)))
		coach := m.coachingLineForCurrentEntry(state)
		if coach != "" {
			lines = append(lines, coach)
		}
		if memory := m.commandMemoryLine(state); memory != "" && (m.currentEntryIsIncident() || coach == "" || !strings.HasPrefix(coach, "Coach: 훈련 키 ")) {
			lines = append(lines, memory)
		}
		if m.hintMessage != "" {
			lines = append(lines, revealedHintLine(m.hintMessage))
		}
	default:
		if state.Runtime.MaxInputs > 0 {
			lines = append(lines, fmt.Sprintf("Inputs left: %d/%d", state.Runtime.InputsLeft, state.Runtime.MaxInputs))
		}
		if memory := m.commandMemoryLine(state); memory != "" && !m.currentEntryIsIncident() {
			lines = append(lines, memory)
		}
		if coach := m.runningCoachingLine(state); coach != "" {
			lines = append(lines, coach)
		}
		if memory := m.commandMemoryLine(state); memory != "" && m.currentEntryIsIncident() {
			lines = append(lines, memory)
		}
		if m.hintMessage != "" {
			lines = append(lines, revealedHintLine(m.hintMessage))
		}
	}
	return lines
}

func revealedHintLine(message string) string {
	return "힌트 내용  " + message + " · 등급에 영향"
}

func (m Model) focusPanelActions(state scenario.State, view tuiadapter.ViewModel) []playableview.ActionLine {
	switch {
	case view.Mode == string(vimengine.ModeVisual),
		view.Mode == string(vimengine.ModeInsert),
		view.Mode == string(vimengine.ModeCommand),
		view.Mode == string(vimengine.ModeSearch):
		return nil
	case state.Status == exerciseruntime.StatusSucceeded:
		return m.successActions()
	case state.Status == exerciseruntime.StatusFailed:
		return failureActions()
	default:
		return runningActions()
	}
}

func runningActions() []playableview.ActionLine {
	return []playableview.ActionLine{
		focusAction("hint", "힌트: ?"),
		focusAction("quit", "종료: q"),
	}
}

func scenarioFeedbackLine(state scenario.State) string {
	message := strings.TrimSpace(state.Message)
	briefing := strings.TrimSpace(state.Briefing)
	if message == "" || message == briefing {
		return ""
	}
	return message
}

func (m Model) runningCoachingLine(state scenario.State) string {
	if m.currentEntryIsIncident() {
		if m.hintMessage != "" {
			return ""
		}
		return "판단: 목표 상태를 보고 이미 배운 Vim 동작을 선택하세요."
	}
	coach := coachingLine(state)
	if coach != "" && strings.HasPrefix(coach, "Coach: 훈련 키 ") && m.commandMemoryLine(state) != "" {
		return ""
	}
	return coach
}

func (m Model) coachingLineForCurrentEntry(state scenario.State) string {
	coach := coachingLine(state)
	if coach == "" || !m.currentEntryIsIncident() {
		return coach
	}
	return strings.Replace(coach, "Coach: 훈련 키", "복구 힌트: 필요한 키", 1)
}

func (m Model) commandMemoryLine(state scenario.State) string {
	entry, ok := m.currentEntry()
	if !ok {
		return ""
	}
	if m.currentEntryIsIncident() {
		if state.Status != exerciseruntime.StatusFailed && m.hintMessage == "" {
			return ""
		}
		commands := entry.TrainedCommands
		if len(commands) == 0 {
			commands = entry.ReviewedCommands
		}
		if memory := commandMemory(commands); memory != "" {
			return "참고 명령: " + memory
		}
		return ""
	}
	commands := entry.TrainedCommands
	if len(commands) == 0 {
		commands = state.Runtime.RequiredKeys
	}
	if memory := commandMemory(commands); memory != "" {
		return "기억할 명령: " + memory
	}
	return ""
}

func (m Model) currentEntryIsIncident() bool {
	entry, ok := m.currentEntry()
	if !ok {
		return false
	}
	return entry.PlaylistCategory == "incident"
}

func (m Model) currentEntryIsTutorial() bool {
	entry, ok := m.currentEntry()
	if !ok {
		return false
	}
	return entry.PlaylistCategory == "tutorial"
}

func e2eSelection(selection *vimengine.Selection) *e2estate.Selection {
	if selection == nil {
		return nil
	}
	return &e2estate.Selection{
		Active: selection.Active,
		Kind:   string(selection.Kind),
		Anchor: e2eCursor(selection.Anchor),
		Head:   e2eCursor(selection.Head),
		Start:  e2eCursor(selection.Start),
		End:    e2eCursor(selection.End),
	}
}

func e2eUI(panel *playableview.FocusPanel) e2estate.UI {
	if panel == nil {
		return e2estate.UI{}
	}
	return e2estate.UI{
		FocusPanel: e2estate.FocusPanel{
			Kind:    panel.Kind,
			Title:   panel.Title,
			Lines:   append([]string(nil), panel.Lines...),
			Actions: e2eActions(panel.Actions),
		},
	}
}

func e2eActions(actions []playableview.ActionLine) []e2estate.ActionLine {
	if len(actions) == 0 {
		return nil
	}
	out := make([]e2estate.ActionLine, 0, len(actions))
	for _, action := range actions {
		out = append(out, e2estate.ActionLine{
			ID:    action.ID,
			Label: action.Label,
		})
	}
	return out
}

func (m Model) e2eReview() e2estate.Review {
	state := e2estate.Review{
		QueueCount: len(m.reviewQueue),
		DailyRoute: m.dailyRouteSummary(),
	}
	if len(m.reviewQueue) == 0 {
		return state
	}
	state.PrimaryExerciseID = m.reviewQueue[0].ExerciseID
	state.PrimaryReason = string(m.reviewQueue[0].Reason)
	return state
}

func e2eCursor(cursor vimengine.Cursor) e2estate.Cursor {
	return e2estate.Cursor{
		Row: cursor.Row,
		Col: cursor.Col,
	}
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
		fmt.Sprintf("이번 복구: grade %s, %d keys", state.Score.Grade, len(state.Runtime.KeyTrace)),
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
		lines = append(lines, fmt.Sprintf("최단 복구: grade %s, %s", bestGrade, bestKeys))
	}
	if state.Score.ExpectedKeyCount > 0 {
		lines = append(lines, fmt.Sprintf("목표 입력: %d keys", state.Score.ExpectedKeyCount))
	}
	if entry, ok := m.currentEntry(); ok {
		completed, total := m.playlistCompletion(entry.PlaylistID)
		lines = append(lines, fmt.Sprintf("Runbook: %d/%d 복구 완료", completed, total))
	}
	if residual := m.residualRiskSummary(); residual != "" {
		lines = append(lines, residual)
	}
	if next := m.nextDispatchSummary(); next != "" {
		lines = append(lines, next)
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

func commandMemory(commands []string) string {
	if len(commands) == 0 {
		return ""
	}
	seen := map[string]bool{}
	out := make([]string, 0, len(commands))
	for _, command := range commands {
		command = strings.TrimSpace(command)
		if command == "" || seen[command] {
			continue
		}
		seen[command] = true
		out = append(out, command)
		if len(out) == 6 {
			break
		}
	}
	return strings.Join(out, " ")
}

var _ tea.Model = Model{}
