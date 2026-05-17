package vimengine

import "testing"

func TestNewStateStartsInNormalMode(t *testing.T) {
	state := NewState([]string{"alpha"})

	if state.Mode != ModeNormal {
		t.Fatalf("Mode = %q, want %q", state.Mode, ModeNormal)
	}
	if state.Cursor.Row != 0 || state.Cursor.Col != 0 {
		t.Fatalf("Cursor = (%d,%d), want (0,0)", state.Cursor.Row, state.Cursor.Col)
	}
}

func TestStateCopiesLines(t *testing.T) {
	lines := []string{"alpha"}
	engine := New(lines)
	lines[0] = "changed"

	state := engine.State()
	state.Lines[0] = "mutated"

	got := engine.State().Lines[0]
	if got != "alpha" {
		t.Fatalf("engine line = %q, want %q", got, "alpha")
	}
}

func TestHorizontalMovement(t *testing.T) {
	engine := New([]string{"abc"})

	assertApply(t, engine, "l", 0, 1, EventMoved)
	assertApply(t, engine, "l", 0, 2, EventMoved)
	assertApply(t, engine, "l", 0, 2, EventBoundary)
	assertApply(t, engine, "h", 0, 1, EventMoved)
	assertApply(t, engine, "h", 0, 0, EventMoved)
	assertApply(t, engine, "h", 0, 0, EventBoundary)
}

func TestVerticalMovementClampsToShorterLinesAndRestoresDesiredColumn(t *testing.T) {
	engine := New([]string{"abcd", "x", "wxyz"})

	assertApply(t, engine, "l", 0, 1, EventMoved)
	assertApply(t, engine, "l", 0, 2, EventMoved)
	assertApply(t, engine, "j", 1, 0, EventMoved)
	assertApply(t, engine, "j", 2, 2, EventMoved)
	assertApply(t, engine, "k", 1, 0, EventMoved)
	assertApply(t, engine, "k", 0, 2, EventMoved)
}

func TestVerticalMovementHandlesEmptyLines(t *testing.T) {
	engine := New([]string{"abc", "", "xyz"})

	assertApply(t, engine, "l", 0, 1, EventMoved)
	assertApply(t, engine, "j", 1, 0, EventMoved)
	assertApply(t, engine, "j", 2, 1, EventMoved)
}

func TestUnsupportedKeyDoesNotMove(t *testing.T) {
	engine := New([]string{"abc"})
	result := engine.Apply("w")

	if result.State.Cursor.Row != 0 || result.State.Cursor.Col != 0 {
		t.Fatalf("Cursor = (%d,%d), want (0,0)", result.State.Cursor.Row, result.State.Cursor.Col)
	}
	assertEvent(t, result, EventUnsupportedKey)
}

func TestApplyIsPureTransition(t *testing.T) {
	state := NewState([]string{"abc"})
	result := Apply(state, "l")

	if state.Cursor.Col != 0 {
		t.Fatalf("input state cursor col = %d, want 0", state.Cursor.Col)
	}
	if result.State.Cursor.Col != 1 {
		t.Fatalf("result cursor col = %d, want 1", result.State.Cursor.Col)
	}
}

func TestUnicodeColumnsUseRunes(t *testing.T) {
	engine := New([]string{"가나"})

	assertApply(t, engine, "l", 0, 1, EventMoved)
	assertApply(t, engine, "l", 0, 1, EventBoundary)
}

func assertApply(t *testing.T, engine *Engine, key string, row int, col int, eventType EventType) {
	t.Helper()

	result := engine.Apply(key)
	if result.State.Cursor.Row != row || result.State.Cursor.Col != col {
		t.Fatalf("after %q cursor = (%d,%d), want (%d,%d)", key, result.State.Cursor.Row, result.State.Cursor.Col, row, col)
	}
	assertEvent(t, result, eventType)
}

func assertEvent(t *testing.T, result Result, eventType EventType) {
	t.Helper()

	if len(result.Events) != 1 {
		t.Fatalf("len(Events) = %d, want 1", len(result.Events))
	}
	if result.Events[0].Type != eventType {
		t.Fatalf("event type = %q, want %q", result.Events[0].Type, eventType)
	}
}
