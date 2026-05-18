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

func TestUnsupportedKeyIsRecorded(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "unsupported",
		Initial: vimengine.NewState([]string{"abc"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 1),
		},
	})

	result := session.ApplyKey("x")
	if result.State.Status != StatusRunning {
		t.Fatalf("status = %q, want %q", result.State.Status, StatusRunning)
	}
	assertTrace(t, result.State.KeyTrace, []string{"x"})
	if len(result.Vim.Events) != 1 || result.Vim.Events[0].Type != vimengine.EventUnsupportedKey {
		t.Fatalf("vim events = %+v, want unsupported key event", result.Vim.Events)
	}
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
