package progressadapter

import (
	"testing"
	"time"

	"github.com/young-st511/advimture/internal/progress"
	exerciseruntime "github.com/young-st511/advimture/internal/runtime"
	"github.com/young-st511/advimture/internal/scenario"
	"github.com/young-st511/advimture/internal/scoring"
	"github.com/young-st511/advimture/internal/vimengine"
)

func TestMissionCompletionFromScenario(t *testing.T) {
	state := completedScenarioState()

	completion, err := MissionCompletionFromScenario("mission-1", state, 1500*time.Millisecond)
	if err != nil {
		t.Fatalf("MissionCompletionFromScenario returned error: %v", err)
	}

	if !completion.Completed {
		t.Fatal("Completed = false, want true")
	}
	if completion.MissionID != "mission-1" {
		t.Fatalf("MissionID = %q, want mission-1", completion.MissionID)
	}
	if completion.Grade != "S" {
		t.Fatalf("Grade = %q, want S", completion.Grade)
	}
	if completion.Keystrokes != 2 {
		t.Fatalf("Keystrokes = %d, want 2", completion.Keystrokes)
	}
	if completion.Attempts != 1 {
		t.Fatalf("Attempts = %d, want 1", completion.Attempts)
	}
	if completion.TimeMs != 1500 {
		t.Fatalf("TimeMs = %d, want 1500", completion.TimeMs)
	}
}

func TestMissionCompletionRejectsIncompleteState(t *testing.T) {
	state := completedScenarioState()
	state.Status = exerciseruntime.StatusRunning

	_, err := MissionCompletionFromScenario("mission-1", state, time.Second)
	if err == nil {
		t.Fatal("error = nil, want error")
	}
}

func TestMissionCompletionRejectsMissingScore(t *testing.T) {
	state := completedScenarioState()
	state.Score = nil

	_, err := MissionCompletionFromScenario("mission-1", state, time.Second)
	if err == nil {
		t.Fatal("error = nil, want error")
	}
}

func TestApplyMissionCompletionReturnsUpdatedCopy(t *testing.T) {
	original := progress.NewProgress()
	completion := MissionCompletion{
		MissionID:  "mission-1",
		Completed:  true,
		Grade:      "A",
		Keystrokes: 3,
		TimeMs:     1200,
		Attempts:   1,
	}

	updated := ApplyMissionCompletion(*original, completion)

	if _, ok := original.Missions["mission-1"]; ok {
		t.Fatal("original progress was mutated")
	}
	got := updated.Missions["mission-1"]
	if !got.Completed {
		t.Fatal("updated mission completed = false, want true")
	}
	if got.BestGrade != "A" {
		t.Fatalf("BestGrade = %q, want A", got.BestGrade)
	}
	if got.BestKeystrokes != 3 {
		t.Fatalf("BestKeystrokes = %d, want 3", got.BestKeystrokes)
	}
	if got.BestTimeMs != 1200 {
		t.Fatalf("BestTimeMs = %d, want 1200", got.BestTimeMs)
	}
}

func completedScenarioState() scenario.State {
	score := scoring.Result{
		Passed: true,
		Grade:  scoring.GradeS,
	}
	return scenario.State{
		ScenarioID: "scenario-1",
		Status:     exerciseruntime.StatusSucceeded,
		Runtime: exerciseruntime.State{
			Status:   exerciseruntime.StatusSucceeded,
			KeyTrace: []string{vimengine.KeyL, vimengine.KeyL},
			Attempts: 1,
		},
		Score: &score,
	}
}
