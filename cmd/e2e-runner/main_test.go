package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestKeyBytes(t *testing.T) {
	tests := map[string]string{
		"enter":  "\r",
		"esc":    "\x1b",
		"ctrl+c": "\x03",
		"ctrl+r": "\x12",
		"right":  "\x1b[C",
		"left":   "\x1b[D",
		"up":     "\x1b[A",
		"down":   "\x1b[B",
		"space":  " ",
		"x":      "x",
	}

	for key, want := range tests {
		if got := keyBytes(key); got != want {
			t.Fatalf("keyBytes(%q) = %q, want %q", key, got, want)
		}
	}
}

func TestCleanTerminal(t *testing.T) {
	raw := []byte("\x1b]11;?\x1b\\\x1b[?1049h\x1b[1;1Hhello\r\n\x1b[31mworld\x1b[0m\x07")
	clean := cleanTerminal(raw)

	if !strings.Contains(clean, "hello\nworld") {
		t.Fatalf("cleaned output = %q", clean)
	}
	if strings.Contains(clean, "\x1b") {
		t.Fatalf("cleaned output still contains escape sequence: %q", clean)
	}
}

func TestCleanFinalScreenKeepsLastAdvimtureFrame(t *testing.T) {
	raw := []byte("ADVIMTURE | Tutorial | Exercise: 1/4\nold frame\nADVIMTURE | Tutorial | Exercise: 2/4\nnew frame\n")

	final := cleanFinalScreen(raw)

	if strings.Contains(final, "old frame") {
		t.Fatalf("cleanFinalScreen = %q, should not include earlier frame", final)
	}
	if !strings.Contains(final, "ADVIMTURE | Tutorial | Exercise: 2/4") || !strings.Contains(final, "new frame") {
		t.Fatalf("cleanFinalScreen = %q, want last frame", final)
	}
}

func TestCleanFinalScreenKeepsPlayableErrorFrame(t *testing.T) {
	raw := []byte("boot\nPlayable error: content missing\nq: quit\n")

	final := cleanFinalScreen(raw)

	if final != "Playable error: content missing\nq: quit" {
		t.Fatalf("cleanFinalScreen = %q, want playable error frame", final)
	}
}

func TestWaitForScreenIgnoresOutputBeforeOffset(t *testing.T) {
	var raw bytes.Buffer
	raw.WriteString("old screen\nNext: enter\n")
	var mu sync.Mutex

	_, err := waitForScreen(&mu, &raw, "Next: enter", time.Now().Add(20*time.Millisecond), raw.Len())
	if err == nil || !strings.Contains(err.Error(), "timed out") {
		t.Fatalf("waitForScreen error = %v, want timeout for stale output", err)
	}
}

func TestWaitForScreenFindsOutputAfterOffset(t *testing.T) {
	var raw bytes.Buffer
	raw.WriteString("old screen\n")
	var mu sync.Mutex
	offset := raw.Len()
	go func() {
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		defer mu.Unlock()
		raw.WriteString("new screen\nNext: enter\n")
	}()

	nextOffset, err := waitForScreen(&mu, &raw, "Next: enter", time.Now().Add(500*time.Millisecond), offset)
	if err != nil {
		t.Fatalf("waitForScreen returned error: %v", err)
	}
	if nextOffset <= offset {
		t.Fatalf("nextOffset = %d, want > %d", nextOffset, offset)
	}
}

func TestScreenContainsNormalizesWrappedText(t *testing.T) {
	screen := "│ 좋습니다. anchor가 오른쪽이어도 선택 범위를 정규화해 정확히 │\n│ 제거했습니다. │"
	want := "좋습니다. anchor가 오른쪽이어도 선택 범위를 정규화해 정확히 제거했습니다."

	if !screenContains(screen, want) {
		t.Fatalf("screenContains(%q, %q) = false, want true", screen, want)
	}
}

func TestAssertScenarioChecksKeyTrace(t *testing.T) {
	sc := scenario{
		Assert: assertionConfig{
			KeyTrace: []string{"l", "l"},
		},
	}
	result := runResult{
		trace: []string{"l", "h"},
	}

	err := assertScenario(sc, result)
	if err == nil || !strings.Contains(err.Error(), "key trace") {
		t.Fatalf("assertScenario error = %v, want key trace error", err)
	}
}

func TestAssertScenarioChecksProgressFileContains(t *testing.T) {
	home := t.TempDir()
	progressDir := filepath.Join(home, ".advimture")
	if err := os.MkdirAll(progressDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(progressDir, "progress.json"), []byte(`{"mission":"m01"}`), 0o644); err != nil {
		t.Fatal(err)
	}

	sc := scenario{
		Assert: assertionConfig{
			ProgressFileContains: []string{`"mission":"m01"`},
		},
	}
	result := runResult{homeDir: home}

	if err := assertScenario(sc, result); err != nil {
		t.Fatalf("assertScenario returned error: %v", err)
	}
}

func TestAssertScenarioChecksAppStateSummary(t *testing.T) {
	home := t.TempDir()
	stateDir := filepath.Join(home, ".advimture")
	if err := os.MkdirAll(stateDir, 0o755); err != nil {
		t.Fatal(err)
	}
	statePath := filepath.Join(stateDir, "e2e_state.json")
	raw := []byte(`{
		"buffer": ["abc"],
		"cursor": {"row": 0, "col": 2},
		"mode": "normal",
		"status": "succeeded",
		"score": {"grade": "S", "passed": true},
		"progress": {"mission_id": "mission-1", "completed": true}
	}`)
	if err := os.WriteFile(statePath, raw, 0o644); err != nil {
		t.Fatal(err)
	}

	row := 0
	col := 2
	passed := true
	completed := true
	sc := scenario{
		Assert: assertionConfig{
			AppState: appStateAssertion{
				Buffer: []string{"abc"},
				Cursor: &cursorAssertion{
					Row: &row,
					Col: &col,
				},
				Mode:   "normal",
				Status: "succeeded",
				Score: &scoreAssertion{
					Grade:  "S",
					Passed: &passed,
				},
				Progress: &progressAssertion{
					MissionID: "mission-1",
					Completed: &completed,
				},
			},
		},
	}
	result := runResult{homeDir: home}

	if err := assertScenario(sc, result); err != nil {
		t.Fatalf("assertScenario returned error: %v", err)
	}
}

func TestAssertScenarioChecksAppStateSelection(t *testing.T) {
	home := t.TempDir()
	stateDir := filepath.Join(home, ".advimture")
	if err := os.MkdirAll(stateDir, 0o755); err != nil {
		t.Fatal(err)
	}
	statePath := filepath.Join(stateDir, "e2e_state.json")
	raw := []byte(`{
		"mode": "visual",
		"selection": {
			"active": true,
			"kind": "charwise",
			"anchor": {"row": 0, "col": 1},
			"head": {"row": 0, "col": 3},
			"start": {"row": 0, "col": 1},
			"end": {"row": 0, "col": 3}
		}
	}`)
	if err := os.WriteFile(statePath, raw, 0o644); err != nil {
		t.Fatal(err)
	}

	active := true
	anchorRow := 0
	anchorCol := 1
	headCol := 3
	sc := scenario{
		Assert: assertionConfig{
			AppState: appStateAssertion{
				Selection: &selectionAssertion{
					Active: &active,
					Kind:   "charwise",
					Anchor: &cursorAssertion{
						Row: &anchorRow,
						Col: &anchorCol,
					},
					Head: &cursorAssertion{
						Col: &headCol,
					},
					End: &cursorAssertion{
						Col: &headCol,
					},
				},
			},
		},
	}

	if err := assertScenario(sc, runResult{homeDir: home}); err != nil {
		t.Fatalf("assertScenario returned error: %v", err)
	}
}

func TestAssertScenarioChecksAppStateReview(t *testing.T) {
	home := t.TempDir()
	stateDir := filepath.Join(home, ".advimture")
	if err := os.MkdirAll(stateDir, 0o755); err != nil {
		t.Fatal(err)
	}
	statePath := filepath.Join(stateDir, "e2e_state.json")
	raw := []byte(`{
		"review": {
			"queue_count": 3,
			"primary_exercise_id": "normal-motion-basic-002",
			"primary_reason": "incomplete",
			"daily_route": "오늘의 복구 루트: 3건 대기"
		}
	}`)
	if err := os.WriteFile(statePath, raw, 0o644); err != nil {
		t.Fatal(err)
	}

	queueCount := 3
	sc := scenario{
		Assert: assertionConfig{
			AppState: appStateAssertion{
				Review: &reviewAssertion{
					QueueCount:        &queueCount,
					PrimaryExerciseID: "normal-motion-basic-002",
					PrimaryReason:     "incomplete",
					DailyRoute:        "오늘의 복구 루트: 3건 대기",
				},
			},
		},
	}

	if err := assertScenario(sc, runResult{homeDir: home}); err != nil {
		t.Fatalf("assertScenario returned error: %v", err)
	}
}

func TestAssertScenarioChecksAppStateUIFocusPanel(t *testing.T) {
	home := t.TempDir()
	stateDir := filepath.Join(home, ".advimture")
	if err := os.MkdirAll(stateDir, 0o755); err != nil {
		t.Fatal(err)
	}
	statePath := filepath.Join(stateDir, "e2e_state.json")
	raw := []byte(`{
		"ui": {
			"focus_panel": {
				"kind": "training",
				"title": "TRAINING BRIEF",
				"lines": ["Coach: 훈련 키 l"],
				"actions": [{"id": "retry", "label": "다시 시도: r 또는 enter"}]
			}
		}
	}`)
	if err := os.WriteFile(statePath, raw, 0o644); err != nil {
		t.Fatal(err)
	}

	sc := scenario{
		Assert: assertionConfig{
			AppState: appStateAssertion{
				UI: &uiAssertion{
					FocusPanel: &focusPanelAssertion{
						Kind:  "training",
						Title: "TRAINING BRIEF",
						Lines: []string{"Coach: 훈련 키 l"},
						Actions: []focusActionAssertion{
							{ID: "retry", Label: "다시 시도: r 또는 enter"},
						},
					},
				},
			},
		},
	}

	if err := assertScenario(sc, runResult{homeDir: home}); err != nil {
		t.Fatalf("assertScenario returned error: %v", err)
	}
}

func TestAssertScenarioReportsAppStateReviewMismatch(t *testing.T) {
	home := t.TempDir()
	stateDir := filepath.Join(home, ".advimture")
	if err := os.MkdirAll(stateDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(stateDir, "e2e_state.json"), []byte(`{"review":{"queue_count":2}}`), 0o644); err != nil {
		t.Fatal(err)
	}

	queueCount := 3
	err := assertScenario(scenario{
		Assert: assertionConfig{
			AppState: appStateAssertion{
				Review: &reviewAssertion{QueueCount: &queueCount},
			},
		},
	}, runResult{homeDir: home})
	if err == nil || !strings.Contains(err.Error(), "review queue_count") {
		t.Fatalf("assertScenario error = %v, want review queue_count mismatch", err)
	}
}

func TestAssertScenarioReportsAppStateSelectionMismatch(t *testing.T) {
	home := t.TempDir()
	stateDir := filepath.Join(home, ".advimture")
	if err := os.MkdirAll(stateDir, 0o755); err != nil {
		t.Fatal(err)
	}
	statePath := filepath.Join(stateDir, "e2e_state.json")
	raw := []byte(`{
		"mode": "visual",
		"selection": {
			"active": true,
			"kind": "charwise",
			"anchor": {"row": 0, "col": 1},
			"head": {"row": 0, "col": 3},
			"start": {"row": 0, "col": 1},
			"end": {"row": 0, "col": 3}
		}
	}`)
	if err := os.WriteFile(statePath, raw, 0o644); err != nil {
		t.Fatal(err)
	}

	wantCol := 2
	err := assertScenario(scenario{
		Assert: assertionConfig{
			AppState: appStateAssertion{
				Selection: &selectionAssertion{
					End: &cursorAssertion{Col: &wantCol},
				},
			},
		},
	}, runResult{homeDir: home})
	if err == nil || !strings.Contains(err.Error(), "selection end col") {
		t.Fatalf("assertScenario error = %v, want selection end col mismatch", err)
	}
}

func TestAssertScenarioFailsWhenAppStateMissing(t *testing.T) {
	sc := scenario{
		Assert: assertionConfig{
			AppState: appStateAssertion{
				Mode: "normal",
			},
		},
	}

	err := assertScenario(sc, runResult{homeDir: t.TempDir()})
	if err == nil || !strings.Contains(err.Error(), "app state") {
		t.Fatalf("assertScenario error = %v, want app state error", err)
	}
}

func TestAssertScenarioReportsAppStateMismatch(t *testing.T) {
	home := t.TempDir()
	stateDir := filepath.Join(home, ".advimture")
	if err := os.MkdirAll(stateDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(stateDir, "e2e_state.json"), []byte(`{"mode":"insert"}`), 0o644); err != nil {
		t.Fatal(err)
	}

	sc := scenario{
		Assert: assertionConfig{
			AppState: appStateAssertion{
				Mode: "normal",
			},
		},
	}

	err := assertScenario(sc, runResult{homeDir: home})
	if err == nil || !strings.Contains(err.Error(), "mode") {
		t.Fatalf("assertScenario error = %v, want mode mismatch", err)
	}
}

func TestSetupHomeRejectsRealHomeByDefault(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}

	_, cleanup, err := setupHome(scenario{
		Setup: setupConfig{
			Home: home,
		},
	})
	defer cleanup()

	if err == nil || !strings.Contains(err.Error(), "unsafe home") {
		t.Fatalf("setupHome error = %v, want unsafe home error", err)
	}
}

func TestSetupHomeWritesProgressFixture(t *testing.T) {
	home, cleanup, err := setupHome(scenario{
		Setup: setupConfig{
			Home:         "temp",
			ProgressFile: `{"missions":{"normal-motion-basic-001":{"completed":true}}}`,
		},
	})
	defer cleanup()
	if err != nil {
		t.Fatalf("setupHome returned error: %v", err)
	}

	raw, err := os.ReadFile(filepath.Join(home, ".advimture", "progress.json"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), "normal-motion-basic-001") {
		t.Fatalf("progress fixture = %s, want normal-motion-basic-001", raw)
	}
}

func TestSetupHomeBuildsProgressBeforeExercise(t *testing.T) {
	home, cleanup, err := setupHome(scenario{
		Setup: setupConfig{
			Home:           "temp",
			CompleteBefore: "survival-save-quit-001",
			ContentRoot:    filepath.Join("..", "..", "content"),
		},
	})
	defer cleanup()
	if err != nil {
		t.Fatalf("setupHome returned error: %v", err)
	}

	raw, err := os.ReadFile(filepath.Join(home, ".advimture", "progress.json"))
	if err != nil {
		t.Fatal(err)
	}
	var decoded struct {
		Missions map[string]struct {
			Completed bool   `json:"completed"`
			BestGrade string `json:"best_grade"`
			Attempts  int    `json:"attempts"`
		} `json:"missions"`
	}
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if len(decoded.Missions) != 4 {
		t.Fatalf("missions = %d, want 4 before survival-save-quit-001: %s", len(decoded.Missions), raw)
	}
	if !decoded.Missions["normal-motion-basic-004"].Completed {
		t.Fatalf("normal-motion-basic-004 fixture = %+v, want completed", decoded.Missions["normal-motion-basic-004"])
	}
	if _, ok := decoded.Missions["survival-save-quit-001"]; ok {
		t.Fatalf("survival-save-quit-001 should not be completed before itself: %s", raw)
	}
}

func TestProgressFixtureRejectsInlineAndBuilderTogether(t *testing.T) {
	_, err := progressFixtureJSON(setupConfig{
		ProgressFile:   `{"missions":{}}`,
		CompleteBefore: "normal-motion-basic-001",
	})
	if err == nil || !strings.Contains(err.Error(), "cannot be used together") {
		t.Fatalf("progressFixtureJSON error = %v, want mutual exclusion", err)
	}
}

func TestProgressFixtureRejectsMissingCompleteBeforeExercise(t *testing.T) {
	_, err := progressFixtureJSON(setupConfig{
		CompleteBefore: "missing-exercise",
		ContentRoot:    filepath.Join("..", "..", "content"),
	})
	if err == nil || !strings.Contains(err.Error(), "not found") {
		t.Fatalf("progressFixtureJSON error = %v, want not found", err)
	}
}

func TestWriteEvidenceWritesSummary(t *testing.T) {
	root := t.TempDir()
	runErr := assertError("screen mismatch")
	result := runResult{
		clean:    "screen",
		exitCode: 1,
		homeDir:  t.TempDir(),
		trace:    []string{"ctrl+c"},
	}

	if err := writeEvidence(root, scenario{ID: "summary"}, result, runErr); err != nil {
		t.Fatalf("writeEvidence returned error: %v", err)
	}

	raw, err := os.ReadFile(filepath.Join(root, "summary", "summary.json"))
	if err != nil {
		t.Fatal(err)
	}
	var summary summaryEvidence
	if err := json.Unmarshal(raw, &summary); err != nil {
		t.Fatal(err)
	}
	if summary.Passed {
		t.Fatal("summary.Passed = true, want false")
	}
	if summary.Error != "screen mismatch" {
		t.Fatalf("summary.Error = %q, want screen mismatch", summary.Error)
	}
	if summary.ExitCode != 1 {
		t.Fatalf("summary.ExitCode = %d, want 1", summary.ExitCode)
	}
}

func TestWriteEvidenceCopiesAppStateAndProgressSnapshots(t *testing.T) {
	root := t.TempDir()
	result := runResult{
		clean:       "screen",
		exitCode:    0,
		homeDir:     t.TempDir(),
		appStateRaw: []byte(`{"mode":"normal"}`),
		progressRaw: []byte(`{"missions":{}}`),
	}
	sc := scenario{
		ID: "snapshots",
		Evidence: evidenceConfig{
			SaveAppState: true,
			SaveProgress: true,
		},
	}

	if err := writeEvidence(root, sc, result, nil); err != nil {
		t.Fatalf("writeEvidence returned error: %v", err)
	}
	for _, name := range []string{"app_state.json", "progress.json"} {
		if _, err := os.Stat(filepath.Join(root, "snapshots", name)); err != nil {
			t.Fatalf("%s was not written: %v", name, err)
		}
	}
}

func TestWriteEvidenceWritesScreenTimeline(t *testing.T) {
	root := t.TempDir()
	result := runResult{
		clean:    "frame one\nframe two",
		exitCode: 0,
		homeDir:  t.TempDir(),
	}
	sc := scenario{
		ID: "timeline",
		Evidence: evidenceConfig{
			SaveScreenTimeline: true,
		},
	}

	if err := writeEvidence(root, sc, result, nil); err != nil {
		t.Fatalf("writeEvidence returned error: %v", err)
	}
	raw, err := os.ReadFile(filepath.Join(root, "timeline", "screen_timeline.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(raw) != result.clean {
		t.Fatalf("screen_timeline.txt = %q, want clean screen timeline", raw)
	}
	summaryRaw, err := os.ReadFile(filepath.Join(root, "timeline", "summary.json"))
	if err != nil {
		t.Fatal(err)
	}
	var summary summaryEvidence
	if err := json.Unmarshal(summaryRaw, &summary); err != nil {
		t.Fatal(err)
	}
	if !summary.ScreenTimeline {
		t.Fatal("summary.ScreenTimeline = false, want true")
	}
}

func TestWriteEvidenceWritesScreenFinal(t *testing.T) {
	root := t.TempDir()
	result := runResult{
		raw:      []byte("ADVIMTURE | old\nold frame\nADVIMTURE | final\nfinal frame"),
		clean:    "ADVIMTURE | old\nold frame\nADVIMTURE | final\nfinal frame",
		exitCode: 0,
		homeDir:  t.TempDir(),
	}
	sc := scenario{
		ID: "final",
		Evidence: evidenceConfig{
			SaveScreenFinal: true,
		},
	}

	if err := writeEvidence(root, sc, result, nil); err != nil {
		t.Fatalf("writeEvidence returned error: %v", err)
	}
	raw, err := os.ReadFile(filepath.Join(root, "final", "screen_final.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), "old frame") || !strings.Contains(string(raw), "final frame") {
		t.Fatalf("screen_final.txt = %q, want only final frame", raw)
	}
	summaryRaw, err := os.ReadFile(filepath.Join(root, "final", "summary.json"))
	if err != nil {
		t.Fatal(err)
	}
	var summary summaryEvidence
	if err := json.Unmarshal(summaryRaw, &summary); err != nil {
		t.Fatal(err)
	}
	if !summary.ScreenFinal {
		t.Fatal("summary.ScreenFinal = false, want true")
	}
}

func TestBuildSummaryRecordsAppStateLoaded(t *testing.T) {
	home := t.TempDir()
	stateDir := filepath.Join(home, ".advimture")
	if err := os.MkdirAll(stateDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(stateDir, "e2e_state.json"), []byte(`{"mode":"normal"}`), 0o644); err != nil {
		t.Fatal(err)
	}

	summary := buildSummary(scenario{ID: "state"}, runResult{homeDir: home}, nil)
	if !summary.AppStateExists {
		t.Fatal("AppStateExists = false, want true")
	}
	if summary.AppStatePath == "" {
		t.Fatal("AppStatePath is empty")
	}
}

func TestGoToolEnvSetsTempGoCacheWhenUnset(t *testing.T) {
	t.Setenv("GOCACHE", "")
	t.Setenv("GOPATH", "/existing/go")
	t.Setenv("GOMODCACHE", "/existing/go/pkg/mod")
	home := t.TempDir()

	env := goToolEnv(home)
	want := "GOCACHE=" + filepath.Join(home, ".cache", "go-build")

	if len(env) != 1 || env[0] != want {
		t.Fatalf("goToolEnv() = %v, want %q", env, want)
	}
}

func TestGoToolEnvRespectsExistingGoCache(t *testing.T) {
	t.Setenv("GOCACHE", "/custom/cache")
	t.Setenv("GOPATH", "/existing/go")
	t.Setenv("GOMODCACHE", "/existing/go/pkg/mod")

	if env := goToolEnv(t.TempDir()); len(env) != 0 {
		t.Fatalf("goToolEnv() = %v, want no override", env)
	}
}

func TestGoToolEnvPinsParentGoModuleCacheWhenHomeChanges(t *testing.T) {
	t.Setenv("GOCACHE", "")
	t.Setenv("GOPATH", "")
	t.Setenv("GOMODCACHE", "")
	previousLookup := lookupGoEnv
	t.Cleanup(func() {
		lookupGoEnv = previousLookup
	})
	lookupGoEnv = func(key string) (string, error) {
		switch key {
		case "GOPATH":
			return "/parent/go\n", nil
		case "GOMODCACHE":
			return "/parent/go/pkg/mod\n", nil
		default:
			return "", nil
		}
	}

	env := goToolEnv(t.TempDir())

	if !containsString(env, "GOPATH=/parent/go") {
		t.Fatalf("goToolEnv() = %v, want GOPATH pin", env)
	}
	if !containsString(env, "GOMODCACHE=/parent/go/pkg/mod") {
		t.Fatalf("goToolEnv() = %v, want GOMODCACHE pin", env)
	}
}

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

type assertError string

func (e assertError) Error() string {
	return string(e)
}
