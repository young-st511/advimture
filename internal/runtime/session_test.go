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

func TestSessionFailsWhenDiscardGoalReceivesWriteQuit(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "discard",
		Initial: vimengine.NewState([]string{"draft"}),
		Goal: Goal{
			Command: CommandGoal(":q!"),
		},
	})

	session.ApplyKey(vimengine.KeyColon)
	session.ApplyKey("w")
	session.ApplyKey("q")
	result := session.ApplyKey(vimengine.KeyEnter)

	if result.State.Status != StatusFailed {
		t.Fatalf("status = %q, want failed", result.State.Status)
	}
	if result.State.Message != ":wq가 아니라 :q!로 버리고 나가세요." {
		t.Fatalf("message = %q, want discard command mismatch feedback", result.State.Message)
	}
}

func TestSessionFailsWhenWriteGoalReceivesDiscardQuit(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "write",
		Initial: vimengine.NewState([]string{"draft"}),
		Goal: Goal{
			Command: CommandGoal(":wq"),
		},
	})

	session.ApplyKey(vimengine.KeyColon)
	session.ApplyKey("q")
	session.ApplyKey("!")
	result := session.ApplyKey(vimengine.KeyEnter)

	if result.State.Status != StatusFailed {
		t.Fatalf("status = %q, want failed", result.State.Status)
	}
	if result.State.Message != ":q!가 아니라 :wq로 저장 후 나가세요." {
		t.Fatalf("message = %q, want write command mismatch feedback", result.State.Message)
	}
}

func TestSessionReplaysLiteralSearchTrace(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "search-timeout",
		Initial: vimengine.NewState([]string{"info ok", "warn timeout", "error timeout"}),
		Goal: Goal{
			Cursor: CursorGoal(1, 5),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeySlash, vimengine.KeyEnter},
		},
	})

	for _, key := range []string{vimengine.KeySlash, "t", "i", "m", "e", "o", "u", "t"} {
		result := session.ApplyKey(key)
		if result.State.Status != StatusRunning {
			t.Fatalf("status after %q = %q, want running", key, result.State.Status)
		}
	}
	result := session.ApplyKey(vimengine.KeyEnter)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after enter = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeySlash, "t", "i", "m", "e", "o", "u", "t", vimengine.KeyEnter})
}

func TestSessionSucceedsWithVisualDeleteTrace(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "visual-delete",
		Initial: vimengine.NewState([]string{"abcdef"}),
		Goal: Goal{
			Lines:  []string{"aef"},
			Cursor: CursorGoal(0, 1),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyV, vimengine.KeyD},
		},
	})

	for _, key := range []string{vimengine.KeyL, vimengine.KeyV, vimengine.KeyL, vimengine.KeyL} {
		result := session.ApplyKey(key)
		if result.State.Status != StatusRunning {
			t.Fatalf("status after %q = %q, want running", key, result.State.Status)
		}
	}
	result := session.ApplyKey(vimengine.KeyD)

	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after d = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyL, vimengine.KeyV, vimengine.KeyL, vimengine.KeyL, vimengine.KeyD})
}

func TestSessionSucceedsWithLinewiseVisualDeleteTrace(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "visual-line-delete",
		Initial: vimengine.NewState([]string{"drop", "drop", "keep"}),
		Goal: Goal{
			Lines:  []string{"keep"},
			Cursor: CursorGoal(0, 0),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyShiftV, vimengine.KeyD},
		},
	})

	for _, key := range []string{vimengine.KeyShiftV, vimengine.KeyJ} {
		result := session.ApplyKey(key)
		if result.State.Status != StatusRunning {
			t.Fatalf("status after %q = %q, want running", key, result.State.Status)
		}
	}
	result := session.ApplyKey(vimengine.KeyD)

	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after d = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyShiftV, vimengine.KeyJ, vimengine.KeyD})
}

func TestSessionKeepsRunningWhenVisualOperatorUnsupported(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "visual-multiline-unsupported",
		Initial: vimengine.NewState([]string{"abc", "def"}),
		Goal: Goal{
			Lines: []string{"unsupported target"},
		},
	})

	for _, key := range []string{vimengine.KeyV, vimengine.KeyJ} {
		result := session.ApplyKey(key)
		if result.State.Status != StatusRunning {
			t.Fatalf("status after %q = %q, want running", key, result.State.Status)
		}
	}
	result := session.ApplyKey(vimengine.KeyD)

	if result.State.Status != StatusRunning {
		t.Fatalf("status after d = %q, want running", result.State.Status)
	}
	if len(result.State.Vim.Lines) != 2 || result.State.Vim.Lines[0] != "abc" || result.State.Vim.Lines[1] != "def" {
		t.Fatalf("lines = %+v, want [abc def]", result.State.Vim.Lines)
	}
	if result.State.Vim.Mode != vimengine.ModeVisual || result.State.Vim.Selection == nil {
		t.Fatalf("mode/selection = %s/%+v, want visual selection preserved", result.State.Vim.Mode, result.State.Vim.Selection)
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

func TestSessionKeepsRunningWhenGoalReachedWithoutRequiredKeyAndInputsRemain(t *testing.T) {
	initial := vimengine.NewState([]string{"api"})
	initial.Cursor.Col = 1
	initial.Cursor.DesiredCol = 1
	session := NewSession(Exercise{
		ID:      "redo-training",
		Initial: initial,
		Goal: Goal{
			Lines:  []string{"ai"},
			Cursor: CursorGoal(0, 1),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			MaxInputs:    3,
			RequiredKeys: []string{vimengine.KeyCtrlR},
		},
	})

	result := session.ApplyKey(vimengine.KeyX)
	if result.State.Status != StatusRunning {
		t.Fatalf("status after x = %q, want running", result.State.Status)
	}
	if result.State.Failure != FailureNone {
		t.Fatalf("failure after x = %q, want none", result.State.Failure)
	}
	if result.MatchedGoal {
		t.Fatal("MatchedGoal after x = true, want false until required key is used")
	}

	result = session.ApplyKey(vimengine.KeyU)
	if result.State.Status != StatusRunning {
		t.Fatalf("status after u = %q, want running", result.State.Status)
	}

	result = session.ApplyKey(vimengine.KeyCtrlR)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after ctrl+r = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyX, vimengine.KeyU, vimengine.KeyCtrlR})
}

func TestSessionFailsWhenRequiredKeyIsPressedAfterShortcutWithoutLeavingGoal(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "shortcut-then-required",
		Initial: vimengine.NewState([]string{"abc"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 2),
		},
		Constraints: Constraints{
			MaxInputs:    2,
			RequiredKeys: []string{vimengine.KeyL},
		},
	})

	result := session.ApplyKey(vimengine.KeyDollar)
	if result.State.Status != StatusRunning {
		t.Fatalf("status after shortcut = %q, want running", result.State.Status)
	}

	result = session.ApplyKey(vimengine.KeyL)
	if result.State.Status != StatusFailed {
		t.Fatalf("status after l = %q, want failed", result.State.Status)
	}
	if result.State.Failure != FailureRequiredKeysMissing {
		t.Fatalf("failure = %q, want %q", result.State.Failure, FailureRequiredKeysMissing)
	}
}

func TestSessionFailsMissingRequiredKeyWhenInputBudgetIsSpent(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "required-key-budget",
		Initial: vimengine.NewState([]string{"abc"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 2),
		},
		Constraints: Constraints{
			MaxInputs:    1,
			RequiredKeys: []string{vimengine.KeyL},
		},
	})

	result := session.ApplyKey(vimengine.KeyDollar)
	if result.State.Status != StatusFailed {
		t.Fatalf("status = %q, want failed", result.State.Status)
	}
	if result.State.Failure != FailureRequiredKeysMissing {
		t.Fatalf("failure = %q, want %q", result.State.Failure, FailureRequiredKeysMissing)
	}
}

func TestStateExposesRequiredKeysForCoaching(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "coaching",
		Initial: vimengine.NewState([]string{"abc"}),
		Goal: Goal{
			Cursor: CursorGoal(0, 1),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyL},
		},
	})

	state := session.State()
	state.RequiredKeys[0] = "mutated"

	next := session.State()
	if got := next.RequiredKeys; len(got) != 1 || got[0] != vimengine.KeyL {
		t.Fatalf("RequiredKeys = %#v, want copy of [l]", got)
	}
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

func TestSessionReplaysChangeWithMotionTrace(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "change-word",
		Initial: vimengine.NewState([]string{"alpha beta"}),
		Goal: Goal{
			Lines:  []string{"omega beta"},
			Cursor: CursorGoal(0, 6),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyC, vimengine.KeyW, vimengine.KeyEsc},
		},
	})

	for _, key := range []string{vimengine.KeyC, vimengine.KeyW, "o", "m", "e", "g", "a", " "} {
		result := session.ApplyKey(key)
		if result.State.Status != StatusRunning {
			t.Fatalf("status after %q = %q, want running", key, result.State.Status)
		}
	}
	result := session.ApplyKey(vimengine.KeyEsc)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after esc = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyC, vimengine.KeyW, "o", "m", "e", "g", "a", " ", vimengine.KeyEsc})
}

func TestSessionReplaysYankPutTrace(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "yank-put-line",
		Initial: vimengine.NewState([]string{"one", "two"}),
		Goal: Goal{
			Lines:  []string{"one", "one", "two"},
			Cursor: CursorGoal(1, 0),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyY, vimengine.KeyP},
		},
	})

	for _, key := range []string{vimengine.KeyY, vimengine.KeyY} {
		result := session.ApplyKey(key)
		if result.State.Status != StatusRunning {
			t.Fatalf("status after %q = %q, want running", key, result.State.Status)
		}
	}

	result := session.ApplyKey(vimengine.KeyP)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after p = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyY, vimengine.KeyY, vimengine.KeyP})
}

func TestSessionReplaysOpenLineBelowTrace(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "open-line-below",
		Initial: vimengine.NewState([]string{"alpha", "omega"}),
		Goal: Goal{
			Lines:  []string{"alpha", "guard", "omega"},
			Cursor: CursorGoal(1, 4),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyO, vimengine.KeyEsc},
		},
	})

	for _, key := range []string{vimengine.KeyO, "g", "u", "a", "r", "d"} {
		result := session.ApplyKey(key)
		if result.State.Status != StatusRunning {
			t.Fatalf("status after %q = %q, want running", key, result.State.Status)
		}
	}
	result := session.ApplyKey(vimengine.KeyEsc)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after esc = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyO, "g", "u", "a", "r", "d", vimengine.KeyEsc})
}

func TestSessionReplaysOpenLineAboveTrace(t *testing.T) {
	initial := vimengine.NewState([]string{"alpha", "omega"})
	initial.Cursor.Row = 1
	initial.Cursor.Col = 2
	initial.Cursor.DesiredCol = 2
	session := NewSession(Exercise{
		ID:      "open-line-above",
		Initial: initial,
		Goal: Goal{
			Lines:  []string{"alpha", "guard", "omega"},
			Cursor: CursorGoal(1, 4),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyShiftO, vimengine.KeyEsc},
		},
	})

	for _, key := range []string{vimengine.KeyShiftO, "g", "u", "a", "r", "d"} {
		result := session.ApplyKey(key)
		if result.State.Status != StatusRunning {
			t.Fatalf("status after %q = %q, want running", key, result.State.Status)
		}
	}
	result := session.ApplyKey(vimengine.KeyEsc)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after esc = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyShiftO, "g", "u", "a", "r", "d", vimengine.KeyEsc})
}

func TestSessionReplaysRepeatLastChangeTrace(t *testing.T) {
	session := NewSession(Exercise{
		ID:      "repeat-last-change",
		Initial: vimengine.NewState([]string{"api", "api"}),
		Goal: Goal{
			Lines:  []string{"api!", "api!"},
			Cursor: CursorGoal(1, 3),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyShiftA, vimengine.KeyEsc, vimengine.KeyDot},
		},
	})

	for _, key := range []string{vimengine.KeyShiftA, "!", vimengine.KeyEsc, vimengine.KeyJ} {
		result := session.ApplyKey(key)
		if result.State.Status != StatusRunning {
			t.Fatalf("status after %q = %q, want running", key, result.State.Status)
		}
	}
	result := session.ApplyKey(vimengine.KeyDot)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after dot = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyShiftA, "!", vimengine.KeyEsc, vimengine.KeyJ, vimengine.KeyDot})
}

func TestSessionReplaysDeleteInnerWordTrace(t *testing.T) {
	initial := vimengine.NewState([]string{"mode=broken"})
	initial.Cursor.Col = 7
	initial.Cursor.DesiredCol = 7
	session := NewSession(Exercise{
		ID:      "delete-inner-word",
		Initial: initial,
		Goal: Goal{
			Lines:  []string{"mode="},
			Cursor: CursorGoal(0, 4),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyD, vimengine.KeyI, vimengine.KeyW},
		},
	})

	for _, key := range []string{vimengine.KeyD, vimengine.KeyI} {
		result := session.ApplyKey(key)
		if result.State.Status != StatusRunning {
			t.Fatalf("status after %q = %q, want running", key, result.State.Status)
		}
	}
	result := session.ApplyKey(vimengine.KeyW)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after w = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyD, vimengine.KeyI, vimengine.KeyW})
}

func TestSessionReplaysChangeInnerWordTrace(t *testing.T) {
	initial := vimengine.NewState([]string{"mode=broken"})
	initial.Cursor.Col = 7
	initial.Cursor.DesiredCol = 7
	session := NewSession(Exercise{
		ID:      "change-inner-word",
		Initial: initial,
		Goal: Goal{
			Lines:  []string{"mode=stable"},
			Cursor: CursorGoal(0, 10),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyC, vimengine.KeyI, vimengine.KeyW, vimengine.KeyEsc},
		},
	})

	for _, key := range []string{vimengine.KeyC, vimengine.KeyI, vimengine.KeyW, "s", "t", "a", "b", "l", "e"} {
		result := session.ApplyKey(key)
		if result.State.Status != StatusRunning {
			t.Fatalf("status after %q = %q, want running", key, result.State.Status)
		}
	}
	result := session.ApplyKey(vimengine.KeyEsc)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after esc = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyC, vimengine.KeyI, vimengine.KeyW, "s", "t", "a", "b", "l", "e", vimengine.KeyEsc})
}

func TestSessionReplaysYankInnerWordPutTrace(t *testing.T) {
	initial := vimengine.NewState([]string{"mode=stable", "mirror="})
	initial.Cursor.Col = 7
	initial.Cursor.DesiredCol = 7
	session := NewSession(Exercise{
		ID:      "yank-inner-word",
		Initial: initial,
		Goal: Goal{
			Lines:  []string{"mode=stable", "mirror=stable"},
			Cursor: CursorGoal(1, 12),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyY, vimengine.KeyI, vimengine.KeyW, vimengine.KeyP},
		},
	})

	for _, key := range []string{vimengine.KeyY, vimengine.KeyI, vimengine.KeyW, vimengine.KeyJ, vimengine.KeyDollar} {
		result := session.ApplyKey(key)
		if result.State.Status != StatusRunning {
			t.Fatalf("status after %q = %q, want running", key, result.State.Status)
		}
	}
	result := session.ApplyKey(vimengine.KeyP)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after p = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyY, vimengine.KeyI, vimengine.KeyW, vimengine.KeyJ, vimengine.KeyDollar, vimengine.KeyP})
}

func TestSessionReplaysQuoteTextObjectTrace(t *testing.T) {
	initial := vimengine.NewState([]string{`status="down"`})
	initial.Cursor.Col = 9
	initial.Cursor.DesiredCol = 9
	session := NewSession(Exercise{
		ID:      "change-inner-quote",
		Initial: initial,
		Goal: Goal{
			Lines:  []string{`status="up"`},
			Cursor: CursorGoal(0, 10),
			Mode:   ModeGoal(vimengine.ModeNormal),
		},
		Constraints: Constraints{
			RequiredKeys: []string{vimengine.KeyC, vimengine.KeyI, vimengine.KeyDoubleQuote, vimengine.KeyEsc},
		},
	})

	for _, key := range []string{vimengine.KeyC, vimengine.KeyI, vimengine.KeyDoubleQuote, "u", "p"} {
		result := session.ApplyKey(key)
		if result.State.Status != StatusRunning {
			t.Fatalf("status after %q = %q, want running", key, result.State.Status)
		}
	}
	result := session.ApplyKey(vimengine.KeyEsc)
	if result.State.Status != StatusSucceeded {
		t.Fatalf("status after esc = %q, want succeeded", result.State.Status)
	}
	assertTrace(t, result.State.KeyTrace, []string{vimengine.KeyC, vimengine.KeyI, vimengine.KeyDoubleQuote, "u", "p", vimengine.KeyEsc})
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

	hint, ok := session.CurrentHint()
	if !ok || hint != "move right" {
		t.Fatalf("hint before input = (%q,%v), want (%q,true)", hint, ok, "move right")
	}

	session.ApplyKey(vimengine.KeyH)
	hint, ok = session.CurrentHint()
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
