package e2estate

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteStateCreatesSummaryFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".advimture", "e2e_state.json")
	state := State{
		Buffer: []string{"abc"},
		Cursor: Cursor{
			Row: 0,
			Col: 2,
		},
		Mode:   "normal",
		Status: "succeeded",
		Score: Score{
			Grade:  "S",
			Passed: true,
		},
		Progress: Progress{
			MissionID: "mission-1",
			Completed: true,
		},
		Review: Review{
			QueueCount:        2,
			PrimaryExerciseID: "mission-2",
			PrimaryReason:     "incomplete",
			DailyRoute:        "오늘의 복구 루트: 2건 대기",
		},
		UI: UI{
			FocusPanel: FocusPanel{
				Kind:    "success",
				Title:   "STEP SEALED",
				Lines:   []string{"이번 복구: grade S, 2 keys"},
				Actions: []ActionLine{{ID: "next", Label: "다음 단계: enter"}},
			},
		},
		Selection: &Selection{
			Active: true,
			Kind:   "charwise",
			Anchor: Cursor{Row: 0, Col: 1},
			Head:   Cursor{Row: 0, Col: 2},
			Start:  Cursor{Row: 0, Col: 1},
			End:    Cursor{Row: 0, Col: 2},
		},
	}

	if err := Write(path, state); err != nil {
		t.Fatalf("Write returned error: %v", err)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var got State
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	if got.Cursor.Col != 2 {
		t.Fatalf("cursor col = %d, want 2", got.Cursor.Col)
	}
	if got.Score.Grade != "S" {
		t.Fatalf("grade = %q, want S", got.Score.Grade)
	}
	if got.Selection == nil || got.Selection.Kind != "charwise" || got.Selection.End.Col != 2 {
		t.Fatalf("selection = %+v, want charwise end col 2", got.Selection)
	}
	if got.Review.QueueCount != 2 || got.Review.PrimaryExerciseID != "mission-2" {
		t.Fatalf("review = %+v, want queue count 2 and primary mission-2", got.Review)
	}
	if got.UI.FocusPanel.Kind != "success" || got.UI.FocusPanel.Title != "STEP SEALED" {
		t.Fatalf("ui focus panel = %+v, want success STEP SEALED", got.UI.FocusPanel)
	}
	if len(got.UI.FocusPanel.Actions) != 1 || got.UI.FocusPanel.Actions[0].ID != "next" {
		t.Fatalf("ui focus panel actions = %+v, want next action", got.UI.FocusPanel.Actions)
	}
}
