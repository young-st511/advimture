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
