package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestKeyBytes(t *testing.T) {
	tests := map[string]string{
		"enter":  "\r",
		"esc":    "\x1b",
		"ctrl+c": "\x03",
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

type assertError string

func (e assertError) Error() string {
	return string(e)
}
