package runtime

import (
	"testing"

	"github.com/young-st511/advimture/internal/vimengine"
)

func TestSessionSucceedsWhenGoalMatches(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "move-to-target",
		Initial: vimengine.NewState([]string{"abc"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 2),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
	})

	result := session.ApplyKey(vimengine.KeyL)
	if result.State.Status != StatusRunning {
		t.Fatalf("status after first key = %q, want %q", result.State.Status, StatusRunning)
	}

	result = session.ApplyKey(vimengine.KeyL)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after second key = %q, want %q", result.State.Status, StatusSucceeded)
	}
	if result.State.Vim.Cursor.Row != 0 || result.State.Vim.Cursor.Col != 2 {
		t.Fatalf("cursor = (%d,%d), want (0,2)", result.State.Vim.Cursor.Row, result.State.Vim.Cursor.Col)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyL, vimengine.KeyL})
}

func TestSessionKeepsRunningBeforeGoalMatches(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "move-one-more",
		Initial: vimengine.NewState([]string{"abcd"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 3),
		},
	})

	result := session.ApplyKey(vimengine.KeyL)
	if result.State.Status != StatusRunning {
		t.Fatalf("status = %q, want %q", result.State.Status, StatusRunning)
	}
	if result.MatchedGoal {
		t.Fatal("MatchedGoal = true, want false")
	}
}

func TestSessionSucceedsWhenCommandGoalMatches(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "write-quit",
		Initial: vimengine.NewState([]string{"draft"}),
		Goal: Goal{
			Command: CommandGoal(":wq"),
		},
	})

	session.ApplyKey(vimengine.KeyColon)
	session.ApplyKey("w")
	session.ApplyKey("q")
	result := session.ApplyKey(vimengine.KeyEnter)

	if result.State.Status != StatusSucceeded {
		t.Fatalf("status = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyColon, "w", "q", vimengine.KeyEnter})
}

func TestUnsupportedKeyIsRecorded(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "unsupported",
		Initial: vimengine.NewState([]string{"abc"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 1),
		},
	})

	result := session.ApplyKey("z")
	if result.State.Status != StatusRunning {
		t.Fatalf("status = %q, want %q", result.State.Status, StatusRunning)
	}
	assertTrace(t, result.State.KeyTrace, []string{"z"})
	if len(result.Vim.Events) != 1 || result.Vim.Events[0].Type != vimengine.EventUnsupportedKey {
		t.Fatalf("vim events = %+v, want unsupported key event", result.Vim.Events)
	}
}

func TestSessionFailsForbiddenInputWithoutMoving(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "forbidden",
		Initial: vimengine.NewState([]string{"abc"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 2),
		},
		Constraints: Constraints{
			ForbiddenKeys: []string{vimengine.KeyW},
		},
	})

	result := session.ApplyKey(vimengine.KeyW)

	if result.State.Status != StatusFailed {
		t.Fatalf("status = %q, want %q", result.State.Status, StatusFailed)
	}
	if result.State.Failure != FailureForbiddenInput {
		t.Fatalf("failure = %q, want %q", result.State.Failure, FailureForbiddenInput)
	}
	if result.State.Vim.Cursor.Col != 0 {
		t.Fatalf("cursor col = %d, want unchanged 0", result.State.Vim.Cursor.Col)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyW})
}

func TestSessionFailsWhenMaxInputsExceededWithoutMoving(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "max-inputs",
		Initial: vimengine.NewState([]string{"abcd"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 3),
		},
		Constraints: Constraints{
			MaxInputs: 2,
		},
	})

	session.ApplyKey(vimengine.KeyH)
	session.ApplyKey(vimengine.KeyH)
	result := session.ApplyKey(vimengine.KeyL)

	if result.State.Status != StatusFailed {
		t.Fatalf("status = %q, want %q", result.State.Status, StatusFailed)
	}
	if result.State.Failure != FailureMaxInputsExceeded {
		t.Fatalf("failure = %q, want %q", result.State.Failure, FailureMaxInputsExceeded)
	}
	if result.State.Vim.Cursor.Col != 0 {
		t.Fatalf("cursor col = %d, want unchanged 0", result.State.Vim.Cursor.Col)
	}
	if result.State.InputsLeft != 0 {
		t.Fatalf("inputs left = %d, want 0", result.State.InputsLeft)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyH, vimengine.KeyH, vimengine.KeyL})
}

func TestSessionFailsWhenGoalReachedWithoutRequiredKey(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "required-key",
		Initial: vimengine.NewState([]string{"abc"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 2),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyL},
		},
	})

	result := session.ApplyKey(vimengine.KeyDollar)

	if result.State.Status != StatusFailed {
		t.Fatalf("status = %q, want %q", result.State.Status, StatusFailed)
	}
	if result.State.Failure != FailureRequiredKeysMissing {
		t.Fatalf("failure = %q, want %q", result.State.Failure, FailureRequiredKeysMissing)
	}
	if result.MatchedGoal {
		t.Fatal("MatchedGoal = true, want false")
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyDollar})
}

func TestSessionReplaysDeleteWithMotionTrace(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "delete-word",
		Initial: vimengine.NewState([]string{"alpha beta"}),
		Goal: Goal{
			Lines:  []string{"beta"},
			Cursor: CursorGoal(0, 0),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyD, vimengine.KeyW},
		},
	})

	result := session.ApplyKey(vimengine.KeyD)
	if result.State.Status != StatusRunning {
		t.Fatalf("status after d = %q, want running", result.State.Status)
	}

	result = session.ApplyKey(vimengine.KeyW)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after dw = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyD, vimengine.KeyW})
}

func TestSessionDoesNotStartSucceededWhenRequiredKeysAreMissing(t *testing.T) {
	initial := vimengine.NewState([]string{"api"})
	initial.Cursor.Col = 1
	initial.Cursor.DesiredCol = 1
	session := NewSession(Exercise{
		ID:      "undo-training",
		Initial: initial,
		Goal: Goal{
			Lines:  []string{"api"},
			Cursor: CursorGoal(0, 1),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyU},
		},
	})

	state := session.State()
	if state.Status != StatusRunning {
		t.Fatalf("status = %q, want %q", state.Status, StatusRunning)
	}

	session.ApplyKey(vimengine.KeyX)
	result := session.ApplyKey(vimengine.KeyU)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after undo = %q, want %q", result.State.Status, StatusSucceeded)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyX, vimengine.KeyU})
}

func TestRetryResetsSessionAndIncrementsAttempts(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "retry",
		Initial: vimengine.NewState([]string{"abc"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 2),
		},
	})

	session.ApplyKey(vimengine.KeyL)
	state := session.Retry()

	if state.Status != StatusRunning {
		t.Fatalf("status = %q, want %q", state.Status, StatusRunning)
	}
	if state.Attempts != 2 {
		t.Fatalf("attempts = %d, want 2", state.Attempts)
	}
	if state.Vim.Cursor.Col != 0 {
		t.Fatalf("cursor col = %d, want 0", state.Vim.Cursor.Col)
	}
	if len(state.KeyTrace) != 0 {
		t.Fatalf("key trace = %+v, want empty", state.KeyTrace)
	}
}

func TestRetryResetsFailedSession(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "retry-failed",
		Initial: vimengine.NewState([]string{"abc"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 2),
		},
		Constraints: Constraints{
			ForbiddenKeys: []string{vimengine.KeyW},
		},
	})

	session.ApplyKey(vimengine.KeyW)
	state := session.Retry()

	if state.Status != StatusRunning {
		t.Fatalf("status = %q, want %q", state.Status, StatusRunning)
	}
	if state.Failure != FailureNone {
		t.Fatalf("failure = %q, want none", state.Failure)
	}
	if state.Message != "" {
		t.Fatalf("message = %q, want empty", state.Message)
	}
	if state.Attempts != 2 {
		t.Fatalf("attempts = %d, want 2", state.Attempts)
	}
	if state.AttemptLimit != 0 {
		t.Fatalf("attempt limit = %d, want 0", state.AttemptLimit)
	}
}

func TestHintIsDeterministicByKeyTraceLength(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "hint",
		Initial: vimengine.NewState([]string{"abc"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 2),
		},
		Hints: []Hint{
			{AfterKeys: 1, Text: "move right"},
			{AfterKeys: 3, Text: "press l twice"},
		},
	})

	if _, ok := session.CurrentHint(); ok {
		t.Fatal("hint before input ok = true, want false")
	}

	session.ApplyKey(vimengine.KeyH)
	hint, ok := session.CurrentHint()
	if !ok || hint != "move right" {
		t.Fatalf("hint = (%q,%v), want (%q,true)", hint, ok, "move right")
	}

	session.ApplyKey(vimengine.KeyH)
	session.ApplyKey(vimengine.KeyH)
	hint, ok = session.CurrentHint()
	if !ok || hint != "press l twice" {
		t.Fatalf("hint = (%q,%v), want (%q,true)", hint, ok, "press l twice")
	}
}

func TestSessionStateIsCopied(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "copy",
		Initial: vimengine.NewState([]string{"abc"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 2),
		},
	})

	state := session.State()
	state.KeyTrace = append(state.KeyTrace, "mutate")
	state.Vim.Lines[0] = "changed"

	got := session.State()
	if len(got.KeyTrace) != 0 {
		t.Fatalf("key trace = %+v, want empty", got.KeyTrace)
	}
	if got.Vim.Lines[0] != "abc" {
		t.Fatalf("line = %q, want %q", got.Vim.Lines[0], "abc")
	}
}

func assertTrace(t *testing.T, got []string, want []string) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("trace len = %d, want %d: %+v", len(got), len(want), got)
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("trace[%d] = %q, want %q", index, got[index], want[index])
		}
	}
}
