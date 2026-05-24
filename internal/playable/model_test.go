package playable

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/young-st511/advimture/internal/content"
	"github.com/young-st511/advimture/internal/progress"
	"github.com/young-st511/advimture/internal/scoring"
)

func TestPlayableStartsWithBriefing(t *testing.T) {
	model := New(Options{ContentRoot: contentRootForTest()})

	if model.State().Status != "running" {
		t.Fatalf("status = %q, want running", model.State().Status)
	}
	if !strings.Contains(model.View(), "터미널 지도에서 목표 문자까지 커서를 이동하세요.") {
		t.Fatalf("view = %q, want briefing", model.View())
	}
	if !strings.Contains(model.View(), "Tutorial 0: 커서 감각 회상") {
		t.Fatalf("view = %q, want tutorial title", model.View())
	}
	if !strings.Contains(model.View(), "Exercise: 1/4") {
		t.Fatalf("view = %q, want episode-local count", model.View())
	}
	if !strings.Contains(model.View(), "재점검 대상: 목표 문자까지 이동하기: 미복구") {
		t.Fatalf("view = %q, want review queue", model.View())
	}
	if !strings.Contains(model.View(), "오늘의 복구 루트: 목표 문자까지 이동하기(미복구) 외 2건 대기") {
		t.Fatalf("view = %q, want daily route", model.View())
	}
	if model.State().Review.QueueCount != 3 || model.State().Review.PrimaryExerciseID != "normal-motion-basic-001" {
		t.Fatalf("review state = %+v, want first review candidate", model.State().Review)
	}
	if !strings.Contains(model.View(), "Coach: 훈련 키 l") {
		t.Fatalf("view = %q, want proactive coaching", model.View())
	}
	if !strings.Contains(model.View(), "TRAINING BRIEF") {
		t.Fatalf("view = %q, want training focus panel", model.View())
	}
	if strings.Contains(model.View(), "ACTION") {
		t.Fatalf("view = %q, should not show debug action label", model.View())
	}
	if model.State().UI.FocusPanel.Kind != "training" || model.State().UI.FocusPanel.Title != "TRAINING BRIEF" {
		t.Fatalf("ui focus panel = %+v, want training TRAINING BRIEF", model.State().UI.FocusPanel)
	}
}

func TestPlayablePassesWindowSizeToRenderer(t *testing.T) {
	model := New(Options{ContentRoot: contentRootForTest()})

	model, _ = updateWithWindowSize(t, model, 100, 30)

	view := model.View()
	if !strings.Contains(view, "MISSION") || !strings.Contains(view, "RUNBOOK CONSOLE") {
		t.Fatalf("view = %q, want mission HUD and console after window size", view)
	}
	if strings.Contains(view, "\n복구 현황\n") {
		t.Fatalf("view = %q, should fold recovery status into mission HUD", view)
	}
	if !strings.Contains(view, "복구 현황: 재점검 대상:") {
		t.Fatalf("view = %q, want recovery status in mission HUD", view)
	}
}

func TestPlayableShowsFailureFeedbackInFocusPanel(t *testing.T) {
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progressCompleteBefore(t, "word-motion-basic-001"),
	})

	model, _ = updateWithKey(t, model, "l")

	if model.State().Status != "failed" {
		t.Fatalf("status = %q, want failed", model.State().Status)
	}
	if !strings.Contains(model.View(), "라우팅 설정 한 줄에서 backend로 바로 이동해야 합니다") {
		t.Fatalf("view = %q, want original briefing", model.View())
	}
	lines := model.State().UI.FocusPanel.Lines
	if !containsLineWith(lines, "한 글자씩 가면 늦습니다.") {
		t.Fatalf("focus panel lines = %v, want failure feedback", lines)
	}
}

func TestPlayableRejectsWriteQuitForDiscardMission(t *testing.T) {
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progressCompleteBefore(t, "survival-save-quit-002"),
	})

	for _, key := range []string{":", "w", "q", "enter"} {
		model, _ = updateWithKey(t, model, key)
	}

	if model.State().Status != "failed" {
		t.Fatalf("status = %q, want failed", model.State().Status)
	}
	if !containsLineWith(model.State().UI.FocusPanel.Lines, ":wq가 아니라 :q!") {
		t.Fatalf("focus panel lines = %v, want :wq rejection", model.State().UI.FocusPanel.Lines)
	}
	if !strings.Contains(model.View(), "임시 설정을 잘못 열었습니다.") {
		t.Fatalf("view = %q, want original briefing after failed command", model.View())
	}
}

func TestPlayableSucceedsAndUpdatesProgress(t *testing.T) {
	saveCalls := 0
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progress.NewProgress(),
		Now: func() time.Time {
			return time.Unix(10, 0)
		},
		SaveProgress: func(*progress.Progress) error {
			saveCalls++
			return nil
		},
	})

	model, _ = updateWithKey(t, model, "l")
	model, _ = updateWithKey(t, model, "l")

	state := model.State()
	if state.Status != "succeeded" {
		t.Fatalf("status = %q, want succeeded", state.Status)
	}
	if state.Score.Grade != string(scoring.GradeS) {
		t.Fatalf("grade = %q, want S", state.Score.Grade)
	}
	if state.Cursor.Col != 2 {
		t.Fatalf("cursor col = %d, want 2", state.Cursor.Col)
	}
	if !state.Progress.Completed {
		t.Fatal("progress completed = false, want true")
	}
	if saveCalls != 1 {
		t.Fatalf("saveCalls = %d, want 1", saveCalls)
	}
}

func containsLineWith(lines []string, want string) bool {
	for _, line := range lines {
		if strings.Contains(line, want) {
			return true
		}
	}
	return false
}

func TestPlayableShowsSuccessDebriefAndBestRecord(t *testing.T) {
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progress.NewProgress(),
		SaveProgress: func(*progress.Progress) error {
			return nil
		},
	})

	model, _ = updateWithKey(t, model, "l")
	model, _ = updateWithKey(t, model, "l")

	view := model.View()
	if !strings.Contains(view, "복구 기록: grade S, 2 keys") {
		t.Fatalf("view = %q, want success debrief", view)
	}
	if !strings.Contains(view, "최단 복구 기록: grade S, 2 keys") {
		t.Fatalf("view = %q, want best record", view)
	}
	if !strings.Contains(view, "Runbook: 1/4 복구 완료") {
		t.Fatalf("view = %q, want playlist completion count", view)
	}
	if !strings.Contains(view, "잔류 리스크: 경고 지점으로 이동하기: 미복구") {
		t.Fatalf("view = %q, want residual risk recommendation", view)
	}
	if !strings.Contains(view, "오늘의 복구 루트: 경고 지점으로 이동하기(미복구) 외 2건 대기") {
		t.Fatalf("view = %q, want daily route after success", view)
	}
}

func TestPlayableShowsReviewQueueForLowGrade(t *testing.T) {
	p := progressWithAllPlayableCompleted(t)
	p.Missions["normal-motion-basic-001"] = progress.MissionProgress{Completed: true, BestGrade: "B", BestKeystrokes: 2, Attempts: 1}

	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    p,
	})

	if !strings.Contains(model.View(), "재점검 대상: 목표 문자까지 이동하기: 복구 등급 B") {
		t.Fatalf("view = %q, want low grade review queue", model.View())
	}
}

func TestIncidentActionPanelUsesJudgementCueWithoutTrainingKeySpoiler(t *testing.T) {
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progressCompleteBefore(t, "incident-hotfix-001"),
	})

	view := model.View()
	if !strings.Contains(view, "판단: 목표 상태를 보고 이미 배운 Vim 동작을 선택하세요.") {
		t.Fatalf("view = %q, want incident judgement cue", view)
	}
	if strings.Contains(view, "Coach: 훈련 키") {
		t.Fatalf("view = %q, should not reveal training keys on running incident", view)
	}
}

func TestIncidentFailureCanRevealRecoveryKeyHint(t *testing.T) {
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progressCompleteBefore(t, "incident-hotfix-001"),
	})

	for i := 0; i < 10; i++ {
		model, _ = updateWithKey(t, model, "x")
	}

	view := model.View()
	if !strings.Contains(view, "복구 힌트: 필요한 키") {
		t.Fatalf("view = %q, want recovery key hint after incident failure", view)
	}
	if strings.Contains(view, "Coach: 훈련 키") {
		t.Fatalf("view = %q, should not use tutorial coaching language in incident failure", view)
	}
}

func TestPlayableAdvancesToNextExerciseAfterSuccess(t *testing.T) {
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progress.NewProgress(),
	})

	model, _ = updateWithKey(t, model, "l")
	model, _ = updateWithKey(t, model, "l")
	if model.State().Status != "succeeded" {
		t.Fatalf("status = %q, want succeeded", model.State().Status)
	}

	model, _ = updateWithSpecialKey(t, model, tea.KeyEnter)

	if model.State().Status != "running" {
		t.Fatalf("status after next = %q, want running", model.State().Status)
	}
	if !strings.Contains(model.View(), "부팅 로그에서 WARN 줄을 놓쳤습니다") {
		t.Fatalf("view = %q, want second exercise briefing", model.View())
	}
	if !strings.Contains(model.View(), "Exercise: 2/4") {
		t.Fatalf("view = %q, want second exercise in first tutorial", model.View())
	}
}

func TestPlayableStartsAtFirstIncompleteExercise(t *testing.T) {
	p := progress.NewProgress()
	p.CompleteMission("normal-motion-basic-001", "S", 2, 1000)

	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    p,
	})

	if !strings.Contains(model.View(), "부팅 로그에서 WARN 줄을 놓쳤습니다") {
		t.Fatalf("view = %q, want first incomplete exercise", model.View())
	}
	if !strings.Contains(model.View(), "Exercise: 2/4") {
		t.Fatalf("view = %q, want episode-local count", model.View())
	}
}

func TestPlayableShowsNextTutorialAtEpisodeBoundary(t *testing.T) {
	p := progress.NewProgress()
	p.CompleteMission("normal-motion-basic-001", "S", 2, 1000)
	p.CompleteMission("normal-motion-basic-002", "S", 1, 1000)
	p.CompleteMission("normal-motion-basic-003", "S", 2, 1000)

	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    p,
	})

	if !strings.Contains(model.View(), "로그를 한 줄 더 내려가 버렸습니다") {
		t.Fatalf("view = %q, want last movement exercise", model.View())
	}
	model, _ = updateWithKey(t, model, "k")

	if !strings.Contains(model.View(), "Next tutorial: enter") {
		t.Fatalf("view = %q, want next tutorial transition", model.View())
	}
	if !strings.Contains(model.View(), "Runbook: 4/4 복구 완료") {
		t.Fatalf("view = %q, want completed playlist debrief", model.View())
	}
	if !strings.Contains(model.View(), "STEP SEALED") {
		t.Fatalf("view = %q, want success focus panel", model.View())
	}

	model, _ = updateWithSpecialKey(t, model, tea.KeyEnter)
	if !strings.Contains(model.View(), "Tutorial 1: Vim 생존 키트") {
		t.Fatalf("view = %q, want survival tutorial title", model.View())
	}
	if !strings.Contains(model.View(), "Exercise: 1/3") {
		t.Fatalf("view = %q, want first survival exercise count", model.View())
	}
	if !strings.Contains(model.View(), "입력 모드에 커서가 묶였습니다") {
		t.Fatalf("view = %q, want first survival exercise", model.View())
	}
}

func TestPlayableAutosavesEachCompletedExercise(t *testing.T) {
	var saved *progress.Progress
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progress.NewProgress(),
		SaveProgress: func(p *progress.Progress) error {
			copy := *p
			saved = &copy
			return nil
		},
	})

	model, _ = updateWithKey(t, model, "l")
	model, _ = updateWithKey(t, model, "l")
	if saved == nil || !saved.Missions["normal-motion-basic-001"].Completed {
		t.Fatalf("saved progress = %+v, want normal-motion-basic-001 completed", saved)
	}

	model, _ = updateWithSpecialKey(t, model, tea.KeyEnter)
	model, _ = updateWithKey(t, model, "j")
	if saved == nil ||
		!saved.Missions["normal-motion-basic-001"].Completed ||
		!saved.Missions["normal-motion-basic-002"].Completed {
		t.Fatalf("saved progress = %+v, want both first exercises completed", saved)
	}
}

func TestPlayableFailsForbiddenInputWithoutSavingAndRetriesWithEnter(t *testing.T) {
	saveCalls := 0
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progress.NewProgress(),
		SaveProgress: func(*progress.Progress) error {
			saveCalls++
			return nil
		},
	})

	model, _ = updateWithKey(t, model, "w")

	if model.State().Status != "failed" {
		t.Fatalf("status = %q, want failed", model.State().Status)
	}
	if saveCalls != 0 {
		t.Fatalf("saveCalls = %d, want 0", saveCalls)
	}
	if !strings.Contains(model.View(), "Retry: r or enter") {
		t.Fatalf("view = %q, want retry prompt", model.View())
	}
	if !strings.Contains(model.View(), "RECOVERY REQUIRED") {
		t.Fatalf("view = %q, want failure focus panel", model.View())
	}
	if !strings.Contains(model.View(), "Attempts: 1/unlimited") {
		t.Fatalf("view = %q, want attempt count", model.View())
	}
	if !strings.Contains(model.View(), "Inputs left: 1/2") {
		t.Fatalf("view = %q, want remaining inputs", model.View())
	}

	model, _ = updateWithSpecialKey(t, model, tea.KeyEnter)

	if model.State().Status != "running" {
		t.Fatalf("status after retry = %q, want running", model.State().Status)
	}
	if !strings.Contains(model.View(), "Exercise: 1/4") {
		t.Fatalf("view = %q, want same exercise after retry", model.View())
	}
}

func TestPlayableFailsArrowKeyShortcutWithoutSaving(t *testing.T) {
	saveCalls := 0
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progress.NewProgress(),
		SaveProgress: func(*progress.Progress) error {
			saveCalls++
			return nil
		},
	})

	model, _ = updateWithSpecialKey(t, model, tea.KeyRight)

	if model.State().Status != "failed" {
		t.Fatalf("status = %q, want failed", model.State().Status)
	}
	if saveCalls != 0 {
		t.Fatalf("saveCalls = %d, want 0", saveCalls)
	}
	if !containsLineWith(model.State().UI.FocusPanel.Lines, "이 입력은 이번 문항에서 사용할 수 없습니다.") {
		t.Fatalf("focus panel lines = %v, want forbidden input message", model.State().UI.FocusPanel.Lines)
	}
	if !strings.Contains(model.View(), "Retry: r or enter") {
		t.Fatalf("view = %q, want retry prompt", model.View())
	}
}

func TestPlayableDoesNotQuitOnCtrlC(t *testing.T) {
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progress.NewProgress(),
	})

	model, cmd := updateWithSpecialKey(t, model, tea.KeyCtrlC)

	if cmd != nil {
		t.Fatal("cmd != nil, want no quit command")
	}
	if model.State().Status != "running" {
		t.Fatalf("status = %q, want running", model.State().Status)
	}
	if !strings.Contains(model.View(), "커서 위치 맞추기") {
		t.Fatalf("view = %q, want still on first exercise", model.View())
	}
}

func TestPlayableRetriesFailedExerciseWithR(t *testing.T) {
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progress.NewProgress(),
	})

	model, _ = updateWithKey(t, model, "w")
	if model.State().Status != "failed" {
		t.Fatalf("status = %q, want failed", model.State().Status)
	}

	model, _ = updateWithKey(t, model, "r")

	if model.State().Status != "running" {
		t.Fatalf("status after retry = %q, want running", model.State().Status)
	}
	if !strings.Contains(model.View(), "Exercise: 1/4") {
		t.Fatalf("view = %q, want same exercise after retry", model.View())
	}
}

func TestPlayableFailsShortcutThatSkipsRequiredInput(t *testing.T) {
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progress.NewProgress(),
	})

	model, _ = updateWithKey(t, model, "$")
	model, _ = updateWithKey(t, model, "l")

	if model.State().Status != "failed" {
		t.Fatalf("status = %q, want failed", model.State().Status)
	}
	if !containsLineWith(model.State().UI.FocusPanel.Lines, "의도한 입력을 사용하지 않았습니다") {
		t.Fatalf("focus panel lines = %v, want required input coaching", model.State().UI.FocusPanel.Lines)
	}
}

func TestPlayableShowsRequestedHintInActionPanel(t *testing.T) {
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progress.NewProgress(),
	})

	model, _ = updateWithKey(t, model, "l")
	model, _ = updateWithKey(t, model, "?")

	if model.State().Status != "running" {
		t.Fatalf("status = %q, want running", model.State().Status)
	}
	if !strings.Contains(model.View(), "Hint: 오른쪽으로 한 칸 더 이동해야 합니다.") {
		t.Fatalf("view = %q, want requested hint", model.View())
	}
	if model.State().Progress.Completed {
		t.Fatal("progress completed = true, want false")
	}
}

func TestPlayableWritesE2EState(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".advimture", "e2e_state.json")
	model := New(Options{
		Progress:     progress.NewProgress(),
		E2EStatePath: path,
		ContentRoot:  contentRootForTest(),
		SaveProgress: func(*progress.Progress) error {
			return nil
		},
	})

	model, _ = updateWithKey(t, model, "l")
	model, _ = updateWithKey(t, model, "l")

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), `"status": "succeeded"`) {
		t.Fatalf("state summary = %s", raw)
	}
}

func TestPlayableReportsContentLoadError(t *testing.T) {
	model := New(Options{ContentRoot: filepath.Join(t.TempDir(), "missing")})

	if !strings.Contains(model.View(), "Playable error:") {
		t.Fatalf("view = %q, want content load error", model.View())
	}
}

func TestPlayableShowsCommandLineInsteadOfQuitHintInCommandMode(t *testing.T) {
	model := New(Options{ContentRoot: contentRootForTest()})

	model, _ = updateWithKey(t, model, ":")

	view := model.View()
	if !strings.Contains(view, ":") {
		t.Fatalf("view = %q, want command prompt", view)
	}
	if !strings.Contains(view, "COMMAND CHANNEL") {
		t.Fatalf("view = %q, want command focus panel", view)
	}
	if strings.Contains(view, "q: quit") {
		t.Fatalf("view = %q, should not show q quit hint in command mode", view)
	}
	if strings.Contains(view, "ctrl+c: quit") {
		t.Fatalf("view = %q, should not show ctrl+c quit hint in command mode", view)
	}
}

func TestPlayableShowsTextEntryHelpInsteadOfHintInInsertMode(t *testing.T) {
	model := New(Options{
		ContentRoot: contentRootForTest(),
		Progress:    progressCompleteBefore(t, "insert-mode-entry-001"),
	})

	model, _ = updateWithKey(t, model, "i")

	view := model.View()
	if model.State().Mode != "insert" {
		t.Fatalf("mode = %q, want insert", model.State().Mode)
	}
	if !strings.Contains(view, "Keys: type text  esc: normal") {
		t.Fatalf("view = %q, want insert action panel", view)
	}
	if strings.Contains(view, "?: hint") {
		t.Fatalf("view = %q, should not show hint prompt in insert mode", view)
	}
	if strings.Contains(view, "q: quit") {
		t.Fatalf("view = %q, should not show quit prompt in insert mode", view)
	}
}

func TestPlayableShowsSearchLineInSearchMode(t *testing.T) {
	model := New(Options{ContentRoot: contentRootForTest()})

	model, _ = updateWithKey(t, model, "/")
	model, _ = updateWithKey(t, model, "a")

	view := model.View()
	if !strings.Contains(view, "/a") {
		t.Fatalf("view = %q, want search prompt", view)
	}
	if !strings.Contains(view, "Keys: type search  enter: find  esc: normal") {
		t.Fatalf("view = %q, want search action panel", view)
	}
	if strings.Contains(view, "?: hint") {
		t.Fatalf("view = %q, should not show hint prompt in search mode", view)
	}
}

func TestPlayableCanQuitFromContentLoadError(t *testing.T) {
	model := New(Options{ContentRoot: filepath.Join(t.TempDir(), "missing")})

	_, cmd := updateWithKey(t, model, "q")

	if cmd == nil {
		t.Fatal("cmd = nil, want quit command")
	}
}

func updateWithKey(t *testing.T, model Model, key string) (Model, tea.Cmd) {
	t.Helper()

	updated, cmd := model.Update(tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune(key),
	})
	next, ok := updated.(Model)
	if !ok {
		t.Fatalf("updated model type = %T, want playable.Model", updated)
	}
	return next, cmd
}

func updateWithSpecialKey(t *testing.T, model Model, key tea.KeyType) (Model, tea.Cmd) {
	t.Helper()

	updated, cmd := model.Update(tea.KeyMsg{Type: key})
	next, ok := updated.(Model)
	if !ok {
		t.Fatalf("updated model type = %T, want playable.Model", updated)
	}
	return next, cmd
}

func updateWithWindowSize(t *testing.T, model Model, width int, height int) (Model, tea.Cmd) {
	t.Helper()

	updated, cmd := model.Update(tea.WindowSizeMsg{Width: width, Height: height})
	next, ok := updated.(Model)
	if !ok {
		t.Fatalf("updated model type = %T, want playable.Model", updated)
	}
	return next, cmd
}

func contentRootForTest() string {
	return filepath.Join("..", "..", "content")
}

func progressWithAllPlayableCompleted(t *testing.T) *progress.Progress {
	t.Helper()
	p := progress.NewProgress()
	library, err := content.LoadLibrary(contentRootForTest())
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}
	entries, err := playlistEntries(library)
	if err != nil {
		t.Fatalf("playlistEntries returned error: %v", err)
	}
	for _, entry := range entries {
		exercise := library.Exercises[entry.ExerciseID]
		p.Missions[entry.ExerciseID] = progress.MissionProgress{
			Completed:      true,
			BestGrade:      "S",
			BestKeystrokes: exercise.Grading.OptimalKeyCount,
			Attempts:       1,
		}
	}
	return p
}

func progressCompleteBefore(t *testing.T, exerciseID string) *progress.Progress {
	t.Helper()
	p := progress.NewProgress()
	library, err := content.LoadLibrary(contentRootForTest())
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}
	entries, err := playlistEntries(library)
	if err != nil {
		t.Fatalf("playlistEntries returned error: %v", err)
	}
	for _, entry := range entries {
		if entry.ExerciseID == exerciseID {
			return p
		}
		exercise := library.Exercises[entry.ExerciseID]
		p.Missions[entry.ExerciseID] = progress.MissionProgress{
			Completed:      true,
			BestGrade:      "S",
			BestKeystrokes: exercise.Grading.OptimalKeyCount,
			Attempts:       1,
		}
	}
	t.Fatalf("exercise %q not found", exerciseID)
	return p
}
