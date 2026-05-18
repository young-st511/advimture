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
