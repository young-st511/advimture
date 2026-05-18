package content

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	exerciseruntime "github.com/young-st511/advimture/internal/runtime"
	"github.com/young-st511/advimture/internal/vimengine"
)

func TestLoadLibraryLoadsRootContent(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	if len(lib.CommandClusters) != 3 {
		t.Fatalf("command clusters = %d, want 3", len(lib.CommandClusters))
	}
	if len(lib.Exercises) != 3 {
		t.Fatalf("exercises = %d, want 3", len(lib.Exercises))
	}
	if len(lib.Scenarios) != 3 {
		t.Fatalf("scenarios = %d, want 3", len(lib.Scenarios))
	}
	if len(lib.Playlists) != 1 {
		t.Fatalf("playlists = %d, want 1", len(lib.Playlists))
	}
}

func TestLoadLibraryFiltersPlayableExercises(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	playable := lib.PlayableExercises()
	if len(playable) != 1 {
		t.Fatalf("playable exercises = %d, want 1: %+v", len(playable), playable)
	}
	if playable[0].ID != "normal-motion-basic-001" {
		t.Fatalf("playable[0].ID = %q, want normal-motion-basic-001", playable[0].ID)
	}
	if playable[0].ReplayStatus != ReplayStatusPass {
		t.Fatalf("playable[0].ReplayStatus = %q, want pass", playable[0].ReplayStatus)
	}
}

func TestCompileLoadedExerciseMatchesPlayableTarget(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	compiled, err := lib.CompileExercise("normal-motion-basic-001")
	if err != nil {
		t.Fatalf("CompileExercise returned error: %v", err)
	}

	session := exerciseruntime.NewSession(compiled.Exercise)
	session.ApplyKey(vimengine.KeyL)
	result := session.ApplyKey(vimengine.KeyL)

	if result.State.Status != exerciseruntime.StatusSucceeded {
		t.Fatalf("status = %q, want succeeded", result.State.Status)
	}
	if result.State.Vim.Cursor.Row != 0 || result.State.Vim.Cursor.Col != 2 {
		t.Fatalf("cursor = %d,%d, want 0,2", result.State.Vim.Cursor.Row, result.State.Vim.Cursor.Col)
	}
	assertStrings(t, compiled.ExpectedKeys, []string{"l", "l"})
	assertStrings(t, compiled.AllowedKeys, []string{"h", "j", "k", "l", "esc"})
}

func TestCoverageReportsMissingCommandsWithoutFailingLoad(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	var normal CoverageReport
	for _, report := range lib.CoverageReports() {
		if report.CommandClusterID == "normal-motion-basic" {
			normal = report
			break
		}
	}

	assertStrings(t, normal.Covered, []string{"l"})
	assertStrings(t, normal.Missing, []string{"h", "j", "k"})
}

func TestLoadLibraryPreservesE2EAssertions(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	assertion := lib.Exercises["normal-motion-basic-001"].E2EAssertions
	assertStrings(t, assertion.Buffer, []string{"abc"})
	if assertion.Cursor == nil {
		t.Fatal("E2E cursor assertion is nil")
	}
	if assertion.Cursor.Row != 0 || assertion.Cursor.Col != 2 {
		t.Fatalf("E2E cursor = %d,%d, want 0,2", assertion.Cursor.Row, assertion.Cursor.Col)
	}
	if assertion.Mode != "normal" || assertion.Status != "succeeded" {
		t.Fatalf("E2E mode/status = %q/%q, want normal/succeeded", assertion.Mode, assertion.Status)
	}
}

func TestLoadLibraryRejectsMissingCommandClusterReference(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "exercises", "bad.yaml"), `
exercises:
  - id: missing-cluster
    status: draft
    command_cluster: no-such-cluster
    engine_support: implemented
    title: Bad
    initial_state:
      mode: normal
      buffer: "abc\n"
    target_state:
      cursor: {row: 0, col: 1}
    optimal_keys: ["l"]
    allowed_keys: ["l"]
    grading:
      pass_condition: "cursor.row == 0 && cursor.col == 1"
      optimal_key_count: 1
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "missing command cluster") {
		t.Fatalf("LoadLibrary error = %v, want missing command cluster", err)
	}
}

func TestLoadLibraryRejectsApprovedImplementedClusterWithoutCoverage(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "command_clusters", "clusters.yaml"), `
command_clusters:
  - id: normal-motion-basic
    status: approved
    compatibility_tier: exact
    engine_support: implemented
    title: Basic motion
    commands: ["h", "j", "k", "l"]
    coverage_required: []
    oracle: optional
    purpose: Move cursor
    prerequisite: []
    difficulty: beginner
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "coverage_required") {
		t.Fatalf("LoadLibrary error = %v, want coverage_required", err)
	}
}

func TestLoadLibraryRejectsApprovedImplementedExerciseWithoutReplayPass(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "exercises", "exercises.yaml"), `
exercises:
  - id: normal-motion-basic-001
    status: approved
    command_cluster: normal-motion-basic
    engine_support: implemented
    replay_status: pending
    title: Move right
    initial_state:
      mode: normal
      buffer: "abc\n"
    target_state:
      mode: normal
      cursor: {row: 0, col: 1}
    optimal_keys: ["l"]
    allowed_keys: ["l"]
    grading:
      pass_condition: "cursor.row == 0 && cursor.col == 1"
      optimal_key_count: 1
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "replay_status") {
		t.Fatalf("LoadLibrary error = %v, want replay_status", err)
	}
}

func TestLoadLibraryRejectsReplayPassWhenOptimalKeysMissGoal(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "exercises", "exercises.yaml"), `
exercises:
  - id: normal-motion-basic-001
    status: approved
    command_cluster: normal-motion-basic
    engine_support: implemented
    replay_status: pass
    title: Move right
    initial_state:
      mode: normal
      buffer: "abc\n"
    target_state:
      mode: normal
      cursor: {row: 0, col: 2}
    optimal_keys: ["l"]
    allowed_keys: ["l"]
    grading:
      pass_condition: "cursor.row == 0 && cursor.col == 2"
      optimal_key_count: 1
    e2e_assertions:
      buffer: ["abc"]
      cursor: {row: 0, col: 2}
      mode: normal
      status: succeeded
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "replay failed") {
		t.Fatalf("LoadLibrary error = %v, want replay failed", err)
	}
}

func TestLoadLibraryRejectsReplayPassWithTrailingOptimalKeys(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "exercises", "exercises.yaml"), `
exercises:
  - id: normal-motion-basic-001
    status: approved
    command_cluster: normal-motion-basic
    engine_support: implemented
    replay_status: pass
    title: Move right
    initial_state:
      mode: normal
      buffer: "abc\n"
    target_state:
      mode: normal
      cursor: {row: 0, col: 1}
    optimal_keys: ["l", "l"]
    allowed_keys: ["l"]
    grading:
      pass_condition: "cursor.row == 0 && cursor.col == 1"
      optimal_key_count: 2
    e2e_assertions:
      buffer: ["abc"]
      cursor: {row: 0, col: 1}
      mode: normal
      status: succeeded
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "key trace") {
		t.Fatalf("LoadLibrary error = %v, want key trace", err)
	}
}

func TestLoadLibraryRejectsReplayPassWithoutE2EAssertions(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "exercises", "exercises.yaml"), `
exercises:
  - id: normal-motion-basic-001
    status: approved
    command_cluster: normal-motion-basic
    engine_support: implemented
    replay_status: pass
    title: Move right
    initial_state:
      mode: normal
      buffer: "abc\n"
    target_state:
      mode: normal
      cursor: {row: 0, col: 1}
    optimal_keys: ["l"]
    allowed_keys: ["l"]
    grading:
      pass_condition: "cursor.row == 0 && cursor.col == 1"
      optimal_key_count: 1
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "e2e_assertions") {
		t.Fatalf("LoadLibrary error = %v, want e2e_assertions", err)
	}
}

func TestLoadLibraryRejectsMissingScenarioExerciseReference(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "scenarios", "bad.yaml"), `
scenarios:
  - id: bad-scenario
    status: draft
    exercise_id: no-such-exercise
    engine_support: implemented
    mission_title: Bad
    briefing: Bad
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "missing exercise") {
		t.Fatalf("LoadLibrary error = %v, want missing exercise", err)
	}
}

func TestLoadLibraryRejectsBlankPlayableScenarioCopy(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "scenarios", "scenarios.yaml"), `
scenarios:
  - id: normal-motion-basic-001-scenario
    status: approved
    exercise_id: normal-motion-basic-001
    engine_support: implemented
    mission_title: Move right
    briefing: Move right
    mentor_success: ""
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "mentor_success") {
		t.Fatalf("LoadLibrary error = %v, want mentor_success", err)
	}
}

func TestLoadLibraryRejectsOutOfRangeCursor(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "exercises", "bad.yaml"), `
exercises:
  - id: bad-cursor
    status: draft
    command_cluster: normal-motion-basic
    engine_support: implemented
    title: Bad
    initial_state:
      mode: normal
      buffer: "abc\n"
    target_state:
      cursor: {row: 0, col: 9}
    optimal_keys: ["l"]
    allowed_keys: ["l"]
    grading:
      pass_condition: "cursor.row == 0 && cursor.col == 9"
      optimal_key_count: 1
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "goal cursor is out of range") {
		t.Fatalf("LoadLibrary error = %v, want out-of-range cursor", err)
	}
}

func TestLoadLibraryRejectsOptimalKeyCountMismatch(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "exercises", "bad.yaml"), `
exercises:
  - id: bad-count
    status: draft
    command_cluster: normal-motion-basic
    engine_support: implemented
    title: Bad
    initial_state:
      mode: normal
      buffer: "abc\n"
    target_state:
      cursor: {row: 0, col: 1}
    optimal_keys: ["l"]
    allowed_keys: ["l"]
    grading:
      pass_condition: "cursor.row == 0 && cursor.col == 1"
      optimal_key_count: 2
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "optimal_key_count") {
		t.Fatalf("LoadLibrary error = %v, want optimal_key_count mismatch", err)
	}
}

func TestLoadLibraryRejectsForbiddenOptimalKey(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "exercises", "bad.yaml"), `
exercises:
  - id: bad-key
    status: draft
    command_cluster: normal-motion-basic
    engine_support: implemented
    title: Bad
    initial_state:
      mode: normal
      buffer: "abc\n"
    target_state:
      cursor: {row: 0, col: 1}
    optimal_keys: ["l"]
    allowed_keys: ["h"]
    forbidden_keys: ["l"]
    grading:
      pass_condition: "cursor.row == 0 && cursor.col == 1"
      optimal_key_count: 1
    e2e_assertions:
      buffer: ["abc"]
      cursor: {row: 0, col: 1}
      mode: normal
      status: succeeded
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "not allowed") {
		t.Fatalf("LoadLibrary error = %v, want not allowed key", err)
	}
}

func validLibraryFixture(t *testing.T) string {
	t.Helper()

	root := t.TempDir()
	for _, dir := range []string{"command_clusters", "exercises", "scenarios", "playlists"} {
		if err := os.MkdirAll(filepath.Join(root, dir), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	writeYAML(t, filepath.Join(root, "command_clusters", "clusters.yaml"), `
command_clusters:
  - id: normal-motion-basic
    status: approved
    compatibility_tier: exact
    engine_support: implemented
    title: Basic motion
    commands: ["h", "j", "k", "l"]
    coverage_required: ["h", "j", "k", "l"]
    oracle: optional
    purpose: Move cursor
    prerequisite: []
    difficulty: beginner
`)
	writeYAML(t, filepath.Join(root, "exercises", "exercises.yaml"), `
exercises:
  - id: normal-motion-basic-001
    status: approved
    command_cluster: normal-motion-basic
    engine_support: implemented
    replay_status: pass
    title: Move right
    initial_state:
      mode: normal
      buffer: "abc\n"
    target_state:
      mode: normal
      cursor: {row: 0, col: 1}
    optimal_keys: ["l"]
    allowed_keys: ["l"]
    grading:
      pass_condition: "cursor.row == 0 && cursor.col == 1"
      optimal_key_count: 1
    e2e_assertions:
      buffer: ["abc"]
      cursor: {row: 0, col: 1}
      mode: normal
      status: succeeded
`)
	writeYAML(t, filepath.Join(root, "scenarios", "scenarios.yaml"), `
scenarios:
  - id: normal-motion-basic-001-scenario
    status: approved
    exercise_id: normal-motion-basic-001
    engine_support: implemented
    mission_title: Move right
    briefing: Move right
    mentor_success: Done
`)
	writeYAML(t, filepath.Join(root, "playlists", "playlists.yaml"), `
playlists:
  - id: first
    status: draft
    title: First
    beats:
      - id: beat-1
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
`)
	return root
}

func writeYAML(t *testing.T, path string, content string) {
	t.Helper()

	if err := os.WriteFile(path, []byte(strings.TrimSpace(content)+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
}
