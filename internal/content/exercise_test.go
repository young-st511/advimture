package content

import (
	"testing"

	exerciseruntime "github.com/young-st511/advimture/internal/runtime"
	"github.com/young-st511/advimture/internal/vimengine"
)

func TestCompileExerciseProducesRuntimeExercise(t *testing.T) {
	compiled, err := CompileExercise(ExerciseSpec{
		ID:               "move-right",
		CommandClusterID: "normal-motion-basic",
		Title:            "Move right twice",
		Initial: StateSpec{
			Lines: []string{"abc"},
			Mode:  "normal",
		},
		Goal: GoalSpec{
			Cursor: CursorSpecPtr(0, 2),
			Mode:   "normal",
		},
		Hints: []HintSpec{
			{AfterKeys: 1, Text: "Use l to move right."},
		},
		Constraints: ConstraintSpec{
			MaxInputs:     2,
			RequiredKeys:  []string{vimengine.KeyL},
			ForbiddenKeys: []string{vimengine.KeyW},
		},
		ExpectedKeys: []string{vimengine.KeyL, vimengine.KeyL},
		AllowedKeys:  []string{vimengine.KeyH, vimengine.KeyJ, vimengine.KeyK, vimengine.KeyL},
	})
	if err != nil {
		t.Fatalf("CompileExercise returned error: %v", err)
	}

	session := exerciseruntime.NewSession(compiled.Exercise)
	session.ApplyKey(vimengine.KeyL)
	result := session.ApplyKey(vimengine.KeyL)

	if result.State.Status != exerciseruntime.StatusSucceeded {
		t.Fatalf("status = %q, want %q", result.State.Status, exerciseruntime.StatusSucceeded)
	}
	assertStrings(t, compiled.ExpectedKeys, []string{vimengine.KeyL, vimengine.KeyL})
	assertStrings(t, compiled.AllowedKeys, []string{vimengine.KeyH, vimengine.KeyJ, vimengine.KeyK, vimengine.KeyL})
	if compiled.Exercise.Constraints.MaxInputs != 2 {
		t.Fatalf("max inputs = %d, want 2", compiled.Exercise.Constraints.MaxInputs)
	}
	assertStrings(t, compiled.Exercise.Constraints.RequiredKeys, []string{vimengine.KeyL})
	assertStrings(t, compiled.Exercise.Constraints.ForbiddenKeys, []string{vimengine.KeyW})
}

func TestCompileExerciseProducesCommandGoal(t *testing.T) {
	compiled, err := CompileExercise(ExerciseSpec{
		ID: "write-quit",
		Initial: StateSpec{
			Lines: []string{"draft"},
			Mode:  "normal",
		},
		Goal: GoalSpec{
			Command: ":wq",
		},
		ExpectedKeys: []string{vimengine.KeyColon, "w", "q", vimengine.KeyEnter},
		AllowedKeys:  []string{vimengine.KeyColon, "w", "q", vimengine.KeyEnter},
	})
	if err != nil {
		t.Fatalf("CompileExercise returned error: %v", err)
	}

	session := exerciseruntime.NewSession(compiled.Exercise)
	session.ApplyKey(vimengine.KeyColon)
	session.ApplyKey("w")
	session.ApplyKey("q")
	result := session.ApplyKey(vimengine.KeyEnter)

	if result.State.Status != exerciseruntime.StatusSucceeded {
		t.Fatalf("status = %q, want succeeded", result.State.Status)
	}
}

func TestCompileExerciseCopiesMetadata(t *testing.T) {
	spec := ExerciseSpec{
		ID: "copy",
		Initial: StateSpec{
			Lines: []string{"abc"},
		},
		Goal: GoalSpec{
			Cursor: CursorSpecPtr(0, 1),
		},
		ExpectedKeys: []string{vimengine.KeyL},
		AllowedKeys:  []string{vimengine.KeyL},
	}

	compiled, err := CompileExercise(spec)
	if err != nil {
		t.Fatalf("CompileExercise returned error: %v", err)
	}

	spec.ExpectedKeys[0] = "changed"
	spec.AllowedKeys[0] = "changed"
	compiled.ExpectedKeys[0] = "mutated"
	compiled.AllowedKeys[0] = "mutated"

	compiledAgain, err := CompileExercise(spec)
	if err != nil {
		t.Fatalf("CompileExercise returned error: %v", err)
	}

	assertStrings(t, compiledAgain.ExpectedKeys, []string{"changed"})
	assertStrings(t, compiledAgain.AllowedKeys, []string{"changed"})
}

func TestCompileExerciseRejectsMissingID(t *testing.T) {
	_, err := CompileExercise(ExerciseSpec{
		Initial: StateSpec{Lines: []string{"abc"}},
		Goal:    GoalSpec{Cursor: CursorSpecPtr(0, 1)},
	})
	if err == nil {
		t.Fatal("CompileExercise error = nil, want error")
	}
}

func TestCompileExerciseRejectsInvalidMode(t *testing.T) {
	_, err := CompileExercise(ExerciseSpec{
		ID: "invalid-mode",
		Initial: StateSpec{
			Lines: []string{"abc"},
			Mode:  "visual",
		},
		Goal: GoalSpec{Cursor: CursorSpecPtr(0, 1)},
	})
	if err == nil {
		t.Fatal("CompileExercise error = nil, want error")
	}
}

func TestCompileExerciseRejectsOutOfRangeInitialCursor(t *testing.T) {
	_, err := CompileExercise(ExerciseSpec{
		ID: "bad-cursor",
		Initial: StateSpec{
			Lines:  []string{"abc"},
			Cursor: CursorSpecPtr(1, 0),
		},
		Goal: GoalSpec{Cursor: CursorSpecPtr(0, 1)},
	})
	if err == nil {
		t.Fatal("CompileExercise error = nil, want error")
	}
}

func TestCompileExerciseRejectsOutOfRangeGoalCursor(t *testing.T) {
	_, err := CompileExercise(ExerciseSpec{
		ID:      "bad-goal",
		Initial: StateSpec{Lines: []string{"abc"}},
		Goal: GoalSpec{
			Cursor: CursorSpecPtr(0, 9),
		},
	})
	if err == nil {
		t.Fatal("CompileExercise error = nil, want error")
	}
}

func TestCompileExerciseRequiresAtLeastOneGoal(t *testing.T) {
	_, err := CompileExercise(ExerciseSpec{
		ID:      "no-goal",
		Initial: StateSpec{Lines: []string{"abc"}},
	})
	if err == nil {
		t.Fatal("CompileExercise error = nil, want error")
	}
}

func TestCompileExerciseRejectsNegativeConstraints(t *testing.T) {
	_, err := CompileExercise(ExerciseSpec{
		ID:      "bad-constraints",
		Initial: StateSpec{Lines: []string{"abc"}},
		Goal:    GoalSpec{Cursor: CursorSpecPtr(0, 1)},
		Constraints: ConstraintSpec{
			MaxInputs: -1,
		},
	})
	if err == nil {
		t.Fatal("CompileExercise error = nil, want error")
	}
}

func assertStrings(t *testing.T, got []string, want []string) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d: %+v", len(got), len(want), got)
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("value[%d] = %q, want %q", index, got[index], want[index])
		}
	}
}
