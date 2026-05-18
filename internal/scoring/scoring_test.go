package scoring

import (
	"testing"

	exerciseruntime "github.com/young-st511/advimture/internal/runtime"
	"github.com/young-st511/advimture/internal/vimengine"
)

func TestEvaluatePerfectRun(t *testing.T) {
	result := Evaluate(Input{
		Status:       exerciseruntime.StatusSucceeded,
		KeyTrace:     []string{vimengine.KeyL, vimengine.KeyL},
		ExpectedKeys: []string{vimengine.KeyL, vimengine.KeyL},
		Attempts:     1,
		HintsUsed:    0,
	})

	if !result.Passed {
		t.Fatal("Passed = false, want true")
	}
	if !result.IntentSatisfied {
		t.Fatal("IntentSatisfied = false, want true")
	}
	if !result.ExactKeys {
		t.Fatal("ExactKeys = false, want true")
	}
	if result.Grade != GradeS {
		t.Fatalf("Grade = %q, want %q", result.Grade, GradeS)
	}
	if result.Efficiency != 1 {
		t.Fatalf("Efficiency = %v, want 1", result.Efficiency)
	}
}

func TestEvaluateFailsWhenRuntimeDidNotSucceed(t *testing.T) {
	result := Evaluate(Input{
		Status:       exerciseruntime.StatusRunning,
		KeyTrace:     []string{vimengine.KeyL, vimengine.KeyL},
		ExpectedKeys: []string{vimengine.KeyL, vimengine.KeyL},
		Attempts:     1,
	})

	if result.Passed {
		t.Fatal("Passed = true, want false")
	}
	if result.Grade != GradeF {
		t.Fatalf("Grade = %q, want %q", result.Grade, GradeF)
	}
}

func TestEvaluateRequiredKeyFailureMarksIntentUnsatisfied(t *testing.T) {
	result := Evaluate(Input{
		Status:       exerciseruntime.StatusFailed,
		Failure:      exerciseruntime.FailureRequiredKeysMissing,
		KeyTrace:     []string{vimengine.KeyDollar},
		ExpectedKeys: []string{vimengine.KeyL, vimengine.KeyL},
		Attempts:     1,
	})

	if result.Passed {
		t.Fatal("Passed = true, want false")
	}
	if result.IntentSatisfied {
		t.Fatal("IntentSatisfied = true, want false")
	}
	if result.Failure != exerciseruntime.FailureRequiredKeysMissing {
		t.Fatalf("Failure = %q, want %q", result.Failure, exerciseruntime.FailureRequiredKeysMissing)
	}
	if result.Grade != GradeF {
		t.Fatalf("Grade = %q, want %q", result.Grade, GradeF)
	}
}

func TestEvaluateExtraKeysLowerEfficiencyAndGrade(t *testing.T) {
	result := Evaluate(Input{
		Status:       exerciseruntime.StatusSucceeded,
		KeyTrace:     []string{vimengine.KeyH, vimengine.KeyL, vimengine.KeyL},
		ExpectedKeys: []string{vimengine.KeyL, vimengine.KeyL},
		Attempts:     1,
	})

	if result.ExactKeys {
		t.Fatal("ExactKeys = true, want false")
	}
	if result.Efficiency != float64(2)/float64(3) {
		t.Fatalf("Efficiency = %v, want 2/3", result.Efficiency)
	}
	if result.Grade != GradeB {
		t.Fatalf("Grade = %q, want %q", result.Grade, GradeB)
	}
}

func TestEvaluateHintAndRetryPenalty(t *testing.T) {
	result := Evaluate(Input{
		Status:       exerciseruntime.StatusSucceeded,
		KeyTrace:     []string{vimengine.KeyL, vimengine.KeyL},
		ExpectedKeys: []string{vimengine.KeyL, vimengine.KeyL},
		Attempts:     2,
		HintsUsed:    1,
	})

	if result.Grade != GradeB {
		t.Fatalf("Grade = %q, want %q", result.Grade, GradeB)
	}
	if result.Penalty != 0.25 {
		t.Fatalf("Penalty = %v, want 0.25", result.Penalty)
	}
}

func TestEvaluateWithoutExpectedKeysStillPassesSucceededState(t *testing.T) {
	result := Evaluate(Input{
		Status:   exerciseruntime.StatusSucceeded,
		KeyTrace: []string{vimengine.KeyL},
		Attempts: 1,
	})

	if !result.Passed {
		t.Fatal("Passed = false, want true")
	}
	if !result.ExactKeys {
		t.Fatal("ExactKeys = false, want true when expected keys are omitted")
	}
	if result.Efficiency != 1 {
		t.Fatalf("Efficiency = %v, want 1", result.Efficiency)
	}
}

func TestEvaluateCopiesInputSlices(t *testing.T) {
	input := Input{
		Status:       exerciseruntime.StatusSucceeded,
		KeyTrace:     []string{vimengine.KeyL},
		ExpectedKeys: []string{vimengine.KeyL},
		Attempts:     1,
	}

	result := Evaluate(input)
	input.KeyTrace[0] = "changed"
	input.ExpectedKeys[0] = "changed"

	if result.KeyCount != 1 || result.ExpectedKeyCount != 1 {
		t.Fatalf("counts = (%d,%d), want (1,1)", result.KeyCount, result.ExpectedKeyCount)
	}
}
