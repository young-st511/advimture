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

	if len(lib.CommandClusters) != 19 {
		t.Fatalf("command clusters = %d, want 19", len(lib.CommandClusters))
	}
	if len(lib.Exercises) != 96 {
		t.Fatalf("exercises = %d, want 96", len(lib.Exercises))
	}
	if len(lib.Scenarios) != 96 {
		t.Fatalf("scenarios = %d, want 96", len(lib.Scenarios))
	}
	if len(lib.Playlists) != 22 {
		t.Fatalf("playlists = %d, want 22", len(lib.Playlists))
	}
}

func TestLoadLibraryFiltersPlayableExercises(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	playable := lib.PlayableExercises()
	if len(playable) != 96 {
		t.Fatalf("playable exercises = %d, want 96: %+v", len(playable), playable)
	}
	if playable[0].ID != "change-with-motion-001" {
		t.Fatalf("playable[0].ID = %q, want change-with-motion-001", playable[0].ID)
	}
	assertPlayableIDs(t, playable, []string{
		"change-with-motion-001",
		"change-with-motion-002",
		"change-with-motion-003",
		"char-find-line-001",
		"char-find-line-002",
		"char-find-line-003",
		"char-find-line-004",
		"char-find-line-005",
		"char-find-line-006",
		"command-choice-inline-target-001",
		"command-choice-repeat-substitute-001",
		"command-choice-scope-001",
		"delete-with-motion-001",
		"delete-with-motion-002",
		"delete-with-motion-003",
		"incident-hotfix-001",
		"incident-hotfix-002",
		"incident-hotfix-003",
		"incident-hotfix-004",
		"incident-hotfix-005",
		"incident-hotfix-006",
		"incident-inline-target-001",
		"incident-inline-target-002",
		"incident-linewise-001",
		"incident-linewise-002",
		"incident-linewise-003",
		"incident-linewise-004",
		"incident-linewise-005",
		"incident-structure-001",
		"incident-structure-002",
		"incident-structure-003",
		"incident-structure-004",
		"incident-structure-005",
		"incident-visual-001",
		"incident-visual-002",
		"incident-visual-003",
		"incident-visual-004",
		"incident-visual-005",
		"insert-mode-entry-001",
		"insert-mode-entry-002",
		"insert-mode-entry-003",
		"normal-motion-basic-001",
		"normal-motion-basic-002",
		"normal-motion-basic-003",
		"normal-motion-basic-004",
		"open-line-edit-001",
		"open-line-edit-002",
		"open-line-edit-003",
		"open-line-edit-004",
		"open-line-edit-005",
		"repeat-last-change-001",
		"repeat-last-change-002",
		"repeat-last-change-003",
		"repeat-last-change-004",
		"search-basic-001",
		"search-basic-002",
		"search-basic-003",
		"search-basic-004",
		"single-char-edit-001",
		"single-char-edit-002",
		"survival-save-quit-001",
		"survival-save-quit-002",
		"survival-save-quit-003",
		"text-object-inner-word-001",
		"text-object-inner-word-002",
		"text-object-inner-word-003",
		"text-object-inner-word-004",
		"text-object-inner-word-005",
		"text-object-inner-word-006",
		"text-object-quote-pair-001",
		"text-object-quote-pair-002",
		"text-object-quote-pair-003",
		"text-object-quote-pair-004",
		"undo-redo-basic-001",
		"undo-redo-basic-002",
		"vim-ex-command-substitute-001",
		"vim-ex-command-substitute-002",
		"vim-ex-command-substitute-003",
		"visual-char-line-001",
		"visual-char-line-002",
		"visual-char-line-003",
		"visual-line-basic-001",
		"visual-line-basic-002",
		"visual-line-basic-003",
		"whole-file-navigation-001",
		"whole-file-navigation-002",
		"whole-file-navigation-003",
		"whole-file-navigation-004",
		"word-motion-basic-001",
		"word-motion-basic-002",
		"word-motion-basic-003",
		"yank-put-basic-001",
		"yank-put-basic-002",
		"yank-put-basic-003",
		"yank-put-basic-004",
		"yank-put-basic-005",
	})
	for _, exercise := range playable {
		if exercise.ReplayStatus != ReplayStatusPass {
			t.Fatalf("exercise %q ReplayStatus = %q, want pass", exercise.ID, exercise.ReplayStatus)
		}
	}
}

func TestLoadLibraryFiltersPlayablePlaylists(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	playlists := lib.PlayablePlaylists()
	got := make([]string, 0, len(playlists))
	for _, playlist := range playlists {
		got = append(got, playlist.ID)
		if len(playlist.Beats) > maxTutorialPlaylistBeats {
			t.Fatalf("playlist %q has %d beats, want at most %d", playlist.ID, len(playlist.Beats), maxTutorialPlaylistBeats)
		}
	}
	assertStrings(t, got, []string{
		"tutorial-0-movement",
		"tutorial-1-survival",
		"tutorial-2-fast-navigation",
		"tutorial-3-small-edits",
		"tutorial-4-ex-command",
		"tutorial-5-operator-grammar",
		"tutorial-6-yank-put",
		"tutorial-7-text-object-inner-word",
		"tutorial-8-open-line-edit",
		"tutorial-9-repeat-last-change",
		"tutorial-90-search-basic",
		"tutorial-91-text-object-quote-pair",
		"tutorial-92-visual-selection",
		"tutorial-93-visual-line",
		"tutorial-94-char-find-line",
		"incident-001-hotfix",
		"incident-002-structure-recovery",
		"incident-003-visual-recovery",
		"incident-004-linewise-block-recovery",
		"incident-005-command-choice",
		"incident-006-inline-target-repair",
	})
}

func TestPlayablePlaylistsUseExplicitCategoryOrder(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "playlists", "playlists.yaml"), `
playlists:
  - id: incident-001
    status: approved
    category: incident
    order: 1
    title: Incident
    beats:
      - id: beat-incident
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
  - id: tutorial-b
    status: approved
    category: tutorial
    order: 2
    title: Tutorial B
    beats:
      - id: beat-b
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
  - id: tutorial-a
    status: approved
    category: tutorial
    order: 1
    title: Tutorial A
    beats:
      - id: beat-a
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
`)

	lib, err := LoadLibrary(root)
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	playlists := lib.PlayablePlaylists()
	got := make([]string, 0, len(playlists))
	for _, playlist := range playlists {
		got = append(got, playlist.ID)
	}
	assertStrings(t, got, []string{"tutorial-a", "tutorial-b", "incident-001"})
}

func TestLoadLibraryRejectsApprovedPlaylistWithoutExplicitOrder(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "playlists", "playlists.yaml"), `
playlists:
  - id: missing-order
    status: approved
    category: tutorial
    title: Missing order
    beats:
      - id: beat-1
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "order") {
		t.Fatalf("LoadLibrary error = %v, want order", err)
	}
}

func TestLoadLibraryAllowsRetiredLegacyPlaylistLongerThanTutorialLimit(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	legacy := lib.Playlists["first-5-minute"]
	if legacy.Status != StatusRetired {
		t.Fatalf("legacy status = %q, want retired", legacy.Status)
	}
	if len(legacy.Beats) <= maxTutorialPlaylistBeats {
		t.Fatalf("legacy beats = %d, want more than %d", len(legacy.Beats), maxTutorialPlaylistBeats)
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

func TestCoverageReportsNormalMotionCommandsCovered(t *testing.T) {
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

	assertStrings(t, normal.Covered, []string{"h", "j", "k", "l"})
	if len(normal.Missing) != 0 {
		t.Fatalf("normal motion missing coverage = %+v, want empty", normal.Missing)
	}
}

func TestCoverageReportsWordMotionCommandsCovered(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	var word CoverageReport
	for _, report := range lib.CoverageReports() {
		if report.CommandClusterID == "word-motion-basic" {
			word = report
			break
		}
	}

	assertStrings(t, word.Covered, []string{"w", "b", "e"})
	if len(word.Missing) != 0 {
		t.Fatalf("word motion missing coverage = %+v, want empty", word.Missing)
	}
}

func TestCoverageReportsSurvivalCommandsCovered(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	var survival CoverageReport
	for _, report := range lib.CoverageReports() {
		if report.CommandClusterID == "survival-save-quit" {
			survival = report
			break
		}
	}

	assertStrings(t, survival.Covered, []string{"esc", ":q!", ":wq"})
	if len(survival.Missing) != 0 {
		t.Fatalf("survival missing coverage = %+v, want empty", survival.Missing)
	}
}

func TestCoverageReportsNavigationCommandsCovered(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	var navigation CoverageReport
	for _, report := range lib.CoverageReports() {
		if report.CommandClusterID == "whole-file-navigation" {
			navigation = report
			break
		}
	}

	assertStrings(t, navigation.Covered, []string{"gg", "G", "0", "$"})
	if len(navigation.Missing) != 0 {
		t.Fatalf("navigation missing coverage = %+v, want empty", navigation.Missing)
	}
}

func TestCoverageReportsExCommandSubstituteCommandsCovered(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	var substitute CoverageReport
	for _, report := range lib.CoverageReports() {
		if report.CommandClusterID == "vim-ex-command-substitute" {
			substitute = report
			break
		}
	}

	assertStrings(t, substitute.Covered, []string{":s", ":%s", ":2,3s"})
	if len(substitute.Missing) != 0 {
		t.Fatalf("substitute missing coverage = %+v, want empty", substitute.Missing)
	}
}

func TestCoverageReportsSmallEditCommandsCovered(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	reports := make(map[string]CoverageReport)
	for _, report := range lib.CoverageReports() {
		reports[report.CommandClusterID] = report
	}

	single := reports["single-char-edit"]
	assertStrings(t, single.Covered, []string{"x", "r"})
	if len(single.Missing) != 0 {
		t.Fatalf("single char edit missing coverage = %+v, want empty", single.Missing)
	}

	insert := reports["insert-mode-entry"]
	assertStrings(t, insert.Covered, []string{"i", "a", "A", "esc"})
	if len(insert.Missing) != 0 {
		t.Fatalf("insert mode entry missing coverage = %+v, want empty", insert.Missing)
	}

	undo := reports["undo-redo-basic"]
	assertStrings(t, undo.Covered, []string{"u", "ctrl+r"})
	if len(undo.Missing) != 0 {
		t.Fatalf("undo redo missing coverage = %+v, want empty", undo.Missing)
	}
}

func TestCoverageReportsOperatorGrammarCommandsCovered(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	reports := make(map[string]CoverageReport)
	for _, report := range lib.CoverageReports() {
		reports[report.CommandClusterID] = report
	}

	deleteReport := reports["delete-with-motion"]
	assertStrings(t, deleteReport.Covered, []string{"dw", "d$", "dd"})
	if len(deleteReport.Missing) != 0 {
		t.Fatalf("delete with motion missing coverage = %+v, want empty", deleteReport.Missing)
	}

	changeReport := reports["change-with-motion"]
	assertStrings(t, changeReport.Covered, []string{"cw", "c$", "cc"})
	if len(changeReport.Missing) != 0 {
		t.Fatalf("change with motion missing coverage = %+v, want empty", changeReport.Missing)
	}

	yankPutReport := reports["yank-put-basic"]
	assertStrings(t, yankPutReport.Covered, []string{"yw", "y$", "yy", "p", "P"})
	if len(yankPutReport.Missing) != 0 {
		t.Fatalf("yank put missing coverage = %+v, want empty", yankPutReport.Missing)
	}

	textObjectReport := reports["text-object-inner-word"]
	assertStrings(t, textObjectReport.Covered, []string{"diw", "ciw", "yiw"})
	if len(textObjectReport.Missing) != 0 {
		t.Fatalf("text object missing coverage = %+v, want empty", textObjectReport.Missing)
	}

	openLineReport := reports["open-line-edit"]
	assertStrings(t, openLineReport.Covered, []string{"o", "O"})
	if len(openLineReport.Missing) != 0 {
		t.Fatalf("open line missing coverage = %+v, want empty", openLineReport.Missing)
	}

	repeatReport := reports["repeat-last-change"]
	assertStrings(t, repeatReport.Covered, []string{"."})
	if len(repeatReport.Missing) != 0 {
		t.Fatalf("repeat last change missing coverage = %+v, want empty", repeatReport.Missing)
	}

	searchReport := reports["search-basic"]
	assertStrings(t, searchReport.Covered, []string{"/", "n", "N"})
	if len(searchReport.Missing) != 0 {
		t.Fatalf("search basic missing coverage = %+v, want empty", searchReport.Missing)
	}

	charFindReport := reports["char-find-line"]
	assertStrings(t, charFindReport.Covered, []string{"f", "t", "df", "dt", "cf", "ct"})
	if len(charFindReport.Missing) != 0 {
		t.Fatalf("char find missing coverage = %+v, want empty", charFindReport.Missing)
	}
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

func TestLoadLibraryPreservesE2ESelectionAssertion(t *testing.T) {
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
      cursor: {row: 0, col: 1}
    optimal_keys: ["v", "l"]
    allowed_keys: ["v", "l"]
    grading:
      pass_condition: "cursor.row == 0 && cursor.col == 1"
      optimal_key_count: 2
    e2e_assertions:
      buffer: ["abc"]
      cursor: {row: 0, col: 1}
      mode: visual
      status: succeeded
      selection:
        active: true
        kind: charwise
        anchor: {row: 0, col: 0}
        head: {row: 0, col: 1}
        start: {row: 0, col: 0}
        end: {row: 0, col: 1}
`)

	lib, err := LoadLibrary(root)
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	selection := lib.Exercises["normal-motion-basic-001"].E2EAssertions.Selection
	if selection == nil {
		t.Fatal("selection assertion is nil")
	}
	if !selection.Active || selection.Kind != "charwise" {
		t.Fatalf("selection = %+v, want active charwise", selection)
	}
	if selection.End == nil || selection.End.Col != 1 {
		t.Fatalf("selection end = %+v, want col 1", selection.End)
	}
}

func TestLoadLibraryRejectsReplayPassWhenSelectionAssertionMissesActualState(t *testing.T) {
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
      cursor: {row: 0, col: 1}
    optimal_keys: ["v", "l"]
    allowed_keys: ["v", "l"]
    grading:
      pass_condition: "cursor.row == 0 && cursor.col == 1"
      optimal_key_count: 2
    e2e_assertions:
      buffer: ["abc"]
      cursor: {row: 0, col: 1}
      mode: visual
      status: succeeded
      selection:
        active: true
        kind: charwise
        anchor: {row: 0, col: 0}
        head: {row: 0, col: 1}
        start: {row: 0, col: 0}
        end: {row: 0, col: 2}
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "selection") {
		t.Fatalf("LoadLibrary error = %v, want selection mismatch", err)
	}
}

func TestLoadLibraryRejectsReplayPassWhenSelectionAssertionRequiresActiveSelection(t *testing.T) {
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
    e2e_assertions:
      buffer: ["abc"]
      cursor: {row: 0, col: 1}
      mode: normal
      status: succeeded
      selection:
        active: true
        kind: charwise
        anchor: {row: 0, col: 0}
        head: {row: 0, col: 1}
        start: {row: 0, col: 0}
        end: {row: 0, col: 1}
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "selection") {
		t.Fatalf("LoadLibrary error = %v, want selection missing", err)
	}
}

func TestLoadLibraryPreservesHintThresholds(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	hints := lib.Exercises["word-motion-basic-001"].Hints
	if len(hints) != 2 {
		t.Fatalf("len(hints) = %d, want 2", len(hints))
	}
	if hints[0].AfterKeys != 1 || hints[1].AfterKeys != 2 {
		t.Fatalf("hint after_keys = %d,%d, want 1,2", hints[0].AfterKeys, hints[1].AfterKeys)
	}
}

func TestLoadLibraryPreservesExerciseConstraints(t *testing.T) {
	lib, err := LoadLibrary(filepath.Join("..", "..", "content"))
	if err != nil {
		t.Fatalf("LoadLibrary returned error: %v", err)
	}

	exercise := lib.Exercises["normal-motion-basic-001"]
	if exercise.Constraints.MaxInputs != 2 {
		t.Fatalf("max inputs = %d, want 2", exercise.Constraints.MaxInputs)
	}
	assertStrings(t, exercise.Constraints.RequiredKeys, []string{"l"})
	assertStrings(t, exercise.Constraints.ForbiddenKeys, []string{"w"})

	compiled, err := lib.CompileExercise("normal-motion-basic-001")
	if err != nil {
		t.Fatalf("CompileExercise returned error: %v", err)
	}
	if compiled.Exercise.Constraints.MaxInputs != 2 {
		t.Fatalf("compiled max inputs = %d, want 2", compiled.Exercise.Constraints.MaxInputs)
	}
	assertStrings(t, compiled.Exercise.Constraints.RequiredKeys, []string{"l"})
	assertStrings(t, compiled.Exercise.Constraints.ForbiddenKeys, []string{"right", "left", "up", "down", "w"})

	word, err := lib.CompileExercise("word-motion-basic-001")
	if err != nil {
		t.Fatalf("CompileExercise word returned error: %v", err)
	}
	if word.Exercise.Constraints.MaxInputs != 2 {
		t.Fatalf("word max inputs = %d, want 2", word.Exercise.Constraints.MaxInputs)
	}
	assertStrings(t, word.Exercise.Constraints.RequiredKeys, []string{"w"})
	assertStrings(t, word.Exercise.Constraints.ForbiddenKeys, []string{"right", "left", "up", "down", "h", "l"})
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

func TestLoadLibraryRejectsApprovedPlaylistOverTutorialLimit(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "playlists", "playlists.yaml"), `
playlists:
  - id: too-long
    status: approved
    title: Too long
    beats:
      - id: beat-1
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
      - id: beat-2
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
      - id: beat-3
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
      - id: beat-4
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
      - id: beat-5
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
      - id: beat-6
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
      - id: beat-7
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
      - id: beat-8
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
      - id: beat-9
        command_cluster: normal-motion-basic
        exercise_id: normal-motion-basic-001
        scenario_id: normal-motion-basic-001-scenario
        engine_support: implemented
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "want at most 8") {
		t.Fatalf("LoadLibrary error = %v, want tutorial limit", err)
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

func TestLoadLibraryRejectsNegativeConstraintValues(t *testing.T) {
	root := validLibraryFixture(t)
	writeYAML(t, filepath.Join(root, "exercises", "exercises.yaml"), `
exercises:
  - id: bad-constraint
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
    constraints:
      max_inputs: -1
    grading:
      pass_condition: "cursor.row == 0 && cursor.col == 1"
      optimal_key_count: 1
`)

	_, err := LoadLibrary(root)
	if err == nil || !strings.Contains(err.Error(), "constraints.max_inputs") {
		t.Fatalf("LoadLibrary error = %v, want constraints.max_inputs", err)
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

func assertPlayableIDs(t *testing.T, exercises []ExerciseDocument, want []string) {
	t.Helper()

	got := make([]string, 0, len(exercises))
	for _, exercise := range exercises {
		got = append(got, exercise.ID)
	}
	assertStrings(t, got, want)
}
