package data

import "testing"

func TestLoadTutorial(t *testing.T) {
	td, err := LoadTutorial("t00_test.yaml")
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if td.ID != "0-0" {
		t.Errorf("expected id '0-0', got '%s'", td.ID)
	}
	if len(td.Substeps) != 2 {
		t.Errorf("expected 2 substeps, got %d", len(td.Substeps))
	}
	if td.Substeps[0].Goal.Type != "cursor_position" {
		t.Errorf("expected goal type 'cursor_position', got '%s'", td.Substeps[0].Goal.Type)
	}
}

func TestLoadAllTutorials(t *testing.T) {
	tutorials, err := LoadAllTutorials()
	if err != nil {
		t.Fatalf("load all failed: %v", err)
	}
	if len(tutorials) < 1 {
		t.Error("expected at least 1 tutorial")
	}
}

func TestGoalChecker_CursorPosition(t *testing.T) {
	gc := NewGoalChecker()
	goal := GoalData{Type: "cursor_position", Row: 0, Col: 4}
	if gc.CheckGoal(goal, "hello", 0, 4, "NORMAL") != true {
		t.Error("cursor_position should match")
	}
	if gc.CheckGoal(goal, "hello", 0, 3, "NORMAL") != false {
		t.Error("cursor_position should not match")
	}
}

func TestGoalChecker_CursorOnChar(t *testing.T) {
	gc := NewGoalChecker()
	goal := GoalData{Type: "cursor_on_char", Char: "o"}
	if gc.CheckGoal(goal, "hello", 0, 4, "NORMAL") != true {
		t.Error("cursor_on_char 'o' should match at col 4")
	}
}

func TestGoalChecker_TextMatch(t *testing.T) {
	gc := NewGoalChecker()
	goal := GoalData{Type: "text_match", Text: "hello vim"}
	if gc.CheckGoal(goal, "hello vim", 0, 0, "NORMAL") != true {
		t.Error("text_match should match")
	}
	if gc.CheckGoal(goal, "hello world", 0, 0, "NORMAL") != false {
		t.Error("text_match should not match")
	}
}

func TestGoalChecker_SaveQuit(t *testing.T) {
	gc := NewGoalChecker()
	goal := GoalData{Type: "save_quit", Text: "done"}

	if gc.CheckGoal(goal, "done", 0, 0, "NORMAL") {
		t.Error("save_quit should not pass without RecordSaveQuit")
	}

	gc.RecordSaveQuit()
	if !gc.CheckGoal(goal, "done", 0, 0, "NORMAL") {
		t.Error("save_quit should pass after RecordSaveQuit with matching text")
	}
}

func TestGoalChecker_ModeIs(t *testing.T) {
	gc := NewGoalChecker()
	goal := GoalData{Type: "mode_is", Mode: "INSERT"}
	if !gc.CheckGoal(goal, "", 0, 0, "INSERT") {
		t.Error("mode_is INSERT should match")
	}
	if gc.CheckGoal(goal, "", 0, 0, "NORMAL") {
		t.Error("mode_is INSERT should not match NORMAL")
	}
}

func TestGoalChecker_CommandUsed(t *testing.T) {
	gc := NewGoalChecker()
	gc.RecordCommand("dd")
	goal := GoalData{Type: "command_used", Cmd: "dd"}
	if !gc.CheckGoal(goal, "", 0, 0, "NORMAL") {
		t.Error("command_used 'dd' should match")
	}
}

func TestGoalChecker_Reset(t *testing.T) {
	gc := NewGoalChecker()
	gc.RecordCommand("dd")
	gc.RecordSaveQuit()
	gc.Reset()
	if gc.LastCommandUsed != "" || gc.SaveQuitCalled {
		t.Error("reset should clear state")
	}
}

// ─── Mission loader tests ─────────────────────────────────────────────────────

func TestLoadMission(t *testing.T) {
	m, err := LoadMission("m00_test.yaml")
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if m.ID != "m-00" {
		t.Errorf("expected id 'm-00', got '%s'", m.ID)
	}
	if m.OptimalKeystrokes != 4 {
		t.Errorf("expected optimal_keystrokes 4, got %d", m.OptimalKeystrokes)
	}
	if len(m.Tips) != 1 {
		t.Errorf("expected 1 tip, got %d", len(m.Tips))
	}
}

func TestLoadAllMissions(t *testing.T) {
	_, err := LoadAllMissions()
	if err != nil {
		t.Fatalf("load all missions failed: %v", err)
	}
}

func TestValidateMission_OK(t *testing.T) {
	m := &MissionData{
		ID: "m-01", Title: "test", InitialText: "a", ExpectedText: "b", OptimalKeystrokes: 5,
	}
	if err := ValidateMission(m); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestValidateMission_MissingID(t *testing.T) {
	m := &MissionData{Title: "test", InitialText: "a", ExpectedText: "b", OptimalKeystrokes: 5}
	if err := ValidateMission(m); err == nil {
		t.Error("expected error for missing ID")
	}
}

func TestValidateMission_ZeroOptimal(t *testing.T) {
	m := &MissionData{ID: "m-01", Title: "test", InitialText: "a", ExpectedText: "b", OptimalKeystrokes: 0}
	if err := ValidateMission(m); err == nil {
		t.Error("expected error for OptimalKeystrokes=0")
	}
}

func TestCompareText_Equal(t *testing.T) {
	ok, diffs, totalDiff := CompareText("hello\nworld", "hello\nworld")
	if !ok {
		t.Error("expected match")
	}
	if totalDiff != 0 {
		t.Errorf("expected 0 total diffs, got %d", totalDiff)
	}
	if len(diffs) != 0 {
		t.Errorf("expected 0 diffs, got %d", len(diffs))
	}
}

func TestCompareText_TrailingWhitespace(t *testing.T) {
	ok, _, _ := CompareText("hello  ", "hello")
	if !ok {
		t.Error("trailing whitespace should be ignored")
	}
}

func TestCompareText_LineMismatch(t *testing.T) {
	ok, diffs, totalDiff := CompareText("a\nb\nc", "a\nb")
	if ok {
		t.Error("expected mismatch for different line counts")
	}
	if totalDiff != 1 {
		t.Errorf("expected 1 totalDiff, got %d", totalDiff)
	}
	if len(diffs) != 1 {
		t.Errorf("expected 1 diff entry, got %d", len(diffs))
	}
}

func TestCompareText_MultipleDiffs(t *testing.T) {
	ok, diffs, totalDiff := CompareText("a\nb\nc\nd\ne", "a\nX\nX\nX\nX")
	if ok {
		t.Error("expected mismatch")
	}
	if totalDiff != 4 {
		t.Errorf("expected 4 total diffs, got %d", totalDiff)
	}
	if len(diffs) != 3 {
		t.Errorf("expected max 3 returned diffs, got %d", len(diffs))
	}
}
