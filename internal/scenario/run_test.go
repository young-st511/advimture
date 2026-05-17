package scenario

import (
	"testing"

	"github.com/young-st511/advimture/internal/content"
	exerciseruntime "github.com/young-st511/advimture/internal/runtime"
	"github.com/young-st511/advimture/internal/scoring"
	"github.com/young-st511/advimture/internal/vimengine"
)

func TestRunStartsWithBriefing(t *testing.T) {
	run := newTestRun(t)
	state := run.State()

	if state.Status != exerciseruntime.StatusRunning {
		t.Fatalf("status = %q, want %q", state.Status, exerciseruntime.StatusRunning)
	}
	if state.Message != "Reach the marked column." {
		t.Fatalf("message = %q, want briefing", state.Message)
	}
	if state.Score != nil {
		t.Fatalf("score = %+v, want nil", state.Score)
	}
}

func TestRunScoresOnSuccess(t *testing.T) {
	run := newTestRun(t)

	run.ApplyKey(vimengine.KeyL)
	result := run.ApplyKey(vimengine.KeyL)

	if result.State.Status != exerciseruntime.StatusSucceeded {
		t.Fatalf("status = %q, want %q", result.State.Status, exerciseruntime.StatusSucceeded)
	}
	if result.State.Message != "Door opened." {
		t.Fatalf("message = %q, want success", result.State.Message)
	}
	if result.State.Score == nil {
		t.Fatal("score = nil, want result")
	}
	if result.State.Score.Grade != scoring.GradeS {
		t.Fatalf("grade = %q, want %q", result.State.Score.Grade, scoring.GradeS)
	}
}

func TestRequestHintRecordsHintUse(t *testing.T) {
	run := newTestRun(t)

	if _, ok := run.RequestHint(); ok {
		t.Fatal("hint before input ok = true, want false")
	}
	run.ApplyKey(vimengine.KeyH)

	hint, ok := run.RequestHint()
	if !ok || hint != "Use l twice." {
		t.Fatalf("hint = (%q,%v), want (%q,true)", hint, ok, "Use l twice.")
	}
	if run.State().HintsUsed != 1 {
		t.Fatalf("HintsUsed = %d, want 1", run.State().HintsUsed)
	}
}

func TestHintPenaltyAffectsScore(t *testing.T) {
	run := newTestRun(t)

	run.ApplyKey(vimengine.KeyL)
	run.RequestHint()
	result := run.ApplyKey(vimengine.KeyL)

	if result.State.Score == nil {
		t.Fatal("score = nil, want result")
	}
	if result.State.Score.Grade != scoring.GradeA {
		t.Fatalf("grade = %q, want %q", result.State.Score.Grade, scoring.GradeA)
	}
}

func TestRetryResetsRunState(t *testing.T) {
	run := newTestRun(t)

	run.ApplyKey(vimengine.KeyL)
	run.ApplyKey(vimengine.KeyL)
	state := run.Retry()

	if state.Status != exerciseruntime.StatusRunning {
		t.Fatalf("status = %q, want %q", state.Status, exerciseruntime.StatusRunning)
	}
	if state.Message != "Reach the marked column." {
		t.Fatalf("message = %q, want briefing", state.Message)
	}
	if state.Score != nil {
		t.Fatalf("score = %+v, want nil", state.Score)
	}
	if state.Runtime.Attempts != 2 {
		t.Fatalf("attempts = %d, want 2", state.Runtime.Attempts)
	}
}

func newTestRun(t *testing.T) *Run {
	t.Helper()

	compiled, err := content.CompileExercise(content.ExerciseSpec{
		ID:               "move-right",
		CommandClusterID: "normal-motion-basic",
		Initial: content.StateSpec{
			Lines: []string{"abc"},
		},
		Goal: content.GoalSpec{
			Cursor: content.CursorSpecPtr(0, 2),
			Mode:   "normal",
		},
		Hints: []content.HintSpec{
			{AfterKeys: 1, Text: "Use l twice."},
		},
		ExpectedKeys: []string{vimengine.KeyL, vimengine.KeyL},
	})
	if err != nil {
		t.Fatalf("CompileExercise returned error: %v", err)
	}

	run, err := NewRun(Spec{
		ID:          "door",
		Title:       "Open the door",
		Briefing:    "Reach the marked column.",
		SuccessText: "Door opened.",
		Exercise:    compiled,
	})
	if err != nil {
		t.Fatalf("NewRun returned error: %v", err)
	}
	return run
}
